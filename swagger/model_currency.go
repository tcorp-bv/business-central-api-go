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

type Currency struct {
	// The id property for the Dynamics 365 Business Central currency entity
	Id string `json:"id,omitempty"`
	// The code property for the Dynamics 365 Business Central currency entity
	Code string `json:"code,omitempty"`
	// The displayName property for the Dynamics 365 Business Central currency entity
	DisplayName string `json:"displayName,omitempty"`
	// The symbol property for the Dynamics 365 Business Central currency entity
	Symbol string `json:"symbol,omitempty"`
	// The amountDecimalPlaces property for the Dynamics 365 Business Central currency entity
	AmountDecimalPlaces string `json:"amountDecimalPlaces,omitempty"`
	// The amountRoundingPrecision property for the Dynamics 365 Business Central currency entity
	AmountRoundingPrecision float64 `json:"amountRoundingPrecision,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central currency entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
}