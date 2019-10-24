package handler

import (
	"context"
	"fmt"
	"testing"

	"github.com/steveoc64/petstore/database/testdb"

	"github.com/sirupsen/logrus"
	pb "github.com/steveoc64/petstore/proto"
)

const (
	testRestPort = 8082
	testRPCPort  = 8083
	testAPIKey   = "ABC123"
)

// TestGetPetByID provides full test coverage for the GetPetByID call
func TestGetPetByID(t *testing.T) {
	t.Log("Testing GetPetByID")

	petServer := NewPetstoreServer(logrus.New(), testdb.New(), testRPCPort, testRestPort, testAPIKey)
	ctx := context.Background()

	for _, i := range []int{1, 2, 3, 4, 5} {
		v := int64(i)
		pet, err := petServer.GetPetByID(ctx, &pb.PetID{PetId: v})
		if err != nil {
			t.Error("GetPetByID 1 unexpected error", err.Error())
		}
		if pet.PetId != v {
			t.Errorf("GetPetByID got ID %v, expected %v", pet.PetId, v)
		}
		expectedName := fmt.Sprintf("Pet number %d", v)
		if pet.Name != expectedName {
			t.Errorf("GetPetByID got Name %#v, want %#v", pet.Name, expectedName)
		}

		t.Logf("Got Pet %#v", pet)
	}
}

func TestUpdatePetWithForm(t *testing.T) {
	t.Log(("Testing UpdatePetWithForm"))

	petServer := NewPetstoreServer(logrus.New(), testdb.New(), testRPCPort, testRestPort, testAPIKey)
	ctx := context.Background()

	// get the old pet
	oldPet, err := petServer.GetPetByID(ctx, &pb.PetID{PetId: 1})
	if err != nil {
		t.Error("GetPetByID 1 unexpected error", err.Error())
		return
	}

	// update the new pet
	newName := "New Name"
	newStatus := "pending"
	petServer.UpdatePetWithForm(ctx, &pb.UpdatePetWithFormReq{
		PetId:  1,
		Name:   newName,
		Status: newStatus,
	})

	newPet, err := petServer.GetPetByID(ctx, &pb.PetID{PetId: 1})
	if newPet.Name != newName {
		t.Errorf("After update, name %#v, want %#v", newPet.Name, newName)
	}
	if newPet.Status != newStatus {
		t.Errorf("After update, status %#v, want %#v", newPet.Status, newStatus)
	}

	// put the pet back the way we started
	petServer.UpdatePetWithForm(ctx, &pb.UpdatePetWithFormReq{
		PetId:  1,
		Name:   oldPet.Name,
		Status: oldPet.Status,
	})
}

func TestAddPet(t *testing.T) {
	t.Log(("Testing TestAddPet"))

	db := testdb.New()
	petServer := NewPetstoreServer(logrus.New(), db, testRPCPort, testRestPort, testAPIKey)
	ctx := context.Background()

	// Add a new pet
	newName := "My New Dog"
	newPet, err := petServer.AddPet(ctx, &pb.Pet{
		PetId: 0,
		Category: &pb.Category{
			Id:   1,
			Name: "dog",
		},
		Name:      newName,
		PhotoUrls: []string{"photos/newdog.jpg"},
		Tags: []*pb.Tag{
			&pb.Tag{
				Id:   1,
				Name: "housetrained",
			},
			{
				Id:   3,
				Name: "newly created",
			},
		},
		Status: "pending",
	})

	if err != nil {
		t.Error("AddPet unexpected error", err.Error())
		return
	}

	t.Logf("Created new pet with ID of %#v", newPet.PetId)
	if newPet.PetId != 6 {
		t.Errorf("New PetID %#v, want 6", newPet.PetId)
		return
	}

	// Delete the new pet
	db.DeletePet(ctx, 6)
}

func TestDeletePet(t *testing.T) {

}

func TestFindPetsByStatus(t *testing.T) {

}

func TestUpdatePet(t *testing.T) {

}

func TestUploadFile(t *testing.T) {

}
