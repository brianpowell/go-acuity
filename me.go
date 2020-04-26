/* me.go
// Acuity API interface
*/

package acuity

// GetMe - pull info
func (a *Acuity) GetMe() (Me, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "me",
	}
	// Get the details
	out := Me{}
	err := a.request(req, &out)
	return out, err
}
