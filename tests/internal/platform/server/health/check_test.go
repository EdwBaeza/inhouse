package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EdwBaeza/inhouse/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestHelthCheckHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", health.CheckHandler())

	t.Run("returns 200", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		if err != nil {
			t.Error(err)
		}
		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
