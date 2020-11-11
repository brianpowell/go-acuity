/* availability.go
// Acuity API interface
*/

package acuity

// GetAvailabilityDates - pull list
func (a *Acuity) GetAvailabilityDates(query map[string]interface{}) ([]DateTime, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "availability/dates",
		Query:  query,
	}
	// Get the list of DateTimes
	out := []DateTime{}
	err := a.request(req, &out)
	return out, err
}

// GetAvailabilityTimes - pull list
func (a *Acuity) GetAvailabilityTimes(query map[string]interface{}) ([]DateTime, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "availability/times",
		Query:  query,
	}
	// Get the list of DateTimes
	out := []DateTime{}
	err := a.request(req, &out)
	return out, err
}

// GetAvailabilityClasses - pull list
func (a *Acuity) GetAvailabilityClasses(query map[string]interface{}) ([]Class, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "availability/classes",
		Query:  query,
	}
	// Get the list of Classes
	out := []Class{}
	err := a.request(req, &out)
	return out, err
}

// AvailablityCheckTime - check single appt time
func (a *Acuity) AvailablityCheckTime(body AvailabilityCheckTimeBody) (AvailabilityCheckTime, error) {
	// Build the Request
	req := Request{
		Method: "POST",
		URL:    "availability/check-times",
		Body:   body,
	}
	// Check the Appointment Time
	out := AvailabilityCheckTime{}
	err := a.request(req, &out)
	return out, err
}

// AvailablityCheckTimes - check the list of appt
func (a *Acuity) AvailablityCheckTimes(body []AvailabilityCheckTimeBody) ([]AvailabilityCheckTime, error) {
	// Build the Request
	req := Request{
		Method: "POST",
		URL:    "availability/check-times",
		Body:   body,
	}
	// Check the Appointments Time
	out := []AvailabilityCheckTime{}
	err := a.request(req, &out)
	return out, err
}
