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
		//
		// The third parameter in the RedirectHandler is the
		// state. If you're storing a state, PLEASE base64 encode
		// it beforehand!
		dc.RedirectHandler(w, r, "") // w: http.ResponseWriter, r: *http.Request, state: string
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
		// Define Variables
		var (
			// Get the code from the redirect parameters (&code=...)
			codeFromURLParamaters = r.URL.Query()["code"][0]

			// Get the access token using the above codeFromURLParamaters
			accessToken, _ = dc.GetOnlyAccessToken(codeFromURLParamaters)

			// Get the authorized user's data using the above accessToken
			userData, _ = discord.GetUserData(accessToken)
		)
		// Print the user data map
		fmt.Println(userData)
	})

	// Listen and Serve to the incoming http requests
	http.ListenAndServe(":8000", nil)
}
