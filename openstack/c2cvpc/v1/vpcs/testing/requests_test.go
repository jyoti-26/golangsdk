package testing

import ("testing"
th "github.com/huaweicloud/golangsdk/testhelper"
	"net/http"
	"fmt"
	fake "github.com/huaweicloud/golangsdk/openstack/networking/v1/common"
	"github.com/huaweicloud/golangsdk/openstack/c2cvpc/v1/vpcs"
)
func TestCreateVpc(t *testing.T){
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1/85636478b0bd8e67e89469c7749d4127/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
 "vpc":
     {
     "name": "terraform-provider-vpctestcreate",
     "cidr": "192.168.0.0/16"
     }
}
			`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "vpc": {
        "id": "97e01fc2-e39e-4cfc-abf6-1d0886d120af",
        "name": "terraform-provider-vpctestcreate",
        "cidr": "192.168.0.0/16",
        "status": "CREATING"
    }
}		`)
	})

	options := vpcs.CreateOpts{
		Name: "terraform-provider-vpctestcreate",
		CIDR: "192.168.0.0/16",
	}
	n, err := vpcs.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "terraform-provider-vpctestcreate", n.Name)
	th.AssertEquals(t, "97e01fc2-e39e-4cfc-abf6-1d0886d120af", n.ID)
	th.AssertEquals(t, "192.168.0.0/16", n.CIDR)
	th.AssertEquals(t, "CREATING", n.Status)
	th.AssertEquals(t, false, n.EnableSharedSnat)
}