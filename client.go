package disgoauth

// Import net/http
import (
	"net/http"
)

// Request Client for sending http requests
var RequestClient *http.Client = &http.Client{}

// The Client struct contains five primary keys
/* ClientID: 	string 	 { "Your Application's Client ID (REQUIRED)" } */
/* ClientSecret: 	string 	 { "Your Application's Client Secret (REQUIRED)" } */
/* Scopes: 	[]string { "Your Application's Permission Scopes (REQUIRED)" } */
/* Prompt: 	string 	 { "The Consent Prompt Parameter for Auth Reapproval" } */
/* RefreshRedirectURI: 	string	 { "The Redirect URI For Token Refreshing" } */
/* Implicit: 	bool 	 { "Whether to Use Discord's Implicit Endpoint For Getting Access Token" } */
/* RedirectURI: 	string 	 { "The Redirect URI (This is where you use the GetUserData functions) (REQUIRED)" } */
/* OAuthURL: 	string 	 { "Your Application's OAuth URL (If none is provided, one will be generated for you)" } */
type Client struct {
	ClientID           string
	ClientSecret       string
	RedirectURI        string
	RefreshRedirectURI string
	Scopes             []string
	OAuthURL           string
	Prompt             string
	Implicit           bool
}

// The checkStructErrors() function is used to check for
// any invalid / empty struct values that are required
// for the discord oauth to work.
func (dc *Client) checkStructErrors() {
	// Make sure the user has provided
	// a valid client id
	if len(dc.ClientID) < 1 {
		panic("DisGOAuth Error: invalid ClientID in Client (ClientID: string)")
	}
	// Make sure the user has provided
	// a valid client secret
	if len(dc.ClientSecret) < 1 {
		panic("DisGOAuth Error: invalid ClientSecret in Client (ClientSecret: string)")
	}
	// Make sure the user has provided
	// a valid redirect uri
	if len(dc.RedirectURI) < 1 {
		panic("DisGOAuth Error: invalid RedirectURI in Client (RedirectURI: string)")
	}
	// Make sure the user has provided
	// a sufficient number of scopes
	if len(dc.Scopes) < 1 {
		panic("DisGOAuth Error: not enough scopes in Client (Scopes: []string)")
	}
}

// The appendScopes() function is used to append
// the provided scopes to the OAuth URL. This function
// is called from the InitOAuthURL() function and is
// only ran if the number of provided scopes is valid.
//
// Using append() and a byte slice is much faster than
// using += to a string!
func (dc *Client) appendScopes(url []byte) string {
	// Append the initial parameter name (scope)
	url = append(url, "&scope="...)

	// For each of the discord client's scopes
	for i := 0; i < len(dc.Scopes); i++ {
		// Append the scope to the OAuth URL
		url = append(url, dc.Scopes[i]...)

		// If there are multiple scopes and the
		// current index isn't the last scope
		if i != len(dc.Scopes)-1 {
			// Append %20 (space)
			url = append(url, "%20"...)
		}
	}
	return string(url)
}

// The initOAuthURL() function is used to initialize
// a discord OAuth URL. This function is called from
// the Init() function and is only ran if there is
// no previously provided OAuth URL.
func (dc *Client) initOAuthURL() string {
	// Non Implicit OAuth
	var tempUrl string = dc.nonImplicitOAuth() // implicit.go

	// Implicit OAuth
	if dc.Implicit {
		tempUrl = dc.implicitOAuth() // implicit.go
	}
	// If user provided scopes
	if len(dc.Scopes) > 0 {
		// Append the scopes to the OAuth URL
		tempUrl = dc.appendScopes([]byte(tempUrl)) // discord_client.go (this file)
	}
	return tempUrl
}

// The Init() function is used to initalize
// the required data for the discord oauth to work
// It panics if required parameters are missing from
// the provided Client struct
func Init(dc *Client) *Client {
	// Check for Client struct errors
	dc.checkStructErrors() // discord_client.go (this file)

	// Initialize the OAuth URL
	if len(dc.OAuthURL) < 40 {
		dc.OAuthURL = dc.initOAuthURL() // discord_client.go (this file)
	}
	// Return the discord client
	return dc
}
