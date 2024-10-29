package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидалось получить статус 200!")
	assert.NotEmpty(t, responseRecorder.Body, "Ожидалось непустое тело ответа!")
}

func TestMainHandlerWhenCityIsNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=testcity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидалось получить статус 400!")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Ожидалось получить сообщение: wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидалось получить статус 200!")
	assert.Len(t, list, totalCount, "Ожидалось получить все кафе!")
}
