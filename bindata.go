// Code generated by go-bindata. DO NOT EDIT.
// sources:
// frontend/app.js
// frontend/index.html

package main


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataFrontendAppjs = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\x41\x6f\xe3\x2c\x10\xbd\xe7\x57\x8c\xd4\x48\xe0\xca\x21\x5f\x7b" +
	"\xfa\xe4\x2a\x95\xbe\x7e\x8d\xd2\x9e\x7a\xd8\x4a\x7b\x5c\x11\x33\xb6\xa9\x30\x58\x80\x93\x76\xab\xfc\xf7\x95\x31" +
	"\x38\xce\x6e\xb5\xd7\x99\xc7\x9b\xc7\x9b\x07\xeb\x6b\xa8\x95\xd9\x73\x05\xcb\xa2\xe2\xca\x21\x3c\x71\x2d\x14\xee" +
	"\xb9\x75\xb1\x70\xbd\x5e\x2c\x4a\xc5\x9d\x83\x1d\x57\x0a\xed\x07\x7c\x2e\x00\x4a\xa3\x9d\xb7\x7d\xe9\x8d\xa5\x59" +
	"\xa8\xc0\xec\x28\xb3\x58\x4b\xe7\xd1\x3e\xa1\xea\xd0\x52\x22\xab\x9d\x27\x39\xf8\x46\x3a\xd6\x4c\xb0\x07\x59\xd7" +
	"\x68\x5f\x1b\xae\xb3\xbb\xc0\x10\xfa\x68\xad\xb1\x3f\x7c\xa7\x60\x33\xa7\x2c\x4d\xdb\x49\x85\x74\x49\xc9\x55\x80" +
	"\xac\x3c\xb6\x9d\xe2\x1e\x49\xc6\x1a\xdf\x2a\x9a\xcd\x59\xb8\xda\xf7\xed\xdf\x59\x02\xe4\x4b\x96\x33\x8d\xe9\xbc" +
	"\x34\xda\xa1\x87\x4d\xbc\x25\xc0\x7a\x0d\x3b\xd4\x68\xb9\x82\xd8\x8e\x0d\x61\x8e\x5a\x19\x2e\x0a\x08\xd6\xe5\xb1" +
	"\xec\x50\xe1\x60\x54\x01\x84\xc9\xb6\x26\xf9\x99\xe7\xb5\xe9\xdb\xbd\xe6\x52\x25\x0a\x9f\x0a\x05\x78\xdb\x4f\x14" +
	"\xa1\xfc\x5d\x0a\xdf\x14\xf0\xef\x3f\xa9\x8a\xef\xe1\xfc\x73\xcb\x6b\x2c\x80\x38\x5b\x4e\xdc\xae\x31\xc7\xd0\x7c" +
	"\xf8\x78\xc4\x8a\xf7\xca\x5f\x88\x3a\x0d\x4e\x9d\x86\x7b\x0a\xe9\x3a\xc5\x3f\xb6\x83\xa3\xb4\x45\xe7\x78\x8d\xb0" +
	"\x01\x42\xd2\x56\x97\x94\xb0\xd2\x68\xcf\xa5\x46\xbb\xaa\x54\x2f\x05\xc9\x18\xef\x3a\xd4\x82\x5e\x2e\x8c\x26\x87" +
	"\x22\x4f\x9c\x35\xee\x25\x4c\xfb\x6a\xf7\xf4\x70\x93\xc3\xe1\x36\x4f\x66\xa6\xc1\xb2\x02\x7a\xb8\x81\x7b\x38\xdc" +
	"\x66\x93\xf7\x16\x7d\x6f\x75\x82\xb2\x4a\x07\x09\x71\xf1\xa7\xc5\x17\x10\xa9\x0f\x68\x1d\x9e\x71\x41\x88\xd4\xd2" +
	"\x4b\xae\xe4\x4f\xa4\xbd\x95\xd3\x5d\x19\x7f\xe3\xef\x43\x25\x9f\x26\x86\xdb\x15\x40\x33\xd8\xdc\x8f\x99\xb8\xb0" +
	"\x8c\xfc\x6f\x7a\x25\x40\x1b\x0f\x15\xfa\xb2\x81\x90\x2a\xc7\x48\x96\x4f\x66\xf8\xc6\x88\x02\xc8\x6e\xfb\x7a\xde" +
	"\x4f\x5f\x96\xe8\x5c\x01\x54\x70\xcf\xcf\xdc\x9d\x35\x43\xfd\xbf\x40\x32\xf6\x92\x8b\x93\xf6\x4b\x4c\x98\xf7\x18" +
	"\x48\x46\xc9\x0a\x7d\x14\x01\x1b\x98\xba\xe3\x83\x70\xa3\x51\x51\xa3\x33\xd6\x53\xca\x73\xd8\x07\x01\x7b\x26\x05" +
	"\xac\x80\x33\x29\xd2\x1b\x88\xc0\xca\xd8\x2d\x2f\x1b\x3a\x0e\x0b\xe0\x64\xcf\x34\x6d\xab\xb0\x45\x3d\xbc\x93\x25" +
	"\xbd\x7c\x81\xf1\x54\x5c\x52\x48\xd4\x55\x3d\xfe\x25\xab\xa3\xe5\xdd\x39\x4e\x73\xa2\x09\x3e\x2f\x32\x25\xeb\xc6" +
	"\xc7\x8f\x68\x1c\xd3\x77\x82\x7b\x14\x2f\x61\xdd\xdf\xd0\x4f\x21\x04\x88\x43\x9e\x45\x31\x92\x30\x29\x92\xff\xa7" +
	"\x24\x67\xe6\xeb\x1f\x4c\xa6\xf3\x83\x87\x9f\xa7\x64\x6d\x8c\xd6\xcb\xfe\x0d\x4b\xcf\xb8\x73\xb2\xd6\xf4\xf3\x94" +
	"\xff\xf6\x57\x84\x28\xa7\xb0\x9d\x16\x8b\x25\xa5\x93\x67\x83\x5f\x51\x17\x6c\x40\xe3\x31\x7d\xab\x34\xe0\x63\x8b" +
	"\xcd\xf2\x49\xe2\x12\xde\x9c\xd1\x24\xbb\x5b\x0c\x8a\x7f\x05\x00\x00\xff\xff\x96\x32\x8b\x33\xb8\x05\x00\x00")

func bindataFrontendAppjsBytes() ([]byte, error) {
	return bindataRead(
		_bindataFrontendAppjs,
		"frontend/app.js",
	)
}



func bindataFrontendAppjs() (*asset, error) {
	bytes, err := bindataFrontendAppjsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "frontend/app.js",
		size: 1464,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataFrontendIndexhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x6b\x53\xdb\xb8\xd7\x7f\xbf\x9f\x42\xeb\xce\x33\x43\x87\xca\x8e" +
	"\x73\x23\xa1\x09\xb3\x10\x4a\xc3\x25\x5c\x02\xb4\xd0\x77\xb2\x2d\xdb\x4a\x64\xc9\x48\x72\x2e\x64\xf8\xee\xcf\xc8" +
	"\x97\xc4\x49\x49\x61\x77\xbb\xf3\x4f\xa7\x10\x4b\x47\x47\xbf\xdf\xb9\xf8\x9c\x43\xe7\x4f\x8f\xbb\x6a\x1e\x63\x10" +
	"\xaa\x88\x1e\xfc\xd1\xd1\xbf\x00\x45\x2c\xe8\x1a\x98\x19\x07\x7f\x00\xd0\x09\x31\xf2\xf4\x17\x00\x3a\x7f\x42\x08" +
	"\x86\xf8\x29\x21\x02\x7b\x20\xc2\x0a\x01\x85\x02\x09\x20\xcc\xf7\xd3\x25\x37\x44\x42\x62\xd5\x35\x12\xe5\xc3\x96" +
	"\x51\xde\x0a\x95\x8a\xa1\x3e\x3f\xe9\x1a\x0f\xf0\xfe\x10\xf6\x78\x14\x23\x45\x1c\x8a\x0d\xe0\x72\xa6\x30\x53\x5d" +
	"\xe3\xf4\x4b\x17\x7b\x01\x5e\x3b\xc9\x50\x84\xbb\xc6\x84\xe0\x69\xcc\x85\x2a\x09\x4f\x89\xa7\xc2\xae\x87\x27\xc4" +
	"\xc5\x30\x7d\xf8\x04\x08\x23\x8a\x20\x0a\xa5\x8b\x28\xee\xda\x9f\x80\x0c\x05\x61\x63\xa8\x38\xf4\x89\xea\x32\x6e" +
	"\x1c\xfc\xb1\x22\x74\xc4\xb9\x92\x4a\xa0\x18\xf4\x6e\x6f\x57\x5c\x28\x61\x63\x20\x30\xed\x1a\x52\xcd\x29\x96\x21" +
	"\xc6\xca\x00\xa1\xc0\x7e\xd7\xd0\x3c\xe4\xbe\x65\xb9\x1e\x1b\x49\xd3\xa5\x3c\xf1\x7c\x8a\x04\x36\x5d\x1e\x59\x68" +
	"\x84\x66\x16\x25\x8e\xb4\xd4\x94\x28\x85\x05\x74\x8a\x1b\xac\xba\x59\x31\x2b\x96\x2b\xa5\xb5\x5c\x33\x23\xc2\x4c" +
	"\x57\x4a\x23\xbd\x36\xfb\x10\xa6\x70\x20\x88\x9a\x77\x0d\x19\xa2\x6a\xa3\x09\x2f\x0e\x5b\xed\xe7\xdd\x71\xdb\x1f" +
	"\x05\x83\xf3\x1b\x6b\xfc\x54\xbf\xba\xaa\x0e\x84\xdf\xfa\x46\xd5\x63\x44\xad\x6f\x5f\x6e\x77\x87\x41\xc5\x0f\xab" +
	"\x95\xae\x01\x5c\xc1\xa5\xe4\x82\x04\x84\x75\x0d\xc4\x38\x9b\x47\x3c\x91\x06\xb0\x7e\x1f\xb7\x14\xff\x14\x29\x37" +
	"\xcc\x49\x79\x48\x8c\xe9\xfc\xef\xf2\x72\xc2\xe3\xba\x75\xec\xf5\x0e\xe3\xcb\xc3\x47\x7a\x15\x4f\xce\xa5\xbc\x47" +
	"\xd5\xf0\xb8\x72\x9c\xb4\x66\x22\xb8\x9a\xa8\xc7\xfa\xb4\xda\x90\xe1\xe0\xd7\xbc\x7e\x1b\x31\x4a\x82\x50\x05\x88" +
	"\x52\x2c\xe6\x96\x6d\x36\xcd\x76\xea\xaf\xf2\xf2\xfb\xa8\xb5\x84\xdf\x77\xce\xc4\xee\x28\xa9\x5d\xb9\x95\x76\x7b" +
	"\x74\x72\x36\x18\xda\xb3\xc3\xeb\xa4\xd5\xbb\xbb\xee\xdf\xb7\x92\xeb\xc6\x59\xed\xc1\xfa\x76\xf8\xf8\xbb\xa8\x25" +
	"\x12\x9b\x3e\x67\x0a\x4d\xb1\xe4\x51\xc6\x4c\x60\x8a\x91\xc4\xd2\x9a\x34\xcc\x8a\x69\x67\xd1\x87\x28\xfd\x35\x81" +
	"\x5a\xab\x0e\x77\xbd\xca\x75\xab\xc6\xda\x63\x74\x33\xe8\x4d\x47\xad\x93\xfa\xf0\xec\xa8\xd9\x54\xcf\xa7\xd3\xab" +
	"\xf3\x48\x78\x4e\xbd\xb9\x1b\x73\x71\x6c\x5d\x4d\xc4\xd9\x6e\x6d\xef\xfb\xd3\xe9\x60\xef\x9e\x1f\xa9\x69\xff\xaa" +
	"\x79\x49\x83\xad\xa4\x0a\x4a\x29\x91\x83\x1c\x83\xe9\x22\xe1\x81\x05\x70\x13\x21\xb9\xd8\x07\x31\xd7\x88\xc4\x67" +
	"\xf0\xb2\x26\x90\xfe\x84\x24\x0a\x20\x9f\x60\x41\xd1\x1c\x2c\x0a\x51\x88\x27\x98\x29\xb9\x0f\x18\x67\xf8\x33\x78" +
	"\x86\x84\x79\x78\xb6\x0f\x6a\x1b\x3a\xcc\xdc\x8d\xfb\xa1\x56\x01\x48\x14\x80\xc5\x4a\xda\x7e\x53\x7a\x9f\xa9\x10" +
	"\xba\x21\xa1\xde\x8e\xfd\xb1\x7c\xb4\xfa\x77\x8e\x56\xf5\x51\x25\x10\x93\x3e\x17\xd1\x3e\x10\x5c\x21\x85\x77\x60" +
	"\xd5\xc3\xc1\xc7\x6c\x83\x22\x85\x1f\x77\xec\x4a\x3c\x2b\x2d\x3c\xec\x40\xbb\x1a\xcf\x3e\xfe\x9d\xbb\x6a\xaf\xdf" +
	"\x05\x1a\x9b\x77\xc1\xd6\xc6\x5d\xe0\xb5\xbb\x40\x71\x99\x26\x4e\x5c\xa4\xb8\x48\xbd\x20\x89\x22\x9c\xed\x03\xe4" +
	"\x48\x4e\x13\x85\x3f\x03\xa1\xb3\x66\x1f\xd8\x38\xfa\x0c\x14\x9e\x29\x28\x43\xe4\xf1\xe9\x3e\xa8\x80\x0a\xa8\xc5" +
	"\x33\xf0\xa1\x5a\xd1\xff\x3e\x03\xc5\xe3\x5c\xee\x2d\x7f\x87\x0d\x4d\xe6\x97\xca\xd6\x55\x78\x64\x62\x66\x4e\xf6" +
	"\x88\x8c\x29\x9a\x17\x11\xb2\x2e\x96\x9a\x8c\xab\x9d\x7d\x9f\x08\xa9\x32\xcb\x7d\xdc\xc2\x6b\xe9\x72\xb8\x19\x2e" +
	"\x79\x38\xa5\x26\xcc\x8f\x2d\x0d\x0f\xcc\x9a\x04\x3a\x19\x21\x4f\xd4\x67\x90\x56\xab\x7d\x60\x57\x2a\xff\x57\x52" +
	"\xc2\xd0\x04\xa6\xb9\x4e\x74\x32\x70\xaa\x73\x41\x04\x0e\xda\xa9\x36\x1a\x9f\x8a\xff\x66\x23\xf5\x49\x7e\xe6\xaf" +
	"\x08\x7b\x04\x81\x9d\x08\xcd\x60\xae\x14\xec\x35\xf7\xcc\x76\xea\xcd\x6d\x9f\x05\x30\xa7\x30\x16\xdc\x05\x8b\x02" +
	"\x4a\x23\x45\xb2\xc4\xf2\x9a\xde\x76\xdb\xce\xf5\x22\xe6\x81\x9d\x88\xb0\x62\x6b\xaf\x99\x5d\xf7\x8a\xde\x5a\xed" +
	"\x2d\xbd\xb6\xdd\x6e\x6f\xd1\xdb\x6e\x57\xb7\xe9\xad\x36\x5e\xd7\xbb\x3a\x6c\x57\x2b\x95\x5f\x19\xe1\x75\x3b\x54" +
	"\xcb\x76\xe8\x58\xf9\x7b\x2a\x7b\x52\x44\x51\x7c\x70\x1d\x72\xc5\x65\xc7\xca\x9e\x74\x9b\x64\x15\x7d\x52\xc7\xe1" +
	"\xde\x3c\x2f\xb4\x0c\x4d\x80\x4b\x91\x94\x5d\x83\xa1\x89\x83\x04\xc8\x7e\x41\x3c\x8b\x11\xf3\x20\x0d\x80\x4f\x66" +
	"\xd8\x83\x0e\x57\x8a\x47\xc5\x6e\x5a\x6c\x80\x13\x64\x5f\x8c\xe2\x15\xd9\x41\xeb\xca\xa0\x23\x10\xf3\x8a\x22\xf0" +
	"\xc1\x58\xa2\x42\xcb\x13\x4e\xa2\x14\x67\x1b\xc7\x14\x0f\x02\x8a\x85\x01\x74\xe7\xd7\x35\x32\x19\x03\x78\x48\xa1" +
	"\x7c\xaf\x6b\xb8\x9c\x52\x14\x4b\x5c\x2e\x11\xfa\x93\x09\x21\x11\xe8\x0e\xef\x43\xa6\xf0\x36\x89\x75\x57\x86\xbd" +
	"\x5e\xd6\x95\x19\x00\x09\x82\xa0\xee\xd1\x04\xa7\xcb\x7b\x7f\x12\xdb\x50\x9d\x1e\xca\x0c\x83\xbd\xae\xe1\x23\x2a" +
	"\x71\xae\x8a\x22\x47\x57\xbe\xbb\x14\x9c\xb6\x12\x09\x90\x4e\xae\xa5\x69\x74\x35\x89\xd1\x16\xa2\x90\xb8\x5a\xb4" +
	"\x63\x69\x91\xa5\x69\xac\x8c\xf7\x41\x91\x45\x1d\x8f\x2c\x9d\x55\xb0\x2f\x1c\xb2\xb4\x06\x20\xde\x56\x3a\x25\x2c" +
	"\x09\xdd\x40\xa2\x03\x21\xa2\x10\x25\x8a\x97\xe4\xd2\xb2\x5e\x92\x84\x44\xe1\xc8\x38\x58\xf3\x73\xfa\x32\xd8\x2c" +
	"\xf4\x28\xa4\x58\x48\x33\xc2\x96\x91\xe9\x49\xe3\xb0\x6b\x7c\xc7\x8e\x24\x0a\x1b\x07\x9d\xa5\x5a\x1f\x49\xe0\x23" +
	"\x18\xf2\x48\x2f\x5b\xe4\x00\xf4\x79\x84\x63\x14\x60\x1d\x27\x1d\x8b\x92\xdf\x80\x67\xcc\x12\x65\x12\x56\x74\xbc" +
	"\x46\x81\xe7\x2e\x7f\x5e\xc3\xe3\x68\x3c\x6a\xb9\xa3\x21\xfd\x75\x91\x3c\x13\x1f\x8b\x9f\x21\x75\xac\x84\xae\x5c" +
	"\xe6\x91\x49\x9e\x59\x16\x43\x93\x22\x25\xd7\x1c\xc7\x14\x22\x0c\x0b\xe8\xd3\x84\x78\xab\xd4\x29\xc9\x08\x3e\x05" +
	"\x91\x82\x35\x10\x39\xb0\x51\xf6\xda\x7a\x00\xc0\xc8\x83\x76\x15\x78\xd0\xa7\x78\x06\x46\x89\x54\xc4\x9f\xc3\x7c" +
	"\xf0\x80\x14\xfb\x0a\xe8\x1d\x38\x15\x28\xce\xe2\xa2\xa8\x8a\xe9\xca\x41\x09\xed\x06\xf4\xf4\xeb\x6a\x00\x19\xdd" +
	"\x24\x58\xcc\x41\x5a\x76\x3e\x01\x15\x62\x06\xae\x79\x1c\x63\x61\x8e\x64\xfe\xbc\x9a\x51\xce\x4a\x23\x8a\x74\x05" +
	"\x89\x15\x90\xc2\x7d\x77\x77\x3b\x7a\xd2\x77\x59\x35\xb3\x66\xda\xf9\x43\xda\xcd\x8e\xe4\x7a\x36\xfe\xdc\xce\x9e" +
	"\x04\x71\xcf\xb1\xce\xcf\x6e\xe8\xc5\xa5\x7f\x95\xb4\x6d\x85\x6a\x55\x6e\x5d\x0e\x7e\xcc\xa8\x9a\x0e\x79\xeb\x46" +
	"\x45\xe3\xc1\xd0\x3b\x4c\x5a\xdb\xdb\x59\x9d\x83\x29\xe8\x7f\xce\x20\x2e\x4c\x63\xd9\xa6\x5d\x37\xed\x62\xe1\x7d" +
	"\x2c\x6e\xef\xaa\x03\xec\x8a\x07\x71\x86\x0e\xe5\x53\xe4\xc7\xe3\xf6\xc3\xf0\xe6\xf4\x8e\x1e\xf3\xf9\x20\xba\x57" +
	"\xc1\xf9\xd1\x17\xe6\x1d\x13\x79\xeb\xfe\xa7\x2c\xb6\x8d\x86\xa3\xcd\xc9\xf0\x6d\x4a\x8d\xdd\x4a\xf5\x39\x69\xdc" +
	"\xdf\x5f\xdc\x8c\xaf\xf6\xa6\xf6\xd7\x53\xd1\x9c\xf8\xaa\x17\x0c\xfc\x13\xef\x87\x7b\xd8\xc7\xc7\xea\x84\x9d\xff" +
	"\x90\x47\xf2\x5d\x94\x56\x81\x79\xa1\xab\xcf\xd7\x2c\xa8\xff\x55\xd4\xbd\x32\x53\x8d\x5e\x19\xa9\xde\xe6\xca\xbe" +
	"\xb8\x23\xef\xa6\xe5\x9c\xd4\xb9\x35\x3b\x3f\xe9\xdd\x7e\x3f\x6a\x1f\x8e\x7f\xec\xf6\xaf\xbf\x47\x3d\x67\xdc\xbf" +
	"\x7c\x4a\xc2\x2f\x8d\xc7\xcb\x0b\xef\x7d\x5c\xff\x31\xa1\x00\x86\x48\x86\x96\x6d\x56\xcc\x7a\xf1\xf4\x4e\x0e\xa2" +
	"\xd7\x8c\x2f\xeb\x43\x9b\x47\x23\xa1\x06\xd3\x4a\x93\x9d\x0c\x7a\x87\xc3\xe7\xc7\xa7\x59\x52\x9f\xa3\x6f\x7c\x78" +
	"\xd7\xff\x81\x76\xa7\x77\xcf\xd3\xff\x9a\x83\x0a\x93\xc8\x61\x88\x50\x9d\x4b\x66\x65\x6d\xe9\x9d\x6c\x7a\x63\xeb" +
	"\xfa\xfb\xcd\x3d\xef\x4b\xb7\xed\x55\x1f\xea\xf3\xfb\xc3\x8b\x89\x90\x7b\xf6\x0d\x6f\xb8\x72\x37\x38\xed\x1f\x7f" +
	"\x3f\xa2\xd5\xaf\x81\xf3\xbe\xd7\xc2\x3f\xa6\x13\x22\xe6\x51\xec\x20\x21\xf5\xbb\x41\x67\x93\x6d\x97\x17\xdf\x47" +
	"\x67\xf7\x6c\xd0\x97\x0f\xc3\x39\xbe\x93\x53\x69\xa9\x67\xe7\x34\x6c\x3c\xf6\x6f\x66\x43\xef\xbc\x97\x5c\x8e\xa2" +
	"\xc9\xa5\x7b\x77\xa2\x1e\x9b\xc7\x17\xef\x7b\x3f\xfc\x4c\x07\xc5\xb1\xc6\xb1\x4d\x46\xd7\x11\x44\x9d\x24\x82\x0a" +
	"\x47\xb1\x1e\xc2\x8a\x36\x4d\x8f\x3c\xd6\x0c\xae\x28\xad\x24\x5e\xab\x72\xe9\x30\x52\xb4\x90\xd9\xbc\x34\x0d\x89" +
	"\xc2\x20\x6f\x76\x7d\xca\x51\x5e\xc6\x16\x8b\x0f\xc4\xff\xaa\x00\x89\x50\x80\xa5\x49\x31\x0b\x54\x08\xec\x97\x97" +
	"\x3c\x39\x17\x0b\x4b\xef\xbf\xbc\x94\x8a\xe5\x62\xf1\x01\x23\x37\xcc\xcf\xbc\xbc\x94\x37\x52\x65\x7f\xa5\x13\x12" +
	"\xa8\x96\xb6\x36\xe1\xe9\xb1\x0e\x38\x5c\x78\xd9\xc0\x5a\xf6\x4d\x6a\xaa\xc5\x62\x19\x8d\x2f\x2f\xe5\x5d\x44\x55" +
	"\xba\xab\x9b\x8c\xf5\x9d\xb4\x3d\xcd\x0f\xfb\x09\xa5\x92\x3c\x6b\x89\xf5\x62\xac\x41\x62\x2a\x71\x19\x9a\x86\xf2" +
	"\xbf\x81\x56\x02\x95\x9b\xb9\xbc\xa2\x8d\xfc\x86\x09\x57\x93\x31\xb4\x2b\x95\xa2\x61\x49\x7b\x13\x97\xd3\x24\x62" +
	"\x3f\x35\x2f\x98\x79\xeb\x5d\x68\xd8\x38\x58\x62\xee\x58\x61\xa3\xbc\xb9\x35\x3a\xca\x0a\xe2\x35\x50\x3a\xda\xc0" +
	"\x4f\x7f\x23\x78\xa5\x2b\xcd\x74\xe6\x4d\xe0\x62\xb1\x76\x85\x46\x12\xaf\x03\xd9\xb4\xcf\xf6\x16\x6b\x6b\x6e\x61" +
	"\x21\xb8\xf8\xb7\xb9\xa5\x3b\xc8\x4d\x9b\x46\x1e\x74\x31\x4b\x5b\xda\xad\x1d\xe5\x4c\x42\xbb\xa2\x87\x7a\x28\x23" +
	"\xd8\x02\x79\x93\xd9\x4c\xbf\xd0\x00\xd6\xd7\x9d\x52\x3a\x8c\x28\x16\x0a\xa4\x3f\xa1\x87\x58\xa0\x6f\x59\x2c\x22" +
	"\x2c\x25\x0a\x52\x97\xad\x45\xf7\xbb\xcc\xd2\xb1\xb2\x39\xb5\x63\x65\x7f\xfa\xff\xff\x00\x00\x00\xff\xff\x17\x1e" +
	"\x36\x86\x0b\x18\x00\x00")

func bindataFrontendIndexhtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataFrontendIndexhtml,
		"frontend/index.html",
	)
}



func bindataFrontendIndexhtml() (*asset, error) {
	bytes, err := bindataFrontendIndexhtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "frontend/index.html",
		size: 6155,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"frontend/app.js":     bindataFrontendAppjs,
	"frontend/index.html": bindataFrontendIndexhtml,
}

//
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}


type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"frontend": {Func: nil, Children: map[string]*bintree{
		"app.js": {Func: bindataFrontendAppjs, Children: map[string]*bintree{}},
		"index.html": {Func: bindataFrontendIndexhtml, Children: map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
