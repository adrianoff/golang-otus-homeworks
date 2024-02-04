package main

import (
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrDirectoryNotSupported = errors.New("directory not supported")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrSameFiles             = errors.New("source and destination files are the same")
	ErrWrongParams           = errors.New("offset and limit should be positive integers")
	ErrOpenSrcFile           = errors.New("open source file error")
	ErrRetrievingSrcFileInfo = errors.New("retrieving source file info error")
	ErrRetrievingDstFileInfo = errors.New("retrieving destination file info error")
	ErrOpenDstFileError      = errors.New("open destination file error")
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
		return ErrOpenSrcFile
	}
	defer sourceFile.Close()

	// Get source file info
	if sourceFileInfo, err = sourceFile.Stat(); err != nil {
		return ErrRetrievingSrcFileInfo
	}
	// check file is not device
	if sourceFileInfo.Mode()&fs.ModeDevice == fs.ModeDevice {
		return ErrUnsupportedFile
	}
	// check file is not directory
	if sourceFileInfo.IsDir() {
		return ErrDirectoryNotSupported
	}

	// Check offset > filesize
	filesize := sourceFileInfo.Size()
	if offset > filesize {
		return ErrOffsetExceedsFileSize
	}

	// Open destination file
	if destinationFile, err = os.Create(destinationPath); err != nil {
		return ErrOpenDstFileError
	}
	defer destinationFile.Close()

	// Check destination file is not the same as source file
	if destinationFileInfo, err = destinationFile.Stat(); err != nil {
		return ErrRetrievingDstFileInfo
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
		errRemove := os.Remove(destinationPath)
		if errRemove != nil {
			return errRemove
		}
		return err
	}

	bar.Finish()

	return nil
}
