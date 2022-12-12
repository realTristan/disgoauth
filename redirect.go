package disgoauth

// Import net/http package
import "net/http"

// The RedirectHandler() function is used to redirect the user
// to the provided Client OAuth URL. If there is a
// provided state, it will add it to said OAuth URL
//
// If using a state, base64encode the data beforehand, else,
// set the state to "" (length: 0)
func (dc *Client) RedirectHandler(w http.ResponseWriter, r *http.Request, state string) {
	// Create a copy of the OAuth URL
	var _url string = dc.OAuthURL

	// If the user provided a state. Make sure
	// the state is base64 encoded. Or else
	// many bugs can arise.
	if len(state) > 0 {
		_url = dc.OAuthURL + "&state=" + state
	}
	// Redirect the user to the OAuth URL
	http.Redirect(w, r, _url, http.StatusTemporaryRedirect)
}
