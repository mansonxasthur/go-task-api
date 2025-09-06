package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
	tests := []struct {
		name       string
		data       interface{}
		status     int
		wantStatus int
		wantBody   ResponseBody
	}{
		{
			name:       "valid string data",
			data:       "success",
			status:     http.StatusOK,
			wantStatus: http.StatusOK,
			wantBody:   ResponseBody{Data: "success"},
		},
		{
			name:       "valid integer data",
			data:       123,
			status:     http.StatusOK,
			wantStatus: http.StatusOK,
			wantBody:   ResponseBody{Data: 123},
		},
		{
			name:       "valid map data",
			data:       map[string]interface{}{"key": "value"},
			status:     http.StatusCreated,
			wantStatus: http.StatusCreated,
			wantBody:   ResponseBody{Data: map[string]interface{}{"key": "value"}},
		},
		{
			name:       "valid array data",
			data:       []string{"item1", "item2"},
			status:     http.StatusAccepted,
			wantStatus: http.StatusAccepted,
			wantBody:   ResponseBody{Data: []string{"item1", "item2"}},
		},
		{
			name:       "nil data",
			data:       nil,
			status:     http.StatusOK,
			wantStatus: http.StatusOK,
			wantBody:   ResponseBody{Data: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			SuccessResponse(recorder, tt.data, tt.status)

			res := recorder.Result()

			if res.StatusCode != tt.wantStatus {
				t.Errorf("unexpected status code: got %d, want %d", res.StatusCode, tt.wantStatus)
			}

			var gotBody ResponseBody
			err := json.NewDecoder(res.Body).Decode(&gotBody)
			if err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if !bytes.Equal(marshalBody(gotBody), marshalBody(tt.wantBody)) {
				t.Errorf("unexpected body: got %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func ExampleSuccessResponse() {
	recorder := httptest.NewRecorder()
	data := map[string]interface{}{
		"message": "Hello, World!",
		"status":  "success",
	}

	SuccessResponse(recorder, data, http.StatusOK)

	// Output:
}

func BenchmarkSuccessResponse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recorder := httptest.NewRecorder()
		data := map[string]interface{}{
			"message": "Hello, World!",
			"status":  "success",
		}

		SuccessResponse(recorder, data, http.StatusOK)
	}
}

func marshalBody(body ResponseBody) []byte {
	data, _ := json.Marshal(body)
	return data
}
