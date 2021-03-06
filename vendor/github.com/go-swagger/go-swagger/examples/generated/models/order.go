// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Order order
// swagger:model Order

type Order struct {

	// complete
	Complete bool `json:"complete,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// pet Id
	PetID int64 `json:"petId,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`

	// ship date
	ShipDate strfmt.DateTime `json:"shipDate,omitempty"`

	// Order Status
	Status string `json:"status,omitempty"`
}

/* polymorph Order complete false */

/* polymorph Order id false */

/* polymorph Order petId false */

/* polymorph Order quantity false */

/* polymorph Order shipDate false */

/* polymorph Order status false */

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
