package integration

import (
	"HomeService/model"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type flatsList struct {
	Flats []model.Flat `json:"flats"`
}

func TestGetFlats(t *testing.T) {
	setup()
	defer teardown()

	house := model.House{
		Id:        1,
		Address:   "123 Test St",
		Year:      2024,
		Developer: "Test Developer",
	}

	houseBody, err := json.Marshal(house)
	if err != nil {
		t.Fatalf("failed to marshal house request body: %v", err)
	}

	moderatorToken, err := generateToken("moderator", user_id)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	clientToken, err := generateToken("client", user_id)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	req, err := authRequest("POST", fmt.Sprintf("%s/house/create", baseURL), houseBody, moderatorToken)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send POST request: %v", err)
	}

	// добавляем квартиры к дому, но они добавляются со статусом created
	flats := []model.Flat{
		{HouseId: 1, Id: 1, Rooms: 2, Price: 100000},
		{HouseId: 1, Id: 2, Rooms: 3, Price: 120000},
		{HouseId: 1, Id: 3, Rooms: 2, Price: 120000},
	}

	for _, flat := range flats {
		flatBody, err := json.Marshal(flat)
		if err != nil {
			t.Fatalf("failed to marshal flat request body: %v", err)
		}

		req, err := authRequest("POST", fmt.Sprintf("%s/flat/create", baseURL), flatBody, moderatorToken)

		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		_, err = client.Do(req)
		if err != nil {
			t.Fatalf("failed to send POST request: %v", err)
		}
	}

	req, err = authRequest("GET", fmt.Sprintf("%s/house/%d", baseURL, 1), nil, moderatorToken)
	if err != nil {
		t.Fatalf("failed to create GET request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseMod map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseMod); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	flatsMod, ok := responseMod["flats"].([]interface{})
	assert.True(t, ok, "Expected response to contain 'flats' field for moderator")
	assert.Equal(t, len(flats), len(flatsMod), "Moderator should see all flats")

	// проверяем получение квартир как клиент (должен увидеть только квартиры со статусом approved)

	// у одной квартиры меняем статус на approved
	flatApproved := model.Flat{Id: 1, HouseId: 1, Status: "approved"}
	flatApprovedBody, err := json.Marshal(flatApproved)
	if err != nil {
		t.Fatalf("failed to marshal flat request body: %v", err)
	}
	req, err = authRequest("POST", fmt.Sprintf("%s/flat/update", baseURL), flatApprovedBody, moderatorToken)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	_, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send POST request: %v", err)
	}

	req, err = authRequest("GET", fmt.Sprintf("%s/house/%d", baseURL, house.Id), nil, clientToken)
	if err != nil {
		t.Fatalf("failed to create GET request: %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseClient map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseClient); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	flatsClient, _ := responseClient["flats"].([]interface{})

	// клиент должен видеть только approved квартиры
	expectedClientFlats := []model.Flat{}
	assert.Equal(t, len(expectedClientFlats), len(flatsClient), "Client should see only approved flats")
}
