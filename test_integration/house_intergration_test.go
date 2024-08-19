package integration

import (
	"HomeService/model"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateHouse(t *testing.T) {
	setup()
	defer teardown()

	client := &http.Client{}

	moderatorToken, err := generateToken("moderator", user_id)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	clientToken, err := generateToken("client", user_id)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	// case 1: успешное создание дома модератором
	house := model.House{
		Id:        1,
		Address:   "123 Test St",
		Year:      2024,
		Developer: "Test Developer",
	}

	body, err := json.Marshal(house)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req, err := authRequest("POST", fmt.Sprintf("%s/house/create", baseURL), body, moderatorToken)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to send POST request: %v", err)
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var createdHouse model.House
	if err := json.NewDecoder(resp.Body).Decode(&createdHouse); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	assert.Equal(t, house.Id, createdHouse.Id)
	assert.Equal(t, house.Address, createdHouse.Address)
	assert.Equal(t, house.Year, createdHouse.Year)
	assert.Equal(t, house.Developer, createdHouse.Developer)

	// case 2: попытка создания дома клиентом (ошибка)
	req, err = authRequest("POST", fmt.Sprintf("%s/house/create", baseURL), body, clientToken)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)

	// case 3: создание дома с некорректными данными (недостаточно информации)
	invalidHouse := model.House{
		Id:        2,
		Address:   "",
		Year:      0,
		Developer: "",
	}

	body, err = json.Marshal(invalidHouse)
	if err != nil {
		t.Fatalf("failed to marshal invalid request body: %v", err)
	}

	req, err = authRequest("POST", fmt.Sprintf("%s/house/create", baseURL), body, moderatorToken)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// case 4: попытка повторного создания дома с тем же ID (ошибка)
	req, err = authRequest("POST", fmt.Sprintf("%s/house/create", baseURL), body, moderatorToken)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode) // либо 409
}
