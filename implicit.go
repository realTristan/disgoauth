package disgoauth

// Import fmt package
import "fmt"

// The implicitOAuth() function uses the implicit
// and less-safe response type for getting the
// users access token
func (dc *DiscordClient) implicitOAuth() string {
	// Return the OAuth URL to a formatted string
	// that contains the client id, redirect uri,
	// and response type.
	return fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=token",
		dc.ClientID,
		dc.RedirectURI,
	)
}

// The nonImplicitOAuth() function uses the default and
// safer response type for getting the users access token
func (dc *DiscordClient) nonImplicitOAuth() string {
	// Establish the prompt parameter
	var prompt string = dc.Prompt
	if len(dc.Prompt) < 1 {
		prompt = "none"
	}

	// Return the OAuth URL to a formatted string
	// that contains the client id, redirect uri,
	// and response type.
	return fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&prompt=%s",
		dc.ClientID,
		dc.RedirectURI,
		prompt,
	)
}
