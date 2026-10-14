package disgoauth

import (
	"net/http"
	"testing"
)

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestAccessTokenRequest_Error(t *testing.T) {
	oldClient := RequestClient
	defer func() { RequestClient = oldClient }()

	RequestClient = &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			return nil, http.ErrHandlerTimeout
		}),
	}

	client := Init(&Client{
		ClientID:     "id",
		ClientSecret: "secret",
		RedirectURI:  "http://localhost",
		Scopes:       []string{"identify"},
	})
	_, _, _, err := client.GetAccessToken("some_code")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetUserData_Error(t *testing.T) {
	oldClient := RequestClient
	defer func() { RequestClient = oldClient }()

	RequestClient = &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			return nil, http.ErrHandlerTimeout
		}),
	}

	_, err := GetUserData("some_token")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
