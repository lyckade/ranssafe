package main

import (
	"io"

	"github.com/as27/ranssafe/fileinfo"
)

// Syncer is a implementation of the Distsyncer interface
type Syncer struct {
	Files []fileinfo.File
}

// AddFile adds one file to the Syncer
func (s *Syncer) AddFile(fp string) error {
	fi, err := fileinfo.New(fp)
	if err != nil {
		return err
	}
	s.Files = append(s.Files, fi)
	return nil
}

// GetSrcFileInfo implements the distsync interface
func (s *Syncer) GetSrcFileInfo() []fileinfo.File {
	return s.Files
}

// GetDistFileInfo implements the distsync interface
func (s *Syncer) GetDistFileInfo() []fileinfo.File {
	return nil
}

// PushFile implements the distsync interface
func (s *Syncer) PushFile(string) error {
	return nil
}

// GetFile implements the distsync interface
func (s *Syncer) GetFile(string) (io.Writer, error) {
	return nil, nil
}
