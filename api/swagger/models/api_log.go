// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APILog api log
// swagger:model apiLog

type APILog struct {

	// Last lines of log file
	Lines []string `json:"lines"`
}

/* polymorph apiLog lines false */

// Validate validates this api log
func (m *APILog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLines(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APILog) validateLines(formats strfmt.Registry) error {

	if swag.IsZero(m.Lines) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APILog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APILog) UnmarshalBinary(b []byte) error {
	var res APILog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
