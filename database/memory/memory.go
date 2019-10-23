package memory

import (
	"context"
	"fmt"
	"strings"
	"sync"

	pb "github.com/steveoc64/petstore/proto"
)

type MemoryDB struct {
	sync.RWMutex
	Pets map[int64]*pb.Pet
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{Pets: make(map[int64]*pb.Pet)}
}

// GetPetByID returns the pet by the given ID, or nil + error if not found
func (db *MemoryDB) GetPetByID(ctx context.Context, id int64) (*pb.Pet, error) {
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

// UpdatePet upadates the name and status of a pet
func (db *MemoryDB) UpdatePet(ctx context.Context, id int64, name string, status string) error {
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
func (db *MemoryDB) AddPet(ctx context.Context, pet *pb.Pet) error {
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
