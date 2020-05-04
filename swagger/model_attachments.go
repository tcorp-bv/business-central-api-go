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

type Attachments struct {
	// The byteSize property for the Dynamics 365 Business Central attachments entity
	ByteSize int32 `json:"byteSize,omitempty"`
	// The content property for the Dynamics 365 Business Central attachments entity
	Content string `json:"content,omitempty"`
	// The fileName property for the Dynamics 365 Business Central attachments entity
	FileName string `json:"fileName,omitempty"`
	// The id property for the Dynamics 365 Business Central attachments entity
	Id string `json:"id,omitempty"`
	// The lastModifiedDateTime property for the Dynamics 365 Business Central attachments entity
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	// The parentId property for the Dynamics 365 Business Central attachments entity
	ParentId string `json:"parentId,omitempty"`
}
