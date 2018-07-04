// Code generated by go-bindata. DO NOT EDIT.
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

var _templateApplicationCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xdd\x6e\xb3\x3a\x10\xbc\xcf\x53\x58\xaa\x8e\xda\x1e\xe1\xc8\x90\x9f\x36\xe4\x19\xce\x43\x18\xbc\xc0\x9e\x1a\x1b\xd9\x4b\x7e\x5a\xf5\xdd\x3f\x61\x03\x49\x9a\xf6\x53\x2f\x33\x6c\xd6\x3b\x33\x3b\xfb\xef\xc7\x82\x31\xc6\x0a\x7b\xe2\x1e\xdf\xd1\xd4\x39\x2b\xac\x53\xe0\x78\x61\x4f\xfb\xc5\xe7\x62\xb1\x24\x38\x11\x2f\xc1\x10\x38\x16\xab\x03\x22\x35\xd6\x26\x67\xf1\x43\xa8\x6c\xa8\xd5\x49\x61\xd5\x79\x2c\x6b\xd1\xf0\x23\x2a\x6a\x72\x96\x0a\x71\x38\xee\x67\xb4\x01\xac\x1b\x8a\x70\x13\xe1\xca\x1a\xe2\x95\x6c\x51\x9f\x73\xf6\xf8\x9f\x35\xe4\xc1\x39\x49\x8f\x09\xf3\xd2\x78\xee\xc1\x61\x15\x4b\x0b\x59\xbe\xd5\xce\xf6\x46\x71\x6c\x65\x0d\x39\xd3\x68\x40\x3a\x5e\x3b\xa9\x10\x0c\x3d\x91\x65\x85\x25\xb2\x2d\xd3\x50\x51\xc2\x1e\x56\xb0\xda\x6e\x65\xc2\x1e\xc4\x2e\x4d\xb3\xea\x39\x36\xe2\x47\x28\xde\x90\x78\x78\xdb\xb7\xd6\x52\x13\x14\x90\x86\x50\x6a\x94\x1e\x54\x2c\x54\xe8\x3b\x2d\xcf\x39\xab\x34\x9c\x22\x14\xe8\x73\x24\x68\xfd\x45\x84\xe1\xc3\xff\xbd\x27\xac\xce\xbc\xb4\x86\xc0\xd0\xad\x42\xe9\xc7\x85\xec\x71\x14\x61\x2d\xc4\xa8\x8c\x74\x35\x1a\x1e\x27\xcf\xd9\xba\x8b\x06\x0c\x6a\x4b\x07\x72\x72\x2a\xb8\x33\x30\xed\x7d\xce\xb6\xdd\x38\x4f\xc4\x73\x66\xac\x81\x88\x74\x52\xa9\x40\x27\x13\x53\x51\x24\x8a\xef\x70\x0d\xde\x18\xf2\x32\xc3\xe4\xa4\xf1\x48\x68\xcd\x50\x2c\x5a\xcf\x40\x7a\xe0\x68\xb8\xed\x29\x96\x5c\xcc\xfd\x27\x4c\x5a\xf4\x44\xd6\x8c\xee\x97\xbd\xf3\xd6\xe5\xac\xb3\x78\x11\x67\xa4\x48\xb6\xcb\xd9\x2a\xbb\x9e\x60\xec\x95\xa5\xf3\x00\xa5\xd5\xc3\xff\x8f\x0d\x12\xec\x2f\x5b\xda\x48\x65\x8f\x39\x13\x2c\xeb\x4e\x03\x7d\x26\x98\xab\x0b\xf9\xb4\x49\xd2\x4d\xb2\x5e\x27\x62\xb9\x79\xfe\x22\x80\xb8\x62\x3b\x5b\x89\x66\xd8\x1a\x5e\x68\x5b\xbe\xed\xbf\x93\x76\xf3\x9d\x6c\xe9\xf6\x06\x9c\x3c\xdc\x4c\x1e\x36\xb3\xa9\x53\x5d\x78\xe5\x1e\xbe\xf7\xeb\xa7\x0c\x06\x2a\x76\xb2\xc2\x81\x96\x84\x07\xf8\x95\x49\x43\x7c\x87\x3d\x94\x68\xe6\xf0\xfe\xd4\x2a\x5a\x33\x88\x25\x7b\xb2\x13\x76\x9a\x8c\xd9\x88\x79\xf4\xdb\x4c\x0f\x6f\x80\x29\xdd\xb9\xa3\x8f\x5f\xc6\xf3\x55\x28\xa8\x87\x30\x0a\x09\x55\x35\xc4\x73\xb5\xdd\xc1\xcb\xf3\x38\xb0\xc6\xae\xb0\xd2\xa9\xdf\xb6\xcb\x5e\xd7\xb1\x5f\x5a\x96\x2f\x4a\x24\xec\x21\x53\x0a\x76\xaf\x63\x3f\xdf\x97\x25\x78\xcf\xc7\x21\x41\xdd\x09\x21\x0b\x6f\x75\x3f\x6d\x59\x58\xce\xd1\x4e\xdb\xc9\x12\xe9\x3c\xff\x3e\xa0\xc7\x02\x75\x80\x1a\x54\x0a\xcc\xe4\x4f\x58\x72\x0e\x07\x30\xe4\xaf\x5d\xfd\x92\x91\xa5\xef\x0b\x42\xd2\x30\xdd\xc8\xab\x40\xac\xe7\x3c\xdc\x1e\x82\x4b\x50\xe6\x71\x96\xbb\x6f\x96\x30\x1c\x92\xc0\x18\x8c\xfa\xcb\xa9\xbe\x0f\xd7\xe7\x9f\x00\x00\x00\xff\xff\x4d\x82\x66\xcc\x07\x06\x00\x00")

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

	info := bindataFileInfo{name: "template/application.css", size: 1543, mode: os.FileMode(420), modTime: time.Unix(1530673112, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4b\x8f\xdb\x36\x10\xbe\xfb\x57\x4c\x98\x43\x2e\x15\x69\xc9\xeb\xc7\x06\x92\x0a\x34\xed\x22\x28\x1a\x24\xc0\xa6\x87\x9e\x0a\x9a\x1c\x5b\xdc\x50\xa2\x42\x8e\xec\x75\x16\xfb\xdf\x0b\xea\xb1\xf1\x6e\xb6\x48\x0e\xf1\x81\x9e\xe7\xa7\xe1\xf0\x1b\x32\x7f\xf1\xfb\xfb\x37\x1f\xff\xf9\xf0\x07\x54\x54\xdb\x72\x96\x0f\x7f\xb3\xbc\x42\xa9\xcb\x19\x00\x40\xfe\x22\x49\xfe\x42\x82\xad\x77\xc7\x80\x1e\x3e\x35\xee\x08\x47\xdc\x06\x43\x08\x26\x80\x6b\xc9\xd4\xe6\x0b\x6a\xd8\x39\x0f\xb5\xdb\x1a\x8b\x49\x32\xe6\xd6\x48\x12\x1a\x59\x63\xc1\x0e\x06\x8f\xad\xf3\xc4\x40\xb9\x86\xb0\xa1\x82\x1d\x8d\xa6\xaa\xd0\x78\x30\x0a\x93\x5e\xf9\x05\x4c\x63\xc8\x48\x9b\x04\x25\x2d\x16\x29\x9f\x33\x10\x23\x98\x35\xcd\x27\xa8\x3c\xee\x0a\x56\x11\xb5\xe1\xb5\x10\x3b\xd7\x50\xe0\x7b\xe7\xf6\x16\x65\x6b\x02\x57\xae\x16\x2a\x84\x5f\x77\xb2\x36\xf6\x54\xbc\x8b\x7e\xf4\x5e\x12\x03\x8f\xb6\x60\x81\x4e\x16\x43\x85\x48\xec\x1c\xf5\xa9\x6f\xfc\x8c\xb8\xbb\x33\x3b\xe0\x6f\x65\xb8\x26\xb9\xc7\xfb\xfb\xbb\x3b\x3e\x4a\xe2\xee\x0e\x1b\x7d\x7f\xdf\x76\x5b\x6b\x94\x90\x6d\x6b\x8d\x92\x64\x5c\xc3\x55\x08\xac\x9c\xe5\x62\xe8\xe1\x2c\xdf\x3a\x7d\x2a\x67\x79\x2d\x4d\x03\xca\xca\x10\x0a\x16\xb0\xd1\x53\x01\xda\x1c\x26\x73\xec\x8c\x34\x0d\xfa\xd1\xd7\xfb\xab\xb4\xfc\xb3\x0b\x04\xef\x1b\x85\xb9\xa8\xd2\x33\x57\xfb\x80\xd7\x6d\xc9\x90\x45\x56\x5e\x63\xa3\x41\x42\x8d\x21\xc8\x3d\x02\x55\x92\x20\xa0\xdd\x81\xc6\x40\xbe\x53\x14\x40\xee\x08\x3d\xb8\x06\x21\x9e\x49\x2e\xda\x33\xc4\xef\x14\xd3\xc7\xec\x9c\xaf\xc1\xe8\xb8\x0b\xe5\x91\xa2\xca\x40\xaa\xb8\xf7\x1f\x6a\x99\xf2\x28\x09\x19\xd4\x48\x95\xd3\x05\xfb\xf0\xfe\xfa\x23\x03\x6c\x14\x9d\x5a\x2c\x58\xdd\x59\x32\xad\xf4\x24\x22\x72\xa2\x25\x49\x56\xce\x1e\x95\xf0\xb4\x54\xd3\xb4\x1d\x25\x3b\x83\x56\x3f\x29\xf6\x21\xfa\x45\x92\xfc\xdd\x5a\x27\x35\x5c\xf7\x55\xc3\x95\xb1\xf8\x1a\xf2\x3e\x35\xee\xe6\xd5\x2e\x12\xb7\x57\x5f\xc1\x50\x49\xb4\xb0\x91\xbe\x83\x1c\xcc\x17\x2c\xb2\x65\x99\x8b\xad\x2f\x23\xcb\x9f\xff\x18\xe1\x2d\x49\x8f\xb2\xef\xd2\xa4\xa4\x13\x54\x1d\xf6\x0c\x5a\x2b\x15\x56\xce\x6a\xf4\x05\xfb\x20\x03\x21\x9c\x5c\xe7\x1f\x0e\xae\x42\x8f\x91\xb3\x9f\x3b\xe3\x51\x17\x6c\x92\x58\x99\x8b\x09\xf1\xdb\xad\xe6\x42\x9b\xc3\x33\xe6\xb3\x5e\x6d\x3b\x22\xd7\xfc\x7b\xf4\xb2\x6d\xbf\x39\xdb\x87\x84\x21\x6a\xca\xc1\x46\xf9\x53\x4b\x6c\xec\x4b\xe8\xb6\xb5\xa1\x69\x3b\xc3\xc9\xb3\xf2\xba\xb7\x3e\x0b\xd7\x43\x86\xc3\x1e\x6e\x6b\xdb\x84\x61\x7c\x5f\x0b\x71\x3c\x1e\xf9\x71\xc1\x9d\xdf\x8b\x6c\x3e\x9f\x8b\x70\xd8\x33\x18\x2e\x05\x96\x66\x0c\x2a\x34\xfb\x8a\x06\x39\x52\xf5\x37\x77\x5b\xb0\x39\xcc\x21\x5b\x6d\xf8\x66\x91\x4d\xff\xac\xcc\x5b\x49\x15\xec\x8c\xb5\x05\x7b\x79\x75\x75\xc5\x40\x17\xec\x5d\xb6\x5a\xf2\x74\x0d\x69\xb6\xe4\xcb\xf5\xda\x26\x9b\x79\xb2\x99\xab\xe4\x82\x6f\x36\xc3\x92\x66\x7c\x7d\xb9\x1a\xe5\x35\x5f\xad\xd7\x30\xef\xb5\x0c\xfa\x75\x94\x87\x30\x98\x43\x1f\xb3\xb1\xcb\x0d\x5f\xad\xa0\x5f\xdf\xa6\x19\x5f\xaa\x64\xc5\x2f\xe7\x0b\x98\x47\xc0\x25\x2c\xf9\xf2\x72\x33\x88\xfd\x32\x87\xc1\x1d\xed\x6b\xf8\x6a\x8f\x4b\x95\xa5\x0b\xbe\x5a\x5e\xd8\xe4\x0c\x34\x1b\x6a\x3c\xab\xe1\x69\x09\x90\xf1\x8b\x8b\x61\x59\xf2\xd5\x05\x2c\x62\xee\x86\x6f\x06\x29\xac\xf8\x22\x56\xc0\xb3\xac\x37\x26\xd1\x68\xd7\x97\xfc\xf2\x72\x1d\x3b\x10\x11\x17\x67\xfb\x5c\x24\x13\xfc\xd0\x85\xcd\x17\x26\xca\x3c\x1e\xc7\xff\xb0\x43\x0c\xf4\xf8\x21\xfe\xe5\xfd\x18\x3f\xb1\x9d\xf1\x51\x9b\x83\xd1\x91\x88\xcf\xe5\x9e\xc5\x85\x4e\x29\x0c\x21\x19\xb9\x88\xcf\x4d\xfa\x79\xbc\xeb\xe8\xbb\x97\xc2\xa3\x39\xed\xbc\x65\x53\xf2\xe4\x88\x13\x28\xb5\x6b\xec\xa9\x20\xdf\xe1\xcf\x99\xbd\x1f\x9b\xb9\x2d\x45\xd1\xb4\x5b\x27\xbd\x66\x10\xaf\xc1\xe4\x41\x4f\x48\xfa\x3d\x52\xc1\x5e\xc6\xaa\x9f\xc7\x8b\xbf\x37\xae\x3d\x01\x39\x78\x33\x25\xfe\x8c\xf3\x7c\x64\x3a\x53\x47\x71\x96\x8b\xf8\xd4\xc5\x17\x70\x7c\xf9\x82\xf2\xa6\xa5\xf1\xf2\x88\x0d\x14\x37\xf2\x20\x07\x2b\x83\xe0\xd5\xd7\x07\x5d\x39\x8d\xfc\xe6\x73\x87\xfe\xd4\xbf\xe5\x83\x98\x64\x3c\xe5\x29\xaf\x4d\xc3\x6f\x42\x64\xca\x90\xfb\x15\xfa\x31\x86\x6e\x6e\x02\x57\xd6\x75\x7a\x67\xa5\xc7\x1e\x48\xde\xc8\x5b\x61\xcd\x36\x88\x87\x26\xf2\x9b\x20\x52\xbe\xe6\xe9\x99\xe9\xdb\x4f\xfc\x17\x00\x00\xff\xff\x04\x6e\x35\xc6\x23\x09\x00\x00")

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

	info := bindataFileInfo{name: "template/index.html", size: 2339, mode: os.FileMode(420), modTime: time.Unix(1530676874, 0)}
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

