package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/mukulmantosh/go-ecommerce-app/internal/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var userJSON = `{
  "username": "mukulbaba123",
  "password": "mukul123",
  "first_name": "mukul",
  "last_name": "mantosh"
}
`

func TestAddUser(t *testing.T) {
	testDB, _ := setup()

	t.Run("should return 200 status ok", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		var service = server.NewServer(testDB)
		// Assertions
		if assert.NoError(t, service.AddUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}

		teardown(testDB)
	})

}
