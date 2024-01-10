/*
	Copyright 2022 Google LLC

#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
*/
package users

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/server"
	"github.com/mukulmantosh/go-ecommerce-app/internal/tests"
	"github.com/mukulmantosh/go-ecommerce-app/internal/utils"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateUserAddress(t *testing.T) {
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

	type UserAddressDataObj struct {
		UserAddressID string `json:"userAddressId"`
		Address       string `json:"address"`
		PostalCode    string `json:"postal_code"`
		Country       string `json:"country"`
		Mobile        string `json:"mobile"`
		UserId        string `json:"user_id"`
	}

	t.Run("update user address by id", func(t *testing.T) {
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
			t.Error(err)
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
		userAddressJSON, err := utils.StructToJSON(userAddr)
		if err != nil {
			t.Error(err.Error())
		}

		newReq := httptest.NewRequest(http.MethodPost, "/user/address", strings.NewReader(userAddressJSON))
		newReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newRec := httptest.NewRecorder()
		getContext := e.NewContext(newReq, newRec)

		if assert.NoError(t, service.AddUserAddress(getContext)) {
			assert.Equal(t, http.StatusCreated, newRec.Code)
		}

		var userAddressResp UserAddressDataObj
		dataAddr, _ := io.ReadAll(newRec.Body)
		err = json.Unmarshal(dataAddr, &userAddressResp)
		userAddressId := userAddressResp.UserAddressID

		UpdateUserAddr := FakeUserAddress{
			Address:    "Libellengasse 15",
			City:       "Sankt Ulrich Am Waasen",
			PostalCode: "8072",
			Country:    "Austria",
			Mobile:     "06888738564",
			UserId:     userId,
		}

		// Serialize the struct to a JSON string
		UpdateUserAddressJSON, err := utils.StructToJSON(UpdateUserAddr)
		if err != nil {
			t.Error(err.Error())
		}

		newAddrReq := httptest.NewRequest(http.MethodPut, "/user/address",
			strings.NewReader(UpdateUserAddressJSON))
		newAddrReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newAddrRec := httptest.NewRecorder()
		UpdateAddrContext := e.NewContext(newAddrReq, newAddrRec)
		UpdateAddrContext.SetParamNames("id")
		UpdateAddrContext.SetParamValues(userAddressId)

		if assert.NoError(t, service.UpdateUserAddress(UpdateAddrContext)) {
			assert.Equal(t, http.StatusOK, newAddrRec.Code)
			assert.Equal(t, "\"User Information Updated!\"\n", newAddrRec.Body.String())
		}

		tests.Teardown(testDB)
	})

}
