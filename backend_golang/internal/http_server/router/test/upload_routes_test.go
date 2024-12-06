package router_test

import (
	"backend-golang/internal/http_server/router"
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type MockBroker struct{}

func (m *MockBroker) SendTask(queue string, body []byte) error {
	// Имитация отправки задачи
	return nil
}

type MockResultSaver struct{}

func (m *MockResultSaver) InsertStatus(processingID, status string) error {
	// Имитация сохранения статуса в базу данных
	return nil
}

func TestRegisterUploadRoutes(t *testing.T) {
	// Инициализация Gin
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	route := r.Group("/")

	// Логгер
	logger, _ := zap.NewProduction()

	// Создание моков
	mockBroker := &MockBroker{}
	mockResultSaver := &MockResultSaver{}

	// Регистрируем маршруты
	router.RegisterUploadRoutes(route, logger, mockBroker, mockResultSaver)

	// Создаем тестовый запрос с файлом
	// Создаем тело запроса с файлом
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "testfile.txt")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	part.Write([]byte("test content"))
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, "/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Отправляем запрос
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	// Проверяем статус
	assert.Equal(t, http.StatusOK, resp.Code)

	// Проверяем ответ
	var response map[string]string
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Проверяем, что в ответе есть processing_id
	assert.NotEmpty(t, response["processing_id"], "Processing ID should not be empty")
}
