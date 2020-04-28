/* orders.go
// Acuity API interface
*/

package acuity

// GetOrders - pull list
func (a *Acuity) GetOrders() ([]Order, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "orders",
	}
	// Get the list of Orders
	out := []Order{}
	err := a.request(req, &out)
	return out, err
}

// GetOrder - get an order
func (a *Acuity) GetOrder(id string) (Order, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "orders/:id",
		Params: map[string]string{"id": id},
	}
	// Get the Order
	out := Order{}
	err := a.request(req, &out)
	return out, err
}
