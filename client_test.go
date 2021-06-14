package gopagseguro

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	timeout := time.Second
	tid := "CHAR_878B0CA5-DE08-4CCC-855D-938DC19C0A50"
	successBody := fmt.Sprintf(`{"id": "%s"}`, tid)
	errorBody := `{"error_messages":[{"code":"00000","message":"error."}]}`

	t.Run("Create charge success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, successBody)
		}))
		defer ts.Close()

		client := NewClient(ts.URL, "token", 0, timeout, timeout, timeout)
		charge := Charge{}

		result, err := client.Charge(&charge)

		assert.Equal(t, tid, result.ID)
		assert.Nil(t, err)
	})

	t.Run("Create charge invalid response status code", func(t *testing.T) {
		expectedError := fmt.Errorf("invalid response status code: 400. %s", errorBody)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, errorBody)
		}))
		defer ts.Close()

		client := NewClient(ts.URL, "token", 0, timeout, timeout, timeout)
		charge := Charge{}

		result, err := client.Charge(&charge)

		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Capture charge success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, successBody)
		}))
		defer ts.Close()

		client := NewClient(ts.URL, "token", 0, timeout, timeout, timeout)
		charge := Charge{}

		result, err := client.Capture(tid, &charge)

		assert.Equal(t, tid, result.ID)
		assert.Nil(t, err)
	})

	t.Run("Capture charge invalid response status code", func(t *testing.T) {
		expectedError := fmt.Errorf("invalid response status code: 400. %s", errorBody)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, errorBody)
		}))
		defer ts.Close()

		client := NewClient(ts.URL, "token", 0, timeout, timeout, timeout)
		charge := Charge{}

		result, err := client.Capture(tid, &charge)

		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Cancel charge success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, successBody)
		}))
		defer ts.Close()

		client := NewClient(ts.URL, "token", 0, timeout, timeout, timeout)
		charge := Charge{}

		result, err := client.Cancel(tid, &charge)

		assert.Equal(t, tid, result.ID)
		assert.Nil(t, err)
	})

	t.Run("Cancel charge invalid response status code", func(t *testing.T) {
		expectedError := fmt.Errorf("invalid response status code: 400. %s", errorBody)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, errorBody)
		}))
		defer ts.Close()

		client := NewClient(ts.URL, "token", 0, timeout, timeout, timeout)
		charge := Charge{}

		result, err := client.Cancel(tid, &charge)

		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})
}
