// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OAuthParamsBody o auth params body
// swagger:model oAuthParamsBody
type OAuthParamsBody struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// grant type
	GrantType string `json:"grant_type,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this o auth params body
func (m *OAuthParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OAuthParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuthParamsBody) UnmarshalBinary(b []byte) error {
	var res OAuthParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
