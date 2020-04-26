/* clients.go
// Acuity API interface
*/

package acuity

// GetClients - pull list
func (a *Acuity) GetClients(query map[string]interface{}) ([]Client, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "clients",
		Query:  query,
	}
	// Get the list of Clients
	out := []Client{}
	err := a.request(req, &out)
	return out, err
}

// PostClient - create an appt
func (a *Acuity) PostClient(body *ClientBody) (Client, error) {
	// Build the Request
	req := Request{
		Method: "POST",
		URL:    "clients",
		Body:   body,
	}
	// Create the Client
	out := Client{}
	err := a.request(req, &out)
	return out, err
}

// PutClient - create an appt
func (a *Acuity) PutClient(body *ClientBody) (Client, error) {
	// Build the Request
	req := Request{
		Method: "POST",
		URL:    "clients",
		Body:   body,
	}
	// Create the Client
	out := Client{}
	err := a.request(req, &out)
	return out, err
}

// DeleteClient - get an appt
func (a *Acuity) DeleteClient(firstName, lastName string) error {
	// Build the Request
	req := Request{
		Method: "DELETE",
		URL:    "clients",
		Params: map[string]string{"firstName": firstName, "lastName": lastName},
	}
	// Delete the Client
	err := a.request(req, nil)
	return err
}
