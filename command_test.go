package changer_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bookun/slack-status-changer"
	"github.com/bookun/slack-status-changer/changerFakes"
)

func TestCommand_Run(t *testing.T) {
	cases := []struct {
		name   string
		ts     *httptest.Server
		expect string
	}{
		{"inOffice", genTestHTTPServer(http.StatusOK), "office"},
		{"inHome", genTestHTTPServer(http.StatusForbidden), "house"},
	}

	t.Helper()
	for _, c := range cases {
		slackClient := new(changerFakes.FakeSlackClient)
		command := &changer.Command{
			InternalURL: c.ts.URL,
			SlackClient: slackClient,
		}
		t.Run(c.name, func(t *testing.T) {
			if err := command.Run(); err != nil {
				t.Error("error should not occur")
			}
			actual := slackClient.SendStatusArgsForCall(0)
			if c.expect != actual {
				t.Errorf("SendStatus's arg expect %s, but actual is %s", c.expect, actual)
			}
		})
	}
}

func genTestHTTPServer(statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusCode)
			return
		},
	))
}
