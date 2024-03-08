package memory

import (
	"testing"

	"github.com/danielokyere/ddd-go/aggregate"
	"github.com/danielokyere/ddd-go/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("John")
	if err != nil {
		t.Fatalf("Error creating customer: %s", err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
        },
	}

	testCases := []testCase{
		{
			name: "No customer found",
			id: uuid.MustParse("f876e3ab-55b3-40ff-9702-74880fbd841c"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "Customer found",
            id:          id,
            expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err!= tc.expectedErr {
				t.Fatalf("Expected error %s, got %s", tc.expectedErr, err)
			}
		})
	}
    
}
