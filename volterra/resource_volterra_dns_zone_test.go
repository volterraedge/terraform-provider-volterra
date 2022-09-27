//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_dns_zone "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_zone"
)

// TestDNSZone object CRUD
func TestDNSZone(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_dns_zone.ObjectType,
	})
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "ves-io")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testDNSZone(name),
			},
		},
	})
}

func testDNSZone(name string) string {
	return fmt.Sprintf(`
	resource "volterra_dns_zone" "system_yblviux3bvh7-co-jp" {
		primary {
		  default_rr_set_group {
		    ttl = 600
		    a_record {
		      name   = ""
		      values = ["1.1.1.1"]
		    }
		  }
		  default_rr_set_group {
		    ttl = 360
		    aaaa_record {
		      name   = ""
		      values = ["AAAA::1:1:1:1"]
		    }
		  }
		  default_rr_set_group {
		    ttl = 360
		    ns_record {
		      name   = ""
		      values = ["ns1.vestest.secondary.com"]
		    }
		  }
		  default_rr_set_group {
		    ttl = 480
		    cname_record {
		      name  = "test"
		      value = "testv1.vestest.io"
		    }
		  }
		  default_rr_set_group {
		    ttl = 120
		    txt_record {
		      name   = ""
		      values = ["docusign=1b0a6754-49b1-4db5-8540-d2c12664b289", "v=spf1 mx -all"]
		    }
		  }
		  default_rr_set_group {
		    ttl = 3600
		    a_record {
		      name   = "testv1"
		      values = ["111.111.111.111"]
		    }
		  }
		  default_rr_set_group {
		    ttl = 3600
		    aaaa_record {
		      name   = "testv1"
		      values = ["AAAA::111:111:111:111"]
		    }
		  }

		  rr_set_group {
		    metadata {
		      description = ""
		      disable     = false
		      name        = "vestest-services"
		    }
		    rr_set {
		      ttl = 480
		      srv_record {
			name = "_ldap._tcp"
			values {
			  port     = 389
			  priority = 10
			  target   = "sso.external.msft.com"
			  weight   = 1
			}
		      }
		    }
		    rr_set {
		      ttl = 420
		      mx_record {
			name = ""
			values {
			  domain   = "mail.google.com"
			  priority = 10
			}
		      }
		    }
		    rr_set {
		      ttl = 360
		      caa_record {
			name = ""
			values {
			  flags = 0
			  tag   = "issue"
			  value = "comodoca.com"
			}
			values {
			  flags = 128
			  tag   = "issuewild"
			  value = "digicert.com; cansignhttpexchanges=yes"
			}
			values {
			  flags = 0
			  tag   = "iodef"
			  value = "mailto:secops@smurugap.com"
			}
		      }
		    }
		  }
		  default_soa_parameters = true
		}
		name      = "%s"
		namespace = "system"
	      }

	`, name)
}
