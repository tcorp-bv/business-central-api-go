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

type SalesQuote struct {
	// The acceptedDate property for the Dynamics 365 Business Central salesQuote entity
	AcceptedDate string `json:"acceptedDate,omitempty"`
	// The billToCustomerId property for the Dynamics 365 Business Central salesQuote entity
	BillToCustomerId string `json:"billToCustomerId,omitempty"`
	// The billToCustomerNumber property for the Dynamics 365 Business Central salesQuote entity
	BillToCustomerNumber string `json:"billToCustomerNumber,omitempty"`
	// The billToName property for the Dynamics 365 Business Central salesQuote entity
	BillToName           string             `json:"billToName,omitempty"`
	BillingPostalAddress *Postaladdresstype `json:"billingPostalAddress,omitempty"`
	// The contactId property for the Dynamics 365 Business Central salesQuote entity
	ContactId string    `json:"contactId,omitempty"`
	Currency  *Currency `json:"currency,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central salesQuote entity
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The currencyId property for the Dynamics 365 Business Central salesQuote entity
	CurrencyId string    `json:"currencyId,omitempty"`
	Customer   *Customer `json:"customer,omitempty"`
	// The customerId property for the Dynamics 365 Business Central salesQuote entity
	CustomerId string `json:"customerId,omitempty"`
	// The customerName property for the Dynamics 365 Business Central salesQuote entity
	CustomerName string `json:"customerName,omitempty"`
	// The customerNumber property for the Dynamics 365 Business Central salesQuote entity
	CustomerNumber string `json:"customerNumber,omitempty"`
	// The discountAmount property for the Dynamics 365 Business Central salesQuote entity
	DiscountAmount float64 `json:"discountAmount,omitempty"`
	// The documentDate property for the Dynamics 365 Business Central salesQuote entity
	DocumentDate string `json:"documentDate,omitempty"`
	// The dueDate property for the Dynamics 365 Business Central salesQuote entity
	DueDate string `json:"dueDate,omitempty"`
	// The email property for the Dynamics 365 Business Central salesQuote entity
	Email string `json:"email,omitempty"`
	// The externalDocumentNumber property for the Dynamics 365 Business Central salesQuote entity
	ExternalDocumentNumber string `json:"externalDocumentNumber,omitempty"`
	// The id property for the Dynamics 365 Business Central salesQuote entity
	Id string `json:"id,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central salesQuote entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The number property for the Dynamics 365 Business Central salesQuote entity
	Number      string       `json:"number,omitempty"`
	PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
	// The paymentTermsId property for the Dynamics 365 Business Central salesQuote entity
	PaymentTermsId string        `json:"paymentTermsId,omitempty"`
	PdfDocument    []PdfDocument `json:"pdfDocument,omitempty"`
	// The phoneNumber property for the Dynamics 365 Business Central salesQuote entity
	PhoneNumber     string           `json:"phoneNumber,omitempty"`
	SalesQuoteLines []SalesQuoteLine `json:"salesQuoteLines,omitempty"`
	// The salesperson property for the Dynamics 365 Business Central salesQuote entity
	Salesperson          string             `json:"salesperson,omitempty"`
	SellingPostalAddress *Postaladdresstype `json:"sellingPostalAddress,omitempty"`
	// The sentDate property for the Dynamics 365 Business Central salesQuote entity
	SentDate string `json:"sentDate,omitempty"`
	// The shipToContact property for the Dynamics 365 Business Central salesQuote entity
	ShipToContact string `json:"shipToContact,omitempty"`
	// The shipToName property for the Dynamics 365 Business Central salesQuote entity
	ShipToName     string          `json:"shipToName,omitempty"`
	ShipmentMethod *ShipmentMethod `json:"shipmentMethod,omitempty"`
	// The shipmentMethodId property for the Dynamics 365 Business Central salesQuote entity
	ShipmentMethodId      string             `json:"shipmentMethodId,omitempty"`
	ShippingPostalAddress *Postaladdresstype `json:"shippingPostalAddress,omitempty"`
	// The status property for the Dynamics 365 Business Central salesQuote entity
	Status string `json:"status,omitempty"`
	// The totalAmountExcludingTax property for the Dynamics 365 Business Central salesQuote entity
	TotalAmountExcludingTax float64 `json:"totalAmountExcludingTax,omitempty"`
	// The totalAmountIncludingTax property for the Dynamics 365 Business Central salesQuote entity
	TotalAmountIncludingTax float64 `json:"totalAmountIncludingTax,omitempty"`
	// The totalTaxAmount property for the Dynamics 365 Business Central salesQuote entity
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
	// The validUntilDate property for the Dynamics 365 Business Central salesQuote entity
	ValidUntilDate string `json:"validUntilDate,omitempty"`
}
