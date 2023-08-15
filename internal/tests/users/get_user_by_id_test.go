package users

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/server"
	"github.com/mukulmantosh/go-ecommerce-app/internal/tests"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetUserById(t *testing.T) {
	testDB, _ := tests.Setup()
	type FakeUser struct {
		Username  string `json:"username" faker:"username"`
		Password  string `json:"password" faker:"word,unique"`
		FirstName string `json:"first_name" faker:"first_name"`
		LastName  string `json:"last_name" faker:"last_name"`
	}

	type DataObj struct {
		UserID  string `json:"user_id"`
		Message string `json:"message"`
	}
	type UserResponse struct {
		Data DataObj `json:"data"`
	}

	t.Run("should return user by id", func(t *testing.T) {
		var service = server.NewServer(testDB)

		var customUser FakeUser
		_ = faker.FakeData(&customUser)
		out, _ := json.Marshal(&customUser)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(out)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, service.AddUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}

		var userResp UserResponse
		data, _ := io.ReadAll(rec.Body)
		err := json.Unmarshal(data, &userResp)
		if err != nil {
			return
		}
		userId := userResp.Data.UserID

		newReq := httptest.NewRequest(http.MethodGet, "/user/"+userId, nil)
		newReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newRec := httptest.NewRecorder()
		getContext := e.NewContext(newReq, newRec)

		// Assertions
		if assert.NoError(t, service.GetUserById(getContext)) {
			assert.Equal(t, http.StatusOK, newRec.Code)
		}
		tests.Teardown(testDB)
	})

}
