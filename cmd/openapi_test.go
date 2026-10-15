package cmd_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/graingo/maltose-quickstart/cmd"
	"github.com/graingo/maltose/frame/m"
	"github.com/stretchr/testify/require"
)

func TestEmbeddedOpenAPIContract(t *testing.T) {
	m.Server().SetConfigWithMap(map[string]any{"openapi_path": "/api.json", "swagger_path": "/swagger"})
	server := cmd.HTTPServer()
	require.NoError(t, server.Prepare(context.Background()))
	require.NoError(t, server.ValidateOpenAPIRoutes())
	response := httptest.NewRecorder()
	server.ServeHTTP(response, httptest.NewRequest("GET", "/api.json", nil))
	require.Equal(t, 200, response.Code)
	require.Contains(t, response.Body.String(), "openapi: 3.1.0")

	response = httptest.NewRecorder()
	server.ServeHTTP(response, httptest.NewRequest("GET", "/api/v1/hello?name=Maltose", nil))
	require.Equal(t, 200, response.Code)
	require.JSONEq(t, `{"code":0,"message":"OK","data":{"name":"Hello, Maltose!","tags":[]}}`, response.Body.String())
}
