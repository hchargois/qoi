package qoi_test

import (
	"bytes"
	"image"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/hchargois/qoi"
	//"github.com/xfmoulet/qoi"
	//"github.com/takeyourhatoff/qoi"
	//qoi "github.com/arian/go-qoi"
)

func findTestFiles(tb testing.TB) []string {
	files, err := filepath.Glob("testdata/*.qoi")
	if err != nil {
		tb.Fatal(err)
	}
	return files
}

func BenchmarkDecode(b *testing.B) {
	files := findTestFiles(b)

	for _, file := range files {
		b.Run(filepath.Base(file), func(b *testing.B) {
			fd, err := os.Open(file)
			if err != nil {
				b.Fatal(err)
			}

			buf, err := io.ReadAll(fd)
			if err != nil {
				b.Fatal(err)
			}

			err = fd.Close()
			if err != nil {
				b.Fatal(err)
			}

			for b.Loop() {
				_, err := qoi.Decode(bytes.NewReader(buf))
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	files := findTestFiles(b)

	for _, file := range files {
		b.Run(filepath.Base(file), func(b *testing.B) {
			fd, err := os.Open(file)
			if err != nil {
				b.Fatal(err)
			}

			buf, err := io.ReadAll(fd)
			if err != nil {
				b.Fatal(err)
			}

			err = fd.Close()
			if err != nil {
				b.Fatal(err)
			}

			img, _, err := image.Decode(bytes.NewReader(buf))
			if err != nil {
				b.Fatal(err)
			}

			var outBuf bytes.Buffer
			for b.Loop() {
				qoi.Encode(&outBuf, img, nil)
				if err != nil {
					b.Fatal(err)
				}
				outBuf.Reset()
			}
		})
	}
}

func TestRoundtrip(t *testing.T) {
	files := findTestFiles(t)

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			fd, err := os.Open(file)
			if err != nil {
				t.Fatal(err)
			}

			qoiBytes, err := io.ReadAll(fd)
			if err != nil {
				t.Fatal(err)
			}

			err = fd.Close()
			if err != nil {
				t.Fatal(err)
			}

			// There's no way to keep track of Channels and Colorspace info
			// in stdlib's image.Image, so we'll extract and copy them explicitly.
			cfg, err := qoi.DecodeConfigExtra(bytes.NewReader(qoiBytes))
			if err != nil {
				t.Fatal(err)
			}

			img, err := qoi.DecodeBytes(qoiBytes)
			if err != nil {
				t.Fatal(err)
			}

			var out bytes.Buffer
			err = qoi.Encode(&out, img, &qoi.Options{
				Channels:   cfg.Channels,
				Colorspace: cfg.Colorspace,
			})
			if err != nil {
				t.Fatal(err)
			}

			outBytes := out.Bytes()
			outLen, expLen := len(outBytes), len(qoiBytes)
			if outLen != expLen {
				t.Errorf("unequal lengths, got %d expected %d", outLen, expLen)
			}

			for i := 0; i < min(outLen, expLen); i++ {
				if outBytes[i] != qoiBytes[i] {
					t.Errorf("first difference at byte %d", i)
					break
				}
			}
		})
	}
}
