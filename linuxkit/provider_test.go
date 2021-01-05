package linuxkit

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testProviders = map[string]terraform.ResourceProvider{
	"linuxkit": Provider(),
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func getID(s *terraform.State, resource string) string {
	return s.RootModule().Resources[resource].Primary.ID
}
