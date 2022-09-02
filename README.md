# DisGOAuth ![Stars](https://img.shields.io/github/stars/realTristan/DisGOAuth?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/realTristan/DisGOAuth?label=Watchers)
![banner](https://user-images.githubusercontent.com/75189508/188035581-008c64d0-d59a-4a95-9e75-55cb3d8f4e79.png)

# About
- DisGOAuth is a light-weight, fast and easy-to-use module that makes using Discord's OAuth2.0 much easier. DisGOAuth uses solely native golang packages which makes it fast and secure.

# Installation
`go get -u github.com/realTristan/DisGOAuth`

# Quick Usage
```go

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

	// Home Page Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Send the user to the discord authentication
		// website. This is where they authorize access.
		dc.RedirectHandler(w, r, "")
	})

	// The OAuth URL Redirect Uri
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
```

# License
MIT License

Copyright (c) 2022 Tristan Simpson

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.