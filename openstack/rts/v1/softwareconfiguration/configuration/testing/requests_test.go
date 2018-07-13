package testing

import (
	"fmt"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/softwareconfiguration/configuration"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
	"net/http"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/software_configs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r,
			`
{
    "name": "SoftwareConfig_test"
  
}	`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "software_config": {
        "inputs": [
            {
              
                "type": "String",
                "name": "foo"
            }
        ],
        "group": "script",
        "name": "SoftwareConfig_test",
        "outputs": [
            {
                "type": "String",
                "name": "result"
                
            }
        ],
       "id": "ddee7aca-aa32-4335-8265-d436b20db4f1"
    }
}	`)
	})

	options := configuration.CreateOpts{
		Name: "SoftwareConfig_test",
	}

	n, err := configuration.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "SoftwareConfig_test", n.Name)

}

func TestDeleteConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/software_configs/dec17794-adfb-4374-9bdd-aaf45ceaa4f0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
	res := configuration.Delete(fake.ServiceClient(), "dec17794-adfb-4374-9bdd-aaf45ceaa4f0")
	th.AssertNoErr(t, res.Err)
}

func TestGetConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/software_configs/a6ff3598-f2e0-4111-81b0-aa3e1cac2529", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "software_config": {
        "inputs": [
            {
                "default": null,
                "type": "String",
                "name": "foo",
                "description": null
            },
            {
                "default": null,
                "type": "String",
                "name": "bar",
                "description": null
            }
        ],
        "group": "script",
        "name": "a-config-we5zpvyu7b5o",
        "outputs": [
            {
                "type": "String",
                "name": "result",
                "error_output": false,
                "description": null
            }
        ],
        "creation_time": "2018-06-12T05:24:48.995227"
       
    }
}`)
	})

	result, err := configuration.Get(fake.ServiceClient(), "a6ff3598-f2e0-4111-81b0-aa3e1cac2529").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "a-config-we5zpvyu7b5o", result.Name)
	fmt.Println(result)

}

func TestListConfig(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/software_configs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "software_configs": [
        {
            "creation_time": "2018-07-10T08:27:43.946413",
            "group": "script",
            "id": "f77b6e3b-74d8-4102-a08b-4792d88ba954",
            "name": "test-cong"
        },
        {
            "creation_time": "2018-07-10T08:23:29.419303",
            "group": "script",
            "id": "871b2f0c-0ccb-4b22-ac07-1a75c8923d89",
            "name": "terraform-provider_test"
        }
    ]
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "software_configs": [
        {
            "creation_time": "2018-07-10T08:27:43.946413",
            "group": "script",
            "id": "f77b6e3b-74d8-4102-a08b-4792d88ba954",
            "name": "test-cong"
        },
        {
            "creation_time": "2018-07-10T08:23:29.419303",
            "group": "script",
            "id": "871b2f0c-0ccb-4b22-ac07-1a75c8923d89",
            "name": "terraform-provider_test"
        }
    ]
}	`)

	})

	pages,err:=configuration.List(fake.ServiceClient(),configuration.ListOpts{}).AllPages()
	th.AssertNoErr(t,err)
	fmt.Println(pages)

}
