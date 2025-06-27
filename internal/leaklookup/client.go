package leaklookup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type APIResponse struct {
	Error   string          `json:"error"`
	Message json.RawMessage `json:"message"`
}

func CheckEmail(email string) (map[string][]interface{}, error) {
	key := os.Getenv("LEAKLOOKUP_KEY")
	if key == "" {
		return nil, fmt.Errorf("LEAKLOOKUP_KEY not set")
	}

	body := fmt.Sprintf("key=%s&type=email_address&query=%s", key, email)
	resp, err := http.Post(
		"https://leak-lookup.com/api/search",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	if r.Error == "true" {
		var errorMsg string
		if err := json.Unmarshal(r.Message, &errorMsg); err != nil {
			errorMsg = "Unknown API error"
		}
		return nil, fmt.Errorf("Leak-Lookup API error: %s", errorMsg)
	}

	var leaks map[string][]interface{}
	if err := json.Unmarshal(r.Message, &leaks); err != nil {
		return nil, fmt.Errorf("не удалось распарсить данные об утечках: %v", err)
	}

	return leaks, nil
}
