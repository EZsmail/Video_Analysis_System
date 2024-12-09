package router_test

import (
	"backend-golang/internal/http_server/router"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type MockStatusGetter struct {
	Status string
	Err    error
}

func (m *MockStatusGetter) GetStatus(processingID string) (string, error) {
	return m.Status, m.Err
}

type MockResultGetter struct {
	Result [][]string
	Err    error
}

func (m *MockResultGetter) GetResult(ctx context.Context, processingID string) ([][]string, error) {
	return m.Result, m.Err
}

func TestRegisterResultRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	logger, _ := zap.NewDevelopment()
	statusMock := &MockStatusGetter{}
	resultMock := &MockResultGetter{}

	r := gin.Default()
	router.RegisterResultRoutes(r, logger, statusMock, resultMock)

	t.Run("Result not ready", func(t *testing.T) {
		statusMock.Status = "in_progress"
		statusMock.Err = nil

		req, _ := http.NewRequest(http.MethodGet, "/result/test-id", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusAccepted, resp.Code)
		assert.JSONEq(t, `{"message": "Result is not ready", "processing_id": "test-id", "status": "in_progress"}`, resp.Body.String())
	})

	t.Run("Result ready", func(t *testing.T) {
		statusMock.Status = "completed"
		statusMock.Err = nil
		resultMock.Result = [][]string{{"Part 1", "00:00", "00:30"}, {"Part 1", "00:00", "00:30"}}
		resultMock.Err = nil

		req, _ := http.NewRequest(http.MethodGet, "/result/test-id", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"processing_id": "test-id", "result": {"field1": "value1", "field2": "value2"}}`, resp.Body.String())
	})

	t.Run("Status retrieval error", func(t *testing.T) {
		statusMock.Err = errors.New("database error")

		req, _ := http.NewRequest(http.MethodGet, "/result/test-id", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.JSONEq(t, `{"error": "Failed to retrieve status"}`, resp.Body.String())
	})

	t.Run("Result retrieval error", func(t *testing.T) {
		statusMock.Status = "completed"
		statusMock.Err = nil
		resultMock.Err = errors.New("database error")

		req, _ := http.NewRequest(http.MethodGet, "/result/test-id", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.JSONEq(t, `{"error": "Failed to retrieve result"}`, resp.Body.String())
	})
}
