package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockOrderRepo struct {
	saveFunc func(Order) error
}

func (m *mockOrderRepo) Save(order Order) error {
	return m.saveFunc(order)
}

func TestOrder(t *testing.T) {
	t.Run("case success", func(t *testing.T) {
		repo := &mockOrderRepo{
			saveFunc: func(o Order) error {
				return nil
			},
		}
		service := NewOrderService(repo)
		err := service.CreateOrder(Order{Total: 200})
		assert.NoError(t, err)
	})

	t.Run("Fail Case total < 0", func(t *testing.T) {
		repo := &mockOrderRepo{
			saveFunc: func(o Order) error {
				return nil
			},
		}
		service := NewOrderService(repo)
		err := service.CreateOrder(Order{Total: -200})
		assert.Error(t, err)
		assert.Equal(t, "Total have to be positive", err.Error())
	})

	t.Run("repo error", func(t *testing.T) {
		repo := &mockOrderRepo{
			saveFunc: func(o Order) error {
				return errors.New("db error")
			},
		}
		service := NewOrderService(repo)
		err := service.CreateOrder(Order{Total: 200})
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
	})

}
