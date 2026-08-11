package cmd_test

import (
	"context"
	"testing"

	"github.com/graingo/maltose-quickstart/api/hello/v1"
	"github.com/graingo/maltose-quickstart/cmd"
	"github.com/graingo/maltose-quickstart/internal/controller/hello"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPServerRegistersHelloLogic(t *testing.T) {
	server := cmd.HTTPServer()
	require.NotNil(t, server)

	response, err := hello.NewV1().Hello(context.Background(), &v1.HelloReq{Name: "Maltose"})
	require.NoError(t, err)
	assert.Equal(t, "Hello, Maltose!", response.Name)
}
