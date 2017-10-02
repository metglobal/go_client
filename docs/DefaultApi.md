# \DefaultApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailabilityProductCodeGet**](DefaultApi.md#AvailabilityProductCodeGet) | **Get** /availability/{product_code} | Availability with Product Code
[**BookProvisionCodePost**](DefaultApi.md#BookProvisionCodePost) | **Post** /book/{provision_code} | Book with Provision Code
[**BookingsBookingCodeGet**](DefaultApi.md#BookingsBookingCodeGet) | **Get** /bookings/{booking_code} | Get Booking Detail
[**BookingsGet**](DefaultApi.md#BookingsGet) | **Get** /bookings/ | Get Booking List
[**CancelBookingCodePost**](DefaultApi.md#CancelBookingCodePost) | **Post** /cancel/{booking_code} | Cancel Booking with Booking Code
[**HotelAvailabilityGet**](DefaultApi.md#HotelAvailabilityGet) | **Get** /hotel-availability/ | Hotel Availability with Hotel Code and Search Code
[**ProvisionProductCodePost**](DefaultApi.md#ProvisionProductCodePost) | **Post** /provision/{product_code} | Provision with Product Code
[**SearchPost**](DefaultApi.md#SearchPost) | **Post** /search/ | Search with Hotel Code(Hotel Code List) or Destination Code or Geolocation


# **AvailabilityProductCodeGet**
> AvailabilityResponse AvailabilityProductCodeGet(ctx, productCode)
Availability with Product Code

Check Availability of Selected Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **productCode** | **string**| product code that returned in Search(or Hotel Availability) Response | 

### Return type

[**AvailabilityResponse**](AvailabilityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookProvisionCodePost**
> BookResponse BookProvisionCodePost(ctx, name, provisionCode)
Book with Provision Code

Returns Book Response

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **name** | [**[]string**](string.md)| A person&#39;s name. | 
  **provisionCode** | **string**|  | 

### Return type

[**BookResponse**](BookResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookingsBookingCodeGet**
> BookResponse BookingsBookingCodeGet(ctx, bookingCode)
Get Booking Detail

Returns past booking(s) data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **bookingCode** | **string**| This is the code that taken from the response of bookings request | 

### Return type

[**BookResponse**](BookResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BookingsGet**
> BookingListResponse BookingsGet(ctx, optional)
Get Booking List

Returns past booking(s) data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromDate** | **string**| This is the booking date for filtering the bookings from the from_date(YYYY-MM-DD). | 
 **toDate** | **string**| This is the booking date for filtering the bookings until the to_date(YYYY-MM-DD). | 
 **format** | **string**| Only JSON supported | 

### Return type

[**BookingListResponse**](BookingListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelBookingCodePost**
> CancelResponse CancelBookingCodePost(ctx, bookingCode)
Cancel Booking with Booking Code

Cancel the Booking

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **bookingCode** | **string**| Booking Code that returned in Book Response | 

### Return type

[**CancelResponse**](CancelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HotelAvailabilityGet**
> HotelAvailabilityResponse HotelAvailabilityGet(ctx, searchCode, hotelCode)
Hotel Availability with Hotel Code and Search Code

Check Availability of Selected Hotel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **searchCode** | **string**| search code that returned in search response | 
  **hotelCode** | **string**| requested hotel code | 

### Return type

[**HotelAvailabilityResponse**](HotelAvailabilityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvisionProductCodePost**
> ProvisionResponse ProvisionProductCodePost(ctx, productCode)
Provision with Product Code

Provision of Selected Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **productCode** | **string**| product code that returned in Search(or Hotel Availability) Response | 

### Return type

[**ProvisionResponse**](ProvisionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchPost**
> SearchResponse SearchPost(ctx, pax, checkin, checkout, clientNationality, currency, optional)
Search with Hotel Code(Hotel Code List) or Destination Code or Geolocation

Returns list of products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **pax** | [**[]string**](string.md)| Number of pax | 
  **checkin** | **string**| Checkin | 
  **checkout** | **string**| Checkout | 
  **clientNationality** | **string**| Client Nationality | 
  **currency** | **string**| Currency (Supported Currencies USD, EUR, GBP, TRY) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pax** | [**[]string**](string.md)| Number of pax | 
 **checkin** | **string**| Checkin | 
 **checkout** | **string**| Checkout | 
 **clientNationality** | **string**| Client Nationality | 
 **currency** | **string**| Currency (Supported Currencies USD, EUR, GBP, TRY) | 
 **hotelCode** | **string**| Requested Hotel Code | 
 **destinationCode** | **string**| Requested Destination Code | 
 **lat** | **string**| Requested Latitude Code(lat, lon and radius should be given together) | 
 **lon** | **string**| Requested Longitude Code(lat, lon and radius should be given together) | 
 **radius** | **string**| Requested Radius Code(lat, lon and radius should be given together) | 
 **maxProduct** | **int32**| Max Product | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

