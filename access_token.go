package DisGOAuth

// Import Packages
import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// The AccessTokenBody() function is used to return
// the request body bytes being used in the
// GetAccessToken() http request
func (dc *DiscordClient) AccessTokenBody(code string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s&scope=identify",
		dc.ClientID, dc.ClientSecret, dc.RedirectURI, code,
	)))
}

// The RefreshAccessTokenBody() function is used to return
// the request body bytes being used in the
// RefreshAccessToken() http request
func (dc *DiscordClient) RefreshAccessTokenBody(refreshToken string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=refresh_token&redirect_uri=%s&refresh_token=%s",
		dc.ClientID, dc.ClientSecret, dc.RefreshRedirectURI, refreshToken,
	)))
}

// The CredentialsAccessTokenBody() function is used to return
// the request body bytes being used in the
// RefreshAccessToken() http request
func CredentialsAccessTokenBody(scopes []string) *bytes.Buffer {
	var _url string = "grant_type=client_credentials&scope="
	// For each of the scopes
	for i := 0; i < len(scopes); i++ {
		// Append the scope to the url
		_url += scopes[i]

		// If there are multiple scopes and the
		// current index isn't the last scope
		if i != len(scopes) {
			// Append %20 (space)
			_url += "%20"
		}
	}
	// Return the url bytes
	return bytes.NewBuffer([]byte(_url))
}

// The _GetAccessToken() function is used to send an api
// request to discord's oauth2/token endpoint.
// The function returns the data required for
// accessing the authorized users data
func (dc *DiscordClient) _GetAccessToken(body *bytes.Buffer, creds bool) map[string]interface{} {
	// Establish a new request object
	var req, _ = http.NewRequest("POST",
		"https://discordapp.com/api/oauth2/token",
		body, // access_token.go (this file)
	)
	// Set the request object's headers
	req.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"Accept":       []string{"application/json"},
	}

	// If using the credentials access token endpoint
	if creds {
		// Base64 encode the client id and secret
		var auth string = base64.StdEncoding.EncodeToString([]byte(dc.ClientID + ":" + dc.ClientSecret))
		// Set the authorization header
		req.Header["Authorization"] = []string{"Basic " + auth}
	}

	// Send the http request using the above request object
	// then decode the response body and marshal it to
	// a readable golang map.
	var (
		// Send the http request
		res, _ = RequestClient.Do(req)
		// Readable golang map used for storing
		// the response body
		data map[string]interface{}
	)
	// Decode the response body and return
	// the data map
	json.NewDecoder(res.Body).Decode(&data)
	return data
}

/////////////////////////////////////////
// Get Access Token
/////////////////////////////////////////

// The GetAccessToken() function is used to get the users
// bearer authorization token.
func (dc *DiscordClient) GetAccessToken(code string) string {
	// Send http request to get token data
	var data map[string]interface{} = dc._GetAccessToken(dc.AccessTokenBody(code), false)
	// Return the bearer token from said data
	return data["token_type"].(string) + " " + data["access_token"].(string)
}

/////////////////////////////////////////
// Get Access Token + Data
/////////////////////////////////////////

// The GetAccessTokenData() function is used to return all
// the map data revolving around getting the access token
func (dc *DiscordClient) GetAccessTokenData(code string) map[string]interface{} {
	return dc._GetAccessToken(dc.AccessTokenBody(code), false)
}

/////////////////////////////////////////
// Refresh Access Token
/////////////////////////////////////////

// The RefreshAccessToken() function is used to refresh
// the users bearer authorization token.
func (dc *DiscordClient) RefreshAccessToken(refreshToken string) map[string]interface{} {
	return dc._GetAccessToken(dc.RefreshAccessTokenBody(refreshToken), false)
}

/////////////////////////////////////////
// Get Credentials Access Token
/////////////////////////////////////////

// The GetCredentialsAccessToken() function is used to get
// the users credentials access bearer auth token.
func (dc *DiscordClient) GetCredentialsAccessToken(scopes []string) string {
	// Send http request to get token data
	var data map[string]interface{} = dc._GetAccessToken(CredentialsAccessTokenBody(scopes), true)
	// Return the bearer token from said data
	return data["token_type"].(string) + " " + data["access_token"].(string)
}

/////////////////////////////////////////
// Get Credentials Access Token + Data
/////////////////////////////////////////

// The GetCredentialsAccessTokenData() function
// is used to return all the map data revolving
// around getting the credentials access token
func (dc *DiscordClient) GetCredentialsAccessTokenData(scopes []string) map[string]interface{} {
	return dc._GetAccessToken(CredentialsAccessTokenBody(scopes), true)
}
