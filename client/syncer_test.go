package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/as27/ranssafe/fileinfo"
)

func TestAddFile(t *testing.T) {
	s := NewSyncer("myServer")
	s.newFileinfo = func(fp string) (fileinfo.File, error) {
		if fp == "ErrorPath" {
			return fileinfo.File{}, errors.New("Error creating fileinfo")
		}
		return fileinfo.File{
			FilePath:  fp,
			Timestamp: 1,
		}, nil
	}
	s.AddFile("filePath1")
	err := s.AddFile("filePath2")
	if err != nil {
		t.Fatal(err)
	}
	expect := []fileinfo.File{
		{"filePath1", 1},
		{"filePath2", 1},
	}
	if reflect.DeepEqual(expect, s.files) != true {
		t.Fatalf("AddFile not correct!\nExpect:%v\nGot:%v", expect, s.files)
	}
	err = s.AddFile("ErrorPath")
	if err == nil {
		t.Fatal("AddFile - expected one error.")
	}
}
func TestGetDisFileInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}
