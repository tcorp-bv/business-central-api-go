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

type CashFlowStatement struct {
	// The lineNumber property for the Dynamics 365 Business Central cashFlowStatement entity
	LineNumber int32 `json:"lineNumber,omitempty"`
	// The display property for the Dynamics 365 Business Central cashFlowStatement entity
	Display string `json:"display,omitempty"`
	// The netChange property for the Dynamics 365 Business Central cashFlowStatement entity
	NetChange float64 `json:"netChange,omitempty"`
	// The lineType property for the Dynamics 365 Business Central cashFlowStatement entity
	LineType string `json:"lineType,omitempty"`
	// The indentation property for the Dynamics 365 Business Central cashFlowStatement entity
	Indentation int32 `json:"indentation,omitempty"`
	// The dateFilter property for the Dynamics 365 Business Central cashFlowStatement entity
	DateFilter time.Time `json:"dateFilter,omitempty"`
}
