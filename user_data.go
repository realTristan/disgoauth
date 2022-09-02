package disgoauth

// Import Packages
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// The GetUserData() function is used to send an api
// request to the discord/users/@me endpoint with
// the provided accessToken.
func GetUserData(token string) (map[string]interface{}, error) {
	// Establish a new request object
	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)

	// Handle the error
	if err != nil {
		return map[string]interface{}{}, err
	}
	// Set the request object's headers
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{token},
	}
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

	// Handle the error
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return map[string]interface{}{}, err
	}
	return data, nil
}
