package handler

import (
	"context"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	pb "github.com/steveoc64/petstore/proto"
)

const (
	testRestPort = 8082
	testRpcPort  = 8083
	testAPIKey   = "ABC123"
)

// TestGetPetByID provides full test coverage for the GetPetByID call
func TestGetPetByID(t *testing.T) {
	t.Log("Testing GetPetByID")

	petServer := NewPetstoreServer(logrus.New(), testRpcPort, testRestPort, testAPIKey)

	for i := range []int{1, 2, 3, 4, 5} {
		v := int64(i)
		pet, err := petServer.GetPetByID(context.TODO(), &pb.GetByIDRequest{
			Id: v,
		})
		if err != nil {
			t.Error("GetPetByID 1 unexpected error", err.Error())
		}
		if pet.Id != v {
			t.Errorf("GetPetByID got ID %v, expected %v", pet.Id, v)
		}
		expectedName := fmt.Sprintf("Pet number %d", v)
		if pet.Name != expectedName {
			t.Errorf("GetPetByID got Name %v, expected %v", pet.Name, expectedName)
		}

		t.Logf("Got Pet %#v", pet)
	}
}
