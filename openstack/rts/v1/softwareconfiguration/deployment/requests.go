package deployment

import (
	"reflect"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// CreateOpts contains all the values needed to create a new Deployment. There are
type CreateOpts struct {
	ServerId     string                 `json:"server_id" required:"true"`
	ConfigId     string                 `json:"config_id" required:"true"`
	Action       string                 `json:"action,omitempty"`
	InputValues  map[string]interface{} `json:"input_values,omitempty"`
	Status       string                 `json:"status,omitempty"`
	StatusReason string                 `json:"status_reason,omitempty"`
}

//  ToDeploymentCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToDeploymentCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToDeploymentCreateMap() (map[string]interface{}, error)
}

func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToDeploymentCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, reqOpt)
	return
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned.
type ListOpts struct {
	// ServerId is the the query attribute .
	ServerId string `q:"server_id"`
	// Status specify the current status  deployment resource
	Status string
	// action that triggers this deployment resource
	Action string
	//Specifies the  unique ID of this deployment resource.
	Id string
	ConfigId string
}

func List(c *golangsdk.ServiceClient, opts ListOpts) ([]SoftwareDeploy, error) {
	u := rootURL(c)
	pages, err := pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return SoftwareDeploymentPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()

	allSoftwareDeploys, err := ExtractSoftwareDeployments(pages)
	if err != nil {
		return nil, err
	}

	return FilterDeployments(allSoftwareDeploys, opts)
}

func FilterDeployments(softwareDeploys []SoftwareDeploy, opts ListOpts) ([]SoftwareDeploy, error) {
	var refinedsoftwareDeploys []SoftwareDeploy
	var matched bool
	m := map[string]interface{}{}

	if opts.ServerId != "" {
		m["ServerId"] = opts.ServerId
	}
	if opts.Action != "" {
		m["Action"] = opts.Action
	}
	if opts.Status != "" {
		m["Status"] = opts.Status
	}
	if opts.Id != "" {
		m["Id"] = opts.Id
	}
	if opts.ConfigId !=""{
		m["ConfigId"] = opts.ConfigId
	}



	if len(m) > 0 && len(softwareDeploys) > 0 {
		for _, vpc := range softwareDeploys {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&vpc, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedsoftwareDeploys = append(refinedsoftwareDeploys, vpc)
			}
		}

	} else {
		refinedsoftwareDeploys = softwareDeploys
	}

	return refinedsoftwareDeploys, nil
}

func getStructField(v *SoftwareDeploy, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

// Get retrieves a particular deployment resource details based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// Delete will permanently delete a particular  deployment resource based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), nil)
	return
}

/*// ToVpcUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToVpcUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}*/

type UpdateOpts struct {
	// Specifies the stack action that triggers this deployment resource.
	Action       string                 `json:"action,omitempty"`
	ConfigId     string                 `json:"config_id"`
	InputValues  map[string]interface{} `json:"input_values,omitempty"`
	OutputValues map[string]interface{} `json:"output_values" required:"true"`
	Status       string                 `json:"status,omitempty"`
	StatusReason string                 `json:"status_reason,omitempty"`
}

// ToDeploymentUpdateMap builds a create request body from CreateOpts.
func (opts UpdateOpts) ToDeploymentUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToDeploymentUpdateMap() (map[string]interface{}, error)
}

func Update(c *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToDeploymentUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
