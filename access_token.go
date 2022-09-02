package DisGOAuth

// Import Packages
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// The GetAccessTokenBody() function is used to return
// the request body bytes being used in the
// GetAccessToken() http request
func (dc *DiscordClient) GetAccessTokenBody(code string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s&scope=identify",
		dc.ClientID,
		dc.ClientSecret,
		dc.RedirectURI,
		code,
	)))
}

// The GetAccessToken() function is used to send an api
// request to discord's oauth2/token endpoint.
// The function returns the token required for
// accessing the authorized users data
func (dc *DiscordClient) GetAccessToken(code string) string {
	// Establish a new request object
	var req, _ = http.NewRequest("POST",
		"https://discordapp.com/api/oauth2/token",
		dc.GetAccessTokenBody(code), // access_token.go (this file)
	)
	// Set the request object's headers
	req.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"Accept":       []string{"application/json"},
	}
	// Send the http request using the above request object
	// then decode the response body and marshal it to
	// a readable golang map. Return the "access_token" value
	// from said golang map
	var (
		// Send the http request
		res, _ = RequestClient.Do(req)
		// Readable golang map used for storing
		// the response body
		data map[string]interface{}
	)
	// Decode the response body and return
	// the access token
	json.NewDecoder(res.Body).Decode(&data)

	// Return the token type and the access token
	// combined together
	return data["token_type"].(string) + " " + data["access_token"].(string)
}
