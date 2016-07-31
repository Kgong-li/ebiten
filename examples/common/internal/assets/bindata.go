package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _arcadefont_png = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x80\x00\x00\x000\b\x03\x00\x00\x00\xee*\x82\x8c\x00\x00\x00\x06PLTE\x00\x00\x00\xff\xff\xff\xa5ٟ\xdd\x00\x00\x00\x01tRNS\x00@\xe6\xd8f\x00\x00\x03\x98IDATx\x9c\xc4X\x89\x96\xac(\f\xe5\xd6\xf3\xff\u007f9\xd3B\x96\x9b\x18ѩ\xe99/ݥ\x82\x81\xec\v~F\x00\xa0\xd7\xf5\x10\xf3t=\xdf#\xe1\xdbCY\xf4\x0e0>\xbb\x97F\xf9\xdc{\x8e\xe6\xbf\b*\xa2\xfc\\%X,\xf7\r\ti\x18\x90\xf9\xbf(bm#sNN|\xdeW\x98\xbes\x80W\xa4\x8d~B];\xc1\t\x89\x13\x14ǧ\xa1\x8e\x9d\x01{P\x04\x10\x8f\xa3\x19\xfbz\x9a@\x1e\x99\xb5\xf5\x0e\xffG\xc2,>\x01\xb2]\xb7o\xe6(\x99@\x12G2\xff\xd6rӌ\x8b!\t\xd1\xe8\x89\tYE\xeb\x04\xbd\xc8\u007f\xef\xc7H\x174a\xf2\x10\x05\xd5i\xff\x85\x9b\xfc\nli\x81\x9d,9\xcfZ\xa63\x92\x87\x11s\xa2\xc3\xceɈ\x884\x83\xb5\xf0ϴm(y\xbd\a*G\x8d\xaf\x8b\xe5\x05\xba[*\x1b\xe9^\x9c?6\xfcy<&\x1f8=B\xc02\x12}\x01q\x1cQOq\xbb\x85\x88N\xf5\xd0I\xcd\xe5Q\t$cJ\"\xc8҇)HA\xc1u\xe3\xda\xd79\x90͏f\xc1\xf0x\x82\xe5*\x962\xe6]A\xb0\x90\xbd\xec\xef\x02\x93\x8f\xa4\xf7\x87\xaa\x9f\xe9',SM#Hp\xd4\x1b\xdb^'\x95\x13\x96\xf1\x85\xa4\xc2\xe4\xdd\xccj\x8a\x8eb1\x9b/\xf4!\rc1\xd8\xc8\xf4;\xf0*׀\x95~\t\u007fR\x84i\xc65\xa4o܉\xd9\t\x91\xe69l\v\xfe\xe9\x84ӟ\xf4\xa5\x19\vP\xf7_\xfe\x91\x8d\xd8\xe9u\x8c\xceM\xf2\x04,\x9c\x89!\xc8g\xe6!wd\xe4Œ*\xfdf\xf3\x16*Jۨ\x1c\xc1\x19\xe3u\x9b\xe7\xfc\xa03ro\xea{\x16\x91\x198ŧ\xfc\xe7q^\xa1\x98`P\xee\xbc+\xa3\xb9\xbe\xab\xa4Y\xde\xcf\xe9\x15K\xd3\"\x97*\x0f\xdc\xcaҙe\xe2rǆ\a\tƩ\x01A\xed\x94\x16\x8e\x00T\fǥ\xdcI\xd3;ώ5R)\xa9\xb6\xe8\u0096\xffω\xe0\r\x1c\x1e\xbe*`\x1a\xa7\f\x99㸉k\x89b\x15\x1b \xf2^nn\x15>\xcb\xf63\fE\xbb\v\x1f\xeb\xb0v\x8b\x05\x90L-\xe4J\"\xe4r\"-\xfe\xe6`\xf2@7U\x81\xeaD\\<kG\x98\xfd\xb7e@\x95\x87\xbeϤ\x13\x93r\xa0? 9\xed\r\a?R\xb1`\a\x19\aI\x80\xb9\xd75\x1f̤\x99\xab\x9bZUz\xfa3\x01t\x82,\xf8\x90\xd5\xc4)c\xdcw\xd9Q\x16.%l\xb4\xf4\xafV(\fؾ\x00\xa7\xb6i\x02\x0e{\x1f\x8b\xd3\xcf\xf7\xe0!\x8eT\xc4\xc1\xf6|\xf4\x17\x80*S)\x9aMv\xaa}un\xe1\xae\rNid\xb7\xf9\xee\xcf(\xca\xe9TRs\xae<\xe0\xe7\xf9\xbd\x92Ϯ\xd8k+\xbbUɄ>\xc6(=\xe2(%\x83\x9e\x904\xb5\xf0y\xd9y\x8f<`\x89P3 \xf8>bl\x89\x11\x1e\xcd\"\xa5\xdc\\m\xa2\xb3\x9eZi\xffc\xa7\xa9\xbe\xa1y\v\xef\x9aރ\xd1+\x89\x8b\b<z\xf1Q\xea\xe9\xd0vB\xce\x03k\x19T]q\x1fMj\x16\xb9(\bcېd\x01\xfcT\xf8-<\xb5\x8e\xbf\a_}\x00\xe4\x95\xd8m\xb3-\xc7\xff\x11\xaa\v\xb5\x1c\x1c).s\x97㫨\xe1I?\xc3\xcd?\xdb\xf0\x15[\xb3-\x8f\x83\x90\xa2p^ٞ\x8b\a\xf5\xae.\xe8\xe0\xc3XƔ\xa1\x9f!#>:\x13\x88m\xf4`|\x17\xb6\x86\x83U\xc9.\x85\xcb\xf0/\x06\xc6@\x0e\xb7`\xa3\x8f\x11o|\xac#BtP\x89\xfe\x068\x80\x81Z\xd9\xea'\x9d\xd4s\xc8`\r'\u007f\x18\xe3\xe2\x03\x86\xd9Z䖁o\xe1\xcb\xf8\xff'\x00\x00\xff\xff\"\xd9\xedҰ\xcb\xea\x95\x00\x00\x00\x00IEND\xaeB`\x82")

func arcadefont_png_bytes() ([]byte, error) {
	return _arcadefont_png, nil
}

func arcadefont_png() (*asset, error) {
	bytes, err := arcadefont_png_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "arcadefont.png", size: 1008, mode: os.FileMode(420), modTime: time.Unix(1463226773, 0)}
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
	"arcadefont.png": arcadefont_png,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"arcadefont.png": {arcadefont_png, map[string]*_bintree_t{}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
