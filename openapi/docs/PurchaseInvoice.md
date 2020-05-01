# PurchaseInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**Number** | Pointer to **string** | The number property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**InvoiceDate** | Pointer to [**time.Time**](time.Time.md) | The invoiceDate property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**DueDate** | Pointer to [**time.Time**](time.Time.md) | The dueDate property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**VendorInvoiceNumber** | Pointer to **string** | The vendorInvoiceNumber property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**VendorId** | Pointer to **string** | The vendorId property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**VendorNumber** | Pointer to **string** | The vendorNumber property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**VendorName** | Pointer to **string** | The vendorName property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**PayToName** | Pointer to **string** | The payToName property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**PayToContact** | Pointer to **string** | The payToContact property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**PayToVendorId** | Pointer to **string** | The payToVendorId property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**PayToVendorNumber** | Pointer to **string** | The payToVendorNumber property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**ShipToName** | Pointer to **string** | The shipToName property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**ShipToContact** | Pointer to **string** | The shipToContact property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**BuyFromAddress** | [**Postaladdresstype**](postaladdresstype.md) |  | [optional] 
**PayToAddress** | [**Postaladdresstype**](postaladdresstype.md) |  | [optional] 
**ShipToAddress** | [**Postaladdresstype**](postaladdresstype.md) |  | [optional] 
**CurrencyId** | Pointer to **string** | The currencyId property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**CurrencyCode** | Pointer to **string** | The currencyCode property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**PricesIncludeTax** | Pointer to **bool** | The pricesIncludeTax property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**DiscountAmount** | Pointer to **float32** | The discountAmount property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**DiscountAppliedBeforeTax** | Pointer to **bool** | The discountAppliedBeforeTax property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**TotalAmountExcludingTax** | Pointer to **float32** | The totalAmountExcludingTax property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**TotalTaxAmount** | Pointer to **float32** | The totalTaxAmount property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**TotalAmountIncludingTax** | Pointer to **float32** | The totalAmountIncludingTax property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**Status** | Pointer to **string** | The status property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | The lastModifiedDateTime property for the Dynamics 365 Business Central purchaseInvoice entity | [optional] 
**PurchaseInvoiceLines** | Pointer to [**[]PurchaseInvoiceLine**](purchaseInvoiceLine.md) |  | [optional] 
**PdfDocument** | Pointer to [**[]PdfDocument**](pdfDocument.md) |  | [optional] 
**Vendor** | [**Vendor**](vendor.md) |  | [optional] 
**Currency** | [**Currency**](currency.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


