// Code generated by go-swagger; DO NOT EDIT.

package security_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListAdvisorsParams creates a new ListAdvisorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAdvisorsParams() *ListAdvisorsParams {
	return &ListAdvisorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAdvisorsParamsWithTimeout creates a new ListAdvisorsParams object
// with the ability to set a timeout on a request.
func NewListAdvisorsParamsWithTimeout(timeout time.Duration) *ListAdvisorsParams {
	return &ListAdvisorsParams{
		timeout: timeout,
	}
}

// NewListAdvisorsParamsWithContext creates a new ListAdvisorsParams object
// with the ability to set a context for a request.
func NewListAdvisorsParamsWithContext(ctx context.Context) *ListAdvisorsParams {
	return &ListAdvisorsParams{
		Context: ctx,
	}
}

// NewListAdvisorsParamsWithHTTPClient creates a new ListAdvisorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAdvisorsParamsWithHTTPClient(client *http.Client) *ListAdvisorsParams {
	return &ListAdvisorsParams{
		HTTPClient: client,
	}
}

/*
ListAdvisorsParams contains all the parameters to send to the API endpoint

	for the list advisors operation.

	Typically these are written to a http.Request.
*/
type ListAdvisorsParams struct {
	// Body.
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list advisors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAdvisorsParams) WithDefaults() *ListAdvisorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list advisors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAdvisorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list advisors params
func (o *ListAdvisorsParams) WithTimeout(timeout time.Duration) *ListAdvisorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list advisors params
func (o *ListAdvisorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list advisors params
func (o *ListAdvisorsParams) WithContext(ctx context.Context) *ListAdvisorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list advisors params
func (o *ListAdvisorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list advisors params
func (o *ListAdvisorsParams) WithHTTPClient(client *http.Client) *ListAdvisorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list advisors params
func (o *ListAdvisorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list advisors params
func (o *ListAdvisorsParams) WithBody(body interface{}) *ListAdvisorsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list advisors params
func (o *ListAdvisorsParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListAdvisorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
