// Code generated by go-swagger; DO NOT EDIT.

package psmdb_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreatePSMDBClusterReader is a Reader for the CreatePSMDBCluster structure.
type CreatePSMDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePSMDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePSMDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePSMDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePSMDBClusterOK creates a CreatePSMDBClusterOK with default headers values
func NewCreatePSMDBClusterOK() *CreatePSMDBClusterOK {
	return &CreatePSMDBClusterOK{}
}

/*CreatePSMDBClusterOK handles this case with default header values.

A successful response.
*/
type CreatePSMDBClusterOK struct {
	Payload interface{}
}

func (o *CreatePSMDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PSMDBCluster/Create][%d] createPsmdbClusterOk  %+v", 200, o.Payload)
}

func (o *CreatePSMDBClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreatePSMDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePSMDBClusterDefault creates a CreatePSMDBClusterDefault with default headers values
func NewCreatePSMDBClusterDefault(code int) *CreatePSMDBClusterDefault {
	return &CreatePSMDBClusterDefault{
		_statusCode: code,
	}
}

/*CreatePSMDBClusterDefault handles this case with default header values.

An unexpected error response.
*/
type CreatePSMDBClusterDefault struct {
	_statusCode int

	Payload *CreatePSMDBClusterDefaultBody
}

// Code gets the status code for the create PSMDB cluster default response
func (o *CreatePSMDBClusterDefault) Code() int {
	return o._statusCode
}

func (o *CreatePSMDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PSMDBCluster/Create][%d] CreatePSMDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePSMDBClusterDefault) GetPayload() *CreatePSMDBClusterDefaultBody {
	return o.Payload
}

func (o *CreatePSMDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreatePSMDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreatePSMDBClusterBody create PSMDB cluster body
swagger:model CreatePSMDBClusterBody
*/
type CreatePSMDBClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// PSMDB cluster name.
	// a DNS-1035 label must consist of lower case alphanumeric characters or '-',
	// start with an alphabetic character, and end with an alphanumeric character
	// (e.g. 'my-name',  or 'abc-123', regex used for validation is '[a-z]([-a-z0-9]*[a-z0-9])?')
	Name string `json:"name,omitempty"`

	// params
	Params *CreatePSMDBClusterParamsBodyParams `json:"params,omitempty"`
}

// Validate validates this create PSMDB cluster body
func (o *CreatePSMDBClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePSMDBClusterBody) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePSMDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePSMDBClusterBody) UnmarshalBinary(b []byte) error {
	var res CreatePSMDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePSMDBClusterDefaultBody create PSMDB cluster default body
swagger:model CreatePSMDBClusterDefaultBody
*/
type CreatePSMDBClusterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this create PSMDB cluster default body
func (o *CreatePSMDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePSMDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CreatePSMDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePSMDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePSMDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreatePSMDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePSMDBClusterParamsBodyParams PSMDBClusterParams represents PSMDB cluster parameters that can be updated.
swagger:model CreatePSMDBClusterParamsBodyParams
*/
type CreatePSMDBClusterParamsBodyParams struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// replicaset
	Replicaset *CreatePSMDBClusterParamsBodyParamsReplicaset `json:"replicaset,omitempty"`
}

// Validate validates this create PSMDB cluster params body params
func (o *CreatePSMDBClusterParamsBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReplicaset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePSMDBClusterParamsBodyParams) validateReplicaset(formats strfmt.Registry) error {

	if swag.IsZero(o.Replicaset) { // not required
		return nil
	}

	if o.Replicaset != nil {
		if err := o.Replicaset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "replicaset")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePSMDBClusterParamsBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePSMDBClusterParamsBodyParams) UnmarshalBinary(b []byte) error {
	var res CreatePSMDBClusterParamsBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePSMDBClusterParamsBodyParamsReplicaset ReplicaSet container parameters.
swagger:model CreatePSMDBClusterParamsBodyParamsReplicaset
*/
type CreatePSMDBClusterParamsBodyParamsReplicaset struct {

	// compute resources
	ComputeResources *CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this create PSMDB cluster params body params replicaset
func (o *CreatePSMDBClusterParamsBodyParamsReplicaset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePSMDBClusterParamsBodyParamsReplicaset) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "replicaset" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePSMDBClusterParamsBodyParamsReplicaset) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePSMDBClusterParamsBodyParamsReplicaset) UnmarshalBinary(b []byte) error {
	var res CreatePSMDBClusterParamsBodyParamsReplicaset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources
*/
type CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this create PSMDB cluster params body params replicaset compute resources
func (o *CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources) UnmarshalBinary(b []byte) error {
	var res CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}