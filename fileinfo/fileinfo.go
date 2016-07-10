// Package fileinfo is for the exange of metadata of each file. The package is
// used by server and client.
package fileinfo

// File is the metadata of a file, which is used to compare files before copying.
type File struct {
	//FilePath is the path to the file starting at the root folder of each package
	FilePath string `json:"filePath"`
	//Timestamp represents the last modification time
	Timestamp int `json:"timestamp"`
}

// The Encoder interface represents a Encode method of any encoding library
// the exchange format is open at this point
type Encoder interface {
	Encode(interface{}) error
}

// The Decoder interface represents a Decode method analog to the Encoder
type Decoder interface {
	Decode(interface{}) error
}
