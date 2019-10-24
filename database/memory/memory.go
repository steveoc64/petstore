package memory

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"

	pb "github.com/steveoc64/petstore/proto"
)

// DB implements a database in memory
type DB struct {
	sync.RWMutex
	Pets map[int64]*pb.Pet
}

// New returns a clean new memoryDB
func New() *DB {
	return &DB{Pets: make(map[int64]*pb.Pet)}
}

// clonePet to create deep copies of Pet objects.
// Always use this when returning a Pet object from a DB call.
//
// We want to do this because if we just return pointers to objects
// that are in a constant state of change, bad things can and will happen.
func (db *DB) clonePet(pet *pb.Pet) (*pb.Pet, error) {
	newPet := &pb.Pet{}
	data, err := proto.Marshal(pet)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(data, newPet)
	if err != nil {
		return nil, err
	}
	return newPet, nil
}

// GetPetByID returns the pet by the given ID, or nil + error if not found
func (db *DB) GetPetByID(ctx context.Context, id int64) (*pb.Pet, error) {
	db.RLock()
	defer db.RUnlock()
	pet, ok := db.Pets[id]
	if !ok {
		return nil, fmt.Errorf("404:Pet %#v not found", id)
	}
	copy, err := db.clonePet(pet)
	if err != nil {
		return nil, fmt.Errorf("500:Failed to get pet %d", id)
	}
	return copy, nil
}

// UpdatePetWithForm updates the name and status of a pet
func (db *DB) UpdatePetWithForm(ctx context.Context, id int64, name string, status string) error {
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

// DeletePet deletes a pet
func (db *DB) DeletePet(ctx context.Context, id int64) error {
	db.Lock()
	defer db.Unlock()
	if _, ok := db.Pets[id]; !ok {
		return fmt.Errorf("404:Pet not found %d", id)
	}
	delete(db.Pets, id)
	return nil
}

// AddPet adds a pet to the database, unless it already exists or is invalid
func (db *DB) AddPet(ctx context.Context, pet *pb.Pet) error {
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

// UpdatePet to the new contents
func (db *DB) UpdatePet(ctx context.Context, pet *pb.Pet) error {
	db.Lock()
	defer db.Unlock()
	// In the SwaggerAPI example, if you enter a pet with ID 0, then it
	// creates a new pet and returns 200.  We will do the same here
	if pet.PetId == 0 {
		return db.AddPet(ctx, pet)
	}
	// if the petID does not exist, then 404
	if _, ok := db.Pets[pet.PetId]; !ok {
		return fmt.Errorf("404:Pet %d not found", pet.PetId)
	}
	db.Pets[pet.PetId] = pet
	return nil
}

// FindPetsByStatus returns a list of pets that match any of the given status codes
func (db *DB) FindPetsByStatus(ctx context.Context, statuses []string) (*pb.Pets, error) {
	db.RLock()
	defer db.RUnlock()
	pets := &pb.Pets{}
	for _, pet := range db.Pets {
		for _, status := range statuses {
			if strings.EqualFold(pet.Status, status) {
				copy, err := db.clonePet(pet)
				if err != nil {
					return nil, fmt.Errorf("500:Error fetching pet")
				}
				pets.Pets = append(pets.Pets, copy)
			}
		}
	}
	return pets, nil
}

// UploadFile records the uploaded file against the pet
func (db *DB) UploadFile(ctx context.Context, id int64, filename string) error {
	db.Lock()
	db.Unlock()
	// TODO - need some hacking to get this working with grpc
	return nil
}
