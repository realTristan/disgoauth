package main

// Import Packages
import (
	"fmt"
	"net/http"

	// Import DisGOAuth
	discord "github.com/realTristan/DisGOAuth"
)

// Main function
func main() {
	// Establish a new discord client
	var dc *discord.DiscordClient = discord.Init(&discord.DiscordClient{
		ClientID:     "CLIENT ID",
		ClientSecret: "CLIENT SECRET",
		RedirectURI:  "localhost:8000/redirect",
		Scopes:       []string{discord.ScopeIdentify},
	})

	////////////////////////////////////////////////////////////////////////
	//
	// Home Page Handler
	//
	// It is suggested to put this in it's own function,
	// I only did it like for the showcase.
	//
	////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Send the user to the discord authentication
		// website. This is where they authorize access.
		dc.RedirectHandler(w, r, "")
	})

	////////////////////////////////////////////////////////////////////////
	//
	// The OAuth URL Redirect Uri
	//
	// It is suggested to put this in it's own function,
	// I only did it like for the showcase.
	//
	////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		// Put this in the handler of the dc.RedirectURI
		// Define Variables
		var (
			// Get the code from the redirect parameters (&code=...)
			codeFromURLParamaters = r.URL.Query()["code"][0]

			// Get the access token using the above codeFromURLParamaters
			accessToken string = dc.GetAccessToken(codeFromURLParamaters)

			// Get the authorized user's data using the above accessToken
			userData map[string]interface{} = discord.GetUserData(accessToken)
		)
		// Print the user data map
		fmt.Println(userData)
	})

	// Listen and Serve to the incoming http requests
	http.ListenAndServe(":8000", nil)
}
