package cache
import (
	"github.com/boltdb/bolt"
	"path/filepath"
	"errors"
	"os"
)

const (
	cache_db = "initialser.db"

)
var (
	cache_bucket = []byte("initialser_bucket")
)

type BoltCache struct {
	Base string
	db   *bolt.DB
}

func NewBoltCache(base string) KV {
	os.Mkdir(base,defaultPathPerm);
	db, err := bolt.Open(filepath.Join(base, cache_db), 0600, nil)
	if err != nil {
		println(err.Error())
		return nil
	}
	//db.NoSync = true;
	return &BoltCache{
		Base:base,
		db:db,
	}
}

func (bc *BoltCache)Get(key []byte) ([]byte, bool) {
	var data []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(cache_bucket)
		if b == nil {
			return errors.New("bucket not exists");
		}
		data = b.Get(key)
		return nil
	})
	return data, err == nil && data != nil
}

func (bc *BoltCache)Set(key []byte, data []byte) error {
	return bc.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(cache_bucket);
		if err != nil {
			return err
		}
		return b.Put(key, data);
	});
}

func (bc *BoltCache)Clear() error {
	return bc.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(cache_bucket);
	});
}
func (bc *BoltCache)Close() {
	bc.db.Sync()
	bc.db.Close()
}