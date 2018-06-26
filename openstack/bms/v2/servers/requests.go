package servers

import (
	"github.com/huaweicloud/golangsdk/pagination"
	"github.com/huaweicloud/golangsdk"
	"reflect"
)

// ListServerOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListServerOpts struct {
	// ID uniquely identifies this server amongst all other servers,
	// including those not accessible to the current tenant.
	ID string `json:"id"`
	// Name contains the human-readable name for the server.
	Name string `json:"name"`
	// Status contains the current operational status of the server,
	// such as IN_PROGRESS or ACTIVE.
	Status string `json:"status"`
	//ID of the user to which the BMS belongs.
	UserID string `json:"user_id"`
	//Contains the nova-compute status
	HostStatus string `json:"host_status"`
	//Contains the host ID of the BMS.
	HostID string `json:"hostid"`
	// KeyName indicates which public key was injected into the server on launch.
	KeyName string `json:"key_name"`
}

// ListServer returns a Pager which allows you to iterate over a collection of
// dedicated hosts Server resources. It accepts a ListServerOpts struct, which allows you to
// filter the returned collection for greater efficiency.
func ListServer(c *golangsdk.ServiceClient, opts ListServerOpts) ([]Server, error) {
	c.Microversion= "2.26"
	pages, err := pagination.NewPager(c,listDetailURL(c), func(r pagination.PageResult) pagination.Page {
		return ServerPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()

	allservers, err := ExtractServers(pages)
	if err != nil {
		return nil, err
	}
	return FilterServers(allservers, opts)
}

func FilterServers(servers []Server, opts ListServerOpts) ([]Server, error) {
	var refinedServers []Server
	var matched bool
	m := map[string]interface{}{}

	if opts.ID != "" {
		m["ID"] = opts.ID
	}
	if opts.Name != "" {
		m["Name"] = opts.Name
	}
	if opts.Status != "" {
		m["Status"] = opts.Status
	}
	if opts.UserID != "" {
		m["UserID"] = opts.Status
	}
	if opts.HostStatus != "" {
		m["HostStatus"] = opts.Status
	}
	if opts.HostID != "" {
		m["HostID"] = opts.Status
	}
	if opts.KeyName != "" {
		m["KeyName"] = opts.Status
	}
	if len(m) > 0 && len(servers) > 0 {
		for _, server := range servers {
			matched = true

			for key, value := range m {
				if sVal := getStructServerField(&server, key); !(sVal == value) {
					matched = false
				}
			}
			if matched {
				refinedServers = append(refinedServers, server)
			}
		}
	} else {
		refinedServers = servers
	}

	return refinedServers, nil
}

func getStructServerField(v *Server, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}
// Get requests details on a single server, by ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{"X-OpenStack-Nova-API-Version": "2.26"},
	})
	return
}
