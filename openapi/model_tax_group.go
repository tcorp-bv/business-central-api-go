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

// TaxGroup struct for TaxGroup
type TaxGroup struct {
	// The id property for the Dynamics 365 Business Central taxGroup entity
	Id string `json:"id,omitempty"`
	// The code property for the Dynamics 365 Business Central taxGroup entity
	Code *string `json:"code,omitempty"`
	// The displayName property for the Dynamics 365 Business Central taxGroup entity
	DisplayName *string `json:"displayName,omitempty"`
	// The taxType property for the Dynamics 365 Business Central taxGroup entity
	TaxType *string `json:"taxType,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central taxGroup entity
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}
