# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePicture**](PictureApi.md#DeletePicture) | **Delete** /companies({company_id})/picture({picture_id}) | Deletes an object of type picture in Dynamics 365 Business Central
[**DeletePictureForCustomer**](PictureApi.md#DeletePictureForCustomer) | **Delete** /companies({company_id})/customers({customer_id})/picture({picture_id}) | Deletes an object of type picture in Dynamics 365 Business Central
[**DeletePictureForEmployee**](PictureApi.md#DeletePictureForEmployee) | **Delete** /companies({company_id})/employees({employee_id})/picture({picture_id}) | Deletes an object of type picture in Dynamics 365 Business Central
[**DeletePictureForItem**](PictureApi.md#DeletePictureForItem) | **Delete** /companies({company_id})/items({item_id})/picture({picture_id}) | Deletes an object of type picture in Dynamics 365 Business Central
[**DeletePictureForVendor**](PictureApi.md#DeletePictureForVendor) | **Delete** /companies({company_id})/vendors({vendor_id})/picture({picture_id}) | Deletes an object of type picture in Dynamics 365 Business Central
[**GetPicture**](PictureApi.md#GetPicture) | **Get** /companies({company_id})/picture({picture_id}) | Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.
[**GetPictureForCustomer**](PictureApi.md#GetPictureForCustomer) | **Get** /companies({company_id})/customers({customer_id})/picture({picture_id}) | Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.
[**GetPictureForEmployee**](PictureApi.md#GetPictureForEmployee) | **Get** /companies({company_id})/employees({employee_id})/picture({picture_id}) | Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.
[**GetPictureForItem**](PictureApi.md#GetPictureForItem) | **Get** /companies({company_id})/items({item_id})/picture({picture_id}) | Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.
[**GetPictureForVendor**](PictureApi.md#GetPictureForVendor) | **Get** /companies({company_id})/vendors({vendor_id})/picture({picture_id}) | Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.
[**ListPicture**](PictureApi.md#ListPicture) | **Get** /companies({company_id})/picture | Returns a list of picture
[**ListPictureForCustomer**](PictureApi.md#ListPictureForCustomer) | **Get** /companies({company_id})/customers({customer_id})/picture | Returns a list of picture
[**ListPictureForEmployee**](PictureApi.md#ListPictureForEmployee) | **Get** /companies({company_id})/employees({employee_id})/picture | Returns a list of picture
[**ListPictureForItem**](PictureApi.md#ListPictureForItem) | **Get** /companies({company_id})/items({item_id})/picture | Returns a list of picture
[**ListPictureForVendor**](PictureApi.md#ListPictureForVendor) | **Get** /companies({company_id})/vendors({vendor_id})/picture | Returns a list of picture
[**PatchPicture**](PictureApi.md#PatchPicture) | **Patch** /companies({company_id})/picture({picture_id}) | Updates an object of type picture in Dynamics 365 Business Central
[**PatchPictureForCustomer**](PictureApi.md#PatchPictureForCustomer) | **Patch** /companies({company_id})/customers({customer_id})/picture({picture_id}) | Updates an object of type picture in Dynamics 365 Business Central
[**PatchPictureForEmployee**](PictureApi.md#PatchPictureForEmployee) | **Patch** /companies({company_id})/employees({employee_id})/picture({picture_id}) | Updates an object of type picture in Dynamics 365 Business Central
[**PatchPictureForItem**](PictureApi.md#PatchPictureForItem) | **Patch** /companies({company_id})/items({item_id})/picture({picture_id}) | Updates an object of type picture in Dynamics 365 Business Central
[**PatchPictureForVendor**](PictureApi.md#PatchPictureForVendor) | **Patch** /companies({company_id})/vendors({vendor_id})/picture({picture_id}) | Updates an object of type picture in Dynamics 365 Business Central

# **DeletePicture**
> DeletePicture(ctx, companyId, pictureId)
Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePictureForCustomer**
> DeletePictureForCustomer(ctx, companyId, customerId, pictureId)
Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePictureForEmployee**
> DeletePictureForEmployee(ctx, companyId, employeeId, pictureId)
Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePictureForItem**
> DeletePictureForItem(ctx, companyId, itemId, pictureId)
Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePictureForVendor**
> DeletePictureForVendor(ctx, companyId, vendorId, pictureId)
Deletes an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **pictureId** | [**string**](.md)| id for picture | 

### Return type

 (empty response body)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPicture**
> Picture GetPicture(ctx, companyId, pictureId, optional)
Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***PictureApiGetPictureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiGetPictureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPictureForCustomer**
> Picture GetPictureForCustomer(ctx, companyId, customerId, pictureId, optional)
Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***PictureApiGetPictureForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiGetPictureForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPictureForEmployee**
> Picture GetPictureForEmployee(ctx, companyId, employeeId, pictureId, optional)
Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***PictureApiGetPictureForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiGetPictureForEmployeeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPictureForItem**
> Picture GetPictureForItem(ctx, companyId, itemId, pictureId, optional)
Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***PictureApiGetPictureForItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiGetPictureForItemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPictureForVendor**
> Picture GetPictureForVendor(ctx, companyId, vendorId, pictureId, optional)
Retrieve the properties and relationships of an object of type picture for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **pictureId** | [**string**](.md)| id for picture | 
 **optional** | ***PictureApiGetPictureForVendorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiGetPictureForVendorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPicture**
> InlineResponse20017 ListPicture(ctx, companyId, optional)
Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***PictureApiListPictureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiListPictureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPictureForCustomer**
> InlineResponse20017 ListPictureForCustomer(ctx, companyId, customerId, optional)
Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
 **optional** | ***PictureApiListPictureForCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiListPictureForCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPictureForEmployee**
> InlineResponse20017 ListPictureForEmployee(ctx, companyId, employeeId, optional)
Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
 **optional** | ***PictureApiListPictureForEmployeeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiListPictureForEmployeeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPictureForItem**
> InlineResponse20017 ListPictureForItem(ctx, companyId, itemId, optional)
Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
 **optional** | ***PictureApiListPictureForItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiListPictureForItemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPictureForVendor**
> InlineResponse20017 ListPictureForVendor(ctx, companyId, vendorId, optional)
Returns a list of picture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
 **optional** | ***PictureApiListPictureForVendorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PictureApiListPictureForVendorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPicture**
> Picture PatchPicture(ctx, body, companyId, pictureId, contentType, ifMatch)
Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPictureForCustomer**
> Picture PatchPictureForCustomer(ctx, body, companyId, customerId, pictureId, contentType, ifMatch)
Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **customerId** | [**string**](.md)| id for customer | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPictureForEmployee**
> Picture PatchPictureForEmployee(ctx, body, companyId, employeeId, pictureId, contentType, ifMatch)
Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **employeeId** | [**string**](.md)| id for employee | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPictureForItem**
> Picture PatchPictureForItem(ctx, body, companyId, itemId, pictureId, contentType, ifMatch)
Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **itemId** | [**string**](.md)| id for item | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPictureForVendor**
> Picture PatchPictureForVendor(ctx, body, companyId, vendorId, pictureId, contentType, ifMatch)
Updates an object of type picture in Dynamics 365 Business Central

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
  **companyId** | [**string**](.md)| id for company | 
  **vendorId** | [**string**](.md)| id for vendor | 
  **pictureId** | [**string**](.md)| id for picture | 
  **contentType** | **string**| application/json | 
  **ifMatch** | **string**| Required. When this request header is included and the eTag provided does not match the current tag on the entity, this will not be updated. | 

### Return type

[**Picture**](picture.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

