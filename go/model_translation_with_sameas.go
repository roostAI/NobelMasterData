/*
 * Nobel Prize Master Data
 *
 * Information about the Nobel Prizes and the Nobel Prize Laureates
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TranslationWithSameas struct {

	En string `json:"en,omitempty"`

	Se string `json:"se,omitempty"`

	No string `json:"no,omitempty"`

	SameAs []Urls `json:"sameAs,omitempty"`
}
