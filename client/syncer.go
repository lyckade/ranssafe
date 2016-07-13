package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/as27/ranssafe/fileinfo"
)

// ServerURL is the default value for every new syncer
//var ServerURL = "http://localhost:1234"

// Syncer is a implementation of the Distsyncer interface
type Syncer struct {
	// ServerAdress is the http adess of the server including
	// the port and the package
	// https://localhost:1234/mypackage
	ServerURL   string
	files       []fileinfo.File
	newFileinfo func(string) (fileinfo.File, error)
}

// NewSyncer takes a serverAdress and returns a pointer to a
// syncer
func NewSyncer(serverURL string) *Syncer {
	s := Syncer{ServerURL: serverURL}
	s.newFileinfo = fileinfo.New
	return &s
}

// AddFile adds one file to the Syncer
func (s *Syncer) AddFile(fp string) error {
	fi, err := s.newFileinfo(fp)
	if err != nil {
		return err
	}
	s.files = append(s.files, fi)
	return nil
}

// GetSrcFileInfo implements the distsync interface
func (s *Syncer) GetSrcFileInfo() []fileinfo.File {
	return s.files
}

// GetDistFileInfo implements the distsync interface
func (s *Syncer) GetDistFileInfo() []fileinfo.File {
	res, err := http.Get(s.ServerURL + "/fileinfo")
	if err != nil {
		log.Fatal(err)
	}
	rbody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", string(rbody))
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
