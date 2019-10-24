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

	// No api_key with valid pet == fail
	_, err := petServer.DeletePet(ctx, &pb.PetID{PetId: 1})
	if err.Error() != expectedErr.Error() {
		t.Errorf("Delete pet with no api_key gets %#v, expecting %#v", err.Error(), expectedErr.Error())
	}

	// Invalid api_key with valid pet == fail
	_, err = petServer.DeletePet(context.WithValue(ctx, contextAPIKey, "ABCxxx"), &pb.PetID{PetId: 1})
	if err.Error() != expectedErr.Error() {
		t.Errorf("Delete pet with invalid key gets %#v, expecting %#v", err.Error(), expectedErr.Error())
	}

	// valid api_key with invalid pet == fail
	expectedErr = errors.New("404:Pet not found 100")
	_, err = petServer.DeletePet(context.WithValue(ctx, contextAPIKey, testAPIKey), &pb.PetID{PetId: 100})
	if err.Error() != expectedErr.Error() {
		t.Errorf("Delete pet with valid key and invalid petID gets %#v, expecting %#v", err.Error(), expectedErr.Error())
	}

	// valid api key with valid pet == pass
	_, err = petServer.DeletePet(context.WithValue(ctx, contextAPIKey, testAPIKey), &pb.PetID{PetId: 1})
	if err != nil {
		t.Errorf("Delete pet with valid key and valid petID gets error %#v, expecting nil", err.Error())
	}

	// check that the pet is indeed deleted
	pet, err := db.GetPetByID(ctx, 1)
	if pet != nil {
		t.Errorf("Pet 1 still exists after being deleted %#v", pet)
	}
	expectedErr = errors.New("404:Pet 1 not found")
	if err != nil && err.Error() != expectedErr.Error() {
		t.Errorf("Deleting pet from DB gets unexpected error %#v, expecting %#v", err.Error(), expectedErr.Error())
	}
}

func TestFindPetsByStatus(t *testing.T) {
	t.Log(("Testing FindPetByStatus"))

	db := testdb.New()
	petServer := NewPetstoreServer(logrus.New(), db, testRPCPort, testRestPort, testAPIKey)
	ctx := context.Background()

	// set 2 to pending and 4 to sold
	db.UpdatePetWithForm(ctx, 2, "", "pending")
	db.UpdatePetWithForm(ctx, 4, "", "sold")

	// Get all available - there should be 3
	testStatus := "available"
	avail, err := petServer.FindPetsByStatus(ctx, &pb.StatusReq{Status: []string{testStatus}})
	if err != nil {
		t.Errorf("Unexpected error getting available %#v", err.Error())
	}
	if len(avail.Pets) != 3 {
		t.Errorf("Got %d avail, expecting 3", len(avail.Pets))
	}
	for _, v := range avail.Pets {
		if v.Status != testStatus {
			t.Errorf("Pet %d has status %v, expecting %s", v.Id, v.Status, testStatus)
		}
	}

	// Get all the pending - there should be 1
	testStatus = "pending"
	pending, err := petServer.FindPetsByStatus(ctx, &pb.StatusReq{Status: []string{testStatus}})
	if err != nil {
		t.Errorf("Unexpected error getting pending %#v", err.Error())
	}
	if len(pending.Pets) != 1 {
		t.Errorf("Got %d pending, expecting 1", len(pending.Pets))
	}
	for _, v := range pending.Pets {
		if v.Status != testStatus {
			t.Errorf("Pet %d has status %v, expecting %s", v.Id, v.Status, testStatus)
		}
	}

	// Get all the sold - there should be 1
	testStatus = "sold"
	sold, err := petServer.FindPetsByStatus(ctx, &pb.StatusReq{Status: []string{testStatus}})
	if err != nil {
		t.Errorf("Unexpected error getting sold %#v", err.Error())
	}
	if len(sold.Pets) != 1 {
		t.Errorf("Got %d sold, expecting 1", len(sold.Pets))
	}
	for _, v := range sold.Pets {
		if v.Status != testStatus {
			t.Errorf("Pet %d has status %v, expecting %s", v.Id, v.Status, testStatus)
		}
	}

	// Get all that are in "pending" or "sold" - there should be 2
	multi, err := petServer.FindPetsByStatus(ctx, &pb.StatusReq{Status: []string{"pending", "sold"}})
	if err != nil {
		t.Errorf("Unexpected error getting pending,sold %#v", err.Error())
	}
	if len(multi.Pets) != 2 {
		t.Errorf("Got %d multi, expecting 2", len(multi.Pets))
	}
	for _, v := range multi.Pets {
		if v.Status != "pending" && v.Status != "sold" {
			t.Errorf("Pet %d has status %v, expecting pending or sold", v.Id, v.Status)
		}
	}

	// test invalid status - should throw an error and return a nil pets struct
	invalid, err := petServer.FindPetsByStatus(ctx, &pb.StatusReq{Status: []string{"pending", "INVALID"}})
	if err == nil {
		t.Error("Expected error getting pending,invalid")
	}
	if invalid != nil {
		t.Errorf("Got %d non-nil invalid results, expecting empty/nil pets", len(invalid.Pets))
	}
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
