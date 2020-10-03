package linuxkit

import (
	"fmt"
	"reflect"
	"testing"

	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestInit(t *testing.T) {
	testCases := []struct {
		ResourceBlock string
		Expected      []string
	}{
		{
			`
			data "linuxkit_init" "test" {
				containers = [
					"a",
					"b",
					"c",
				]
			}
			`,
			[]string{"a", "b", "c"},
		},
	}

	for _, tt := range testCases {
		r.UnitTest(t, r.TestCase{
			Providers: testProviders,
			Steps: []r.TestStep{
				{
					Config: tt.ResourceBlock,
					Check: func(s *terraform.State) error {
						id := getID(s, "data.linuxkit_init.test")
						if !reflect.DeepEqual(tt.Expected, globalCache.inits[id]) {
							return fmt.Errorf("Expected %v to match %v", globalCache.inits[id], tt.Expected)
						}

						return nil
					},
				},
			},
		})
	}
}
