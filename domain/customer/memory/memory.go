package memory

import (
	"fmt"
	"sync"

	"github.com/danielokyere/ddd-go/aggregate"
	"github.com/danielokyere/ddd-go/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (r *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := r.customers[id]; ok {
		return customer, nil
    }
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (r *MemoryRepository) Add(customer aggregate.Customer) error {
	if r.customers == nil {
		r.Lock()
		r.customers = make(map[uuid.UUID]aggregate.Customer)
		r.Unlock()
	}
	if _, ok := r.customers[customer.GetID()]; ok {
		return fmt.Errorf("customer with id %s already exists", customer.GetID())
	}
	r.Lock()
	r.customers[customer.GetID()] = customer
	r.Unlock()
	return nil
}

func (r *MemoryRepository) Update(customer aggregate.Customer) error {
	if _, ok := r.customers[customer.GetID()];!ok {
		return fmt.Errorf("customer with id %s not found", customer.GetID())
	}
	r.Lock()
	r.customers[customer.GetID()] = customer
	r.Unlock()
	return nil
}
