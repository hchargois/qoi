package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/hchargois/qoi"
	"golang.org/x/sync/errgroup"
)

func usage() {
	fmt.Println(`qoiconv <infile> <outfile>
qoiconv [infiles...]

qoiconv converts between QOI and PNG or JPEG.

If exactly two arguments are passed and the second one does not point to an
existing file, the first form is used and qoiconv converts infile to outfile,
inferring the output format based on the extension of outfile.

In all other cases, the second form is used, and all infiles are converted
to output files consisting of the same name with a modified extension. QOI
input files are converted to PNG, and PNG or JPEG input files are converted to
QOI. The output files must not already exist. Conversions are done in parallel
in as many workers as the number of CPUs.`)
}

func fileExists(fp string) bool {
	_, error := os.Stat(fp)
	return !errors.Is(error, os.ErrNotExist)
}

type encoder func(w io.Writer, img image.Image) error

func encoderFromExt(fp string) (encoder, error) {
	ext := filepath.Ext(fp)
	switch ext {
	case ".png":
		return png.Encode, nil
	case ".jpeg", ".jpg":
		return func(w io.Writer, img image.Image) error {
			return jpeg.Encode(w, img, nil)
		}, nil
	case ".qoi":
		return func(w io.Writer, img image.Image) error {
			return qoi.Encode(w, img, nil)
		}, nil
	default:
		return nil, fmt.Errorf("unknown output file extension %s", ext)
	}
}

func conv(fromPath, toPath string) error {
	from, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("opening source file: %w", err)
	}
	defer from.Close()

	img, _, err := image.Decode(from)
	if err != nil {
		return fmt.Errorf("decoding source file: %w", err)
	}

	enc, err := encoderFromExt(toPath)
	if err != nil {
		return err
	}

	to, err := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return fmt.Errorf("opening output file: %w", err)
	}

	err = enc(to, img)
	if err != nil {
		err = fmt.Errorf("encoding to output file: %w", err)
	}
	return errors.Join(err, to.Close())
}

type task struct {
	in, out string
}

func newTask(in string) task {
	ext := filepath.Ext(in)
	var outExt string
	if ext == ".qoi" {
		outExt = ".png"
	} else {
		outExt = ".qoi"
	}
	out := in[:len(in)-len(ext)] + outExt
	return task{in, out}
}

func checkConflicts(tasks []task) {
	// check that no tasks are conflicting i.e. one wants to convert
	// img.qoi to img.png, and another img.png to img.qoi; as these two tasks
	// when executing concurrently could do strange things and maybe even cause
	// data loss
	seen := make(map[string]struct{})
	for _, task := range tasks {
		seen[task.in] = struct{}{}
	}
	for _, task := range tasks {
		if _, ok := seen[task.out]; ok {
			fmt.Fprintln(os.Stderr, "conflicting conversions to and from", task.out)
			os.Exit(1)
		}
	}
}

func checkNotExist(tasks []task) {
	for _, task := range tasks {
		if fileExists(task.out) {
			fmt.Fprintln(os.Stderr, "output file already exists", task.out)
			os.Exit(1)
		}
	}
}

func main() {
	if len(os.Args) == 1 || len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		usage()
		return
	}

	if len(os.Args) == 3 && !fileExists(os.Args[2]) {
		err := conv(os.Args[1], os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	tasks := make([]task, 0, len(os.Args)-1)
	for _, fp := range os.Args[1:] {
		tasks = append(tasks, newTask(fp))
	}

	checkConflicts(tasks)
	checkNotExist(tasks)

	var eg errgroup.Group
	eg.SetLimit(runtime.NumCPU())
	for _, task := range tasks {
		eg.Go(func() error {
			err := conv(task.in, task.out)
			if err != nil {
				fmt.Printf("%s -> %s: error: %v\n", task.in, task.out, err)
			} else {
				fmt.Printf("%s -> %s\n", task.in, task.out)
			}
			return err
		})
	}
	eg.Wait()
}
