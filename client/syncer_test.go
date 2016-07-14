package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/as27/ranssafe/distsync"
	"github.com/as27/ranssafe/fileinfo"
)

var _ distsync.Distsyncer = NewSyncer("")

func TestGetSrcInfo(t *testing.T) {
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
	f := s.GetSrcFileInfo()
	if reflect.DeepEqual(expect, f) != true {
		t.Fatalf("AddFile not correct!\nExpect:%v\nGot:%v", expect, f)
	}
	err = s.AddFile("ErrorPath")
	if err == nil {
		t.Fatal("AddFile - expected one error.")
	}
}

func TestGetDistFileInfo(t *testing.T) {
	var urlStr string
	fi := []fileinfo.File{
		{"filePath1", 1},
		{"filePath2", 1},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlStr = r.URL.String()
		b, _ := json.Marshal(fi)
		w.Write(b)
	}))
	defer ts.Close()
	s := NewSyncer(ts.URL + "/package")
	respFi, err := s.GetDistFileInfo()
	if err != nil {
		t.Fatalf("No error expected!\nGot: %s", err)
	}
	// Check requested URL
	expect := "/package/fileinfo"
	if urlStr != expect {
		t.Fatalf("Wrong URL!\nExpectes:%v\nGot:%v", expect, urlStr)
	}
	// Check response
	if reflect.DeepEqual(fi, respFi) != true {
		t.Fatal("Response is different then definition at the server.")
	}

}

func TestPushFile(t *testing.T) {
	var urlStr string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlStr = r.URL.String()
		//b, _ := json.Marshal(fi)
		//w.Write(b)
	}))
	defer ts.Close()
}
