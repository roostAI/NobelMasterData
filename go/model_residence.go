/*
 * Nobel Prize Master Data
 *
 * Information about the Nobel Prizes and the Nobel Prize Laureates
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Residence struct {

	City *Translation `json:"city,omitempty"`

	Country *Translation `json:"country,omitempty"`

	CityNow *TranslationWithSameas `json:"cityNow,omitempty"`

	CountryNow *TranslationWithSameas `json:"countryNow,omitempty"`

	LocationString *Translation `json:"locationString,omitempty"`
}
