// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

type IPAM interface {
	CreateNetworkPolicy(*CreateNetworkPolicyParams, runtime.ClientAuthInfoWriter) (*CreateNetworkPolicyOK, error)
	DeleteNetworkPolicy(*DeleteNetworkPolicyParams, runtime.ClientAuthInfoWriter) (*DeleteNetworkPolicyNoContent, error)
	GetNetworkPolicies(*GetNetworkPoliciesParams, runtime.ClientAuthInfoWriter) (*GetNetworkPoliciesOK, error)
	GetNetworkPolicy(*GetNetworkPolicyParams, runtime.ClientAuthInfoWriter) (*GetNetworkPolicyOK, error)
	UpdateNetworkPolicy(*UpdateNetworkPolicyParams, runtime.ClientAuthInfoWriter) (*UpdateNetworkPolicyOK, error)
}

// NewIPAM creates a new IPAM API client.
func NewIPAM(transport runtime.ClientTransport, formats strfmt.Registry) IPAM {
	return &ipam{transport: transport, formats: formats}
}

/*
ipam for network policy API
*/
type ipam struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateNetworkPolicy creates a new network policy
*/
func (a *ipam) CreateNetworkPolicy(params *CreateNetworkPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNetworkPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateNetworkPolicy",
		Method:             "POST",
		PathPattern:        "/ipam/v1/stacks/{network_policy.stack_id}/network_policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNetworkPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateNetworkPolicyOK), nil

}

/*
DeleteNetworkPolicy deletes a network policy
*/
func (a *ipam) DeleteNetworkPolicy(params *DeleteNetworkPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworkPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNetworkPolicy",
		Method:             "DELETE",
		PathPattern:        "/ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNetworkPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNetworkPolicyNoContent), nil

}

/*
GetNetworkPolicies gets a list of network policies by stack id
*/
func (a *ipam) GetNetworkPolicies(params *GetNetworkPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworkPolicies",
		Method:             "GET",
		PathPattern:        "/ipam/v1/stacks/{stack_id}/network_policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNetworkPoliciesOK), nil

}

/*
GetNetworkPolicy gets a network policy
*/
func (a *ipam) GetNetworkPolicy(params *GetNetworkPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNetworkPolicy",
		Method:             "GET",
		PathPattern:        "/ipam/v1/stacks/{stack_id}/network_policies/{network_policy_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNetworkPolicyOK), nil

}

/*
UpdateNetworkPolicy updates a network policy
*/
func (a *ipam) UpdateNetworkPolicy(params *UpdateNetworkPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateNetworkPolicy",
		Method:             "PUT",
		PathPattern:        "/ipam/v1/stacks/{network_policy.stack_id}/network_policies/{network_policy.id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateNetworkPolicyOK), nil

}

// SetTransport changes the transport on the client
func (a *ipam) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
