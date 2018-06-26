package flavors

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
	"reflect"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFlavorListQuery() (string, error)
}

/*
	AccessType maps to OpenStack's Flavor.is_public field. Although the is_public
	field is boolean, the request options are ternary, which is why AccessType is
	a string. The following values are allowed:

	The AccessType arguement is optional, and if it is not supplied, OpenStack
	returns the PublicAccess flavors.
*/
type AccessType string

const (
	// PublicAccess returns public flavors and private flavors associated with
	// that project.
	PublicAccess AccessType = "true"

	// PrivateAccess (admin only) returns private flavors, across all projects.
	PrivateAccess AccessType = "false"

	// AllAccess (admin only) returns public and private flavors across all
	// projects.
	AllAccess AccessType = "None"
)

// SortDir is a type for specifying in which direction to sort a list of flavor.
type SortDir string

// SortKey is a type for specifying by which key to sort a list of flavor.
type SortKey string

var (
	// SortAsc is used to sort a list of flavors in ascending order.
	SortAsc SortDir = "asc"
	// SortDesc is used to sort a list of flavors in descending order.
	SortDesc SortDir = "desc"
	// SortId is used to sort a list of flavors by flavorid.
	SortId SortKey = "id"
	// SortName is used to sort a list of flavors by name.
	SortName SortKey = "name"
	// SortRAM is used to sort a list of flavors by memory_mb.
	SortRAM SortKey = "memory_mb"
	// SortVCPUs is used to sort a list of flavors by vcpus.
	SortVCPUs SortKey = "vcpus"
	// SortDisk is used to sort a list of flavors by root_gb.
	SortDisk SortKey = "root_gb"
)

/*
	ListOpts filters the results returned by the List() function.
	For example, a flavor with a minDisk field of 10 will not be returned if you
	specify MinDisk set to 20.

	Typically, software will use the last ID of the previous call to List to set
	the Marker for the current call.
*/

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the flavor attributes you want to see returned.
type ListOpts struct {
	//Specifies the name of the BMS flavor
	Name string

	//Specifies the ID of the BMS flavor
	ID string

	// MinDisk and MinRAM, if provided, elides flavors which do not meet your
	// criteria.
	MinDisk int `q:"minDisk"`
	MinRAM  int `q:"minRam"`

	// AccessType, if provided, instructs List which set of flavors to return.
	// If IsPublic not provided, flavors for the current project are returned.
	AccessType AccessType `q:"is_public"`

	//SortKey allows you to sort by a particular attribute
	SortKey SortKey `q:"sort_key"`

	//SortDir sets the direction, and is either `asc' or `desc'
	SortDir SortDir `q:"sort_dir"`
}

func List(c *golangsdk.ServiceClient, opts ListOpts) ([]Flavor, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	u := listURL(c) + q.String()
	pages, err := pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return FlavorPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()

	allShares, err := ExtractFlavors(pages)
	if err != nil {
		return nil, err
	}

	return FilterFlavors(allShares, opts)
}

func FilterFlavors(shares []Flavor, opts ListOpts) ([]Flavor, error) {

	var refinedShares []Flavor
	var matched bool
	m := map[string]interface{}{}

	if opts.ID != "" {
		m["ID"] = opts.ID
	}
	if opts.Name != "" {
		m["Name"] = opts.Name
	}
	if len(m) > 0 && len(shares) > 0 {
		for _, share := range shares {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&share, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedShares = append(refinedShares, share)
			}
		}
	} else {
		refinedShares = shares
	}
	return refinedShares, nil
}

func getStructField(v *Flavor, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}
