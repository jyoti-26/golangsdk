package testing

import (
	"fmt"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/softwareconfiguration/deployment"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
	"net/http"
	"testing"
)

func TestCreateDeployment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/software_deployments", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r,
			`
{

    "server_id": "ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5",
    "config_id": "69070672-d37d-4095-a19c-52ab1fde9a24"
  
}	`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "software_deployment": {
        "status": null,
        "server_id": "ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5",
        "config_id": "69070672-d37d-4095-a19c-52ab1fde9a24",
        "output_values": null,
        "input_values": {},
        "action": null,
        "status_reason": null,
        "id": "ecbe324b-7f7c-4b6c-9558-32b648a5624a"
    }
}
	`)
	})
	options := deployment.CreateOpts{
		ServerId: "ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5",
		ConfigId: "69070672-d37d-4095-a19c-52ab1fde9a24",
	}
	actual, err := deployment.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5", actual.ServerId)
	th.AssertEquals(t, "69070672-d37d-4095-a19c-52ab1fde9a24", actual.ConfigId)
	th.AssertEquals(t, "ecbe324b-7f7c-4b6c-9558-32b648a5624a", actual.Id)

}

func TestGetDeploy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/software_deployments/343fef6c-1266-407b-bcb4-220b5bb085c6", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "software_deployment": {
        "status": "COMPLETE",
        "server_id": "ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5",
        "config_id": "69070672-d37d-4095-a19c-52ab1fde9a24",
        "action": "CREATE",
        "status_reason": null,
        "id": "343fef6c-1266-407b-bcb4-220b5bb085c6"
    }

}
	`)
	})
	actual, err := deployment.Get(fake.ServiceClient(), "343fef6c-1266-407b-bcb4-220b5bb085c6").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5", actual.ServerId)
	th.AssertEquals(t, "69070672-d37d-4095-a19c-52ab1fde9a24", actual.ConfigId)
	th.AssertEquals(t, "343fef6c-1266-407b-bcb4-220b5bb085c6", actual.Id)
	th.AssertEquals(t, "CREATE", actual.Action)
	th.AssertEquals(t, "COMPLETE", actual.Status)
	th.AssertEquals(t, "", actual.StatusReason)
}

func TestUpdateDeploy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/software_deployments/e99846aa-5d8b-4d1d-b324-4648952354a0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "software_deployment": {
        "status": "COMPLETE",
        "server_id": "ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5",
        "config_id": "69070672-d37d-4095-a19c-52ab1fde9a24",
		"output_values": {
            "firstname":"jyoti"
        },
        "action": "CREATE",
        "status_reason": "Outputs received",
        "id": "e99846aa-5d8b-4d1d-b324-4648952354a0"
    }
}
	`)
	})
	output := map[string]interface{}{"firstname": "jyoti"}
	options := deployment.UpdateOpts{OutputValues: output}
	actual, err := deployment.Update(fake.ServiceClient(), "e99846aa-5d8b-4d1d-b324-4648952354a0", options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "ec14c864-096e-4e27-bb8a-2c2b4dc6f3f5", actual.ServerId)
	th.AssertEquals(t, "69070672-d37d-4095-a19c-52ab1fde9a24", actual.ConfigId)
	th.AssertEquals(t, "e99846aa-5d8b-4d1d-b324-4648952354a0", actual.Id)
	th.AssertEquals(t, "COMPLETE", actual.Status)
	th.AssertEquals(t, "CREATE", actual.Action)
	th.AssertDeepEquals(t, output, actual.OutputValues)

}

func TestListDeploy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/software_deployments", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
    "software_deployments": [
        {
            "status": "IN_PROGRESS",
            "server_id": "e4b191b0-b80b-4782-994c-02abb094480e",
            "config_id": "a6ff3598-f2e0-4111-81b0-aa3e1cac2529"
        },
      {
            "status": "IN_PROGRESS",
            "server_id": "e4b191b0-b80b-4782-994c-02abb094480e",
            "config_id": "a6ff3598-f2e0-4111-81b0-aa3e1cac2529"
     }]
  }`)

	})

	listOpt := deployment.ListOpts{}
	actual, err := deployment.List(fake.ServiceClient(), listOpt)
	expected := []deployment.SoftwareDeploy{
		{
			Status:   "IN_PROGRESS",
			ServerId: "e4b191b0-b80b-4782-994c-02abb094480e",
			ConfigId: "a6ff3598-f2e0-4111-81b0-aa3e1cac2529",
		},
		{
			Status:   "IN_PROGRESS",
			ServerId: "e4b191b0-b80b-4782-994c-02abb094480e",
			ConfigId: "a6ff3598-f2e0-4111-81b0-aa3e1cac2529",
		},
	}
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, actual, expected)

}
