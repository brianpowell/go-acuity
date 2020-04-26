/* calendars.go
// Acuity API interface
*/

package acuity

// GetCalendars - pull list
func (a *Acuity) GetCalendars() ([]Calendar, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "calendars",
	}
	// Get the list of Calendars
	out := []Calendar{}
	err := a.request(req, &out)
	return out, err
}
