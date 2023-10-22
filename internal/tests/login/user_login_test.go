package login

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/server"
	"github.com/mukulmantosh/go-ecommerce-app/internal/tests"
	"github.com/mukulmantosh/go-ecommerce-app/internal/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserLogin(t *testing.T) {
	testDB, _ := tests.Setup()
	type FakeUser struct {
		Username  string `json:"username" faker:"username"`
		Password  string `json:"password" faker:"word,unique"`
		FirstName string `json:"first_name" faker:"first_name"`
		LastName  string `json:"last_name" faker:"last_name"`
	}
	type UserLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	t.Run("should return JWT token for a valid user", func(t *testing.T) {

		var customUser FakeUser
		_ = faker.FakeData(&customUser)
		out, _ := json.Marshal(&customUser)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(out)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		var service = server.NewServer(testDB)
		// Assertions
		if assert.NoError(t, service.AddUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}

		userLoginInfo := UserLogin{Username: customUser.Username, Password: customUser.Password}
		userLoginJSON, err := utils.StructToJSON(userLoginInfo)
		if err != nil {
			t.Error(err.Error())
		}
		newReq := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(userLoginJSON))
		newReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newRec := httptest.NewRecorder()
		loginContext := e.NewContext(newReq, newRec)
		// Assertions
		if assert.NoError(t, service.UserLogin(loginContext)) {
			assert.Equal(t, http.StatusOK, newRec.Code)
		}

		tests.Teardown(testDB)
	})

}
