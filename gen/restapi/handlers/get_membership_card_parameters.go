// Code generated by go-swagger; DO NOT EDIT.

package handlers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetMembershipCardParams creates a new GetMembershipCardParams object
//
// There are no default values defined in the spec.
func NewGetMembershipCardParams() GetMembershipCardParams {

	return GetMembershipCardParams{}
}

// GetMembershipCardParams contains all the bound params for the get membership card operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetMembershipCard
type GetMembershipCardParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	MemberID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMembershipCardParams() beforehand.
func (o *GetMembershipCardParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMemberID, rhkMemberID, _ := route.Params.GetOK("memberId")
	if err := o.bindMemberID(rMemberID, rhkMemberID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMemberID binds and validates parameter MemberID from path.
func (o *GetMembershipCardParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.MemberID = raw

	return nil
}
