package deployment

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type SoftwareDeploy struct {
	Action       string                 `json:"action"`
	ConfigId     string                 `json:"config_id"`
	Id           string                 `json:"id"`
	CreationTime string                 `json:"creation_time"`
	InputValues  map[string]interface{} `json:"input_values"`
	OutputValues map[string]interface{} `json:"output_values"`
	ServerId     string                 `json:"server_id"`
	Status       string                 `json:"status"`
	StatusReason string                 `json:"status_reason"`
	UpdatedTime  string                 `json:"updated_time"`
}

func (r SoftwareDeploymentPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"software_deployments_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// SoftwareDeploymentPage is the page returned by a pager when traversing over a
// collection of vpcs.
type SoftwareDeploymentPage struct {
	pagination.LinkedPageBase
}

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Vpc.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Vpc.
type UpdateResult struct {
	commonResult
}

func (r commonResult) Extract() (*SoftwareDeploy, error) {
	var s struct {
		SoftwareCofig *SoftwareDeploy `json:"software_deployment"`
	}
	err := r.ExtractInto(&s)
	return s.SoftwareCofig, err
}

func ExtractSoftwareDeployments(r pagination.Page) ([]SoftwareDeploy, error) {
	var s struct {
		SoftwareDeploys []SoftwareDeploy `json:"software_deployments"`
	}
	err := (r.(SoftwareDeploymentPage)).ExtractInto(&s)
	return s.SoftwareDeploys, err
}

// LastMarker returns the last service in a ListResult.
func (r SoftwareDeploymentPage) IsEmpty() (bool, error) {
	is, err := ExtractSoftwareDeployments(r)
	return len(is) == 0, err
}
