/* appointment.go
// Acuity API interface
*/

package acuity

import (
	"strconv"
	"time"
)

// GetAppointments - pull list
func (a *Acuity) GetAppointments(query map[string]interface{}) ([]Appointment, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "appointments",
		Query:  query,
	}
	// Get the list of Appointments
	out := []Appointment{}
	err := a.request(req, &out)
	return out, err
}

// PostAppointment - create an appt
func (a *Acuity) PostAppointment(body AppointmentBody) (Appointment, error) {
	// Build the Request
	req := Request{
		Method: "POST",
		URL:    "appointments",
		Body:   body,
	}
	// Create the Appointment
	out := Appointment{}
	err := a.request(req, &out)
	return out, err
}

// PutAppointment - update an appt
func (a *Acuity) PutAppointment(id int, body AppointmentBody) (Appointment, error) {
	// Build the Request
	req := Request{
		Method: "PUT",
		URL:    "appointments/:id",
		Params: map[string]string{"id": strconv.Itoa(id)},
		Body:   body,
	}
	// Create the Appointment
	out := Appointment{}
	err := a.request(req, &out)
	return out, err
}

// GetAppointment - get an appt
func (a *Acuity) GetAppointment(id string) (Appointment, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "appointments/:id",
		Params: map[string]string{"id": id},
	}
	// Get the Appointment
	out := Appointment{}
	err := a.request(req, &out)
	return out, err
}

// CancelAppointment - cancel an appt
func (a *Acuity) CancelAppointment(id string) (Appointment, error) {
	// Build the Request
	req := Request{
		Method: "PUT",
		URL:    "appointments/:id/cancel",
		Params: map[string]string{"id": id},
	}
	out := Appointment{}
	err := a.request(req, &out)
	return out, err
}

// RescheduleAppointment - reschedule an appt
func (a *Acuity) RescheduleAppointment(id string, datetime *time.Time) error {
	// Build the Request
	req := Request{
		Method: "PUT",
		URL:    "appointments/:id/reschedule",
		Params: map[string]string{"id": id},
		Body: map[string]interface{}{
			"datetime": datetime,
		},
	}
	out := map[string]interface{}{}
	return a.request(req, &out)
}

// GetAppointmentPayments - get an appt's payments
func (a *Acuity) GetAppointmentPayments(id string) ([]Payment, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "appointments/:id/payments",
		Params: map[string]string{"id": id},
	}
	// Get the Payments
	out := []Payment{}
	err := a.request(req, &out)
	return out, err
}

// GetAppointmentAddOns - get an appt's payments
func (a *Acuity) GetAppointmentAddOns(id string) ([]AddOn, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "appointment-addons",
	}
	// Get the AddOns
	out := []AddOn{}
	err := a.request(req, &out)
	return out, err
}

// GetAppointmentTypes - get an appt's payments
func (a *Acuity) GetAppointmentTypes(id string) ([]AppointmentType, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "appointment-types",
	}
	// Get the Types
	out := []AppointmentType{}
	err := a.request(req, &out)
	return out, err
}
