package oktaasa

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"oktaasa": testAccProvider,
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
	if v := os.Getenv("OKTAASA_KEY"); v == "" {
		t.Fatal("OKTAASA_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("OKTAASA_KEY_SECRET"); v == "" {
		t.Fatal("OKTAASA_KEY_SECRET must be set for acceptance tests")
	}
	if v := os.Getenv("OKTAASA_TEAM"); v == "" {
		t.Fatal("OKTAASA_TEAM must be set for acceptance tests")
	}
}
