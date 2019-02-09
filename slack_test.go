package changer_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bookun/slack-status-changer"
	"github.com/pkg/errors"
)

func TestSlack_SendStatus(t *testing.T) {
	cases := []struct {
		name   string
		ts     *httptest.Server
		expect error
	}{
		{"slack api returned ok", genTestServer("ok"), nil},
		{"slack api returned false", genTestServer("error"), errors.New("some error")},
	}

	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			config := changer.Config{
				Endpoint: c.ts.URL,
			}
			tokens := []string{"token1"}
			slackClient := changer.NewSlackClient(config, tokens)
			actual := slackClient.SendStatus("home")
			if actual != nil {
				if c.expect.Error() != actual.Error() {
					t.Errorf("expected: %v, actual is %v", c.expect, actual)
				}
			}
		})
		c.ts.Close()
	}
}

func genTestServer(s string) *httptest.Server {
	if s == "error" {
		return httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"ok": false, "error": "some error"}`)
				return
			},
		))
	}
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, `{"ok": true}`)
			return
		},
	))
}
