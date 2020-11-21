package api_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/testProj/api"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStatus(t *testing.T) {
	t.Run("key is empty", func(t *testing.T) {
		a := newApi()

		a.r.GET("//status").SetDebug(true).
			Run(a.e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, r.Code)
			})
	})

	t.Run("succeeded", func(t *testing.T) {
		a := newApi()

		a.m.On("GetURLStatusStatus", mock.Anything).Return(api.Res{
			Key: "123",
		}, nil)

		a.r.GET("/media/123/status").SetDebug(true).
			Run(a.e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)

				assertID(t, "123", r.Body, "key")
			})
	})
}

func newApi() apiMocks {
	r := gofight.New()

	e := echo.New()
	m := &apimocks.Manager{}

	api.Assemble(e, m)

	return apiMocks{
		r: r,
		e: e,
		m: m,
	}
}

func assertID(t assert.TestingT, expected string, body *bytes.Buffer, keys ...string) {
	data := body.Bytes()
	id, _ := jsonparser.GetString(data, keys...)
	assert.Equal(t, expected, id)
}
