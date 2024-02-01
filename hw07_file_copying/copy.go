package main

import (
	"errors"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	inFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	data := make([]byte, 10)

	for {
		n, err := inFile.ReadAt(data, offset)
		if err != nil {
			return err
		}
		offset = offset + int64(n)
	}

	//read, err = inFile.ReadAt(data, offset)
	//if err != nil {
	//	return err
	//}
	//
	//count := 100000
	//
	//// create and start new bar
	//bar := pb.StartNew(count)
	//
	//// start bar from 'default' template
	//// bar := pb.Default.Start(count)
	//
	//// start bar from 'simple' template
	//// bar := pb.Simple.Start(count)
	//
	//// start bar from 'full' template
	//// bar := pb.Full.Start(count)
	//
	//for i := 0; i < count; i++ {
	//	bar.Increment()
	//	time.Sleep(time.Millisecond)
	//}
	//
	//// finish bar
	//bar.Finish()

	// Place your code here.
	return nil
}
