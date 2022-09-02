package DisGOAuth

// Import Packages
import (
	"fmt"
	"net/http"
)

// Request Client for sending http requests
var RequestClient *http.Client = &http.Client{}

// The DiscordClient struct contains five primary keys
/* ClientID: 	string 	 { "Your Application's Client ID" } */
/* ClientSecret: 	string 	 { "Your Application's Client Secret" } */
/* Scopes: 	[]string { "Your Application's Permission Scopes (REQUIRED)" } */
/* RedirectURI: 	string 	 { "The Redirect URI (This is where you use the GetUserData functions)" } */
/* OAuthURL: 	string 	 { "Your Application's OAuth URL (If none is provided, one will be generated for you)" } */
type DiscordClient struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []string
	OAuthURL     string
}

// The AppendScopes() function is used to append
// the provided scopes to the OAuth URL. This function
// is called from the InitOAuthURL() function and is
// only ran if the number of provided scopes is valid.
func (dc *DiscordClient) AppendScopes() {
	// Append the initial parameter name (scope)
	dc.OAuthURL += "&scope="

	// For each of the discord client's scopes
	for i := 0; i < len(dc.Scopes); i++ {
		// Append the scope to the OAuth URL
		dc.OAuthURL += dc.Scopes[i]

		// If there are multiple scopes and the
		// current index isn't the last scope
		if i != len(dc.Scopes) {
			// Append %20 (space)
			dc.OAuthURL += "%20"
		}
	}
}

// The InitOAuthURL() function is used to initialize
// a discord OAuth URL. This function is called from
// the Init() function and is only ran if there is
// no previously provided OAuth URL.
func (dc *DiscordClient) InitOAuthURL() {
	// Set the OAuth URL to a formatted string
	// that contains the client id, redirect uri,
	// and response type.
	dc.OAuthURL = fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&prompt=consent",
		dc.ClientID,
		dc.RedirectURI,
	)
	// Append the scopes to the OAuth URL
	dc.AppendScopes()
}

// The CheckStructErrors() function is used to check for
// any invalid / empty struct values that are required
// for the discord oauth to work.
func (dc *DiscordClient) CheckStructErrors() {
	// Make sure the user has provided
	// a valid client id
	if len(dc.ClientID) < 1 {
		panic("DisGOAuth Error: invalid ClientID in DiscordClient (ClientID: string)")
	}
	// Make sure the user has provided
	// a valid client secret
	if len(dc.ClientSecret) < 1 {
		panic("DisGOAuth Error: invalid ClientSecret in DiscordClient (ClientSecret: string)")
	}
	// Make sure the user has provided
	// a valid redirect uri
	if len(dc.RedirectURI) < 1 {
		panic("DisGOAuth Error: invalid RedirectURI in DiscordClient (RedirectURI: string)")
	}
	// Make sure the user has provided
	// a sufficient number of scopes
	if len(dc.Scopes) < 1 {
		panic("DisGOAuth Error: not enough scopes in DiscordClient (Scopes: []string)")
	}
}

// The Init() function is used to initalize
// the required data for the discord oauth to work
// It panics if required parameters are missing from
// the provided DiscordClient struct
func Init(dc *DiscordClient) *DiscordClient {
	// Check for DiscordClient struct errors
	dc.CheckStructErrors()

	// Initialize the OAuth URL
	if len(dc.OAuthURL) < 40 {
		dc.InitOAuthURL()
	}
	// Return the discord client
	return dc
}