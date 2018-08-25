// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// template/application.css
// template/index.html
package web

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
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
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
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateApplicationCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xef\x72\xbb\x38\x0c\xfc\x9e\xa7\xd0\x4c\xe6\xa6\xed\x0d\xee\x38\x7f\x9b\x90\x67\xb8\x87\x30\x58\x80\xae\xc6\x66\x6c\xd1\x24\xed\xf4\xdd\x6f\xc0\x40\x92\xa6\xbd\x5f\xbf\x2e\x8b\x2c\xed\xae\xf4\xf7\xc7\x0c\x00\x20\x73\x27\x11\xe8\x9d\x6c\x99\x42\xe6\xbc\x46\x2f\x32\x77\x3a\xcc\x3e\x67\xb3\x67\xc6\x13\x8b\x1c\x2d\xa3\x87\xc8\xee\x11\x65\xa8\xb4\x29\xc4\x0f\x3d\xb3\xe2\xda\x24\x99\xd3\xe7\x81\x56\x93\x15\x47\xd2\x5c\xa5\xb0\x90\xf2\xed\x78\x98\xd0\x0a\xa9\xac\x38\xc2\x55\x84\x0b\x67\x59\x14\xaa\x26\x73\x4e\xe1\xe1\x1f\x67\x39\xa0\xf7\x8a\x1f\x12\x08\xca\x06\x11\xd0\x53\x11\xa9\x99\xca\x5f\x4b\xef\x5a\xab\x05\xd5\xaa\xc4\x14\x0c\x59\x54\x5e\x94\x5e\x69\x42\xcb\x8f\xec\x20\x73\xcc\xae\x06\x83\x05\x27\x30\xdf\xad\xb6\x32\x5f\x25\x30\x5f\x62\x56\xec\x17\x4f\xb1\x90\x38\x62\xf6\x4a\x2c\xfa\xb7\x43\xed\x1c\x57\xbd\x02\xca\x32\x29\x43\x2a\xa0\x8e\x44\x4d\xa1\x31\xea\x9c\x42\x61\xf0\x14\xa1\x7e\x7c\x41\x8c\x75\xb8\x88\xd0\x7d\xf8\xb7\x0d\x4c\xc5\x59\xe4\xce\x32\x5a\xbe\x55\x68\xf1\x71\x19\xf6\x38\x88\xb0\x96\x72\x50\x46\xf9\x92\xac\x88\x9d\xa7\xb0\x6e\xa2\x01\x9d\xda\xca\xa3\x1a\x9d\xea\xdd\xe9\x26\x6d\x43\x0a\xdb\x66\xe8\x27\xe2\x29\x58\x67\x31\x22\x8d\xd2\xba\x1f\x67\x29\x47\x52\x1c\x94\xde\xf1\x1a\xbc\x31\xe4\x65\x82\xd9\x2b\x1b\x88\xc9\xd9\x8e\x2c\xeb\x00\xa8\x02\x0a\xb2\xc2\xb5\x1c\x29\x17\x73\xff\x3a\xdc\x05\x63\xd0\xe1\xeb\x07\x61\x54\xe0\xb4\xf7\xa5\x9f\x2e\x6b\x99\x9d\x1d\x12\x93\xb7\x3e\x38\x9f\x42\xe3\xe8\x22\xe8\x20\x0b\xbb\x26\x85\xd5\xf2\xba\xeb\xe1\xfd\xe5\x62\x6a\x3a\x77\xa6\xfb\xff\x58\x11\xe3\xe1\x92\xec\x4a\x69\x77\x4c\x41\xc2\xb2\x39\x75\x92\x81\x04\x5f\x66\xea\x71\x93\x2c\x36\xc9\x7a\x9d\xc8\xe7\xcd\xd3\x17\xd1\xe4\x95\x42\x93\xfd\x64\xbb\xa4\x89\xcc\xb8\xfc\xf5\xf0\x9d\x1d\x9b\xef\xa4\x5e\x6c\x6f\xc0\xd1\xf7\xcd\xe8\x7b\x35\x05\x61\xe4\xf5\xaf\xdc\xc3\xf7\x1e\xff\xb4\xb7\xfd\x28\x6e\xb4\xcf\xa3\x51\x4c\x6f\xf8\x2b\x63\xbb\x95\xef\xb2\xab\xc8\x4e\x0b\xff\x53\xa9\x68\x4d\x27\x96\x6a\xd9\x8d\xd8\x69\x34\x66\x23\xa7\xd6\x6f\xef\x40\xf7\x06\xda\xdc\x9f\x1b\xfe\xf8\xe5\x4a\xef\xa4\xc6\x32\x81\xb9\x94\x0a\x8b\x22\x81\xf9\x6a\xb5\xdd\xe3\xcb\xd3\xd0\xb0\xa1\x26\x73\xca\xeb\xdf\x96\x5b\xee\xd6\xb1\xde\x22\xcf\x5f\xb4\xec\x2e\x83\xd6\xb8\xdf\x0d\xf5\x42\x9b\xe7\x18\x82\x18\x9a\x44\xfd\x07\x21\x6e\x36\xe1\xb3\xfb\x3f\x63\x62\x83\xe3\x25\xbc\x8a\xf0\x7a\x4a\xf0\xed\xba\x5f\xa2\xed\x1a\x95\x13\x9f\x53\x78\xde\x7f\x13\x9b\xfe\x5c\xf4\x3d\xa2\xd5\xff\x73\x90\xef\xd7\xe1\x73\xf6\x5f\x00\x00\x00\xff\xff\x6b\x96\x84\xea\xee\x05\x00\x00")

func templateApplicationCssBytes() ([]byte, error) {
	return bindataRead(
		_templateApplicationCss,
		"template/application.css",
	)
}

func templateApplicationCss() (*asset, error) {
	bytes, err := templateApplicationCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/application.css", size: 1518, mode: os.FileMode(436), modTime: time.Unix(1535158570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x5b\x6f\xdb\xb8\x12\x7e\xf7\xaf\x98\xb2\x0f\x3d\x07\x38\xa2\x2e\x8e\x2f\xe9\x91\xb4\x40\xb3\x1b\x2c\x16\x2d\x5a\x6c\xda\x87\x7d\x5a\x50\xe4\xd8\x62\x4a\x91\x2a\x49\xd9\x71\x8d\xfc\xf7\x05\x25\x2b\x51\x52\x17\xdb\x87\xea\x81\x1a\xce\xe5\xe3\xf0\xe3\x0c\x99\xbf\xf8\xf5\xfd\xd5\xc7\xbf\x3e\xfc\x06\xb5\x6f\x54\x39\xcb\x87\xdf\x2c\xaf\x91\x89\x72\x06\x00\x90\xbf\x88\xa2\xb7\xe8\xa1\xb2\x66\xef\xd0\xc2\x67\x6d\xf6\xb0\xc7\xca\x49\x8f\x20\x1d\x98\xd6\xcb\x46\x7e\x45\x01\x1b\x63\xa1\x31\x95\x54\x18\x45\xa7\xd8\x06\x3d\x03\xcd\x1a\x2c\xc8\x4e\xe2\xbe\x35\xd6\x13\xe0\x46\x7b\xd4\xbe\x20\x7b\x29\x7c\x5d\x08\xdc\x49\x8e\x51\x3f\xf9\x1f\x48\x2d\xbd\x64\x2a\x72\x9c\x29\x2c\x52\x9a\x10\x88\x4f\x60\x4a\xea\xcf\x50\x5b\xdc\x14\xa4\xf6\xbe\x75\xaf\xe3\x78\x63\xb4\x77\x74\x6b\xcc\x56\x21\x6b\xa5\xa3\xdc\x34\x31\x77\xee\x97\x0d\x6b\xa4\x3a\x14\xef\x82\x1d\xad\x65\x9e\x80\x45\x55\x10\xe7\x0f\x0a\x5d\x8d\xe8\xc9\x14\xf5\xb9\xed\xb4\xcc\xf1\x48\xdf\x30\x87\x9f\xfe\x7c\x7b\x7f\x1f\xb7\x5d\xa5\x24\x8f\x59\xdb\x2a\xc9\x99\x97\x46\x53\xee\x1c\x29\x67\x79\x3c\xb0\x35\xcb\x2b\x23\x0e\xe5\x2c\x6f\x98\xd4\xc0\x15\x73\xae\x20\x0e\xb5\x18\x97\x12\x72\x37\xaa\x03\x07\x4c\x6a\xb4\x27\x5b\x6f\xaf\xd3\xf2\x8f\xce\x79\x78\xaf\x39\xe6\x71\x9d\x4e\x4c\xed\x03\x5e\x57\x79\xe9\x15\x92\xf2\x06\xb5\x00\x06\x0d\x3a\xc7\xb6\x08\xbe\x66\x1e\x1c\xaa\x0d\x08\x74\xde\x76\xdc\x3b\x60\x1b\x8f\x16\x8c\x46\x08\xec\xe7\x71\x3b\x41\xfc\x4e\x32\x0f\x0e\xe1\x3b\x1e\xe5\x06\xe8\x15\xd3\x57\x16\x99\xc7\xfb\xfb\x27\xd6\x7c\x63\x6c\x03\x52\x84\x4d\x72\x8b\x3e\x4c\x09\x30\x1e\xa8\x79\xc6\x1d\xef\xe3\x09\x34\xe8\x6b\x23\x0a\xf2\xe1\xfd\xcd\x47\x02\xa8\xb9\x3f\xb4\x58\x90\xa6\x53\x5e\xb6\xcc\xfa\x38\x60\x44\x82\x79\xf6\x3c\x95\xe7\x39\x4b\xdd\x76\x3e\xda\x48\x54\x62\x42\xe1\x13\xef\x17\x51\xf4\xa9\x55\x86\x09\xb8\xe9\xf3\x83\x6b\xa9\xf0\x35\xe4\x7d\x68\xc8\xfb\xd5\x26\xd4\x6a\x3f\x7d\x05\x43\x26\x41\x43\x4e\x15\x3b\xc8\x4e\x7e\xc5\x22\x5b\x94\x79\x5c\xd9\x32\x14\xf6\xf9\xc5\x3c\xde\x79\x66\x91\xf5\x7c\x8c\x93\x74\x84\x6a\xdc\x96\x40\xab\x18\xc7\xda\x28\x81\xb6\x20\x1f\x98\xf3\x08\x07\xd3\xd9\x87\x13\xac\xd1\x62\x28\xd3\x2f\x9d\xb4\x28\x0a\x32\x4a\xa4\xcc\xe3\x11\xf1\xdb\xad\xe6\xb1\x90\xbb\x33\xea\x09\x57\x55\xe7\xbd\xd1\x7f\xef\x2d\x6b\xdb\x27\x15\xf7\x24\x60\xf0\x1a\x63\x50\x73\x7b\x68\x3d\x39\xf1\xe2\xba\xaa\x91\x7e\xdc\xce\x70\xc6\xa4\xbc\xe9\xb5\x67\xe1\x7a\x48\xb7\xdb\xc2\x5d\xa3\xb4\x1b\x3a\xf6\x75\x1c\xef\xf7\x7b\xba\x9f\x53\x63\xb7\x71\x96\x24\x49\xec\x76\x5b\x02\xc3\x3d\x40\xd2\x8c\x40\x8d\x72\x5b\xfb\x41\x0e\x35\xfb\xc6\xdc\x15\x24\x81\x04\xb2\xe5\x9a\xae\xe7\xd9\xf8\x27\x65\xde\x32\x5f\xc3\x46\x2a\x55\x90\x97\xd7\xd7\xd7\x04\x44\x41\xde\x65\xcb\x05\x4d\x57\x90\x66\x0b\xba\x58\xad\x54\xb4\x4e\xa2\x75\xc2\xa3\x0b\xba\x5e\x0f\x43\x9a\xd1\xd5\xe5\xf2\x24\xaf\xe8\x72\xb5\x82\xa4\x9f\x65\xd0\x8f\x27\x79\x70\x83\x04\x7a\x9f\xb5\x5a\xac\xe9\x72\x09\xfd\xf8\x7b\x9a\xd1\x05\x8f\x96\xf4\x32\x99\x43\x12\x00\x17\xb0\xa0\x8b\xcb\xf5\x20\xf6\x43\x02\x83\x39\xe8\x57\xf0\xa8\x0f\x43\x9d\xa5\x73\xba\x5c\x5c\xa8\x68\x02\x9a\x0d\x39\x4e\x72\x78\x9e\x02\x64\xf4\xe2\x62\x18\x16\x74\x79\x01\xf3\x10\xbb\xa6\xeb\x41\x72\x4b\x3a\x0f\x19\xd0\x2c\xeb\x95\x51\x50\xaa\xd5\x25\xbd\xbc\x5c\x05\x06\x02\xe2\x7c\xb2\xcf\x79\x34\xc2\x0f\x2c\xac\xbf\x92\xb8\xcc\xc3\x71\x7c\xa7\x3a\xe2\xa1\x3c\x7e\xa8\xfe\xf2\xbe\x8d\x9f\xe9\x26\xf5\x28\xe4\x4e\x8a\x50\x88\x67\x62\x8f\x47\x54\xee\x9b\x9b\x66\x12\xec\x3a\xce\xd1\xb9\xe8\x54\xa0\x78\xae\xfd\xa7\xfe\xa6\xf3\xff\x7a\x53\x3c\x69\xde\xce\x2a\x32\x06\x8f\x86\xd0\x96\x4c\x18\xad\x0e\x85\xb7\x1d\x96\xc7\x23\x7d\xe7\xb6\xf7\xf7\x3f\xa5\x31\x7f\xac\x21\x2b\x1f\x44\xd9\x56\x86\x59\x41\x20\xdc\x91\xd1\xc3\x3c\xf2\xcc\x6e\xd1\x17\xe4\x65\xc8\xfe\x3c\x5e\xf8\xae\x4c\x7b\x00\x6f\xe0\x6a\x0c\xfc\x19\x87\x7d\xee\x0c\xb5\x98\x1c\xe1\xc4\xe5\x24\xce\xf2\x38\x3c\x92\xe1\xed\x3c\xbd\x99\x8e\x5b\xd9\xfa\xd3\x6d\x13\x48\x8d\x6f\xd9\x8e\x0d\x5a\x02\xce\xf2\xc7\x47\x9f\x1b\x81\xf4\xf6\x4b\x87\xf6\xd0\xbf\xf7\x83\x18\x65\x34\xa5\x29\x6d\xa4\xa6\xb7\x2e\x94\xd6\x10\xfb\x08\xfd\x14\x43\xe8\x5b\x47\xb9\x32\x9d\xd8\x28\x66\xb1\x07\x62\xb7\xec\x2e\x56\xb2\x72\xf1\x03\xb1\xf4\xd6\xc5\x29\x5d\xd1\x74\xa2\xfa\xfe\x12\xe7\xb3\x1f\x76\xae\x71\xff\x48\xfb\x7f\x08\xad\xbc\x26\xff\xfd\xff\xec\x11\x64\xf6\x4f\x00\x00\x00\xff\xff\x17\x5e\x93\xef\x8d\x09\x00\x00")

func templateIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateIndexHtml,
		"template/index.html",
	)
}

func templateIndexHtml() (*asset, error) {
	bytes, err := templateIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/index.html", size: 2445, mode: os.FileMode(436), modTime: time.Unix(1535158258, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/application.css": templateApplicationCss,
	"template/index.html": templateIndexHtml,
}

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
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
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
var _bintree = &bintree{nil, map[string]*bintree{
	"template": &bintree{nil, map[string]*bintree{
		"application.css": &bintree{templateApplicationCss, map[string]*bintree{}},
		"index.html": &bintree{templateIndexHtml, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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

