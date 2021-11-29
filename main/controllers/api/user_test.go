package api

/*
import (
	"crud_echo/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	MockDB = map[string]*models.User{
		"hassan": &models.User{"hassan" , "helfi", nil,"09303639881", time.Now(), 0 , "123456789"},
	}
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)
func TestUser(t *testing.T) error {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost , "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Create) {

	}
}
*/