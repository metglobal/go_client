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

type Product struct {

	// product code
	Code string `json:"code,omitempty"`

	Offer bool `json:"offer,omitempty"`

	PayAtHotel bool `json:"pay_at_hotel,omitempty"`

	Nonrefundable bool `json:"nonrefundable,omitempty"`

	Price string `json:"price,omitempty"`

	Currency string `json:"currency,omitempty"`

	Rooms []Room `json:"rooms,omitempty"`

	SupportsCancellation bool `json:"supports_cancellation,omitempty"`

	HotelCurrency string `json:"hotel_currency,omitempty"`

	HotelPrice string `json:"hotel_price,omitempty"`

	MealType string `json:"meal_type,omitempty"`

	Policies []Policy `json:"policies,omitempty"`

	MinimumSellingPrice string `json:"minimum_selling_price,omitempty"`

	View bool `json:"view,omitempty"`
}
