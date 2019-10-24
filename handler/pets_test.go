package handler

import (
	"context"
	"errors"
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
		if pet.Id != v {
			t.Errorf("GetPetByID got ID %v, expected %v", pet.Id, v)
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
		Id: 0,
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

	t.Logf("Created new pet with ID of %#v", newPet.Id)
	if newPet.Id != 6 {
		t.Errorf("New PetID %#v, want 6", newPet.Id)
		return
	}

	// Delete the new pet
	db.DeletePet(ctx, 6)
}

func TestDeletePet(t *testing.T) {
	t.Log(("Testing DeletePet"))

	db := testdb.New()
	petServer := NewPetstoreServer(logrus.New(), db, testRPCPort, testRestPort, testAPIKey)
	ctx := context.Background()

	expectedErr := errors.New("400:Invalid API_KEY Supplied")
	contextApiKey := "api_key"

	// No api_key with valid pet == fail
	_, err := petServer.DeletePet(ctx, &pb.PetID{PetId: 1})
	if err.Error() != expectedErr.Error() {
		t.Errorf("Delete pet with no api_key gets %#v, expecting %#v", err.Error(), expectedErr.Error())
	}

	// Invalid api_key with valid pet == fail
	_, err = petServer.DeletePet(context.WithValue(ctx, contextApiKey, "ABCxxx"), &pb.PetID{PetId: 1})
	if err.Error() != expectedErr.Error() {
		t.Errorf("Delete pet with invalid key gets %#v, expecting %#v", err.Error(), expectedErr.Error())
	}

	// valid api_key with invalid pet == fail
	expectedErr = errors.New("404:Pet not found 100")
	_, err = petServer.DeletePet(context.WithValue(ctx, contextApiKey, testAPIKey), &pb.PetID{PetId: 100})
	if err.Error() != expectedErr.Error() {
		t.Errorf("Delete pet with valid key and invalid petID gets %#v, expecting %#v", err.Error(), expectedErr.Error())
	}

	// valid api key with valid pet == pass
	_, err = petServer.DeletePet(context.WithValue(ctx, contextApiKey, testAPIKey), &pb.PetID{PetId: 1})
	if err != nil {
		t.Errorf("Delete pet with valid key and valid petID gets error %#v, expecting nil", err.Error())
	}

	// check that the pet is indeed deleted
	pet, err := db.GetPetByID(ctx, 1)
	if pet != nil {
		t.Errorf("Pet 1 still exists after being deleted %#v", pet)
	}
	expectedErr = errors.New("404:Pet 1 not found")
	if err.Error() != expectedErr.Error() {
		t.Errorf("Deleting pet from DB gets unexpected error %#v, expecting %#v", err.Error(), expectedErr.Error())
	}
}

func TestFindPetsByStatus(t *testing.T) {

}

func TestUpdatePet(t *testing.T) {
	t.Log(("Testing DeletePet"))

	db := testdb.New()
	petServer := NewPetstoreServer(logrus.New(), db, testRPCPort, testRestPort, testAPIKey)
	ctx := context.Background()

	// get pet 1
	pet1, err := db.GetPetByID(ctx, 1)

	// Update pet 1
	myPet := &pb.Pet{
		Id: 1,
		Category: &pb.Category{
			Id:   1,
			Name: "dog",
		},
		Name:      "test dog",
		PhotoUrls: []string{"photos/test1.jpg"},
		Tags: []*pb.Tag{
			&pb.Tag{
				Id:   1,
				Name: "housetrained",
			},
		},
		Status: "available",
	}

	pet, err := petServer.UpdatePet(ctx, myPet)
	if err != nil {
		t.Errorf("Updating pet 1 gets unexpected error %#v", err.Error())
	}
	if pet.Name != myPet.Name {
		t.Errorf("Updated pet returns with a name of %#v, expecting %#v", pet.Name, myPet.Name)
	}
	if pet.Id != 1 {
		t.Errorf("Updated pet returns with an ID of %#v, expecting 1", pet.Id)
	}

	// now get pet1 again, and confirm that its changed
	pet1updated, err := db.GetPetByID(ctx, 1)
	if pet1updated.Name != myPet.Name {
		t.Errorf("Updated pet1 (old name %#v) has new name %#v, expecting %#v", pet1.Name, pet1updated.Name, myPet.Name)
	}

	// Now update again, with ID = 0, and observe that it inserts a new one
	myPet.Id = 0
	pet, err = petServer.UpdatePet(ctx, myPet)
	if err != nil {
		t.Errorf("Updating new pet 0 gets unexpected error %#v", err.Error())
	}
	if pet.Name != myPet.Name {
		t.Errorf("Updated new pet returns with a name of %#v, expecting %#v", pet.Name, myPet.Name)
	}
	if pet.Id != 6 {
		t.Errorf("Updated new pet returns with an ID of %#v, expecting 6", pet.Id)
	}
}

func TestUploadFile(t *testing.T) {

}
