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

func TestUpdateCategoryById(t *testing.T) {
	testDB, _ := tests.Setup()
	type FakeCategory struct {
		Name        string `json:"name" faker:"word,unique"`
		Description string `json:"description" faker:"word,unique"`
	}

	type UpdateCategory struct {
		Name        string `json:"name"`
		Description string `json:"description"`
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

	t.Run("update category by id", func(t *testing.T) {

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

		updateCategory := UpdateCategory{
			Name:        "Electronic Equipments",
			Description: "SmartPhones",
		}

		updateCategoryJSON, err := utils.StructToJSON(updateCategory)
		if err != nil {
			t.Error(err.Error())
		}

		newReq := httptest.NewRequest(http.MethodPut, "/category", strings.NewReader(updateCategoryJSON))
		newReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		newRec := httptest.NewRecorder()
		getContext := e.NewContext(newReq, newRec)
		getContext.SetParamNames("id")
		getContext.SetParamValues(categoryId)
		if assert.NoError(t, service.UpdateCategory(getContext)) {
			assert.Equal(t, http.StatusOK, newRec.Code)
			assert.Equal(t, "{\"message\":\"Category Information Updated!\"}\n", newRec.Body.String())
		}
		tests.Teardown(testDB)
	})
}
