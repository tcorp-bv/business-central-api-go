# {{classname}}

All URIs are relative to *https://api.businesscentral.dynamics.com/v2.0/sandbox/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPdfDocument**](PdfDocumentApi.md#GetPdfDocument) | **Get** /companies({company_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**GetPdfDocumentForPurchaseInvoice**](PdfDocumentApi.md#GetPdfDocumentForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**GetPdfDocumentForSalesCreditMemo**](PdfDocumentApi.md#GetPdfDocumentForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**GetPdfDocumentForSalesInvoice**](PdfDocumentApi.md#GetPdfDocumentForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**GetPdfDocumentForSalesQuote**](PdfDocumentApi.md#GetPdfDocumentForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/pdfDocument({pdfDocument_id}) | Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.
[**ListPdfDocument**](PdfDocumentApi.md#ListPdfDocument) | **Get** /companies({company_id})/pdfDocument | Returns a list of pdfDocument
[**ListPdfDocumentForPurchaseInvoice**](PdfDocumentApi.md#ListPdfDocumentForPurchaseInvoice) | **Get** /companies({company_id})/purchaseInvoices({purchaseInvoice_id})/pdfDocument | Returns a list of pdfDocument
[**ListPdfDocumentForSalesCreditMemo**](PdfDocumentApi.md#ListPdfDocumentForSalesCreditMemo) | **Get** /companies({company_id})/salesCreditMemos({salesCreditMemo_id})/pdfDocument | Returns a list of pdfDocument
[**ListPdfDocumentForSalesInvoice**](PdfDocumentApi.md#ListPdfDocumentForSalesInvoice) | **Get** /companies({company_id})/salesInvoices({salesInvoice_id})/pdfDocument | Returns a list of pdfDocument
[**ListPdfDocumentForSalesQuote**](PdfDocumentApi.md#ListPdfDocumentForSalesQuote) | **Get** /companies({company_id})/salesQuotes({salesQuote_id})/pdfDocument | Returns a list of pdfDocument

# **GetPdfDocument**
> PdfDocument GetPdfDocument(ctx, companyId, pdfDocumentId, optional)
Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***PdfDocumentApiGetPdfDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiGetPdfDocumentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfDocumentForPurchaseInvoice**
> PdfDocument GetPdfDocumentForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, pdfDocumentId, optional)
Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***PdfDocumentApiGetPdfDocumentForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiGetPdfDocumentForPurchaseInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfDocumentForSalesCreditMemo**
> PdfDocument GetPdfDocumentForSalesCreditMemo(ctx, companyId, salesCreditMemoId, pdfDocumentId, optional)
Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***PdfDocumentApiGetPdfDocumentForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiGetPdfDocumentForSalesCreditMemoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfDocumentForSalesInvoice**
> PdfDocument GetPdfDocumentForSalesInvoice(ctx, companyId, salesInvoiceId, pdfDocumentId, optional)
Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***PdfDocumentApiGetPdfDocumentForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiGetPdfDocumentForSalesInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfDocumentForSalesQuote**
> PdfDocument GetPdfDocumentForSalesQuote(ctx, companyId, salesQuoteId, pdfDocumentId, optional)
Retrieve the properties and relationships of an object of type pdfDocument for Dynamics 365 Business Central.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
  **pdfDocumentId** | [**string**](.md)| id for pdfDocument | 
 **optional** | ***PdfDocumentApiGetPdfDocumentForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiGetPdfDocumentForSalesQuoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**PdfDocument**](pdfDocument.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocument**
> InlineResponse20010 ListPdfDocument(ctx, companyId, optional)
Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
 **optional** | ***PdfDocumentApiListPdfDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiListPdfDocumentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocumentForPurchaseInvoice**
> InlineResponse20010 ListPdfDocumentForPurchaseInvoice(ctx, companyId, purchaseInvoiceId, optional)
Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **purchaseInvoiceId** | [**string**](.md)| id for purchaseInvoice | 
 **optional** | ***PdfDocumentApiListPdfDocumentForPurchaseInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiListPdfDocumentForPurchaseInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocumentForSalesCreditMemo**
> InlineResponse20010 ListPdfDocumentForSalesCreditMemo(ctx, companyId, salesCreditMemoId, optional)
Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesCreditMemoId** | [**string**](.md)| id for salesCreditMemo | 
 **optional** | ***PdfDocumentApiListPdfDocumentForSalesCreditMemoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiListPdfDocumentForSalesCreditMemoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocumentForSalesInvoice**
> InlineResponse20010 ListPdfDocumentForSalesInvoice(ctx, companyId, salesInvoiceId, optional)
Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesInvoiceId** | [**string**](.md)| id for salesInvoice | 
 **optional** | ***PdfDocumentApiListPdfDocumentForSalesInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiListPdfDocumentForSalesInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPdfDocumentForSalesQuote**
> InlineResponse20010 ListPdfDocumentForSalesQuote(ctx, companyId, salesQuoteId, optional)
Returns a list of pdfDocument

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **companyId** | [**string**](.md)| id for company | 
  **salesQuoteId** | [**string**](.md)| id for salesQuote | 
 **optional** | ***PdfDocumentApiListPdfDocumentForSalesQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PdfDocumentApiListPdfDocumentForSalesQuoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **optional.Int32**| Number of items to return from the top of the list | 
 **skip** | **optional.Int32**| Number of items to skip from the list | 
 **limit** | **optional.Int32**| Number of items to return from the list | 
 **filter** | **optional.String**| Filtering expression | 
 **select_** | [**optional.Interface of []string**](string.md)| Selected properties to be retrieved | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth](../README.md#oAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

