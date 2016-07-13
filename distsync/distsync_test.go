package distsync

import (
	"io"
	"reflect"
	"testing"

	"github.com/as27/ranssafe/fileinfo"
)

type testDistsyncer struct {
	pushFiles []string
	getFiles  []string
}

func (ds *testDistsyncer) GetSrcFileInfo() []fileinfo.File {
	testFiles := []fileinfo.File{
		{"a/b/c/file1", 2016010131000000},
		{"a/b/c/sametimestamp", 2016010131000000},
		{"a/b/newer", 2016010131000000},
		{"a/b/c/older", 2015010131000000},
	}
	return testFiles
}

func (ds *testDistsyncer) GetDistFileInfo() ([]fileinfo.File, error) {
	testFiles := []fileinfo.File{
		{"a/b/c/file2", 2016010131000000},
		{"a/b/c/sametimestamp", 2016010131000000},
		{"a/b/newer", 2015010131000000},
		{"a/b/c/older", 2016010131000000},
	}
	return testFiles, nil
}

/*func (ds *testDistsyncer) SkipFile(fp string) bool {
	return false
}*/

func (ds *testDistsyncer) PushFile(fp string) error {
	ds.pushFiles = append(ds.pushFiles, fp)
	return nil
}

func (ds *testDistsyncer) GetFile(fp string) (io.Writer, error) {
	ds.getFiles = append(ds.getFiles, fp)
	return nil, nil
}

// Check the correct implementation
var _ Distsyncer = &testDistsyncer{}

func TestDistsync(t *testing.T) {
	ds := &testDistsyncer{}
	Distsync(ds)
	expectPush := []string{
		"a/b/c/file1",
		"a/b/newer",
	}
	if reflect.DeepEqual(ds.pushFiles, expectPush) != true {
		t.Fatalf("Push files:\nExpect: %#v\nGot:%#v\n", expectPush, ds.pushFiles)
	}
	expectGet := []string{
		"a/b/c/older",
	}
	if reflect.DeepEqual(ds.getFiles, expectGet) != true {

		t.Fatalf("Get files:\nExpect: %#v\nGot:%#v\n", expectGet, ds.getFiles)
	}
}
