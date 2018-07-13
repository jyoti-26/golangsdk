package vpcs

import (
	"github.com/huaweicloud/golangsdk"
	)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVpcCreateMap() (map[string]interface{}, error)
}

// ToVpcCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToVpcCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "vpc")
}


type CreateOpts struct {
	Name string `json:"name,omitempty"`
	CIDR string `json:"cidr,omitempty"`
}

func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVpcCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular vpc based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}


// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Vpc.
type GetResult struct {
	commonResult
}