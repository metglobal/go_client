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

type BookResponse struct {

	Code string `json:"code"`

	CreatedAt string `json:"created_at"`

	Checkin string `json:"checkin"`

	Checkout string `json:"checkout"`

	HotelCode string `json:"hotel_code"`

	DestinationCode string `json:"destination_code"`

	ClientNationality string `json:"client_nationality"`

	PayAtHotel bool `json:"pay_at_hotel"`

	Currency string `json:"currency"`

	MealtypeCode string `json:"mealtype_code"`

	Nonrefundable string `json:"nonrefundable"`

	View bool `json:"view,omitempty"`

	Price string `json:"price"`

	Policies []Policy `json:"policies"`

	Rooms []Room `json:"rooms"`

	Status string `json:"status"`

	ConfirmationNumbers []BookResponseConfirmationNumbers `json:"confirmation_numbers"`

	HotelPaymentInfo []BookResponseHotelPaymentInfo `json:"hotel_payment_info"`

	MinimumSellingPrice string `json:"minimum_selling_price,omitempty"`

	SpecialRequest string `json:"special_request,omitempty"`
}
