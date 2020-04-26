/* products.go
// Acuity API interface
*/

package acuity

// GetProducts - pull list
func (a *Acuity) GetProducts() ([]Product, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "products",
	}
	// Get the list of Products
	out := []Product{}
	err := a.request(req, &out)
	return out, err
}
