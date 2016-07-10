package distsync

import "github.com/as27/ranssafe/fileinfo"

type testDistsyncer struct{}

func (ds *testDistsyncer) GetSrcFileInfo() []fileinfo.File {
	return nil
}

func (ds *testDistsyncer) GetDistFileInfo() []fileinfo.File {
	return nil
}

func (ds *testDistsyncer) SkipFile(fp string) bool {
	return false
}

func (ds *testDistsyncer) PushFile(fp string) error {
	return nil
}

func (ds *testDistsyncer) GetFile(fp string) error {
	return nil
}

var _ Distsyncer = &testDistsyncer{}
