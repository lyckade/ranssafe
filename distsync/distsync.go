package distsync

import (
	"io"

	"github.com/as27/ranssafe/fileinfo"
)

// Distsyncer interface is used for syncing files between
// server and client. The interface lets the implementation
// of the protocol open.
type Distsyncer interface {
	GetSrcFileInfo() []fileinfo.File
	GetDistFileInfo() []fileinfo.File
	SkipFile(string) bool
	PushFile(string) error
	GetFile(string) (io.Writer, error)
}
