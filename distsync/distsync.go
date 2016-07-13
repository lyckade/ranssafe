package distsync

import (
	"io"

	"github.com/as27/ranssafe/fileinfo"
)

// Distsyncer interface is used for syncing files between
// server and client. The interface lets the implementation
// of the protocol open.
// The sync process is very simple. First the files which have
// to be synced will be returned by GetSrcFileInfo(). All files
// inside this slice are going to be synced. Patterns for skipping
// folders or files have to be includes inside that method.
// The next step gets all the fileInfos from the dist location.
// If the local file is newer it will be pushed, if not the file
// is going to be loaded from the dist path.
type Distsyncer interface {
	GetSrcFileInfo() []fileinfo.File
	GetDistFileInfo() ([]fileinfo.File, error)
	//SkipFile(string) bool
	PushFile(string) error
	GetFile(string) (io.Writer, error)
}

// Distsync uses the Distsyncer interface to sync between different
// locations.
func Distsync(ds Distsyncer) error {
	distFileInfos, err := ds.GetDistFileInfo()
	if err != nil {
		return err
	}
	for _, fi := range ds.GetSrcFileInfo() {
		switch {
		case isPushFile(&fi, &distFileInfos):
			ds.PushFile(fi.FilePath)
		case isGetFile(&fi, &distFileInfos):
			_, err := ds.GetFile(fi.FilePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isPushFile(fi *fileinfo.File, fInfos *[]fileinfo.File) bool {
	df, ok := getDistFile(fi, fInfos)
	if ok == false {
		return true
	}
	if fi.Timestamp > df.Timestamp {
		return true
	}
	return false

}

func isGetFile(fi *fileinfo.File, fInfos *[]fileinfo.File) bool {
	df, ok := getDistFile(fi, fInfos)
	if ok == false {
		return false
	}
	if fi.Timestamp < df.Timestamp {
		return true
	}
	return false
}

func getDistFile(fi *fileinfo.File, fInfos *[]fileinfo.File) (*fileinfo.File, bool) {
	for _, f := range *fInfos {
		if fi.FilePath == f.FilePath {
			return &f, true
		}
	}
	return nil, false
}
