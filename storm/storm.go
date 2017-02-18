package storm

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/gob"
	srs "github.com/nii236/srs/models"
)

var dbclient *DB

type DB struct {
	DB *storm.DB
	*UserService
}

// UserService is the storm implementation of srs.UserService
type UserService struct {
	*storm.DB
}

// New returns a singleton instance of the DB
func New() (*DB, error) {
	if dbclient == nil {
		var err error
		db, err := storm.Open("./srs.db", storm.Codec(gob.Codec))
		if err != nil {
			return nil, err
		}

		dbclient = &DB{
			DB:          db,
			UserService: &UserService{DB: db},
		}
	}

	dbclient.DB.Init(&srs.User{})
	return dbclient, nil
}

// Close will close the DB connection
func Close() error {
	return dbclient.DB.Close()
}
