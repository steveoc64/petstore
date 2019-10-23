package testdb

import (
	"context"
	"fmt"
	"strings"
	"sync"

	pb "github.com/steveoc64/petstore/proto"
)

type TestDB struct {
	sync.RWMutex
	Pets map[int64]*pb.Pet
}

func NewTestDB() *TestDB {
	db := &TestDB{Pets: make(map[int64]*pb.Pet)}

	for v := range []int{1, 2, 3, 4, 5} {
		i := int64(v)
		db.Pets[i] = &pb.Pet{
			PetId: i,
			Category: &pb.Category{
				Id:   1,
				Name: "dog",
			},
			Name:      fmt.Sprintf("Pet number %d", i),
			PhotoUrls: []string{"photos/234.jpg", "photos/345.jpg"},
			Tags: []*pb.Tag{
				{
					Id:   1,
					Name: "housetrained",
				},
				{
					Id:   2,
					Name: "good-with-kids",
				},
			},
			Status: "available",
		}
	}
	return db
}

// GetPetByID returns the pet by the given ID, or nil + error if not found
func (db *TestDB) GetPetByID(ctx context.Context, id int64) (*pb.Pet, error) {
	// TODO - handle context timeouts.
	// .. not that they are likely in an in-Memory DB driver, but needed for completeness
	db.RLock()
	defer db.RUnlock()
	pet, ok := db.Pets[id]
	if !ok {
		return nil, fmt.Errorf("404:Pet %#v not found", id)
	}
	return pet, nil
}

// UpdatePet updates the name and status of a pet
func (db *TestDB) UpdatePet(ctx context.Context, id int64, name string, status string) error {
	db.Lock()
	db.Unlock()
	pet, ok := db.Pets[id]
	if !ok {
		return fmt.Errorf("405:No Pet %d to update", id)
	}
	status = strings.ToLower(status)
	if _, ok := pb.Status_value[status]; !ok {
		return fmt.Errorf("405:Invalid status value %#v", status)
	}
	pet.Name = name
	pet.Status = status
	return nil
}

// AddPet adds a pet to the database, unless it already exists or is invalid
func (db *TestDB) AddPet(ctx context.Context, pet *pb.Pet) error {
	db.Lock()
	defer db.Unlock()
	// If the PetID is not specified, use an auto-increment
	if pet.PetId == 0 {
		pet.PetId = int64(len(db.Pets) + 1)
	}
	if _, ok := db.Pets[pet.PetId]; ok {
		return fmt.Errorf("405:Pet already exists %d", pet.PetId)
	}
	db.Pets[pet.PetId] = pet
	return nil
}
