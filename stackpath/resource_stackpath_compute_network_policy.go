package stackpath

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/internal/client"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/internal/models"
)

func resourceComputeNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkPolicyCreate,
		Read:   resourceComputeNetworkPolicyRead,
		Update: resourceComputeNetworkPolicyUpdate,
		Delete: resourceComputeNetworkPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkPolicyImportState,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"instance_selector": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceComputeMatchExpressionSchema(),
			},
			"network_selector": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceComputeMatchExpressionSchema(),
			},
			"policy_types": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"egress": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ah": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"esp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"gre": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"icmp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"tcp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"tcp_udp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"udp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"to": resourceComputeNetworkPolicyHostRuleSchema(),
					},
				},
			},
			"ingress": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ah": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"esp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"gre": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"icmp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"tcp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"tcp_udp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"udp": &schema.Schema{
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"from": resourceComputeNetworkPolicyHostRuleSchema(),
					},
				},
			},
		},
	}
}

func resourceComputeNetworkPolicyHostRuleSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"instance_selector": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     resourceComputeMatchExpressionSchema(),
				},
				"network_selector": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     resourceComputeMatchExpressionSchema(),
				},
				"ip_block": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"cidr": &schema.Schema{
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validateSubnet,
							},
							"except": &schema.Schema{
								Type:     schema.TypeList,
								Optional: true,
								Elem: &schema.Schema{
									Type:         schema.TypeString,
									ValidateFunc: validateSubnet,
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceComputeNetworkPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	resp, err := config.ipam.CreateNetworkPolicy(&client.CreateNetworkPolicyParams{
		Context:              context.Background(),
		NetworkPolicyStackID: config.StackID,
		Body: &models.V1CreateNetworkPolicyRequest{
			NetworkPolicy: convertComputeNetworkPolicy(data),
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to create network policy: %v", NewStackPathError(err))
	}

	data.SetId(resp.Payload.NetworkPolicy.ID)
	return resourceComputeNetworkPolicyRead(data, meta)
}

func resourceComputeNetworkPolicyRead(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	resp, err := config.ipam.GetNetworkPolicy(&client.GetNetworkPolicyParams{
		StackID:         config.StackID,
		NetworkPolicyID: data.Id(),
		Context:         context.Background(),
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to read network policy: %v", NewStackPathError(err))
	}

	data.Set("name", resp.Payload.NetworkPolicy.Name)
	data.Set("slug", resp.Payload.NetworkPolicy.Slug)
	data.Set("description", resp.Payload.NetworkPolicy.Description)

	if err := data.Set("labels", flattenStringMap(resp.Payload.NetworkPolicy.Metadata.Labels)); err != nil {
		return fmt.Errorf("Error setting labels: %v", err)
	}

	if err := data.Set("annotations", flattenStringMap(resp.Payload.NetworkPolicy.Metadata.Annotations)); err != nil {
		return fmt.Errorf("Error setting annotations: %v", err)
	}

	if err := data.Set("instance_selector", flattenComputeMatchExpressionsOrdered("instance_selector", data, resp.Payload.NetworkPolicy.Spec.InstanceSelectors)); err != nil {
		return fmt.Errorf("Error setting instance_selector: %v", err)
	}

	if err := data.Set("network_selector", flattenComputeMatchExpressionsOrdered("network_selector", data, resp.Payload.NetworkPolicy.Spec.NetworkSelectors)); err != nil {
		return fmt.Errorf("Error setting network_selector: %v", err)
	}

	if err := data.Set("policy_types", flattenComputeNetworkPolicyTypes(resp.Payload.NetworkPolicy.Spec.PolicyTypes)); err != nil {
		return fmt.Errorf("Error setting policy_types: %v", err)
	}

	data.Set("priority", resp.Payload.NetworkPolicy.Spec.Priority)

	if err := data.Set("ingress", flattenComputeNetworkPolicyIngress(resp.Payload.NetworkPolicy.Spec.Ingress)); err != nil {
		return fmt.Errorf("Error setting ingress: %v", err)
	}
	if err := data.Set("egress", flattenComputeNetworkPolicyEgress(resp.Payload.NetworkPolicy.Spec.Egress)); err != nil {
		return fmt.Errorf("Error setting egress: %v", err)
	}

	return nil
}

func resourceComputeNetworkPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkPolicy := convertComputeNetworkPolicy(data)
	networkPolicy.ID = data.Id()

	_, err := config.ipam.UpdateNetworkPolicy(&client.UpdateNetworkPolicyParams{
		Context:              context.Background(),
		NetworkPolicyStackID: config.StackID,
		Body: &models.V1UpdateNetworkPolicyRequest{
			NetworkPolicy: networkPolicy,
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to update network policy: %v", NewStackPathError(err))
	}

	return resourceComputeNetworkPolicyRead(data, meta)
}

func resourceComputeNetworkPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	_, err := config.ipam.DeleteNetworkPolicy(&client.DeleteNetworkPolicyParams{
		Context:         context.Background(),
		StackID:         config.StackID,
		NetworkPolicyID: data.Id(),
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to delete network policy: %v", NewStackPathError(err))
	}

	data.SetId("")
	return nil
}

func resourceComputeNetworkPolicyImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// full UUID of the network policy they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
