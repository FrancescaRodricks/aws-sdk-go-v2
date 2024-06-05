// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about all of the versions for a specified traffic policy.
//
// Traffic policy versions are listed in numerical order by VersionNumber .
func (c *Client) ListTrafficPolicyVersions(ctx context.Context, params *ListTrafficPolicyVersionsInput, optFns ...func(*Options)) (*ListTrafficPolicyVersionsOutput, error) {
	if params == nil {
		params = &ListTrafficPolicyVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrafficPolicyVersions", params, optFns, c.addOperationListTrafficPolicyVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrafficPolicyVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains the information about the request to list your
// traffic policies.
type ListTrafficPolicyVersionsInput struct {

	// Specify the value of Id of the traffic policy for which you want to list all
	// versions.
	//
	// This member is required.
	Id *string

	// The maximum number of traffic policy versions that you want Amazon Route 53 to
	// include in the response body for this request. If the specified traffic policy
	// has more than MaxItems versions, the value of IsTruncated in the response is
	// true , and the value of the TrafficPolicyVersionMarker element is the ID of the
	// first version that Route 53 will return if you submit another request.
	MaxItems *int32

	// For your first request to ListTrafficPolicyVersions , don't include the
	// TrafficPolicyVersionMarker parameter.
	//
	// If you have more traffic policy versions than the value of MaxItems ,
	// ListTrafficPolicyVersions returns only the first group of MaxItems versions. To
	// get more traffic policy versions, submit another ListTrafficPolicyVersions
	// request. For the value of TrafficPolicyVersionMarker , specify the value of
	// TrafficPolicyVersionMarker in the previous response.
	TrafficPolicyVersionMarker *string

	noSmithyDocumentSerde
}

// A complex type that contains the response information for the request.
type ListTrafficPolicyVersionsOutput struct {

	// A flag that indicates whether there are more traffic policies to be listed. If
	// the response was truncated, you can get the next group of traffic policies by
	// submitting another ListTrafficPolicyVersions request and specifying the value
	// of NextMarker in the marker parameter.
	//
	// This member is required.
	IsTruncated bool

	// The value that you specified for the maxitems parameter in the
	// ListTrafficPolicyVersions request that produced the current response.
	//
	// This member is required.
	MaxItems *int32

	// A list that contains one TrafficPolicy element for each traffic policy version
	// that is associated with the specified traffic policy.
	//
	// This member is required.
	TrafficPolicies []types.TrafficPolicy

	// If IsTruncated is true , the value of TrafficPolicyVersionMarker identifies the
	// first traffic policy that Amazon Route 53 will return if you submit another
	// request. Call ListTrafficPolicyVersions again and specify the value of
	// TrafficPolicyVersionMarker in the TrafficPolicyVersionMarker request parameter.
	//
	// This element is present only if IsTruncated is true .
	//
	// This member is required.
	TrafficPolicyVersionMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrafficPolicyVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTrafficPolicyVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTrafficPolicyVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTrafficPolicyVersions"); err != nil {
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
	if err = addOpListTrafficPolicyVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficPolicyVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTrafficPolicyVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTrafficPolicyVersions",
	}
}
