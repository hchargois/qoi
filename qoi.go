// Package qoi implements a QOI image encoder and decoder.
//
// The QOI specification is at https://qoiformat.org/qoi-specification.pdf
//
// This package registers the decoder on import to be used by the standard
// library's [image.Decode].
package qoi

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
)

const Magic = "qoif"
const ChannelsRGB = 3
const ChannelsRGBA = 4
const ColorspaceSRGB = 0
const ColorspaceLinear = 1

var qoiEndStream = []byte{0, 0, 0, 0, 0, 0, 0, 1}

type qoiHeader struct {
	Magic      [4]byte
	Width      uint32
	Height     uint32
	Channels   uint8
	Colorspace uint8
}

func decodeHeader(r io.Reader) (qoiHeader, error) {
	var hdr qoiHeader
	err := binary.Read(r, binary.BigEndian, &hdr)
	if err != nil {
		return hdr, fmt.Errorf("decoding header: %w", err)
	}
	if string(hdr.Magic[:]) != Magic {
		return hdr, fmt.Errorf("invalid magic number in header: %v", hdr.Magic)
	}
	return hdr, nil
}

// DecodeConfig returns the dimensions of a QOI image without decoding the
// entire image. The color model is always color.NRGBAModel as QOI always uses
// this model when decoding.
func DecodeConfig(r io.Reader) (image.Config, error) {
	hdr, err := decodeHeader(r)
	if err != nil {
		return image.Config{}, err
	}
	return image.Config{
		ColorModel: color.NRGBAModel,
		Width:      int(hdr.Width),
		Height:     int(hdr.Height),
	}, nil
}

// ConfigExtra represents the standard library's [image.Config] extended with
// QOI-specific metadata.
type ConfigExtra struct {
	image.Config
	Channels   int
	Colorspace int
}

// DecodeConfigExtra returns the extra config data embedded in a QOI header.
func DecodeConfigExtra(r io.Reader) (ConfigExtra, error) {
	hdr, err := decodeHeader(r)
	if err != nil {
		return ConfigExtra{}, err
	}
	return ConfigExtra{
		Config: image.Config{
			ColorModel: color.NRGBAModel,
			Width:      int(hdr.Width),
			Height:     int(hdr.Height),
		},
		Channels:   int(hdr.Channels),
		Colorspace: int(hdr.Colorspace),
	}, nil
}

type pixel struct {
	r, g, b, a uint8
}

func (p pixel) hash() int {
	return int(p.r*3+p.g*5+p.b*7+p.a*11) % 64
}

type decoder struct {
	seen   [64]pixel
	prev   pixel
	pixLen int
	pix    []uint8
}

func (d *decoder) push(p pixel) {
	d.pix = append(d.pix, p.r, p.g, p.b, p.a)
	d.prev = p
	d.seen[p.hash()] = p
}

func (d *decoder) decode(stream []byte) error {
	d.pix = make([]uint8, 0, d.pixLen)

	for len(d.pix) < d.pixLen {
		if len(stream) < 5 {
			// the largest OP (RGBA) requires 5 bytes; it is always safe to
			// request that many bytes since the stream should always end with
			// an 8-byte tail anyway
			return errors.New("stream too short")
		}
		b := stream[0]
		stream = stream[1:]

		if b == 0xfe {
			// OP_RGB
			d.push(pixel{stream[0], stream[1], stream[2], d.prev.a})
			stream = stream[3:]
			continue
		}
		if b == 0xff {
			// OP_RGBA
			d.push(pixel{stream[0], stream[1], stream[2], stream[3]})
			stream = stream[4:]
			continue
		}
		op := b >> 6
		b &= 0x3f
		if op == 0 {
			// OP_INDEX
			p := d.seen[b]
			// not using push() to avoid recomputing the hash to put the same
			// pixel in "seen" right back where it's coming from...
			d.pix = append(d.pix, p.r, p.g, p.b, p.a)
			d.prev = p
			continue
		}
		if op == 1 {
			// OP_DIFF
			dr := b>>4 - 2
			dg := b>>2&3 - 2
			db := b&3 - 2
			d.push(pixel{d.prev.r + dr, d.prev.g + dg, d.prev.b + db, d.prev.a})
			continue
		}
		if op == 2 {
			// OP_LUMA
			b2 := stream[0]
			dg := b - 32
			drdg := b2>>4 - 8
			dbdg := b2&0xf - 8

			dr := drdg + dg
			db := dbdg + dg
			d.push(pixel{d.prev.r + dr, d.prev.g + dg, d.prev.b + db, d.prev.a})
			stream = stream[1:]
			continue
		}
		// op == 3
		// OP_RUN
		run := int(b + 1)
		// we need to update the seen index even for OP_RUN,
		// cf. https://github.com/phoboslab/qoi/issues/258
		d.seen[d.prev.hash()] = d.prev
		for range run {
			// not using push() since we don't need to update prev/seen
			d.pix = append(d.pix, d.prev.r, d.prev.g, d.prev.b, d.prev.a)
		}
	}

	if len(d.pix) > d.pixLen {
		// an invalid OP_RUN may have caused too many pixels to be output
		return errors.New("too many pixels in stream")
	}

	if !bytes.Equal(stream, qoiEndStream) {
		return errors.New("invalid stream end")
	}

	return nil
}

func decodeStream(hdr qoiHeader, stream []byte) (image.Image, error) {
	dec := decoder{
		prev:   pixel{0, 0, 0, 255},
		pixLen: int(hdr.Width * hdr.Height * 4),
	}

	err := dec.decode(stream)
	if err != nil {
		return nil, err
	}

	return &image.NRGBA{
		Pix:    dec.pix,
		Stride: int(hdr.Width * 4),
		Rect: image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{int(hdr.Width), int(hdr.Height)},
		},
	}, nil

}

// Decode decodes a QOI image. The returned [image.Image] is always an
// *[image.NRGBAImage].
func Decode(r io.Reader) (image.Image, error) {
	// the header decode is done before reading the rest of the data
	// to fail fast if the header is invalid
	hdr, err := decodeHeader(r)
	if err != nil {
		return nil, err
	}

	stream, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return decodeStream(hdr, stream)
}

// DecodeBytes decodes a QOI image. It is similar to [Decode] but is faster and
// should be preferred if you already have the data in a byte slice.
func DecodeBytes(b []byte) (image.Image, error) {
	if len(b) < 14+8 {
		return nil, errors.New("too few bytes")
	}
	hdr, err := decodeHeader(bytes.NewReader(b[:14]))
	if err != nil {
		return nil, err
	}
	stream := b[14:]

	return decodeStream(hdr, stream)
}

func serialize(img *image.NRGBA) []byte {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	if bounds.Min.X == 0 && bounds.Min.Y == 0 && len(img.Pix) == width*height*4 {
		return img.Pix
	}

	pix := make([]byte, 0, width*height*4)
	for y := range height {
		start := y * img.Stride
		end := start + width*4
		pix = append(pix, img.Pix[start:end]...)
	}
	return pix
}

// Options are encoding options.
type Options struct {
	// Channels should be [ChannelsRGB] or [ChannelsRGBA]. It is only used to
	// set the metadata in the header and has no impact on the encoding.
	Channels int
	// Colorspace should be [ColorspaceSRGB] or [ColorspaceLinear]. It is only
	// used to set the metadata in the header and has no impact on the encoding.
	Colorspace int
}

type encoder struct {
	w    io.Writer
	seen [64]pixel
	prev pixel
}

const encBufSize = 4 * 1024
const encBufLimit = encBufSize - 5

func (e *encoder) encode(pix []uint8) error {
	if len(pix)%4 != 0 {
		return errors.New("invalid pixel data")
	}
	buf := make([]byte, 0, encBufSize)
	var run byte
	for len(pix) >= 4 {
		if len(buf) >= encBufLimit {
			_, err := e.w.Write(buf)
			if err != nil {
				return err
			}
			buf = buf[:0]
		}

		p := pixel{pix[0], pix[1], pix[2], pix[3]}
		pix = pix[4:]
		if p == e.prev {
			if run == 0 {
				// we need to update the seen index even for OP_RUN,
				// cf. https://github.com/phoboslab/qoi/issues/258
				e.seen[p.hash()] = p
			}
			if run == 62 {
				// OP_RUN
				buf = append(buf, (run-1)|0xc0)
				run = 0
			}
			run += 1
			continue
		}
		if run > 0 {
			// OP_RUN
			buf = append(buf, (run-1)|0xc0)
			run = 0
		}
		h := p.hash()
		if e.seen[h] == p {
			// OP_INDEX
			buf = append(buf, byte(h))
			e.prev = p
			continue
		}
		e.seen[h] = p
		if p.a != e.prev.a {
			// The remaining ops (DIFF, LUMA and RGB) all need the alpha to be
			// unchanged, so if it's not the case we have no choice but to
			// use OP_RGBA
			buf = append(buf, 0xff, p.r, p.g, p.b, p.a)
			e.prev = p
		} else {
			dr := int8(p.r - e.prev.r)
			dg := int8(p.g - e.prev.g)
			db := int8(p.b - e.prev.b)
			dr_dg := dr - dg
			db_dg := db - dg
			e.prev = p

			if uint8((dr+2)|(dg+2)|(db+2))&0xfc == 0 {
				// same as dr >= -2 && dr <= 1 && dg >= -2 && dg <= 1 && db >= -2 && db <= 1
				// but with a single branch
				// OP_DIFF
				buf = append(buf, 0x40|uint8((dr+2)<<4|(dg+2)<<2|(db+2)))
			} else if (uint8(dg+32)&0xc0)|(uint8((dr_dg+8)|(db_dg+8))&0xf0) == 0 {
				// same as dg >= -32 && dg <= 31 && dr_dg >= -8 && dr_dg <= 7 && db_dg >= -8 && db_dg <= 7
				// OP_LUMA
				buf = append(buf,
					0x80|uint8(dg+32),
					uint8(dr_dg+8)<<4|uint8(db_dg+8),
				)
			} else {
				// OP_RGB
				buf = append(buf, 0xfe, p.r, p.g, p.b)
			}
		}
	}
	if run > 0 {
		// OP_RUN
		buf = append(buf, (run-1)|0xc0)
	}
	if len(buf) != 0 {
		_, err := e.w.Write(buf)
		if err != nil {
			return err
		}
	}
	e.w.Write(qoiEndStream)

	return nil
}

func encodePixels(w io.Writer, pix []uint8) error {
	enc := encoder{
		w:    w,
		prev: pixel{0, 0, 0, 255},
	}
	return enc.encode(pix)
}

// Encode encodes the image to its QOI representation. If nil Options are
// passed, defaults of ChannelsRGBA and ColorspaceSRGB are used.
func Encode(w io.Writer, img image.Image, o *Options) error {
	if o == nil {
		o = &Options{
			Channels:   ChannelsRGBA,
			Colorspace: ColorspaceSRGB,
		}
	} else {
		if o.Channels != ChannelsRGB && o.Channels != ChannelsRGBA {
			return errors.New("invalid channels option")
		}
		if o.Colorspace != ColorspaceSRGB && o.Colorspace != ColorspaceLinear {
			return errors.New("invalid colorspace option")
		}
	}

	nrgba := convertToNRGBA(img)
	hdr := qoiHeader{
		Magic:      [4]byte{'q', 'o', 'i', 'f'},
		Width:      uint32(img.Bounds().Dx()),
		Height:     uint32(img.Bounds().Dy()),
		Channels:   uint8(o.Channels),
		Colorspace: uint8(o.Colorspace),
	}
	err := binary.Write(w, binary.BigEndian, hdr)
	if err != nil {
		return fmt.Errorf("writing header: %w", err)
	}

	pix := serialize(nrgba)
	err = encodePixels(w, pix)
	return err
}

// convertToNRGBA converts any image.Image to *image.NRGBA
func convertToNRGBA(src image.Image) *image.NRGBA {
	if nrgba, ok := src.(*image.NRGBA); ok {
		return nrgba
	}

	bounds := src.Bounds()
	dst := image.NewNRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}

func init() {
	image.RegisterFormat("qoi", Magic, Decode, DecodeConfig)
}
