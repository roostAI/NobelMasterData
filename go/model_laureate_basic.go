/*
 * Nobel Prize Master Data
 *
 * Information about the Nobel Prizes and the Nobel Prize Laureates
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Short information about the Laureate
type LaureateBasic struct {

	Id int32 `json:"id,omitempty"`

	Name *Translation `json:"name,omitempty"`

	Portion string `json:"portion,omitempty"`

	SortOrder string `json:"sortOrder,omitempty"`

	Motivation *Translation `json:"motivation,omitempty"`

	Links []ItemLinks `json:"links,omitempty"`
}