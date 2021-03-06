// Code generated by go-bindata. DO NOT EDIT.
// sources:
// graphiql.html
// schema.gql

package static

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

var _bindataGraphiqlhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4f\x6f\x13\x3f\x10\x3d\x6f\x3e\x85\x7f\x96\x7e\xd2\x46\x2a\x76" +
		"\x52\x24\x0e\x9b\x4d\x0e\xd0\xaa\x02\x15\x4a\x81\x0b\x47\xd7\x9e\x5d\x3b\x78\xed\xed\xd8\x9b\x36\xaa\xf2\xdd\x91" +
		"\xf7\x4f\x28\x7f\x2a\x21\x04\x17\xaf\x3d\x7e\xf3\xde\xf3\xcc\x68\xcb\xff\xce\xae\x5e\x7d\xfa\xfc\xfe\x9c\xe8\xd8" +
		"\xd8\xcd\xac\x1c\x3e\x59\xa9\x41\xa8\xcd\x2c\xcb\x4a\x6b\xdc\x17\x82\x60\xd7\x34\xc4\xbd\x85\xa0\x01\x22\x25\x1a" +
		"\xa1\x5a\x53\x1d\x63\x1b\x0a\xce\xa5\x72\xdb\xc0\xa4\xf5\x9d\xaa\xac\x40\x60\xd2\x37\x5c\x6c\xc5\x3d\xb7\xe6\x26" +
		"\xf0\x1a\x45\xab\xcd\xad\xe5\x0b\xb6\x5c\xb2\xe5\xf2\x18\x60\x32\x04\xca\x7b\x99\x20\xd1\xb4\x91\x04\x94\xbf\x4d" +
		"\x5b\x41\x94\x9a\x9f\xb2\x05\x7b\x3e\xec\x59\x63\x1c\xdb\x06\xba\x29\xf9\x40\xf7\xa7\xcc\x08\x42\x46\xbe\x7c\xc1" +
		"\x4e\xd9\x82\x77\x8d\x1a\x02\xac\x45\xaf\x3a\x19\x8d\x77\x7f\x57\xe9\x99\xf2\xcd\x4f\x6a\x29\xf8\x2f\x14\x9f\x6e" +
		"\xc6\x2f\x14\x4a\x3e\xce\x41\x79\xe3\xd5\x9e\xf4\x13\xb0\xa6\x77\x46\x45\x5d\x90\xe5\x62\xf1\xff\x8a\x68\x30\xb5" +
		"\x8e\xd3\xa9\x11\x58\x1b\x57\x90\xc5\x8a\xf8\x1d\x60\x65\xfd\x5d\x41\xb4\x51\x0a\xdc\x8a\xf6\x96\x95\xd9\x11\xa3" +
		"\xd6\x74\x92\xa5\x13\xeb\x23\xa2\x9d\x5e\xd1\xcd\xa5\x17\xca\xb8\x9a\x31\x56\x72\x65\x76\x8f\xde\x9b\xb6\x59\xd5" +
		"\xb9\xbe\x30\xa4\x6f\xfd\xc5\xf5\x65\xde\x0a\x14\x4d\x98\x93\x87\x74\x9d\x21\xc4\x0e\xc7\xdb\x9c\x0e\xaf\xbc\xb5" +
		"\xf4\x64\xbc\xce\x1a\x88\xda\xab\x82\xd0\xd6\x87\x48\x4f\x86\x60\x7a\x65\x41\xde\x7c\xbc\x7a\xc7\x42\x44\xe3\x6a" +
		"\x53\xed\x27\xde\x11\x22\x11\x14\xb8\x68\x84\x0d\x05\xa1\xc6\x49\xdb\x29\x18\xf3\x0f\x73\x16\x35\xb8\xfc\xe8\x2d" +
		"\x47\x08\xed\xe4\x68\xb2\x94\x62\x2c\xc2\x7d\xcc\xe7\xab\x27\xd2\x92\x8f\x63\x5a\xc4\xfd\xb4\x9d\x28\x7a\x87\xad" +
		"\xc0\x00\x03\x74\xe0\xc9\x0e\x44\x8a\x28\x35\xc9\x01\xd1\xe3\xfc\xc7\xac\x04\x9d\x90\xa3\x70\x7f\x3c\xcc\xd2\xfa" +
		"\x21\x4d\xdd\xd9\xd5\x5b\x86\xe0\x14\x60\xde\x23\xfa\x20\x93\x08\x22\xc2\xb9\x85\x06\x5c\xcc\x2f\xfa\xce\x5d\x5f" +
		"\x9e\x90\x87\xbe\xba\x80\xc5\xb1\x09\x87\xb1\x4c\xca\xcb\x2e\x81\x59\x0d\x71\xcc\x7b\xb9\x7f\xad\xf2\x6f\x6d\x9f" +
		"\x27\x5c\x5a\xbe\x1b\xb7\x64\x71\x33\x2b\xf9\xf0\x1b\xfa\x1a\x00\x00\xff\xff\xdb\x8e\x2c\x18\x9e\x04\x00\x00")

func bindataGraphiqlhtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataGraphiqlhtml,
		"graphiql.html",
	)
}

func bindataGraphiqlhtml() (*asset, error) {
	bytes, err := bindataGraphiqlhtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "graphiql.html",
		size:        1182,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1556057220, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataSchemagql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x54\x4f\x6f\xe3\xb6\x13\x3d\x4b\x9f\x62\x82\xbd\x24\x17\x1f\x7e\xf8" +
		"\x9d\x84\xb6\x80\x93\xb4\x68\xd0\x78\xbb\x5d\x67\x8b\x02\x41\x51\x8c\xc5\xb1\x3c\x30\x45\x6a\x87\xa4\x13\x63\x91" +
		"\xef\x5e\x90\x92\x1d\x4a\x4a\xd3\x7b\x4f\xd2\xfc\x79\xc3\x99\xc7\xc7\x71\xf5\x8e\x5a\x84\x6f\x65\xf1\x35\x90\x1c" +
		"\x2b\x28\x7e\x8b\xdf\xf2\xa5\x2c\xfd\xb1\x23\x48\x56\x0c\x7f\x00\x21\x2f\x4c\x07\x02\xd4\x1a\x0e\xa8\x59\xa1\x27" +
		"\x05\xe8\x1c\x79\x07\xd6\x80\xdf\x11\xac\x3d\x69\x8d\x02\x86\xfc\x93\x95\xfd\xa2\x2c\xfa\x78\x05\x8f\xcb\xf8\x73" +
		"\xf1\xe7\x45\xf9\x4e\x31\x76\x2e\x90\xbc\x53\x6d\x48\xa8\xe0\xf1\x2e\xfd\xcd\xea\x79\x41\x45\xe0\x3c\x7a\x07\x5b" +
		"\xb1\x6d\xaa\xa3\xd1\x79\xf8\xce\x84\xf6\x67\x1b\xc4\x2d\x1b\xfb\x03\xec\xe2\x5f\x44\x5e\x2a\xda\x62\xd0\x1e\xbe" +
		"\x87\xff\xfd\xbf\x77\x5f\x2d\xc0\x76\x9e\xad\x41\xad\x8f\xd0\x89\x3d\xb0\x22\xa8\x6d\x30\x9e\x04\xd0\xa8\x88\xdb" +
		"\xa0\xa3\x7e\x78\x60\xb3\xb5\xb0\xb5\x02\x5b\xd6\x9e\x84\x4d\xb3\x28\x8b\x16\x65\x4f\xde\x5d\x96\x45\x11\x53\xd3" +
		"\xf4\x37\x56\x51\x05\x6b\x1f\x53\x72\x7f\x3f\x4b\x16\x19\xce\x7a\x0b\x94\x87\x66\xb8\x6c\xc4\x0a\xee\x8c\x2f\x8b" +
		"\xab\x0a\x1e\x57\xa9\x95\x19\xf3\x4d\x23\xd4\x24\xda\x47\xa4\x59\xf9\x07\xce\x22\x3a\xf1\xf3\x26\x3d\x11\x63\xb0" +
		"\x25\xb0\xdb\xf4\xdf\xd7\xec\x90\x05\x2e\x69\x11\x19\xf9\x00\x7f\xdc\xaf\xfe\xba\x7e\xb8\xb9\x1a\x93\x05\x42\x2e" +
		"\x68\xef\x16\x65\xe1\xb9\xde\x93\x44\xce\x22\xf0\x23\xb6\xf4\xaf\xc3\x2d\xcf\x63\x9c\xc7\x7c\x29\x4b\x57\x63\x14" +
		"\xce\x35\x37\x31\x71\xb0\x1e\xb8\xa5\x41\xd7\x89\xbe\xa8\xeb\x3a\x63\xf7\xe2\xa4\xaf\x65\x9d\x58\xce\xfc\x11\x94" +
		"\x99\x26\xb4\x43\x8e\x4b\xad\x5c\x94\x05\x06\xbf\xfb\x4c\x5f\x03\x0b\xa9\x0a\xae\xad\xd5\x84\xe6\xec\x3f\xd8\x1a" +
		"\x37\x9a\x46\x81\xb6\x3f\xe3\x27\x6d\x31\x15\xe8\x2f\xdb\x78\xb1\x5a\x93\xba\x3e\xde\xda\x16\xd9\x8c\x20\xa6\xde" +
		"\xd9\xb9\x2a\xc6\x91\x87\x71\xab\xec\x92\x77\x99\x12\xc6\xad\x29\x76\x9d\xc6\xe3\x2d\xd5\xdc\xa2\x76\xd5\x40\x57" +
		"\x9c\x2f\x63\x3e\x26\x92\xab\x33\xb3\xb6\x46\x71\x14\x80\xcb\x9c\x5b\x7e\x26\xf5\x31\xb4\x9b\x28\xc8\x73\xa1\x16" +
		"\x9f\x67\x3e\x76\x5f\x8c\xe6\x96\xfd\xb8\x1b\x21\x45\x6d\xd2\xd5\x9d\x71\x5e\x42\x3d\x3d\xa1\xb6\x5a\xa3\x27\x41" +
		"\xbd\x54\x4a\xc8\x39\x7a\x37\xba\xe6\xc6\xa0\x0f\x32\xc9\x0a\x26\xea\x3f\xf7\x45\xdd\x07\x37\x13\xc1\xdd\xed\x70" +
		"\xb5\xa7\x5d\xd8\xeb\x2b\x8a\x26\x69\xfb\x13\xb2\x64\xa0\x37\x1f\x79\xee\x1f\x3f\xd6\x53\x2f\x6f\x3c\xf2\x49\x68" +
		"\x86\x8b\x15\x7f\xb7\x3a\xc4\x2b\x3a\x89\x67\x00\x4c\xdd\xa9\xd1\x9b\x5e\x67\x3d\xf9\xb6\x23\xf3\x1a\xd7\xf6\xe9" +
		"\xd5\xd8\x71\xb3\xcb\x2a\xee\xd0\x34\xf9\x09\xda\xba\xcc\xe4\x78\xdc\x01\xf5\xda\xa3\xf8\x2a\x3d\xad\x24\x02\x71" +
		"\xfe\x9e\x54\x43\x72\x13\xf3\xa3\xfb\x1c\xb4\xa2\x48\x36\xd6\xee\xd7\x71\xd1\x54\xf0\xeb\xc8\x7e\xe5\x79\xfa\xa2" +
		"\xdf\x63\xfc\xbf\xcc\xc3\xd8\x0f\xdf\x4a\x28\x36\xac\x86\x29\xce\xaf\x69\xc3\x6a\x3a\xed\x86\xd5\x0a\x9f\xf3\xcd" +
		"\xb2\x9f\xa2\xd0\xed\xa7\x28\x74\xfb\x15\x67\x9c\xb8\x4e\x08\xd5\xd4\x5e\xb1\xfa\x64\x39\xdb\x5b\xa7\x6e\x7b\x99" +
		"\xc6\xbb\xea\xc2\x46\x73\xfd\x0b\x1d\xf3\x85\x39\x5e\x28\x41\x74\xbe\x5c\x6d\xab\xbf\x7c\xbe\xcf\x97\x09\x29\x12" +
		"\x8c\x0b\x60\x4d\x72\x18\xa9\x3f\xee\xd3\x99\xd3\x0b\x1a\xb7\x25\x99\x05\x9e\x68\xb3\x0c\x7e\xf7\xa3\x51\x5d\xdf" +
		"\x75\xb6\xd3\x3a\xeb\xd8\xcf\x10\x56\x9a\x87\x27\xf6\x3e\x77\xbe\x94\x7f\x07\x00\x00\xff\xff\x3f\x67\xd7\xeb\x28" +
		"\x09\x00\x00")

func bindataSchemagqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataSchemagql,
		"schema.gql",
	)
}

func bindataSchemagql() (*asset, error) {
	bytes, err := bindataSchemagqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "schema.gql",
		size:        2344,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1556718973, 0),
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
	"graphiql.html": bindataGraphiqlhtml,
	"schema.gql":    bindataSchemagql,
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
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
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
	"graphiql.html": {Func: bindataGraphiqlhtml, Children: map[string]*bintree{}},
	"schema.gql":    {Func: bindataSchemagql, Children: map[string]*bintree{}},
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
