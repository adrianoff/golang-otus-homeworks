package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrSameFiles             = errors.New("source and destination files are the same")
	ErrWrongParams           = errors.New("offset and limit should be positive integers")
)

func Copy(sourcePath, destinationPath string, offset, limit int64) error {
	// Check offset and limit should be positive integers
	if offset < 0 || limit < 0 {
		return ErrWrongParams
	}

	var sourceFile, destinationFile *os.File
	var sourceFileInfo, destinationFileInfo os.FileInfo
	var err error

	// Open source file
	if sourceFile, err = os.Open(sourcePath); err != nil {
		return fmt.Errorf("open source file error: %w", err)
	}
	defer sourceFile.Close()

	// Get source file info and check file is not device
	if sourceFileInfo, err = sourceFile.Stat(); err != nil {
		return fmt.Errorf("retrieving source file info error: %w", err)
	}
	if isSpecificMode(sourceFileInfo.Mode(), fs.ModeDevice) {
		return ErrUnsupportedFile
	}

	// Check offset > filesize
	filesize := sourceFileInfo.Size()
	if offset > filesize {
		return ErrOffsetExceedsFileSize
	}

	// Open destination file
	if destinationFile, err = os.Create(destinationPath); err != nil {
		return fmt.Errorf("open destination file error: %w", err)
	}
	defer destinationFile.Close()

	// Check destination file is not the same as source file
	if destinationFileInfo, err = destinationFile.Stat(); err != nil {
		return fmt.Errorf("retrieving destination file info error: %w", err)
	}
	if os.SameFile(sourceFileInfo, destinationFileInfo) {
		return ErrSameFiles
	}

	if offset != 0 { // Let`s avoid Seek if offset is 0. Cursor is already on 0.
		_, err = sourceFile.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}

	if limit == 0 {
		limit = filesize - offset // If limit == 0 than copy 'filesize-offset' bytes max
	} else if offset+limit > filesize {
		limit = filesize - offset // if offset+limit > filesize than copy 'filesize - offset' bytes max
	}
	reader := io.LimitReader(sourceFile, limit)
	writer := destinationFile

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(reader)

	_, err = io.Copy(writer, barReader)
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}

func isSpecificMode(mode, targetMode os.FileMode) bool {
	return mode&targetMode == targetMode
}
