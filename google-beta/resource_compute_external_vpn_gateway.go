// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"google.golang.org/api/compute/v1"
)

func resourceComputeExternalVpnGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeExternalVpnGatewayCreate,
		Read:   resourceComputeExternalVpnGatewayRead,
		Delete: resourceComputeExternalVpnGatewayDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeExternalVpnGatewayImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"interface": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"redundancy_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"FOUR_IPS_REDUNDANCY", "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY", ""}, false),
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeExternalVpnGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeExternalVpnGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeExternalVpnGatewayName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	redundancyTypeProp, err := expandComputeExternalVpnGatewayRedundancyType(d.Get("redundancy_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redundancy_type"); !isEmptyValue(reflect.ValueOf(redundancyTypeProp)) && (ok || !reflect.DeepEqual(v, redundancyTypeProp)) {
		obj["redundancyType"] = redundancyTypeProp
	}
	interfacesProp, err := expandComputeExternalVpnGatewayInterface(d.Get("interface"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("interface"); !isEmptyValue(reflect.ValueOf(interfacesProp)) && (ok || !reflect.DeepEqual(v, interfacesProp)) {
		obj["interfaces"] = interfacesProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/externalVpnGateways")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ExternalVpnGateway: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ExternalVpnGateway: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating ExternalVpnGateway",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ExternalVpnGateway: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating ExternalVpnGateway %q: %#v", d.Id(), res)

	return resourceComputeExternalVpnGatewayRead(d, meta)
}

func resourceComputeExternalVpnGatewayRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/externalVpnGateways/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeExternalVpnGateway %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}

	if err := d.Set("description", flattenComputeExternalVpnGatewayDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("name", flattenComputeExternalVpnGatewayName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("redundancy_type", flattenComputeExternalVpnGatewayRedundancyType(res["redundancyType"], d)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("interface", flattenComputeExternalVpnGatewayInterface(res["interfaces"], d)); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading ExternalVpnGateway: %s", err)
	}

	return nil
}

func resourceComputeExternalVpnGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/externalVpnGateways/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ExternalVpnGateway %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ExternalVpnGateway")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting ExternalVpnGateway",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ExternalVpnGateway %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeExternalVpnGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/externalVpnGateways/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeExternalVpnGatewayDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeExternalVpnGatewayName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeExternalVpnGatewayRedundancyType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeExternalVpnGatewayInterface(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"id":         flattenComputeExternalVpnGatewayInterfaceId(original["id"], d),
			"ip_address": flattenComputeExternalVpnGatewayInterfaceIpAddress(original["ipAddress"], d),
		})
	}
	return transformed
}
func flattenComputeExternalVpnGatewayInterfaceId(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeExternalVpnGatewayInterfaceIpAddress(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeExternalVpnGatewayDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeExternalVpnGatewayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeExternalVpnGatewayRedundancyType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeExternalVpnGatewayInterface(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandComputeExternalVpnGatewayInterfaceId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["id"] = transformedId
		}

		transformedIpAddress, err := expandComputeExternalVpnGatewayInterfaceIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !isEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeExternalVpnGatewayInterfaceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeExternalVpnGatewayInterfaceIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
