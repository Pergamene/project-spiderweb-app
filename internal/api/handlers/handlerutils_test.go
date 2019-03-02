package handlerutils

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Pergamene/project-spiderweb-service/internal/api"
)

func TestAPI(t *testing.T) {
	cases := []struct {
		name                 string
		method               string
		endpoint             string
		headers              map[string]string
		authN                api.AuthN
		authZ                api.AuthZ
		datacenter           string
		expectedResponseBody string
		expectedStatusCode   int
	}{
		{
			name:     "bad method call",
			method:   http.MethodPost,
			endpoint: "healthcheck",
			authN: api.AuthN{
				Datacenter:      "LOCAL",
				AdminAuthSecret: "SECRET",
			},
			authZ: api.AuthZ{
				APIPath: "api/test",
			},
			expectedResponseBody: "{\"meta\":{\"httpStatus\":\"405 - Method Not Allowed\",\"message\":\"method not allowed\"}}\n",
			expectedStatusCode:   405,
		},
		{
			name:     "bad endpoint call",
			method:   http.MethodGet,
			endpoint: "doesnotexist",
			authN: api.AuthN{
				Datacenter:      "LOCAL",
				AdminAuthSecret: "SECRET",
			},
			authZ: api.AuthZ{
				APIPath: "api/test",
			},
			expectedResponseBody: "{\"meta\":{\"httpStatus\":\"404 - Not Found\",\"message\":\"not found\"}}\n",
			expectedStatusCode:   404,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf(tc.name), func(t *testing.T) {
			resp, respBody := HandleTestRequest(HandleTestRequestParams{
				Method:         tc.method,
				Endpoint:       tc.endpoint,
				Headers:        tc.headers,
				Body:           nil,
				RouterHandlers: GetBaseRouterHandlers(),
				AuthZ:          tc.authZ,
				AuthN:          tc.authN,
			})
			require.Equal(t, tc.expectedResponseBody, respBody)
			require.Equal(t, tc.expectedStatusCode, resp.StatusCode)
		})
	}
}