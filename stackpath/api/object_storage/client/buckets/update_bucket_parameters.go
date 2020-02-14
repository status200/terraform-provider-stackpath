// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/object_storage/models"
)

// NewUpdateBucketParams creates a new UpdateBucketParams object
// with the default values initialized.
func NewUpdateBucketParams() *UpdateBucketParams {
	var ()
	return &UpdateBucketParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBucketParamsWithTimeout creates a new UpdateBucketParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBucketParamsWithTimeout(timeout time.Duration) *UpdateBucketParams {
	var ()
	return &UpdateBucketParams{

		timeout: timeout,
	}
}

// NewUpdateBucketParamsWithContext creates a new UpdateBucketParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBucketParamsWithContext(ctx context.Context) *UpdateBucketParams {
	var ()
	return &UpdateBucketParams{

		Context: ctx,
	}
}

// NewUpdateBucketParamsWithHTTPClient creates a new UpdateBucketParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBucketParamsWithHTTPClient(client *http.Client) *UpdateBucketParams {
	var ()
	return &UpdateBucketParams{
		HTTPClient: client,
	}
}

/*UpdateBucketParams contains all the parameters to send to the API endpoint
for the update bucket operation typically these are written to a http.Request
*/
type UpdateBucketParams struct {

	/*Body*/
	Body *models.UpdateBucketParamsBody
	/*BucketID
	  The ID for the bucket to update

	*/
	BucketID string
	/*StackID
	  The ID for the stack on which the bucket belongs to

	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update bucket params
func (o *UpdateBucketParams) WithTimeout(timeout time.Duration) *UpdateBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update bucket params
func (o *UpdateBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update bucket params
func (o *UpdateBucketParams) WithContext(ctx context.Context) *UpdateBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update bucket params
func (o *UpdateBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update bucket params
func (o *UpdateBucketParams) WithHTTPClient(client *http.Client) *UpdateBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update bucket params
func (o *UpdateBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update bucket params
func (o *UpdateBucketParams) WithBody(body *models.UpdateBucketParamsBody) *UpdateBucketParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update bucket params
func (o *UpdateBucketParams) SetBody(body *models.UpdateBucketParamsBody) {
	o.Body = body
}

// WithBucketID adds the bucketID to the update bucket params
func (o *UpdateBucketParams) WithBucketID(bucketID string) *UpdateBucketParams {
	o.SetBucketID(bucketID)
	return o
}

// SetBucketID adds the bucketId to the update bucket params
func (o *UpdateBucketParams) SetBucketID(bucketID string) {
	o.BucketID = bucketID
}

// WithStackID adds the stackID to the update bucket params
func (o *UpdateBucketParams) WithStackID(stackID string) *UpdateBucketParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the update bucket params
func (o *UpdateBucketParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bucket_id
	if err := r.SetPathParam("bucket_id", o.BucketID); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}