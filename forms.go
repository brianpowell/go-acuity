/* forms.go
// Acuity API interface
*/

package acuity

// GetForms - pull list
func (a *Acuity) GetForms() ([]Form, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "forms",
	}
	// Get the list of Forms
	out := []Form{}
	err := a.request(req, &out)
	return out, err
}
