package akamai

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"time"
)

type SensorInput struct {
	Abck      string `json:"abck"`
	Bmsz      string `json:"bmsz"`
	Version   string `json:"version"`
	PageUrl   string `json:"pageUrl"`
	UserAgent string `json:"userAgent"`

	ScriptHash string `json:"scriptHash"`
}

// GenerateSensorData returns the sensor data required to generate valid akamai cookies using the Hyper Solutions API.
func (s *Session) GenerateSensorData(ctx context.Context, input *SensorInput) (string, error) {
	const sensorEndpoint = "https://akm.justhyped.dev/sensor"
	return s.sendRequest(ctx, sensorEndpoint, input)
}

type PixelInput struct {
	UserAgent string `json:"userAgent"`
	HTMLVar   string `json:"htmlVar"`
	ScriptVar string `json:"scriptVar"`
}

// GeneratePixelData returns the pixel data using the Hyper Solutions API.
func (s *Session) GeneratePixelData(ctx context.Context, input *PixelInput) (string, error) {
	const pixelEndpoint = "https://akm.justhyped.dev/pixel"
	return s.sendRequest(ctx, pixelEndpoint, input)
}

func (s *Session) sendRequest(ctx context.Context, url string, input any) (string, error) {
	if s.apiKey == "" {
		return "", errors.New("missing api key")
	}

	payload, err := jsoniter.Marshal(input)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept-encoding", "gzip")
	req.Header.Set("x-api-key", s.apiKey)

	if s.jwtKey != "" {
		signature, err := s.generateSignature()
		if err != nil {
			return "", err
		}
		req.Header.Set("x-signature", signature)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response struct {
		Payload string `json:"payload"`
		Error   string `json:"error"`
	}
	if err := jsoniter.Unmarshal(respBody, &response); err != nil {
		return "", err
	}

	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}

func (s *Session) generateSignature() (string, error) {
	claims := jwt.MapClaims{
		"key": s.apiKey,
		"exp": time.Now().Add(time.Second * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
