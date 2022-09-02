package DisGOAuth

// Import Packages
import (
	"encoding/json"
	"net/http"
)

// The GetUserData() function is used to send an api
// request to the discord/users/@me endpoint with
// the provided accessToken.
func GetUserData(token string) map[string]interface{} {
	// Establish a new request object
	var req, _ = http.NewRequest("GET", "https://discord.com/api/users/@me", nil)

	// Set the request object's headers
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
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
	// Decode the response body as a golang map
	json.NewDecoder(res.Body).Decode(&data)
	return data
}
