package octopusdeploy

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"octopusdeploy": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("OCTOPUS_URL"); isEmpty(v) {
		t.Fatal("OCTOPUS_URL must be set for acceptance tests")
	}
	if v := os.Getenv("OCTOPUS_APIKEY"); isEmpty(v) {
		t.Fatal("OCTOPUS_APIKEY must be set for acceptance tests")
	}
}
