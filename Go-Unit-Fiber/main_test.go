package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserRoute(t *testing.T) {
	app := setup()

	testCases := []struct {
		description  string
		body         User
		expectStatus int
	}{
		{
			description:  "Valid input",
			body:         User{"palm.tiger@example.com", "Tiger Palm", 30},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Invalid email",
			body:         User{"invalid-email", "Tiger Palm", 30},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid fullname",
			body:         User{"palm.tiger@example.com", "Palm12345", 30},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid age",
			body:         User{"palm.tigere@example.com", "Tiger Palm", -5},
			expectStatus: fiber.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)

			assert.Equal(t, tc.expectStatus, res.StatusCode)
		})
	}
}
