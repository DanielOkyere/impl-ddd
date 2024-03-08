package customer

import (
	"errors"

	"github.com/danielokyere/ddd-go/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("Customer not found")
	ErrFailedToAddCustomer = errors.New("Failed to add the customer")
	ErrUpdateFailed        = errors.New("Failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}


