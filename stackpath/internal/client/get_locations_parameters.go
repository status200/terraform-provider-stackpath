// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLocationsParams creates a new GetLocationsParams object
// with the default values initialized.
func NewGetLocationsParams() *GetLocationsParams {
	var ()
	return &GetLocationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLocationsParamsWithTimeout creates a new GetLocationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLocationsParamsWithTimeout(timeout time.Duration) *GetLocationsParams {
	var ()
	return &GetLocationsParams{

		timeout: timeout,
	}
}

// NewGetLocationsParamsWithContext creates a new GetLocationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLocationsParamsWithContext(ctx context.Context) *GetLocationsParams {
	var ()
	return &GetLocationsParams{

		Context: ctx,
	}
}

// NewGetLocationsParamsWithHTTPClient creates a new GetLocationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLocationsParamsWithHTTPClient(client *http.Client) *GetLocationsParams {
	var ()
	return &GetLocationsParams{
		HTTPClient: client,
	}
}

/*GetLocationsParams contains all the parameters to send to the API endpoint
for the get locations operation typically these are written to a http.Request
*/
type GetLocationsParams struct {

	/*PageRequestAfter
	  The cursor value after which data will be returned.

	*/
	PageRequestAfter *string
	/*PageRequestFilter
	  SQL-style constraint filters.

	*/
	PageRequestFilter *string
	/*PageRequestFirst
	  The number of items desired.

	*/
	PageRequestFirst *string
	/*PageRequestSortBy
	  Sort the response by the given field.

	*/
	PageRequestSortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get locations params
func (o *GetLocationsParams) WithTimeout(timeout time.Duration) *GetLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get locations params
func (o *GetLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get locations params
func (o *GetLocationsParams) WithContext(ctx context.Context) *GetLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get locations params
func (o *GetLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get locations params
func (o *GetLocationsParams) WithHTTPClient(client *http.Client) *GetLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get locations params
func (o *GetLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageRequestAfter adds the pageRequestAfter to the get locations params
func (o *GetLocationsParams) WithPageRequestAfter(pageRequestAfter *string) *GetLocationsParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get locations params
func (o *GetLocationsParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get locations params
func (o *GetLocationsParams) WithPageRequestFilter(pageRequestFilter *string) *GetLocationsParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get locations params
func (o *GetLocationsParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get locations params
func (o *GetLocationsParams) WithPageRequestFirst(pageRequestFirst *string) *GetLocationsParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get locations params
func (o *GetLocationsParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get locations params
func (o *GetLocationsParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetLocationsParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get locations params
func (o *GetLocationsParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageRequestAfter != nil {

		// query param page_request.after
		var qrPageRequestAfter string
		if o.PageRequestAfter != nil {
			qrPageRequestAfter = *o.PageRequestAfter
		}
		qPageRequestAfter := qrPageRequestAfter
		if qPageRequestAfter != "" {
			if err := r.SetQueryParam("page_request.after", qPageRequestAfter); err != nil {
				return err
			}
		}

	}

	if o.PageRequestFilter != nil {

		// query param page_request.filter
		var qrPageRequestFilter string
		if o.PageRequestFilter != nil {
			qrPageRequestFilter = *o.PageRequestFilter
		}
		qPageRequestFilter := qrPageRequestFilter
		if qPageRequestFilter != "" {
			if err := r.SetQueryParam("page_request.filter", qPageRequestFilter); err != nil {
				return err
			}
		}

	}

	if o.PageRequestFirst != nil {

		// query param page_request.first
		var qrPageRequestFirst string
		if o.PageRequestFirst != nil {
			qrPageRequestFirst = *o.PageRequestFirst
		}
		qPageRequestFirst := qrPageRequestFirst
		if qPageRequestFirst != "" {
			if err := r.SetQueryParam("page_request.first", qPageRequestFirst); err != nil {
				return err
			}
		}

	}

	if o.PageRequestSortBy != nil {

		// query param page_request.sort_by
		var qrPageRequestSortBy string
		if o.PageRequestSortBy != nil {
			qrPageRequestSortBy = *o.PageRequestSortBy
		}
		qPageRequestSortBy := qrPageRequestSortBy
		if qPageRequestSortBy != "" {
			if err := r.SetQueryParam("page_request.sort_by", qPageRequestSortBy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
