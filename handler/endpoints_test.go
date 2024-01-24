package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/prapsky/sawitpro/common/validator"
	"github.com/prapsky/sawitpro/handler"
	"github.com/prapsky/sawitpro/service"
)

type Executor struct {
	handler *handler.Server
}

type RequestContext struct {
	context  echo.Context
	recorder *httptest.ResponseRecorder
}

func TestRegistration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("bad request", func(t *testing.T) {
		request := handler.RegistrationRequest{}
		reqByte, _ := json.Marshal(&request)

		url := "/registration"
		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqByte))
		rec := httptest.NewRecorder()

		e := echo.New()
		e.Validator = validator.NewFormValidator()

		ctx := e.NewContext(req, rec)

		ret := &RequestContext{ctx, rec}
		exec := buildExecutor(ctrl)
		err := exec.handler.Register(ret.context)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Nil(t, err)
	})
}

func buildExecutor(ctrl *gomock.Controller) *Executor {
	s := service.NewMockService(ctrl)
	h := handler.NewServer(s)

	return &Executor{
		handler: h,
	}
}
