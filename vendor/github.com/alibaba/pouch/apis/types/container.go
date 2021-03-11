// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Container an array of Container contains response of Engine API:
// GET "/containers/json"
//
// swagger:model Container
type Container struct {

	// command
	Command string `json:"Command,omitempty"`

	// Created time of container in daemon.
	Created int64 `json:"Created,omitempty"`

	// In Moby's API, HostConfig field in Container struct has following type
	// struct { NetworkMode string `json:",omitempty"` }
	// In Pouch, we need to pick runtime field in HostConfig from daemon side to judge runtime type,
	// So Pouch changes this type to be the complete HostConfig.
	// Incompatibility exists, ATTENTION.
	//
	HostConfig *HostConfig `json:"HostConfig,omitempty"`

	// Container ID
	ID string `json:"Id,omitempty"`

	// image
	Image string `json:"Image,omitempty"`

	// image ID
	ImageID string `json:"ImageID,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// Set of mount point in a container.
	Mounts []MountPoint `json:"Mounts"`

	// names
	Names []string `json:"Names"`

	// network settings
	NetworkSettings *ContainerNetworkSettings `json:"NetworkSettings,omitempty"`

	// size root fs
	SizeRootFs int64 `json:"SizeRootFs,omitempty"`

	// size rw
	SizeRw int64 `json:"SizeRw,omitempty"`

	// state
	State string `json:"State,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this container
func (m *Container) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Container) validateHostConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.HostConfig) { // not required
		return nil
	}

	if m.HostConfig != nil {
		if err := m.HostConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HostConfig")
			}
			return err
		}
	}

	return nil
}

func (m *Container) validateMounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Mounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Mounts); i++ {

		if err := m.Mounts[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Mounts" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Container) validateNetworkSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Container) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Container) UnmarshalBinary(b []byte) error {
	var res Container
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContainerNetworkSettings container network settings
// swagger:model ContainerNetworkSettings
type ContainerNetworkSettings struct {

	// networks
	Networks map[string]*EndpointSettings `json:"Networks,omitempty"`
}

// Validate validates this container network settings
func (m *ContainerNetworkSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerNetworkSettings) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for k := range m.Networks {

		if err := validate.Required("NetworkSettings"+"."+"Networks"+"."+k, "body", m.Networks[k]); err != nil {
			return err
		}
		if val, ok := m.Networks[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerNetworkSettings) UnmarshalBinary(b []byte) error {
	var res ContainerNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}