package DisGOAuth

// Import fmt package
import "fmt"

// The ImplicitOAuth() function uses the implicit
// and less-safe response type for getting the
// users access token
func (dc *DiscordClient) ImplicitOAuth() {
	// Set the OAuth URL to a formatted string
	// that contains the client id, redirect uri,
	// and response type.
	dc.OAuthURL = fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=token",
		dc.ClientID,
		dc.RedirectURI,
	)
}

// The NonImplicitOAuth() function uses the default and
// safer response type for getting the users access token
func (dc *DiscordClient) NonImplicitOAuth() {
	// Establish the prompt parameter
	var prompt string = dc.Prompt
	if len(dc.Prompt) < 1 {
		prompt = "none"
	}

	// Set the OAuth URL to a formatted string
	// that contains the client id, redirect uri,
	// and response type.
	dc.OAuthURL = fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&prompt=%s",
		dc.ClientID,
		dc.RedirectURI,
		prompt,
	)
}
