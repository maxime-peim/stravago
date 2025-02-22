/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stravago

type Split struct {
	// The average speed of this split, in meters per second
	AverageSpeed float32 `json:"average_speed,omitempty"`
	// The distance of this split, in meters
	Distance float32 `json:"distance,omitempty"`
	// The elapsed time of this split, in seconds
	ElapsedTime int32 `json:"elapsed_time,omitempty"`
	// The elevation difference of this split, in meters
	ElevationDifference float32 `json:"elevation_difference,omitempty"`
	// The pacing zone of this split
	PaceZone int32 `json:"pace_zone,omitempty"`
	// The moving time of this split, in seconds
	MovingTime int32 `json:"moving_time,omitempty"`
	// N/A
	Split int32 `json:"split,omitempty"`
}
