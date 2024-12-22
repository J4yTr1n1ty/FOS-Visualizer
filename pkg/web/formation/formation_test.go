package formation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFormation(t *testing.T) {
	handler := NewHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.GetFormation))
	defer server.Close()

	data := DescribeFormationRequest{
		FormationString: "@A,F,F,FA:8#VH;NF",
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("GET", server.URL+"/formation/describe", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
