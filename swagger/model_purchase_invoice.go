/*
 * Dynamics 365 Business Central
 *
 * Business Central Standard APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type PurchaseInvoice struct {
	// The id property for the Dynamics 365 Business Central purchaseInvoice entity
	Id string `json:"id,omitempty"`
	// The number property for the Dynamics 365 Business Central purchaseInvoice entity
	Number string `json:"number,omitempty"`
	// The invoiceDate property for the Dynamics 365 Business Central purchaseInvoice entity
	InvoiceDate time.Time `json:"invoiceDate,omitempty"`
	// The dueDate property for the Dynamics 365 Business Central purchaseInvoice entity
	DueDate time.Time `json:"dueDate,omitempty"`
	// The vendorInvoiceNumber property for the Dynamics 365 Business Central purchaseInvoice entity
	VendorInvoiceNumber string `json:"vendorInvoiceNumber,omitempty"`
	// The vendorId property for the Dynamics 365 Business Central purchaseInvoice entity
	VendorId string `json:"vendorId,omitempty"`
	// The vendorNumber property for the Dynamics 365 Business Central purchaseInvoice entity
	VendorNumber string `json:"vendorNumber,omitempty"`
	// The vendorName property for the Dynamics 365 Business Central purchaseInvoice entity
	VendorName string `json:"vendorName,omitempty"`
	// The payToName property for the Dynamics 365 Business Central purchaseInvoice entity
	PayToName string `json:"payToName,omitempty"`
	// The payToContact property for the Dynamics 365 Business Central purchaseInvoice entity
	PayToContact string `json:"payToContact,omitempty"`
	// The payToVendorId property for the Dynamics 365 Business Central purchaseInvoice entity
	PayToVendorId string `json:"payToVendorId,omitempty"`
	// The payToVendorNumber property for the Dynamics 365 Business Central purchaseInvoice entity
	PayToVendorNumber string `json:"payToVendorNumber,omitempty"`
	// The shipToName property for the Dynamics 365 Business Central purchaseInvoice entity
	ShipToName string `json:"shipToName,omitempty"`
	// The shipToContact property for the Dynamics 365 Business Central purchaseInvoice entity
	ShipToContact string `json:"shipToContact,omitempty"`
	BuyFromAddress *Postaladdresstype `json:"buyFromAddress,omitempty"`
	PayToAddress *Postaladdresstype `json:"payToAddress,omitempty"`
	ShipToAddress *Postaladdresstype `json:"shipToAddress,omitempty"`
	// The currencyId property for the Dynamics 365 Business Central purchaseInvoice entity
	CurrencyId string `json:"currencyId,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central purchaseInvoice entity
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The pricesIncludeTax property for the Dynamics 365 Business Central purchaseInvoice entity
	PricesIncludeTax bool `json:"pricesIncludeTax,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central purchaseInvoice entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The discountAppliedBeforeTax property for the Dynamics 365 Business Central purchaseInvoice entity
	DiscountAppliedBeforeTax bool `json:"discountAppliedBeforeTax,omitempty"`
	// The totalAmountExcludingTax property for the Dynamics 365 Business Central purchaseInvoice entity
	TotalAmountExcludingTax float64 `json:"totalAmountExcludingTax,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central purchaseInvoice entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The totalAmountIncludingTax property for the Dynamics 365 Business Central purchaseInvoice entity
	TotalAmountIncludingTax float64 `json:"totalAmountIncludingTax,omitempty"`
	// The status property for the Dynamics 365 Business Central purchaseInvoice entity
	Status string `json:"status,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central purchaseInvoice entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	PurchaseInvoiceLines []PurchaseInvoiceLine `json:"purchaseInvoiceLines,omitempty"`
	PdfDocument []PdfDocument `json:"pdfDocument,omitempty"`
	Vendor *Vendor `json:"vendor,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
}
