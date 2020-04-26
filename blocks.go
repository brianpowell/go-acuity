/* blocks.go
// Acuity API interface
*/

package acuity

import "strconv"

// GetBlocks - pull list
func (a *Acuity) GetBlocks(query map[string]interface{}) ([]Block, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "blocks",
		Query:  query,
	}
	// Get the list of Blocks
	out := []Block{}
	err := a.request(req, &out)
	return out, err
}

// PostBlock - createa block
func (a *Acuity) PostBlock(body BlockBody) (Block, error) {
	// Build the Request
	req := Request{
		Method: "POST",
		URL:    "blocks",
		Body:   body,
	}
	// Create the Block
	out := Block{}
	err := a.request(req, &out)
	return out, err
}

// GetBlock - get a block
func (a *Acuity) GetBlock(id int) (Block, error) {
	// Build the Request
	req := Request{
		Method: "GET",
		URL:    "blocks/:id",
		Params: map[string]string{"id": strconv.Itoa(id)},
	}
	// Get the Block
	out := Block{}
	err := a.request(req, &out)
	return out, err
}

// DeleteBlock - get a block
func (a *Acuity) DeleteBlock(id int) error {
	// Build the Request
	req := Request{
		Method: "DELETE",
		URL:    "blocks/:id",
		Params: map[string]string{"id": strconv.Itoa(id)},
	}
	// Delete the Block
	err := a.request(req, nil)
	return err
}
