// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	Gets the aggregated profile of a profiling group for a specified time range.
//
// Amazon CodeGuru Profiler collects posted agent profiles for a profiling group
// into aggregated profiles.
//
// Because aggregated profiles expire over time GetProfile is not idempotent.
//
// Specify the time range for the requested aggregated profile using 1 or 2 of the
// following parameters: startTime , endTime , period . The maximum time range
// allowed is 7 days. If you specify all 3 parameters, an exception is thrown. If
// you specify only period , the latest aggregated profile is returned.
//
// Aggregated profiles are available with aggregation periods of 5 minutes, 1
// hour, and 1 day, aligned to UTC. The aggregation period of an aggregated profile
// determines how long it is retained. For more information, see [AggregatedProfileTime]
// AggregatedProfileTime . The aggregated profile's aggregation period determines
// how long
//
// it is retained by CodeGuru Profiler.
//
//   - If the aggregation period is 5 minutes, the aggregated profile is retained
//     for 15 days.
//
//   - If the aggregation period is 1 hour, the aggregated profile is retained for
//     60 days.
//
//   - If the aggregation period is 1 day, the aggregated profile is retained for
//     3 years.
//
// There are two use cases for calling GetProfile .
//
//   - If you want to return an aggregated profile that already exists, use [ListProfileTimes]
//     ListProfileTimes to view the time ranges of existing aggregated profiles. Use
//     them in a GetProfile request to return a specific, existing aggregated
//     profile.
//
//   - If you want to return an aggregated profile for a time range that doesn't
//     align with an existing aggregated profile, then CodeGuru Profiler makes a best
//     effort to combine existing aggregated profiles from the requested time range and
//     return them as one aggregated profile.
//
// If aggregated profiles do not exist for the full time range requested, then
//
//	aggregated profiles for a smaller time range are returned. For example, if the
//	requested time range is from 00:00 to 00:20, and the existing aggregated
//	profiles are from 00:15 and 00:25, then the aggregated profiles from 00:15 to
//	00:20 are returned.
//
// [ListProfileTimes]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ListProfileTimes.html
// [AggregatedProfileTime]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_AggregatedProfileTime.html
func (c *Client) GetProfile(ctx context.Context, params *GetProfileInput, optFns ...func(*Options)) (*GetProfileOutput, error) {
	if params == nil {
		params = &GetProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProfile", params, optFns, c.addOperationGetProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the getProfileRequest.
type GetProfileInput struct {

	// The name of the profiling group to get.
	//
	// This member is required.
	ProfilingGroupName *string

	//  The format of the returned profiling data. The format maps to the Accept and
	// Content-Type headers of the HTTP request. You can specify one of the following:
	// or the default .
	//
	//   - application/json — standard JSON format
	//
	//   - application/x-amzn-ion — the Amazon Ion data format. For more information,
	//   see [Amazon Ion].
	//
	// [Amazon Ion]: http://amzn.github.io/ion-docs/
	Accept *string

	//  The end time of the requested profile. Specify using the ISO 8601 format. For
	// example, 2020-06-01T13:15:02.001Z represents 1 millisecond past June 1, 2020
	// 1:15:02 PM UTC.
	//
	// If you specify endTime , then you must also specify period or startTime , but
	// not both.
	EndTime *time.Time

	//  The maximum depth of the stacks in the code that is represented in the
	// aggregated profile. For example, if CodeGuru Profiler finds a method A , which
	// calls method B , which calls method C , which calls method D , then the depth is
	// 4. If the maxDepth is set to 2, then the aggregated profile contains
	// representations of methods A and B .
	MaxDepth *int32

	//  Used with startTime or endTime to specify the time range for the returned
	// aggregated profile. Specify using the ISO 8601 format. For example, P1DT1H1M1S .
	//
	// To get the latest aggregated profile, specify only period .
	Period *string

	// The start time of the profile to get. Specify using the ISO 8601 format. For
	// example, 2020-06-01T13:15:02.001Z represents 1 millisecond past June 1, 2020
	// 1:15:02 PM UTC.
	//
	// If you specify startTime , then you must also specify period or endTime , but
	// not both.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// The structure representing the getProfileResponse.
type GetProfileOutput struct {

	// The content type of the profile in the payload. It is either application/json
	// or the default application/x-amzn-ion .
	//
	// This member is required.
	ContentType *string

	// Information about the profile.
	//
	// This member is required.
	Profile []byte

	// The content encoding of the profile.
	ContentEncoding *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetProfile"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addOpGetProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetProfile(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetProfile",
	}
}
