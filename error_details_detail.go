/* 
 * Hotelspro Api Client
 *
 * Hotelspro Api Client
 *
 * OpenAPI spec version: 2.0.0
 * Contact: clientintegration@hotelspro.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package hotelspro_client

type ErrorDetailsDetail struct {

	ClientIp string `json:"client_ip,omitempty"`

	Channel string `json:"channel,omitempty"`

	Name []string `json:"name,omitempty"`

	CardType string `json:"card_type,omitempty"`

	CardNumber string `json:"card_number,omitempty"`

	DestinationCode string `json:"destination_code,omitempty"`

	Checkin []string `json:"checkin,omitempty"`

	Checkout []string `json:"checkout,omitempty"`

	NonFieldErrors []string `json:"non_field_errors,omitempty"`
}
