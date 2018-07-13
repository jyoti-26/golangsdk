package configuration

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)


// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToSoftwareConfigListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Marker and Limit are used for pagination.
type ListOpts struct {
	Marker string `q:"marker"`
	Limit  int    `q:"limit"`
}

// ToSoftwareConfigListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSoftwareConfigListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// Software Config services. It accepts a ListOpts struct, which allows for pagination via
// marker and limit.
func List(c *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToSoftwareConfigListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {

		p := SoftwareCofigPage{pagination.MarkerPageBase{PageResult: r}}
		p.MarkerPageBase.Owner = p
		return p
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSoftwareConfigCreateMap() (map[string]interface{}, error)
}
// CreateOpts contains all the values needed to create a new Software Config. There are
// no required values.
type CreateOpts struct {
	// Specifies the script used for defining the configuration.
	Config 		string 						`json:"config, omitempty"`
	//Specifies the name of the software configuration group.
	Group 		string 						`json:"group, omitempty"`
	//Specifies the name of the software configuration.
	Name 		string 						`json:"name, omitempty"`
	//Specifies the software configuration input.
	Inputs 		[]map[string]interface{} 	`json:"inputs, omitempty"`
	//Specifies the software configuration output.
	Outputs 	[]map[string]interface{} 	`json:"outputs, omitempty"`
	//Specifies options used by a software configuration management tool.
	Options 	string 						`json:"options, omitempty"`
}

// ToSoftwareConfigCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToSoftwareConfigCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and uses the values to create a new Software config
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSoftwareConfigCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular software config based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// Delete will permanently delete a particular Software Config based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), nil)
	return
}

