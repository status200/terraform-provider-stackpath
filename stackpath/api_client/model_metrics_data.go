/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

// MetricsData The data points in a metrics collection
type MetricsData struct {
	Matrix InlineResponse2004DataMatrix `json:"matrix,omitempty"`
	Vector InlineResponse2004DataVector `json:"vector,omitempty"`
}