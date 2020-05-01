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
	"time"
)

// BalanceSheet struct for BalanceSheet
type BalanceSheet struct {
	// The lineNumber property for the Dynamics 365 Business Central balanceSheet entity
	LineNumber int32 `json:"lineNumber,omitempty"`
	// The display property for the Dynamics 365 Business Central balanceSheet entity
	Display *string `json:"display,omitempty"`
	// The balance property for the Dynamics 365 Business Central balanceSheet entity
	Balance *float32 `json:"balance,omitempty"`
	// The lineType property for the Dynamics 365 Business Central balanceSheet entity
	LineType *string `json:"lineType,omitempty"`
	// The indentation property for the Dynamics 365 Business Central balanceSheet entity
	Indentation *int32 `json:"indentation,omitempty"`
	// The dateFilter property for the Dynamics 365 Business Central balanceSheet entity
	DateFilter *time.Time `json:"dateFilter,omitempty"`
}
