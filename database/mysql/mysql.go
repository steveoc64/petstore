package mysql

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	pb "github.com/steveoc64/petstore/proto"
)

// TODO - full and correct SQL for mapping the pet the the various DB tables

// DB is an INCOMPLETE implementation of a mySQL driver
type DB struct {
	sql *sqlx.DB
	log *logrus.Logger
}

// New returns a new mysql connection for the given DSN
func New(log *logrus.Logger, dsn string) (*DB, error) {
	sql, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{sql: sql, log: log}, nil
}

// GetPetByID returns the pet by the given ID, or nil + error if not found
func (db *DB) GetPetByID(ctx context.Context, id int64) (*pb.Pet, error) {
	pet := &pb.Pet{}
	err := db.sql.SelectContext(ctx, pet, "select * from pets where id=$1", id)
	if err != nil {
		// Log the error, but return a clean 404 to the client
		logrus.WithField("id", id).WithError(err).Error("SQL error looking up pet")
		return nil, fmt.Errorf("404:Pet %v not found", id)
	}
	return pet, nil
}

// UpdatePetWithForm updates the name and status of a pet
func (db *DB) UpdatePetWithForm(ctx context.Context, id int64, name string, status string) error {
	_, err := db.sql.NamedExecContext(ctx, "update pets set name=:name,status=:status where id=:id",
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
func (db *DB) AddPet(ctx context.Context, pet *pb.Pet) error {
	// If the PetID is not specified, use an auto-increment
	if pet.PetId == 0 {
		_, err := db.sql.ExecContext(ctx, "insert into pets (...everything but the id) values (...)")
		return err
	}
	var count int
	if err := db.sql.SelectContext(ctx, &count, "select count(*) from pets where id=$1", pet.PetId); err != nil {
		return err
	}
	if count == 1 {
		return fmt.Errorf("405:Pet already exists %d", pet.PetId)
	}

	_, err := db.sql.NamedExecContext(ctx, "insert into pets (...) values (...)", pet)
	if err != nil {
		db.log.WithError(err).Error("SQL error inserting new pet")
		return errors.New("405:Invalid input data")
	}
	return nil
}

// DeletePet deletes a pet
func (db *DB) DeletePet(ctx context.Context, id int64) error {
	var count int
	if err := db.sql.SelectContext(ctx, &count, "select count(*) from pets where id=$1", id); err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("404:Pet not found %d", id)
	}
	_, err := db.sql.ExecContext(ctx, "delete from pets where id=$1 limit 1", id)
	return err
}

// UpdatePet to the new contents
func (db *DB) UpdatePet(ctx context.Context, pet *pb.Pet) error {
	// In the SwaggerAPI example, if you enter a pet with ID 0, then it
	// creates a new pet and returns 200.  We will do the same here
	if pet.PetId == 0 {
		return db.AddPet(ctx, pet)
	}
	// if the petID does not exist, then 404
	var count int
	err := db.sql.SelectContext(ctx, &count, "select count(*) from pets where id=$1", pet.PetId)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("404:Pet %d not found", pet.PetId)
	}
	_, err = db.sql.ExecContext(ctx, "delete from pets where id=$1", pet.PetId)
	return err
}

// FindPetsByStatus returns a list of pets that match any of the given status codes
func (db *DB) FindPetsByStatus(ctx context.Context, statuses []string) (*pb.Pets, error) {
	pets := &pb.Pets{}
	err := db.sql.SelectContext(ctx, pets.Pets, "select * from pets where status in (?)", statuses)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

// UploadFile records the uploaded file against the pet
func (db *DB) UploadFile(ctx context.Context, id int64, filename string) error {
	// TODO - need some hacking to get this working with grpc
	return nil
}
