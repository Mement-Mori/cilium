// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DNSLookup An IP -> DNS mapping, with metadata
//
// swagger:model DNSLookup
type DNSLookup struct {

	// The endpoint that made this lookup, or 0 for the agent itself.
	EndpointID int64 `json:"endpoint-id,omitempty"`

	// The absolute time when this data will expire in this cache
	// Format: date-time
	ExpirationTime strfmt.DateTime `json:"expiration-time,omitempty"`

	// DNS name
	Fqdn string `json:"fqdn,omitempty"`

	// IP addresses returned in this lookup
	Ips []string `json:"ips"`

	// The absolute time when this data was recieved
	// Format: date-time
	LookupTime strfmt.DateTime `json:"lookup-time,omitempty"`

	// The reason this FQDN IP association exists. Either a DNS lookup or an ongoing connection to an IP that was created by a DNS lookup.
	Source string `json:"source,omitempty"`

	// The TTL in the DNS response
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this DNS lookup
func (m *DNSLookup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLookupTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSLookup) validateExpirationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration-time", "body", "date-time", m.ExpirationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DNSLookup) validateLookupTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LookupTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lookup-time", "body", "date-time", m.LookupTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DNSLookup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSLookup) UnmarshalBinary(b []byte) error {
	var res DNSLookup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
