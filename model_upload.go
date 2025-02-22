/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stravago

type Upload struct {
	// The unique identifier of the upload
	Id int64 `json:"id,omitempty"`
	// The unique identifier of the upload in string format
	IdStr string `json:"id_str,omitempty"`
	// The external identifier of the upload
	ExternalId string `json:"external_id,omitempty"`
	// The error associated with this upload
	Error_ string `json:"error,omitempty"`
	// The status of this upload
	Status string `json:"status,omitempty"`
	// The identifier of the activity this upload resulted into
	ActivityId int64 `json:"activity_id,omitempty"`
}
