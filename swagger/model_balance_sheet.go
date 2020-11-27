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

type BalanceSheet struct {
	// The balance property for the Dynamics 365 Business Central balanceSheet entity
	Balance float64 `json:"balance,omitempty"`
	// The dateFilter property for the Dynamics 365 Business Central balanceSheet entity
	DateFilter time.Time `json:"dateFilter,omitempty"`
	// The display property for the Dynamics 365 Business Central balanceSheet entity
	Display string `json:"display,omitempty"`
	// The indentation property for the Dynamics 365 Business Central balanceSheet entity
	Indentation int32 `json:"indentation,omitempty"`
	// The lineNumber property for the Dynamics 365 Business Central balanceSheet entity
	LineNumber int32 `json:"lineNumber,omitempty"`
	// The lineType property for the Dynamics 365 Business Central balanceSheet entity
	LineType string `json:"lineType,omitempty"`
}
