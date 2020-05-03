# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRetainedEarningsStatement**](RetainedEarningsStatementApi.md#GetRetainedEarningsStatement) | **Get** /companies({company_id})/retainedEarningsStatement({retainedEarningsStatement_lineNumber}) | Retrieve the properties and relationships of an object of type retainedEarningsStatement for Dynamics 365 Business Central.
[**ListRetainedEarningsStatement**](RetainedEarningsStatementApi.md#ListRetainedEarningsStatement) | **Get** /companies({company_id})/retainedEarningsStatement | Returns a list of retainedEarningsStatement

# **GetRetainedEarningsStatement**
> RetainedEarningsStatement GetRetainedEarningsStatement(ctx, companyId, retainedEarningsStatementLineNumber, optional)
Retrieve the properties and relationships of an object of type retainedEarningsStatement for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **retainedEarningsStatementLineNumber** | **int32**| lineNumber for retainedEarningsStatement | 
 **optional** | ***RetainedEarningsStatementApiGetRetainedEarningsStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetainedEarningsStatementApiGetRetainedEarningsStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**RetainedEarningsStatement**](retainedEarningsStatement.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRetainedEarningsStatement**
> InlineResponse20033 ListRetainedEarningsStatement(ctx, companyId, optional)
Returns a list of retainedEarningsStatement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***RetainedEarningsStatementApiListRetainedEarningsStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetainedEarningsStatementApiListRetainedEarningsStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

