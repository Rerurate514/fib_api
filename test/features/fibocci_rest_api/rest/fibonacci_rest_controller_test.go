package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	fibonacci_rest_api "example.com/mod/lib/features/fibonacci_rest_api/rest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFibonacci struct {
	mock.Mock
}

func (m *MockFibonacci) GetNumberAt(index int) (uint64, error) {
	args := m.Called(index)
	return args.Get(0).(uint64), args.Error(1)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/fib", fibonacci_rest_api.GetNumberAt)
	return r
}

func TestGetNumberAt_ValidIndex(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/fib?n=10", nil)
	w := httptest.NewRecorder()

	//send request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, float64(55), response["result"])
}

func TestGetNumberAt_InvalidIndex_Negative(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/fib?n=-1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Contains(t, response["message"], "インデックス -1 は範囲外です（0-255）")
}

func TestGetNumberAt_NonNumericIndex(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/fib?n=abc", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "無効なインデックスです", response["message"])
}

func TestGetNumberAt_LargeIndex(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/fib?n=100", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "result")
	} else {
		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}

func TestGetNumberAt_MissingQueryParam(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/fib", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "クエリパラメータが指定されていません", response["message"])
}
