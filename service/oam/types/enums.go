// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeAwsCloudwatchMetric ResourceType = "AWS::CloudWatch::Metric"
	ResourceTypeAwsLogsLoggroup     ResourceType = "AWS::Logs::LogGroup"
	ResourceTypeAwsXrayTrace        ResourceType = "AWS::XRay::Trace"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"AWS::CloudWatch::Metric",
		"AWS::Logs::LogGroup",
		"AWS::XRay::Trace",
	}
}
