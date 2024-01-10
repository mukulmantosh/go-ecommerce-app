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
package productCategory

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/models"
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

func TestUpdateProductById(t *testing.T) {
	testDB, _ := tests.Setup()
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

	t.Run("update product by id", func(t *testing.T) {

		var customCategory FakeCategory
		_ = faker.FakeData(&customCategory)
		out, _ := json.Marshal(&customCategory)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/category", strings.NewReader(string(out)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		var service = server.NewServer(testDB)

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

		type UpdateProduct struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}
		updateProductInfo := UpdateProduct{Name: "IPHONE 15", Description: "Newphoria. Dynamic Island. " +
			"48MP Main camera with 2x Telephoto"}

		updateProductJSON, err := utils.StructToJSON(updateProductInfo)
		if err != nil {
			t.Error(err.Error())
		}
		newReq = httptest.NewRequest(http.MethodPut, "/products", strings.NewReader(updateProductJSON))
		newReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newRec = httptest.NewRecorder()
		getContext = e.NewContext(newReq, newRec)
		getContext.SetParamNames("id")
		getContext.SetParamValues(productId)
		if assert.NoError(t, service.UpdateProduct(getContext)) {
			assert.Equal(t, http.StatusOK, newRec.Code)
			assert.Equal(t, "{\"message\":\"Updated!\"}\n", newRec.Body.String())
		}
		tests.Teardown(testDB)
	})

}
