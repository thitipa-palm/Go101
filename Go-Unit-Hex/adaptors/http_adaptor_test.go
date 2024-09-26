package adaptors

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thitipa-palm/go-Unit-Hex/core"
)

type mockOrderService struct {
	mock.Mock
}

func (m *mockOrderService) CreateOrder(order core.Order) error {
	arg := m.Called(order)
	return arg.Error(0)
}

func TestOrderHandler(t *testing.T) {
	mockService := new(mockOrderService)
	handle := NewHttpOrderHandler(mockService)

	app := SetUp(handle)
	t.Run("success", func(t *testing.T) {
		mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(nil)
		reqBody, _ := json.Marshal(core.Order{Total: 150})
		req := httptest.NewRequest("POST", "/order", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, res.StatusCode)

		mockService.AssertExpectations(t)
	})

	// t.Run("fail due to total < 0", func(t *testing.T) {
	// 	mockService.ExpectedCalls = nil //clear expected call ของ mockService เพื่อ ไม่ให้เกิด cache ของ function ก่อนหน้านี้
	// 	mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("Total have to be positive"))

	// 	reqBody, _ := json.Marshal(core.Order{Total: -150})
	// 	req := httptest.NewRequest("POST", "/order", bytes.NewReader(reqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	res, err := app.Test(req)

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)

	// 	//incase want to also test response
	// 	// Read the response body
	// 	body, err := io.ReadAll(res.Body)
	// 	if err != nil {
	// 		t.Fatalf("Failed to read response body: %v", err)
	// 	}
	// 	defer res.Body.Close()

	// 	// Convert the response body to a map
	// 	var responseMap map[string]interface{}
	// 	if err := json.Unmarshal(body, &responseMap); err != nil {
	// 		t.Fatalf("Failed to parse JSON response: %v", err)
	// 	}

	// 	// Extract specific keys
	// 	errMsg := responseMap["error"].(string)

	// 	assert.Equal(t, "Total have to be positive", errMsg)

	// 	mockService.AssertExpectations(t)
	// })

	t.Run("fail due to invalid body", func(t *testing.T) {

		req := httptest.NewRequest("POST", "/order", bytes.NewBufferString(`{"total": "invalid"}`))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)

	})

	t.Run("fail service just error", func(t *testing.T) {
		mockService.ExpectedCalls = nil //clear expected call ของ mockService เพื่อ ไม่ให้เกิด cache ของ function ก่อนหน้านี้
		mockService.On("CreateOrder", mock.AnythingOfType("core.Order")).Return(errors.New("service error"))

		reqBody, _ := json.Marshal(core.Order{Total: 150})
		req := httptest.NewRequest("POST", "/order", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)

		mockService.AssertExpectations(t)
	})
}
