package common

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/testhelper/client"
	"strings"
)

const TokenID = client.TokenID

func ServiceClient() *golangsdk.ServiceClient {
	sc := client.ServiceClient()
	e := strings.Replace(sc.Endpoint, "v2", "v2.1", 1)
	sc.ResourceBase = e
	return sc
}
