# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIncomeStatement**](IncomeStatementApi.md#GetIncomeStatement) | **Get** /companies({company_id})/incomeStatement({incomeStatement_lineNumber}) | Retrieve the properties and relationships of an object of type incomeStatement for Dynamics 365 Business Central.
[**ListIncomeStatement**](IncomeStatementApi.md#ListIncomeStatement) | **Get** /companies({company_id})/incomeStatement | Returns a list of incomeStatement

# **GetIncomeStatement**
> IncomeStatement GetIncomeStatement(ctx, companyId, incomeStatementLineNumber, optional)
Retrieve the properties and relationships of an object of type incomeStatement for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **incomeStatementLineNumber** | **int32**| lineNumber for incomeStatement | 
 **optional** | ***IncomeStatementApiGetIncomeStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IncomeStatementApiGetIncomeStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**IncomeStatement**](incomeStatement.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIncomeStatement**
> InlineResponse20039 ListIncomeStatement(ctx, companyId, optional)
Returns a list of incomeStatement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***IncomeStatementApiListIncomeStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IncomeStatementApiListIncomeStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

