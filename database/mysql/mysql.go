package mysql

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	pb "github.com/steveoc64/petstore/proto"
)

// MysqlDB is an INCOMPLETE implementation of a mySQL driver
type MysqlDB struct {
	sync.RWMutex
	sql *sqlx.DB
	log *logrus.Logger
}

// NewMysqlDB returns a new mysql connection for the given DSN
func NewMysqlDB(log *logrus.Logger, dsn string) (*MysqlDB, error) {
	sql, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MysqlDB{sql: sql, log: log}, nil
}

// GetPetByID returns the pet by the given ID, or nil + error if not found
func (db *MysqlDB) GetPetByID(ctx context.Context, id int64) (*pb.Pet, error) {
	// TODO - handle context timeouts.
	db.RLock()
	defer db.RUnlock()
	pet := &pb.Pet{}
	err := db.sql.Select(pet, "select * from pets where id=$1", id)
	if err != nil {
		// Log the error, but return a clean 404 to the client
		logrus.WithField("id", id).WithError(err).Error("SQL error looking up pet")
		return nil, fmt.Errorf("404:Pet %v not found", id)
	}
	return pet, nil
}

// UpdatePet upadates the name and status of a pet
func (db *MysqlDB) UpdatePet(ctx context.Context, id int64, name string, status string) error {
	db.Lock()
	db.Unlock()
	_, err := db.sql.NamedExec("update pets set name=:name,status=:status where id=:id",
		map[string]interface{}{
			"name":   name,
			"status": status,
		})
	if err != nil {
		// Log the error, but return a clean 405 to the client
		logrus.WithField("id", id).WithError(err).Error("SQL error updating pet")
		return fmt.Errorf("405:Pet %v invalid entry", id)
	}
	return nil
}

// AddPet adds a pet to the database, unless it already exists or is invalid
func (db *MysqlDB) AddPet(ctx context.Context, pet *pb.Pet) error {
	db.Lock()
	defer db.Unlock()
	var count int
	if err := db.sql.Select(&count, "select count(*) from pets where id=$1", pet.PetId); err != nil {
		return err
	}
	if count == 1 {
		return fmt.Errorf("405:Pet already exists %d", pet.PetId)
	}

	// TODO - full and correct SQL for mapping the pet the the various DB tables
	_, err := db.sql.NamedExec("insert into pets (...) values (...)", pet)
	if err != nil {
		db.log.WithError(err).Error("SQL error inserting new pet")
		return errors.New("405:Invalid input data")
	}
	return nil
}
