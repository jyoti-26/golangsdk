package testing

import (
	"net/http"
	"testing"

	"fmt"
	"github.com/huaweicloud/golangsdk/openstack/deh/v1/hosts"
	fake "github.com/huaweicloud/golangsdk/openstack/deh/v1/hosts/common"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/17fbda95add24720a4038ba4b1c705ed/dedicated-hosts/66156a61-27c2-4169-936b-910dd9c73da3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "dedicated_host": {
        "allocated_at": "2018-06-13T07:44:55Z",
        "availability_zone": "eu-de-02",
        "csg_host": "pod01.eu-de-02",
        "name": "test-aj2",
        "available_memory": 270336,
        "released_at": "",
        "auto_placement": "off",
        "available_vcpus": 36,
        "dedicated_host_id": "66156a61-27c2-4169-936b-910dd9c73da3",
        "state": "available",
        "instance_total": 0,
        "host_properties": {           
            "host_type": "h1",
            "vcpus": 36,
            "memory": 270336,
            "cores": 12,
            "sockets": 2,
            "host_type_name": "High performance"
        },
        "csd_host": "fc-nova-compute010#8120665",
        "instance_uuids": [],
        "project_id": "17fbda95add24720a4038ba4b1c705ed"
    }
}
		`)
	})

	s, err := hosts.Get(fake.ServiceClient(), "66156a61-27c2-4169-936b-910dd9c73da3").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "66156a61-27c2-4169-936b-910dd9c73da3", s.ID)
	th.AssertEquals(t, "test-aj2", s.Name)
	th.AssertEquals(t, "eu-de-02", s.Az)
	th.AssertEquals(t, "available", s.State)
	th.AssertDeepEquals(t, hosts.HostPropertiesOpts{
		HostTypeName: "High performance",
		HostType:     "h1",
		Vcpus:        36,
		Memory:       270336,
		Cores:        12,
		Sockets:      2,
	}, s.HostProperties)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/17fbda95add24720a4038ba4b1c705ed/dedicated-hosts", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "dedicated_hosts": [ {
            "availability_zone": "eu-de-01",
            "name": "c2c-deh-test",
            "available_memory": 262144,
            "auto_placement": "off",
            "available_vcpus": 70,
            "dedicated_host_id": "671611d2-b45c-4648-9e78-06eb24522291",
            "state": "available",
            "instance_total": 2,
            "host_properties": {                
                "host_type": "general",
                "vcpus": 72,
                "memory": 270336,
                "cores": 12,
                "sockets": 2,
                "host_type_name": "General computing"
            },
            "instance_uuids": [
                "3de1ce75-2550-4a46-a689-dd33ca2b62d6",
                "885dc71d-905d-48b5-bae7-db66801dc175"
            ],
            "project_id": "17fbda95add24720a4038ba4b1c705ed"
        }]
}
			`)
	})

	//count := 0

	actual, err := hosts.List(fake.ServiceClient(), hosts.ListOpts{})
	if err != nil {
		t.Errorf("Failed to extract hosts: %v", err)
	}

	expected := []hosts.Host{
		{
			Az:              "eu-de-01",
			Name:            "c2c-deh-test",
			AvailableMemory: 262144,
			AvailableVcpus:  70,
			ID:              "671611d2-b45c-4648-9e78-06eb24522291",
			State:           "available",
			InstanceTotal:   2,
			AutoPlacement:   "off",
			TenantId:        "17fbda95add24720a4038ba4b1c705ed",
			HostProperties: hosts.HostPropertiesOpts{
				HostType:     "general",
				Vcpus:        72,
				Memory:       270336,
				Cores:        12,
				Sockets:      2,
				HostTypeName: "General computing",
			},
			InstanceUuids: []string{"3de1ce75-2550-4a46-a689-dd33ca2b62d6",
				"885dc71d-905d-48b5-bae7-db66801dc175"},
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestListServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/17fbda95add24720a4038ba4b1c705ed/dedicated-hosts/671611d2-b45c-4648-9e78-06eb24522291/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
     "servers": [ {
            "status": "ACTIVE", 
            "flavor": {
                "id": "normal1"
            },
 			"addresses": {
                "0b98c646-617f-4d90-9ca5-385f0cd73ea7": [
                    {
                        "version": 4,
                        "addr": "192.168.3.133"
                    }
                ]
            },
            "id": "3de1ce75-2550-4a46-a689-dd33ca2b62d6",
            "user_id": "6d78fa8550ae45d6932a1fadfb1fa552",
            "name": "c2c-ecs-test-2",
            "tenant_id": "17fbda95add24720a4038ba4b1c705ed",
            "metadata": {
                "metering.image_id": "c0ea3ff1-432e-4650-8a1b-372a80b2d2be",
                "metering.imagetype": "gold",
                "metering.resourcespeccode": "deh.linux",
                "metering.cloudServiceType": "sys.service.type.ec2",
                "image_name": "Standard_CentOS_7_latest",
                "metering.resourcetype": "1",
                "os_bit": "64",
                "vpc_id": "0b98c646-617f-4d90-9ca5-385f0cd73ea7",
                "os_type": "Linux",
                "charging_mode": "0"
            }
        }]
}
			`)
	})

	actual, err := hosts.ListServer(fake.ServiceClient(), "671611d2-b45c-4648-9e78-06eb24522291", hosts.ListServerOpts{})
	th.AssertNoErr(t, err)

	expected := []hosts.Server{
		{
			Status: "ACTIVE",
			Addresses: map[string]interface{}{
				"0b98c646-617f-4d90-9ca5-385f0cd73ea7": []interface{}{
					map[string]interface{}{
						"version": float64(4),
						"addr":    "192.168.3.133",
					},
				},
			},
			Flavor: map[string]interface{}{
				"id": "normal1",
			},
			ID:       "3de1ce75-2550-4a46-a689-dd33ca2b62d6",
			UserID:   "6d78fa8550ae45d6932a1fadfb1fa552",
			Name:     "c2c-ecs-test-2",
			TenantID: "17fbda95add24720a4038ba4b1c705ed",
			Metadata: map[string]string{
				"metering.image_id":         "c0ea3ff1-432e-4650-8a1b-372a80b2d2be",
				"metering.imagetype":        "gold",
				"metering.resourcespeccode": "deh.linux",
				"metering.cloudServiceType": "sys.service.type.ec2",
				"image_name":                "Standard_CentOS_7_latest",
				"metering.resourcetype":     "1",
				"os_bit":                    "64",
				"vpc_id":                    "0b98c646-617f-4d90-9ca5-385f0cd73ea7",
				"os_type":                   "Linux",
				"charging_mode":             "0",
			},
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}


func TestAllocateDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/dedicated-hosts", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, allocateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, allocateResponse)
	})

	c := client.ServiceClient()
	allocateOpts := hosts.AllocateOpts{Name: "Test-1",
		AvailabilityZone: "eu-de-02",
		HostType:         "h1",
		AutoPlacement:    "off",
		Quantity:         2}
	s, err := hosts.Allocate(c, allocateOpts).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, s, &hosts.AllocatedHosts{
		AllocatedHostIds: []string{"fb4733fd-70a3-44e1-a1cb-0311f028d7e5",
			"7408f985-047d-4313-b3c8-8e12bef01d12"},
	})
}

func TestUpdateDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/dedicated-hosts/"+HostID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, updateRequest)
		w.WriteHeader(http.StatusNoContent)
	})

	c := client.ServiceClient()
	updateOpts := hosts.UpdateOpts{Name: "Test-2",
		AutoPlacement: "off",
	}
	s := hosts.Update(c, HostID, updateOpts)
	th.AssertNoErr(t, s.Err)
}

func TestDeleteDeH(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/dedicated-hosts/"+HostID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})

	result := hosts.Delete(client.ServiceClient(), HostID)
	th.AssertNoErr(t, result.Err)
}
