package cache
import (
	"sync"
	"path/filepath"
	"strings"
	"io/ioutil"
	"os"
	"encoding/hex"
)

const (
	defaultFilePerm os.FileMode = 0600
	defaultPathPerm os.FileMode = 0777
)

type KV interface {
	Get(key []byte) ([]byte, bool)
	Set(key []byte, data []byte) error
	Clear() error
	Close()
}
//TransformFunc parse key to path ,eg a_b_c_d ===> ["a","b","c","d"]
// the final path will be <base>/a/b/c/d
type TransformFunc func(s string) []string

//SimpleDiskCache persists to disk,reduce the number of generation
type SimpleDiskCache struct {
	Base      string
	Transform TransformFunc
	mu        sync.RWMutex
}

func NewSimpleDiskCache(base string, transforms ...TransformFunc) KV {
	transform := func(s string) []string {
		return strings.Split(s, "_")
	}
	if len(transforms) > 0 {
		transform = transforms[0]
	}
	return &SimpleDiskCache{
		Base:base,
		Transform:transform,
	}
}

func (sdc *SimpleDiskCache)Get(key []byte) ([]byte, bool) {
	fp := filepath.Join(sdc.Base, filepath.Join(sdc.Transform(hex.EncodeToString(key))...))
	abs, _ := filepath.Abs(fp)
	data, err := ioutil.ReadFile(abs)
	return data, err == nil
}
//Set if exists, override
func (sdc *SimpleDiskCache)Set(key []byte, data []byte) error {
	fp := filepath.Join(sdc.Base, filepath.Join(sdc.Transform(hex.EncodeToString(key))...))
	os.MkdirAll(filepath.Dir(fp), defaultPathPerm)
	return ioutil.WriteFile(fp, data, defaultFilePerm);
}

func (sdc *SimpleDiskCache)Clear() error {
	return os.RemoveAll(sdc.Base)
}

func (sdc *SimpleDiskCache) Close() {

}


