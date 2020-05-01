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
	"os"
)

// Picture struct for Picture
type Picture struct {
	// The id property for the Dynamics 365 Business Central picture entity
	Id string `json:"id,omitempty"`
	// The width property for the Dynamics 365 Business Central picture entity
	Width *int32 `json:"width,omitempty"`
	// The height property for the Dynamics 365 Business Central picture entity
	Height *int32 `json:"height,omitempty"`
	// The contentType property for the Dynamics 365 Business Central picture entity
	ContentType *string `json:"contentType,omitempty"`
	// The content property for the Dynamics 365 Business Central picture entity
	Content **os.File `json:"content,omitempty"`
}
