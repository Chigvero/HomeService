package integration

import (
	"HomeService/internal/service"
	"bytes"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const baseURL = "http://localhost:8081"

var (
	jwtToken string
	user_id  uuid.UUID = uuid.New()
)

func setup() {
	dockerComposePath := "../test_integration/docker-compose.yml"

	if _, err := os.Stat(dockerComposePath); os.IsNotExist(err) {
		log.Fatalf("docker-compose.yml file not found: %v", err)
	}

	// Запускаем Docker Compose через внешний скрипт
	if err := runScript("start_docker.sh"); err != nil {
		log.Fatalf("failed to start docker-compose: %v", err)
	}

	time.Sleep(10 * time.Second)
	log.Println("Starting...")

	token, err := generateToken("moderator", user_id)
	if err != nil {
		log.Fatalf("failed to get JWT token: %v", err)
	}
	jwtToken = token
}

func teardown() {
	if err := runScript("stop_docker.sh"); err != nil {
		log.Fatalf("failed to stop docker-compose: %v", err)
	}
}

func runScript(scriptName string) error {
	cmd := exec.Command("/bin/bash", scriptName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func authRequest(method, url string, body []byte, token string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func generateToken(user_type string, userId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &service.CustomClaim{
		UserType: user_type,
		UserId:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte("myKey"))
}
