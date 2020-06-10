package transip

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/transip/gotransip/v6/haip"
	"github.com/transip/gotransip/v6/repository"
)

func dataSourceHaip() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaipRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				Type:     schema.TypeString,
			},
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"status": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"is_load_balancing_enabled": {
				Computed: true,
				Type:     schema.TypeBool,
			},
			"load_balancing_mode": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"sticky_cookie_name": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"health_check_interval": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"http_health_check_path": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"http_health_check_port": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"http_health_check_ssl": {
				Computed: true,
				Type:     schema.TypeBool,
			},
			"ip_setup": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"ptr_record": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"ipv4_address": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"ipv6_address": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"attached_ipv4_addresses": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attached_ipv6_addresses": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceHaipRead(d *schema.ResourceData, m interface{}) error {
	// Properties
	name := d.Get("name").(string)

	// Build a client and obtain the HA-IP by name
	client := m.(repository.Client)
	repository := haip.Repository{Client: client}
	v, err := repository.GetByName(name)
	if err != nil {
		return fmt.Errorf("failed to lookup ha-ip %q: %s", name, err)
	}

	// Convert and extract all attached IP addresses
	var ipv4Addresses []string
	var ipv6Addresses []string
	for _, address := range v.IPAddresses {
		if address.To4() != nil {
			ipv4Addresses = append(ipv4Addresses, address.String())
		} else {
			ipv6Addresses = append(ipv6Addresses, address.String())
		}
	}

	d.SetId(v.Name)

	d.Set("description", v.Description)
	d.Set("status", v.Status)
	d.Set("is_load_balancing_enabled", v.IsLoadBalancingEnabled)
	d.Set("load_balancing_mode", v.LoadBalancingMode)
	d.Set("sticky_cookie_name", v.StickyCookieName)
	d.Set("health_check_interval", v.HealthCheckInterval)
	d.Set("http_health_check_path", v.HTTPHealthCheckPath)
	d.Set("http_health_check_port", v.HTTPHealthCheckPort)
	d.Set("http_health_check_ssl", v.HTTPHealthCheckSsl)
	d.Set("ip_setup", v.IPSetup)
	d.Set("ptr_record", v.PtrRecord)
	d.Set("ipv4_address", v.IPv4Address.String())
	d.Set("ipv6_address", v.IPv6Address.String())
	d.Set("attached_ipv4_addresses", ipv4Addresses)
	d.Set("attached_ipv6_addresses", ipv6Addresses)

	return nil
}
