package vpcs

import (
	"net/http"
	"github.com/huaweicloud/golangsdk"
     )


type Result struct {
	// Body is the payload of the HTTP response from the server. In most cases,
	// this will be the deserialized JSON structure.
	Body interface{}

	// Header contains the HTTP header structure from the original response.
	Header http.Header

	// Err is an error that occurred during the operation. It's deferred until
	// extraction to make it easier to chain the Extract call.
	Err error
}


// Route is a possible route in a vpc.
type Route struct {
	NextHop         string `json:"nexthop"`
	DestinationCIDR string `json:"destination"`
}


type Vpc struct {
	// ID is the unique identifier for the vpc.
	ID string `json:"id"`

	// Name is the human readable name for the vpc. It does not have to be
	// unique.
	Name string `json:"name"`

	//Specifies the range of available subnets in the VPC.
	CIDR string `json:"cidr"`

	// Status indicates whether or not a vpc is currently operational.
	Status string `json:"status"`

	// Routes are a collection of static routes that the vpc will host.
	Routes []Route `json:"routes"`

	//Provides informaion about shared snat
	EnableSharedSnat bool `json:"enable_shared_snat"`
}

type commonResult struct {
	golangsdk.Result
}


// Extract is a function that accepts a result and extracts a vpc.
func (r commonResult) Extract() (*Vpc, error) {
	var s struct {
		Vpc *Vpc `json:"vpc"`
	}
	err := r.ExtractInto(&s)
	return s.Vpc, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Vpc.
type CreateResult struct {
	commonResult
}