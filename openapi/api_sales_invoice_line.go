/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// SalesInvoiceLineApiService SalesInvoiceLineApi service
type SalesInvoiceLineApiService service

/*
DeleteSalesInvoiceLine Deletes an object of type salesInvoiceLine in Dynamics 365 Business Central
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param salesInvoiceLineId id for salesInvoiceLine
*/
func (a *SalesInvoiceLineApiService) DeleteSalesInvoiceLine(ctx _context.Context, companyId string, salesInvoiceLineId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoiceLines('{salesInvoiceLine_id}')"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoiceLine_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceLineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
DeleteSalesInvoiceLineForSalesInvoice Deletes an object of type salesInvoiceLine in Dynamics 365 Business Central
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param salesInvoiceId id for salesInvoice
 * @param salesInvoiceLineId id for salesInvoiceLine
*/
func (a *SalesInvoiceLineApiService) DeleteSalesInvoiceLineForSalesInvoice(ctx _context.Context, companyId string, salesInvoiceId string, salesInvoiceLineId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines('{salesInvoiceLine_id}')"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoice_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoiceLine_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceLineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// GetSalesInvoiceLineOpts Optional parameters for the method 'GetSalesInvoiceLine'
type GetSalesInvoiceLineOpts struct {
	Expand  optional.Interface
	Select_ optional.Interface
}

/*
GetSalesInvoiceLine Retrieve the properties and relationships of an object of type salesInvoiceLine for Dynamics 365 Business Central.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param salesInvoiceLineId id for salesInvoiceLine
 * @param optional nil or *GetSalesInvoiceLineOpts - Optional Parameters:
 * @param "Expand" (optional.Interface of []string) -  Entities to expand
 * @param "Select_" (optional.Interface of []string) -  Selected properties to be retrieved
@return SalesInvoiceLine
*/
func (a *SalesInvoiceLineApiService) GetSalesInvoiceLine(ctx _context.Context, companyId string, salesInvoiceLineId string, localVarOptionals *GetSalesInvoiceLineOpts) (SalesInvoiceLine, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SalesInvoiceLine
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoiceLines('{salesInvoiceLine_id}')"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoiceLine_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceLineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("$expand", parameterToString(localVarOptionals.Expand.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("$select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetSalesInvoiceLineForSalesInvoiceOpts Optional parameters for the method 'GetSalesInvoiceLineForSalesInvoice'
type GetSalesInvoiceLineForSalesInvoiceOpts struct {
	Expand  optional.Interface
	Select_ optional.Interface
}

/*
GetSalesInvoiceLineForSalesInvoice Retrieve the properties and relationships of an object of type salesInvoiceLine for Dynamics 365 Business Central.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param salesInvoiceId id for salesInvoice
 * @param salesInvoiceLineId id for salesInvoiceLine
 * @param optional nil or *GetSalesInvoiceLineForSalesInvoiceOpts - Optional Parameters:
 * @param "Expand" (optional.Interface of []string) -  Entities to expand
 * @param "Select_" (optional.Interface of []string) -  Selected properties to be retrieved
@return SalesInvoiceLine
*/
func (a *SalesInvoiceLineApiService) GetSalesInvoiceLineForSalesInvoice(ctx _context.Context, companyId string, salesInvoiceId string, salesInvoiceLineId string, localVarOptionals *GetSalesInvoiceLineForSalesInvoiceOpts) (SalesInvoiceLine, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SalesInvoiceLine
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines('{salesInvoiceLine_id}')"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoice_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoiceLine_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceLineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("$expand", parameterToString(localVarOptionals.Expand.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("$select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListSalesInvoiceLinesOpts Optional parameters for the method 'ListSalesInvoiceLines'
type ListSalesInvoiceLinesOpts struct {
	Top     optional.Int32
	Skip    optional.Int32
	Limit   optional.Int32
	Filter  optional.String
	Expand  optional.Interface
	Select_ optional.Interface
}

/*
ListSalesInvoiceLines Returns a list of salesInvoiceLines
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param optional nil or *ListSalesInvoiceLinesOpts - Optional Parameters:
 * @param "Top" (optional.Int32) -  Number of items to return from the top of the list
 * @param "Skip" (optional.Int32) -  Number of items to skip from the list
 * @param "Limit" (optional.Int32) -  Number of items to return from the list
 * @param "Filter" (optional.String) -  Filtering expression
 * @param "Expand" (optional.Interface of []string) -  Entities to expand
 * @param "Select_" (optional.Interface of []string) -  Selected properties to be retrieved
@return InlineResponse2009
*/
func (a *SalesInvoiceLineApiService) ListSalesInvoiceLines(ctx _context.Context, companyId string, localVarOptionals *ListSalesInvoiceLinesOpts) (InlineResponse2009, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2009
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoiceLines"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Top.IsSet() {
		localVarQueryParams.Add("$top", parameterToString(localVarOptionals.Top.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("$skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("$limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("$filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("$expand", parameterToString(localVarOptionals.Expand.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("$select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListSalesInvoiceLinesForSalesInvoiceOpts Optional parameters for the method 'ListSalesInvoiceLinesForSalesInvoice'
type ListSalesInvoiceLinesForSalesInvoiceOpts struct {
	Top     optional.Int32
	Skip    optional.Int32
	Limit   optional.Int32
	Filter  optional.String
	Expand  optional.Interface
	Select_ optional.Interface
}

/*
ListSalesInvoiceLinesForSalesInvoice Returns a list of salesInvoiceLines
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param salesInvoiceId id for salesInvoice
 * @param optional nil or *ListSalesInvoiceLinesForSalesInvoiceOpts - Optional Parameters:
 * @param "Top" (optional.Int32) -  Number of items to return from the top of the list
 * @param "Skip" (optional.Int32) -  Number of items to skip from the list
 * @param "Limit" (optional.Int32) -  Number of items to return from the list
 * @param "Filter" (optional.String) -  Filtering expression
 * @param "Expand" (optional.Interface of []string) -  Entities to expand
 * @param "Select_" (optional.Interface of []string) -  Selected properties to be retrieved
@return InlineResponse2009
*/
func (a *SalesInvoiceLineApiService) ListSalesInvoiceLinesForSalesInvoice(ctx _context.Context, companyId string, salesInvoiceId string, localVarOptionals *ListSalesInvoiceLinesForSalesInvoiceOpts) (InlineResponse2009, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2009
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoice_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Top.IsSet() {
		localVarQueryParams.Add("$top", parameterToString(localVarOptionals.Top.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("$skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("$limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("$filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("$expand", parameterToString(localVarOptionals.Expand.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("$select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PatchSalesInvoiceLine Updates an object of type salesInvoiceLine in Dynamics 365 Business Central
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param salesInvoiceLineId id for salesInvoiceLine
 * @param contentType application/json
 * @param ifMatch Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated.
 * @param uNKNOWNBASETYPE
@return SalesInvoiceLine
*/
func (a *SalesInvoiceLineApiService) PatchSalesInvoiceLine(ctx _context.Context, companyId string, salesInvoiceLineId string, contentType string, ifMatch string, uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) (SalesInvoiceLine, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SalesInvoiceLine
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoiceLines('{salesInvoiceLine_id}')"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoiceLine_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceLineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["If-Match"] = parameterToString(ifMatch, "")
	// body params
	localVarPostBody = &uNKNOWNBASETYPE
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PatchSalesInvoiceLineForSalesInvoice Updates an object of type salesInvoiceLine in Dynamics 365 Business Central
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param salesInvoiceId id for salesInvoice
 * @param salesInvoiceLineId id for salesInvoiceLine
 * @param contentType application/json
 * @param ifMatch Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated.
 * @param uNKNOWNBASETYPE
@return SalesInvoiceLine
*/
func (a *SalesInvoiceLineApiService) PatchSalesInvoiceLineForSalesInvoice(ctx _context.Context, companyId string, salesInvoiceId string, salesInvoiceLineId string, contentType string, ifMatch string, uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) (SalesInvoiceLine, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SalesInvoiceLine
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines('{salesInvoiceLine_id}')"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoice_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoiceLine_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceLineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["If-Match"] = parameterToString(ifMatch, "")
	// body params
	localVarPostBody = &uNKNOWNBASETYPE
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PostSalesInvoiceLine Creates an object of type salesInvoiceLine in Dynamics 365 Business Central
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param contentType application/json
 * @param uNKNOWNBASETYPE
@return SalesInvoiceLine
*/
func (a *SalesInvoiceLineApiService) PostSalesInvoiceLine(ctx _context.Context, companyId string, contentType string, uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) (SalesInvoiceLine, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SalesInvoiceLine
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoiceLines"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	// body params
	localVarPostBody = &uNKNOWNBASETYPE
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PostSalesInvoiceLineForSalesInvoice Creates an object of type salesInvoiceLine in Dynamics 365 Business Central
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param companyId id for company
 * @param salesInvoiceId id for salesInvoice
 * @param contentType application/json
 * @param uNKNOWNBASETYPE
@return SalesInvoiceLine
*/
func (a *SalesInvoiceLineApiService) PostSalesInvoiceLineForSalesInvoice(ctx _context.Context, companyId string, salesInvoiceId string, contentType string, uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) (SalesInvoiceLine, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SalesInvoiceLine
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/companies({company_id})/salesInvoices({salesInvoice_id})/salesInvoiceLines"
	localVarPath = strings.Replace(localVarPath, "{"+"company_id"+"}", _neturl.QueryEscape(parameterToString(companyId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"salesInvoice_id"+"}", _neturl.QueryEscape(parameterToString(salesInvoiceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	// body params
	localVarPostBody = &uNKNOWNBASETYPE
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
