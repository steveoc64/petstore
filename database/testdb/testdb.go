package testdb

import (
	"fmt"

	"github.com/steveoc64/petstore/database/memory"
	pb "github.com/steveoc64/petstore/proto"
)

// New creates a new memory DB and seeds it with test data
func New() *memory.DB {
	db := memory.New()

	// Seed the DB with pets 1-5
	for _, v := range []int{1, 2, 3, 4, 5} {
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
