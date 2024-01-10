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
package cart

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
	"github.com/mukulmantosh/go-ecommerce-app/internal/server"
	"github.com/mukulmantosh/go-ecommerce-app/internal/tests"
	"github.com/mukulmantosh/go-ecommerce-app/internal/utils"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAddItemsToCart(t *testing.T) {
	testDB, _ := tests.Setup()
	type FakeUser struct {
		Username  string `json:"username" faker:"username"`
		Password  string `json:"password" faker:"word,unique"`
		FirstName string `json:"first_name" faker:"first_name"`
		LastName  string `json:"last_name" faker:"last_name"`
	}

	type FakeCategory struct {
		Name        string `json:"name" faker:"word,unique"`
		Description string `json:"description" faker:"word,unique"`
	}

	type FakeProduct struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		CategoryID  string  `json:"category_id"`
	}

	type CategoryDataObj struct {
		CreatedAt   string           `json:"CreatedAt"`
		UpdatedAt   string           `json:"UpdatedAt"`
		DeletedAt   string           `json:"DeletedAt"`
		ID          string           `json:"ID"`
		Name        string           `json:"name"`
		Description string           `json:"description"`
		Product     []models.Product `json:"product"`
	}

	type LoginResponse struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}

	type ProductDataForCart struct {
		ProductID string `json:"productID"`
	}

	t.Run("add items to cart", func(t *testing.T) {
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

		var customCategory FakeCategory
		_ = faker.FakeData(&customCategory)
		out, _ = json.Marshal(&customCategory)
		req = httptest.NewRequest(http.MethodPost, "/category", strings.NewReader(string(out)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		if assert.NoError(t, service.AddCategory(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}

		var newCategoryResp CategoryDataObj
		data, _ := io.ReadAll(rec.Body)
		err := json.Unmarshal(data, &newCategoryResp)
		if err != nil {
			t.Error(err)
		}
		categoryId := newCategoryResp.ID

		addProduct := FakeProduct{
			Name:        "IPhone 15",
			Description: "IPhone 15, the latest available phone from Apple",
			Price:       3500.00,
			CategoryID:  categoryId,
		}

		AddProductJSON, err := utils.StructToJSON(addProduct)
		if err != nil {
			t.Error(err.Error())
		}

		newReq := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(AddProductJSON))
		newReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newRec := httptest.NewRecorder()
		getContext := e.NewContext(newReq, newRec)
		if assert.NoError(t, service.AddProduct(getContext)) {
			assert.Equal(t, http.StatusCreated, newRec.Code)
		}

		type ProductInfo struct {
			CreatedAt   string  `json:"CreatedAt"`
			UpdatedAt   string  `json:"UpdatedAt"`
			DeletedAt   string  `json:"DeletedAt"`
			ID          string  `json:"ID"`
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
			CategoryID  string  `json:"category_id"`
		}

		var newProductResp ProductInfo
		data, _ = io.ReadAll(newRec.Body)
		err = json.Unmarshal(data, &newProductResp)
		if err != nil {
			t.Error(err)
		}
		productId := newProductResp.ID

		//User Login
		type UserLogin struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		userLoginInfo := UserLogin{Username: customUser.Username, Password: customUser.Password}
		userLoginJSON, err := utils.StructToJSON(userLoginInfo)
		if err != nil {
			t.Error(err.Error())
		}
		newReq = httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(userLoginJSON))
		newReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newRec = httptest.NewRecorder()
		loginContext := e.NewContext(newReq, newRec)
		// Assertions
		if assert.NoError(t, service.UserLogin(loginContext)) {
			assert.Equal(t, http.StatusOK, newRec.Code)
		}

		ProductJSON, err := utils.StructToJSON(ProductDataForCart{ProductID: productId})
		if err != nil {
			t.Error(err.Error())
		}

		var LoginResp LoginResponse
		data, _ = io.ReadAll(newRec.Body)
		err = json.Unmarshal(data, &LoginResp)
		if err != nil {
			t.Error(err)
		}
		AuthToken := LoginResp.Token
		// Parse Token
		parsedToken, err := jwt.Parse(AuthToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		newCartReq := httptest.NewRequest(http.MethodPost, "/cart/", strings.NewReader(ProductJSON))
		newCartReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newCartRec := httptest.NewRecorder()
		cartContext := e.NewContext(newCartReq, newCartRec)
		cartContext.Set("user", parsedToken)
		// Assertions
		if assert.NoError(t, service.AddItemToCart(cartContext)) {
			assert.Equal(t, http.StatusCreated, newCartRec.Code)
			assert.Equal(t, "{\"message\":\"Item Added to Cart!\"}\n", newCartRec.Body.String())

		}

		defer tests.Teardown(testDB)
	})

}
