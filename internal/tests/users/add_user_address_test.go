package users

import (
	"encoding/json"
	"fmt"
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

func structToJSONString(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func TestAddUserAddress(t *testing.T) {
	testDB, _ := tests.Setup()
	type FakeUser struct {
		Username  string `json:"username" faker:"username"`
		Password  string `json:"password" faker:"word,unique"`
		FirstName string `json:"first_name" faker:"first_name"`
		LastName  string `json:"last_name" faker:"last_name"`
	}

	type FakeUserAddress struct {
		Address    string `json:"address"`
		City       string `json:"city"`
		PostalCode string `json:"postal_code"`
		Country    string `json:"country"`
		Mobile     string `json:"mobile"`
		UserId     string `json:"user_id"`
	}

	type DataObj struct {
		UserID  string `json:"user_id"`
		Message string `json:"message"`
	}
	type UserResponse struct {
		Data DataObj `json:"data"`
	}

	t.Run("create new user address", func(t *testing.T) {
		var service = server.NewServer(testDB)

		var customUser FakeUser
		_ = faker.FakeData(&customUser)
		out, _ := json.Marshal(&customUser)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(out)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

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

		userAddr := FakeUserAddress{
			Address:    "132, My Street, Kingston, New York",
			City:       "New York",
			PostalCode: "12401",
			Country:    "USA",
			Mobile:     "9898989898",
			UserId:     userId,
		}

		// Serialize the struct to a JSON string
		userAddressJSON, err := structToJSONString(userAddr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		newReq := httptest.NewRequest(http.MethodPost, "/user/address", strings.NewReader(userAddressJSON))
		newReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newRec := httptest.NewRecorder()
		getContext := e.NewContext(newReq, newRec)

		if assert.NoError(t, service.AddUserAddress(getContext)) {
			assert.Equal(t, http.StatusCreated, newRec.Code)
		}
		tests.Teardown(testDB)
	})

}
