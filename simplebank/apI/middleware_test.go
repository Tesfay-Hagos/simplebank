package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"tesfayprep/token"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderkey, authorizationHeader)

}
func TestAuthMiddleware(t *testing.T) {
	testcases := []struct {
		name string

		setupAuth     func(t *testing.T, request *http.Request, tokenmaker token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{

		{
			name: "ok",
			setupAuth: func(t *testing.T, request *http.Request, tokenmaker token.Maker) {
				addAuthorization(t, request, tokenmaker, authorizationTypeBearer, "user", time.Minute)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},

		{
			name: "noAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenmaker token.Maker) {

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)

			},
		},

		{
			name: "unSupportedAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenmaker token.Maker) {
				addAuthorization(t, request, tokenmaker, "Unsupported", "user", time.Minute)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)

			},
		},

		{
			name: "invalidAuthorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request, tokenmaker token.Maker) {
				addAuthorization(t, request, tokenmaker, "", "user", time.Minute)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)

			},
		},

		{
			name: "ExpiredToken",
			setupAuth: func(t *testing.T, request *http.Request, tokenmaker token.Maker) {
				addAuthorization(t, request, tokenmaker, authorizationTypeBearer, "user", -time.Minute)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)

			},
		},
	}

	for i := range testcases {
		tc := testcases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)
			authpath := "/auth"
			server.router.GET(
				authpath,
				authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
					//ctx.Json(http.StatusOK, gin.H{})

				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authpath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})

	}
}
