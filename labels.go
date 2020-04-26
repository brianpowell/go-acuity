/* labels.go
// Acuity API interface
*/

package acuity

// GetLabels - pull list
func (a *Acuity) GetLabels() ([]Label, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "labels",
	}
	// Get the list of Labels
	out := []Label{}
	err := a.request(req, &out)
	return out, err
}
