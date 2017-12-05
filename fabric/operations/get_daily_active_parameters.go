// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetDailyActiveParams creates a new GetDailyActiveParams object
// with the default values initialized.
func NewGetDailyActiveParams() *GetDailyActiveParams {
	var (
		authorizationDefault    = string("Bearer {access_token}")
		responseBodyPathDefault = string("{organization_id}")
		buildDefault            = string("3.0.4 (71)")
		endDefault              = string("1481328000")
		startDefault            = string("1478736000")
	)
	return &GetDailyActiveParams{
		Authorization:    &authorizationDefault,
		ResponseBodyPath: responseBodyPathDefault,
		Build:            &buildDefault,
		End:              &endDefault,
		Start:            &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDailyActiveParamsWithTimeout creates a new GetDailyActiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDailyActiveParamsWithTimeout(timeout time.Duration) *GetDailyActiveParams {
	var (
		authorizationDefault    = string("Bearer {access_token}")
		responseBodyPathDefault = string("{organization_id}")
		buildDefault            = string("3.0.4 (71)")
		endDefault              = string("1481328000")
		startDefault            = string("1478736000")
	)
	return &GetDailyActiveParams{
		Authorization:    &authorizationDefault,
		ResponseBodyPath: responseBodyPathDefault,
		Build:            &buildDefault,
		End:              &endDefault,
		Start:            &startDefault,

		timeout: timeout,
	}
}

// NewGetDailyActiveParamsWithContext creates a new GetDailyActiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDailyActiveParamsWithContext(ctx context.Context) *GetDailyActiveParams {
	var (
		authorizationDefault    = string("Bearer {access_token}")
		responseBodyPathDefault = string("{organization_id}")
		buildDefault            = string("3.0.4 (71)")
		endDefault              = string("1481328000")
		startDefault            = string("1478736000")
	)
	return &GetDailyActiveParams{
		Authorization:    &authorizationDefault,
		ResponseBodyPath: responseBodyPathDefault,
		Build:            &buildDefault,
		End:              &endDefault,
		Start:            &startDefault,

		Context: ctx,
	}
}

// NewGetDailyActiveParamsWithHTTPClient creates a new GetDailyActiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDailyActiveParamsWithHTTPClient(client *http.Client) *GetDailyActiveParams {
	var (
		authorizationDefault    = string("Bearer {access_token}")
		responseBodyPathDefault = string("{organization_id}")
		buildDefault            = string("3.0.4 (71)")
		endDefault              = string("1481328000")
		startDefault            = string("1478736000")
	)
	return &GetDailyActiveParams{
		Authorization:    &authorizationDefault,
		ResponseBodyPath: responseBodyPathDefault,
		Build:            &buildDefault,
		End:              &endDefault,
		Start:            &startDefault,
		HTTPClient:       client,
	}
}

/*GetDailyActiveParams contains all the parameters to send to the API endpoint
for the get daily active operation typically these are written to a http.Request
*/
type GetDailyActiveParams struct {

	/*Authorization*/
	Authorization *string
	/*ResponseBodyPath*/
	ResponseBodyPath string
	/*AppID*/
	AppID string
	/*Build*/
	Build *string
	/*End*/
	End *string
	/*Start*/
	Start *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get daily active params
func (o *GetDailyActiveParams) WithTimeout(timeout time.Duration) *GetDailyActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get daily active params
func (o *GetDailyActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get daily active params
func (o *GetDailyActiveParams) WithContext(ctx context.Context) *GetDailyActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get daily active params
func (o *GetDailyActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get daily active params
func (o *GetDailyActiveParams) WithHTTPClient(client *http.Client) *GetDailyActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get daily active params
func (o *GetDailyActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get daily active params
func (o *GetDailyActiveParams) WithAuthorization(authorization *string) *GetDailyActiveParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get daily active params
func (o *GetDailyActiveParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithResponseBodyPath adds the responseBodyPath to the get daily active params
func (o *GetDailyActiveParams) WithResponseBodyPath(responseBodyPath string) *GetDailyActiveParams {
	o.SetResponseBodyPath(responseBodyPath)
	return o
}

// SetResponseBodyPath adds the responseBodyPath to the get daily active params
func (o *GetDailyActiveParams) SetResponseBodyPath(responseBodyPath string) {
	o.ResponseBodyPath = responseBodyPath
}

// WithAppID adds the appID to the get daily active params
func (o *GetDailyActiveParams) WithAppID(appID string) *GetDailyActiveParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get daily active params
func (o *GetDailyActiveParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithBuild adds the build to the get daily active params
func (o *GetDailyActiveParams) WithBuild(build *string) *GetDailyActiveParams {
	o.SetBuild(build)
	return o
}

// SetBuild adds the build to the get daily active params
func (o *GetDailyActiveParams) SetBuild(build *string) {
	o.Build = build
}

// WithEnd adds the end to the get daily active params
func (o *GetDailyActiveParams) WithEnd(end *string) *GetDailyActiveParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get daily active params
func (o *GetDailyActiveParams) SetEnd(end *string) {
	o.End = end
}

// WithStart adds the start to the get daily active params
func (o *GetDailyActiveParams) WithStart(start *string) *GetDailyActiveParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get daily active params
func (o *GetDailyActiveParams) SetStart(start *string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetDailyActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	// path param ResponseBodyPath
	if err := r.SetPathParam("ResponseBodyPath", o.ResponseBodyPath); err != nil {
		return err
	}

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	if o.Build != nil {

		// query param build
		var qrBuild string
		if o.Build != nil {
			qrBuild = *o.Build
		}
		qBuild := qrBuild
		if qBuild != "" {
			if err := r.SetQueryParam("build", qBuild); err != nil {
				return err
			}
		}

	}

	if o.End != nil {

		// query param end
		var qrEnd string
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := qrEnd
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}