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

type AgedAccountsPayable struct {
	// The vendorId property for the Dynamics 365 Business Central agedAccountsPayable entity
	VendorId string `json:"vendorId,omitempty"`
	// The vendorNumber property for the Dynamics 365 Business Central agedAccountsPayable entity
	VendorNumber string `json:"vendorNumber,omitempty"`
	// The name property for the Dynamics 365 Business Central agedAccountsPayable entity
	Name string `json:"name,omitempty"`
	// The currencyCode property for the Dynamics 365 Business Central agedAccountsPayable entity
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The balanceDue property for the Dynamics 365 Business Central agedAccountsPayable entity
	BalanceDue float64 `json:"balanceDue,omitempty"`
	// The currentAmount property for the Dynamics 365 Business Central agedAccountsPayable entity
	CurrentAmount float64 `json:"currentAmount,omitempty"`
	// The period1Amount property for the Dynamics 365 Business Central agedAccountsPayable entity
	Period1Amount float64 `json:"period1Amount,omitempty"`
	// The period2Amount property for the Dynamics 365 Business Central agedAccountsPayable entity
	Period2Amount float64 `json:"period2Amount,omitempty"`
	// The period3Amount property for the Dynamics 365 Business Central agedAccountsPayable entity
	Period3Amount float64 `json:"period3Amount,omitempty"`
	// The agedAsOfDate property for the Dynamics 365 Business Central agedAccountsPayable entity
	AgedAsOfDate time.Time `json:"agedAsOfDate,omitempty"`
	// The periodLengthFilter property for the Dynamics 365 Business Central agedAccountsPayable entity
	PeriodLengthFilter string `json:"periodLengthFilter,omitempty"`
}
