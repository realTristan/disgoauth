package disgoauth

// Import Packages
import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// The accessTokenBody() function is used to return
// the request body bytes being used in the
// GetAccessToken() function
func (dc *DiscordClient) accessTokenBody(code string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s&scope=identify",
		dc.ClientID, dc.ClientSecret, dc.RedirectURI, code,
	)))
}

// The refreshAccessTokenBody() function is used to return
// the request body bytes being used in the
// RefreshAccessToken() function
func (dc *DiscordClient) refreshAccessTokenBody(refreshToken string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=refresh_token&redirect_uri=%s&refresh_token=%s",
		dc.ClientID, dc.ClientSecret, dc.RefreshRedirectURI, refreshToken,
	)))
}

// The credentialsAccessTokenBody() function is used to return
// the request body bytes being used in the
// GetCredentialsAccessToken() function
func credentialsAccessTokenBody(scopes []string) *bytes.Buffer {
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

// The accessTokenRequestObject() function is used to establish
// a new request object that will be used for sending
// the api request to the discord oauth token endpoint.
func (dc *DiscordClient) accessTokenRequestObject(body *bytes.Buffer, creds bool) (*http.Request, error) {
	// Establish a new request object
	var req, err = http.NewRequest("POST",
		"https://discordapp.com/api/oauth2/token", body,
	)
	// Handle the error
	if err != nil {
		return req, err
	}

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
	return req, nil
}

// The accessTokenRequest() function is used to send an api
// request to discord's oauth2/token endpoint.
// The function returns the data required for
// accessing the authorized users data
func (dc *DiscordClient) accessTokenRequest(req *http.Request) (map[string]interface{}, error) {
	// Send the http request
	resp, err := RequestClient.Do(req)

	// Handle the error
	// If the response status isn't a success
	if resp.StatusCode != 200 || err != nil {
		// Read the http body
		body, _err := io.ReadAll(resp.Body)

		// Handle the read body error
		if _err != nil {
			return map[string]interface{}{}, _err
		}
		// Handle http response error
		return map[string]interface{}{},
			fmt.Errorf("status: %d, code: %v, body: %s",
				resp.StatusCode, err, string(body))
	}

	// Readable golang map used for storing
	// the response body
	var data map[string]interface{}

	// Decode the data and handle errors
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return map[string]interface{}{}, err
	}
	return data, nil
}

/////////////////////////////////////////
// Get Access Token
/////////////////////////////////////////

// The GetAccessToken() function is used to get the users
// bearer access token, refresh token, the token expiry
// and any errors that occured.
func (dc *DiscordClient) GetAccessToken(code string) (string, string, int, error) {
	// Define Variables
	var (
		// The Access Token Request Body
		tokenBody *bytes.Buffer = dc.accessTokenBody(code)
		// The Access Token Request Object
		tokenReq, err = dc.accessTokenRequestObject(tokenBody, false)
	)
	// Handle the token request object error
	if err != nil {
		return "", "", -1, err
	}
	// Get the token data map
	var data, _err = dc.accessTokenRequest(tokenReq)

	// Handle the token request error
	if _err != nil {
		return "", "", -1, _err
	}
	// The Bearer access token
	var accessToken string = data["token_type"].(string) + " " + data["access_token"].(string)

	// Return the bearer token from said data
	return accessToken, data["refresh_token"].(string), data["expires_in"].(int), nil
}

/////////////////////////////////////////
// Get Only Access Token
/////////////////////////////////////////

// The GetOnlyAccessToken() function is used to get
// the users bearer access token.
func (dc *DiscordClient) GetOnlyAccessToken(code string) (string, error) {
	// Define Variables
	var (
		// The Access Token Request Body
		tokenBody *bytes.Buffer = dc.accessTokenBody(code)
		// The Access Token Request Object
		tokenReq, err = dc.accessTokenRequestObject(tokenBody, false)
	)
	// Handle the token request object error
	if err != nil {
		return "", err
	}
	// Get the token data map
	var data, _err = dc.accessTokenRequest(tokenReq)

	// Handle the token request error
	if _err != nil {
		return "", _err
	}
	// The Bearer access token
	var accessToken string = data["token_type"].(string) + " " + data["access_token"].(string)

	// Return the bearer token from said data
	return accessToken, nil
}

/////////////////////////////////////////
// Get Access Token + Data
/////////////////////////////////////////

// The GetAccessTokenMap() function is used to return all
// the map data revolving around the access token
func (dc *DiscordClient) GetAccessTokenMap(code string) (map[string]interface{}, error) {
	// Define Variables
	var (
		// The Access Token Request Body
		tokenBody *bytes.Buffer = dc.accessTokenBody(code)
		// The Access Token Request Object
		tokenReq, err = dc.accessTokenRequestObject(tokenBody, false)
	)
	// Handle the token request object error
	if err != nil {
		return map[string]interface{}{}, err
	}
	return dc.accessTokenRequest(tokenReq)
}

/////////////////////////////////////////
// Refresh Access Token
/////////////////////////////////////////

// The RefreshAccessToken() function is used to refresh
// the users bearer authorization token.
func (dc *DiscordClient) RefreshAccessToken(refreshToken string) (map[string]interface{}, error) {
	// Define Variables
	var (
		// The Access Token Request Body
		tokenBody *bytes.Buffer = dc.refreshAccessTokenBody(refreshToken)
		// The Access Token Request Object
		tokenReq, err = dc.accessTokenRequestObject(tokenBody, false)
	)
	// Handle the token request object error
	if err != nil {
		return map[string]interface{}{}, err
	}
	return dc.accessTokenRequest(tokenReq)
}

/////////////////////////////////////////
// Get Credentials Access Token
/////////////////////////////////////////

// The GetCredentialsAccessToken() function is used to get
// the credentials auth token, refresh token, the token expiry
// timing, and any errors that had occured.
func (dc *DiscordClient) GetCredentialsAccessToken(scopes []string) (string, string, int, error) {
	// Define Variables
	var (
		// The Access Token Request Body
		tokenBody *bytes.Buffer = credentialsAccessTokenBody(scopes)
		// The Access Token Request Object
		tokenReq, err = dc.accessTokenRequestObject(tokenBody, true)
	)
	// Handle the error
	if err != nil {
		return "", "", -1, err
	}
	// Send http request to get token data
	var data, _err = dc.accessTokenRequest(tokenReq)

	// Handle the token request error
	if _err != nil {
		return "", "", -1, _err
	}
	// The Bearer access token
	var accessToken string = data["token_type"].(string) + " " + data["access_token"].(string)

	// Return the bearer token from said data
	return accessToken, data["refresh_token"].(string), data["expires_in"].(int), nil
}

/////////////////////////////////////////
// Get Only Credentials Access Token
/////////////////////////////////////////

// The GetOnlyCredentialsAccessToken() function is used to get
// the users credentials access bearer auth token.
func (dc *DiscordClient) GetOnlyCredentialsAccessToken(scopes []string) (string, error) {
	// Define Variables
	var (
		// The Access Token Request Body
		tokenBody *bytes.Buffer = credentialsAccessTokenBody(scopes)
		// The Access Token Request Object
		tokenReq, err = dc.accessTokenRequestObject(tokenBody, true)
	)
	// Handle the error
	if err != nil {
		return "", err
	}

	// Send http request to get token data
	var data, _err = dc.accessTokenRequest(tokenReq)

	// Handle the error
	if _err != nil {
		return "", _err
	}
	// The Bearer access token
	var accessToken string = data["token_type"].(string) + " " + data["access_token"].(string)

	// Return the bearer token from said data
	return accessToken, nil
}

/////////////////////////////////////////
// Get Credentials Access Token + Data
/////////////////////////////////////////

// The GetCredentialsAccessTokenMap() function
// is used to return all the map data revolving
// around the credentials access token
func (dc *DiscordClient) GetCredentialsAccessTokenMap(scopes []string) (map[string]interface{}, error) {
	// Define Variables
	var (
		// The Access Token Request Body
		tokenBody *bytes.Buffer = credentialsAccessTokenBody(scopes)
		// The Access Token Request Object
		tokenReq, err = dc.accessTokenRequestObject(tokenBody, true)
	)
	// Handle the token request object error
	if err != nil {
		return map[string]interface{}{}, err
	}
	return dc.accessTokenRequest(tokenReq)
}
