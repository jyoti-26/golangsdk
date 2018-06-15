package hosts
import (
	"reflect"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)


// AllocateOptsBuilder allows extensions to add additional parameters to the
// Allocate request.
type AllocateOptsBuilder interface {
	ToDeHAllocateMap() (map[string]interface{}, error)
}

// AllocateOpts contains all the values needed to allocate a new DeH.
type AllocateOpts struct {
	Name             string `json:"name"`
	AvailabilityZone string `json:"availability_zone"`
	AutoPlacement    string `json:"auto_placement"`
	HostType         string `json:"host_type"`
	Quantity         int    `json:"quantity"`
}

// ToDeHAllocateMap builds a allocate request body from AllocateOpts.
func (opts AllocateOpts) ToDeHAllocateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Allocate accepts a AllocateOpts struct and uses the values to allocate a new DeH.
func Allocate(c *golangsdk.ServiceClient, opts AllocateOptsBuilder) (r AllocateResult) {
	b, err := opts.ToDeHAllocateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{200, 201}}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, reqOpt)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToDeHUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a DeH.
type UpdateOpts struct {
	Name          string `json:"name"`
	AutoPlacement string `json:"auto_placement,omitempty"`
}

// ToDeHUpdateMap builds a update request body from UpdateOpts.
func (opts UpdateOpts) ToDeHUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "dedicated_host")
}

// Update accepts a UpdateOpts struct and uses the values to update a DeH.
func Update(c *golangsdk.ServiceClient, hostID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToDeHUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{204}}
	_, r.Err = c.Put(CommonURL(c, hostID), b, nil, reqOpt)
	return
}

//Deletes the DeH using the specified hostID.
func Delete(c *golangsdk.ServiceClient, hostid string) (r DeleteResult) {
	_, r.Err = c.Delete(CommonURL(c, hostid), nil)
	return
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	// Specifies Dedicated Host ID.
	ID string `q:"dedicated_host_id"`
	// Specifies the Dedicated Host name.
	Name string `q:"name"`
	// Specifes the Dedicated Host type.
	HostType string `q:"host_type"`
	// Specifes the Dedicated Host name of type.
	HostTypeName string `q:"host_type_name"`
	// Specifies flavor ID.
	Flavor string `q:"flavor"`
	// Specifies the Dedicated Host status.
	// The value can be available, fault or released.
	State string `q:"state"`
	// Specifies the AZ to which the Dedicated Host belongs.
	Az string `q:"availability_zone"`
	// Specifies the number of entries displayed on each page.
	Limit string `q:"limit"`
	// 	The value is the ID of the last record on the previous page.
	Marker string `q:"marker"`
	// Filters the response by a date and time stamp when the dedicated host last changed status.
	ChangesSince string `q:"changes-since"`
	// Specifies the UUID of the tenant in a multi-tenancy cloud.
	TenantId string `json:"tenant_id"`
}

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToHostListQuery() (string, error)
}
// Filters out hosts parameters
func FilterHostParam(opts ListOpts) (filter ListOpts) {

	if opts.ID != "" {
		filter.ID = opts.ID
	}
	if opts.Name != "" {
		filter.Name = opts.Name
	}
	if opts.Az != "" {
		filter.Az = opts.Az
	}

	filter.HostType = opts.HostType
	filter.State = opts.State

	return filter
}

// List returns a Pager which allows you to iterate over a collection of
// dedicated hosts resources. It accepts a ListOpts struct, which allows you to
// filter the returned collection for greater efficiency.
func List(c *golangsdk.ServiceClient, opts ListOpts) ([]Host, error) {
	filter := FilterHostParam(opts)
	q, err := golangsdk.BuildQueryString(&filter)
	if err != nil {
		return nil, err
	}
	u := rootURL(c) + q.String()
	pages, err := pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return HostPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()

	allhosts, err := ExtractHosts(pages)
	if err != nil {
		return nil, err
	}

	return FilterHosts(allhosts, opts)
}

func FilterHosts(hosts []Host, opts ListOpts) ([]Host, error) {

	var refinedHosts []Host
	var matched bool
	m := map[string]interface{}{}

	if opts.ID != "" {
		m["ID"] = opts.ID
	}
	if opts.Name != "" {
		m["Name"] = opts.Name
	}
	if opts.Az != "" {
		m["Az"] = opts.Az
	}

	if len(m) > 0 && len(hosts) > 0 {
		for _, host := range hosts {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&host, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedHosts = append(refinedHosts, host)
			}
		}
	} else {
		refinedHosts = hosts
	}
	return refinedHosts, nil
}

func getStructField(v *Host, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

// Get retrieves a particular host based on its unique ID.
func Get(c *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// ListServerOptsBuilder allows extensions to add parameters to the List request.
type ListServerOptsBuilder interface {
	ToServerListQuery() (string, error)
}

// ListServerOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListServerOpts struct {
	// Specifies the number of entries displayed on each page.
	Limit int `q:"limit"`
	// The value is the ID of the last record on the previous page.
	// If the marker value is invalid, error code 400 will be returned.
	Marker string `q:"marker"`
}

// ToServerListQuery formats a ListServerOpts into a query string.
func (opts ListServerOpts) ToServerListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// ListServer makes a request against the API to list servers accessible to you.
func ListServer(client *golangsdk.ServiceClient, id string, opts ListServerOptsBuilder) pagination.Pager {
	url := listServerURL(client, id)
	if opts != nil {
		query, err := opts.ToServerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ServerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}
