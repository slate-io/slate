package deserialize

import (
	"fmt"
	"github.com/rightlag/slate/spec"
	"net/url"
	"testing"
)

var mockIpBrief = `
Interface             IP-Address      OK?    Method Status      Protocol
GigabitEthernet0/1    unassigned      YES    unset  up          up
GigabitEthernet0/2    192.168.190.235 YES    unset  up          up
GigabitEthernet0/3    unassigned      YES    unset  up          up
GigabitEthernet0/4    192.168.191.2   YES    unset  up          up`

func TestDeserialize(t *testing.T) {
	u, err := url.Parse("http://petstore.swagger.io/v2/swagger.json")
	if err != nil {
		t.Error(err)
	}

	schema, err := spec.LoadSchema(u)
	if err != nil {
		t.Error(err)
	}

	reg := schema.Paths["/show ip interface"].Get.Responses.Schema.Pattern
	d := deserialize(mockIpBrief, reg)
	fmt.Printf("%#v", d)
	if len(d) != 4 {
		t.Error("Length is not 4")
	}
	if d[0].Protocol !== 'up' {
		t.Error("Protocol does not match")
	}
}
