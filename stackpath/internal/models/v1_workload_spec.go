// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1WorkloadSpec The specification for the desired state of a workload
// swagger:model v1WorkloadSpec
type V1WorkloadSpec struct {

	// containers
	Containers V1ContainerSpecMapEntry `json:"containers,omitempty"`

	// Credentials that should be used to pull workload images
	ImagePullCredentials []*V1ImagePullCredential `json:"imagePullCredentials"`

	// Network interfaces to bind to the workload's instances
	NetworkInterfaces []*V1NetworkInterface `json:"networkInterfaces"`

	// virtual machines
	VirtualMachines V1VirtualMachineSpecMapEntry `json:"virtualMachines,omitempty"`

	// A list of claims that instances may reference
	//
	// The slug of the claim will be used in combination with the name of the instance to create a stable identifier. The slug should be used in the volume mount specifications for containers and VMs.
	VolumeClaimTemplates []*V1VolumeClaim `json:"volumeClaimTemplates"`
}

// Validate validates this v1 workload spec
func (m *V1WorkloadSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImagePullCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeClaimTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkloadSpec) validateContainers(formats strfmt.Registry) error {

	if swag.IsZero(m.Containers) { // not required
		return nil
	}

	if err := m.Containers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("containers")
		}
		return err
	}

	return nil
}

func (m *V1WorkloadSpec) validateImagePullCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.ImagePullCredentials) { // not required
		return nil
	}

	for i := 0; i < len(m.ImagePullCredentials); i++ {
		if swag.IsZero(m.ImagePullCredentials[i]) { // not required
			continue
		}

		if m.ImagePullCredentials[i] != nil {
			if err := m.ImagePullCredentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imagePullCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1WorkloadSpec) validateNetworkInterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkInterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkInterfaces); i++ {
		if swag.IsZero(m.NetworkInterfaces[i]) { // not required
			continue
		}

		if m.NetworkInterfaces[i] != nil {
			if err := m.NetworkInterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkInterfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1WorkloadSpec) validateVirtualMachines(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMachines) { // not required
		return nil
	}

	if err := m.VirtualMachines.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("virtualMachines")
		}
		return err
	}

	return nil
}

func (m *V1WorkloadSpec) validateVolumeClaimTemplates(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeClaimTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeClaimTemplates); i++ {
		if swag.IsZero(m.VolumeClaimTemplates[i]) { // not required
			continue
		}

		if m.VolumeClaimTemplates[i] != nil {
			if err := m.VolumeClaimTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeClaimTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkloadSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkloadSpec) UnmarshalBinary(b []byte) error {
	var res V1WorkloadSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
