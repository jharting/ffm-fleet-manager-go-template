/*
 * Dinosaur Service Fleet Manager
 *
 * Dinosaur Service Fleet Manager is a Rest API to manage Dinosaur instances.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// InstantQuery struct for InstantQuery
type InstantQuery struct {
	Metric    map[string]string `json:"metric,omitempty"`
	Timestamp int64             `json:"timestamp,omitempty"`
	Value     float64           `json:"value"`
}