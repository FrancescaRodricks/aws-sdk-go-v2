// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This operation has been expanded to use with the Amazon GameLift containers
//
// feature, which is currently in public preview.
//
// Retrieves properties for a compute resource in an Amazon GameLift fleet. To get
// a list of all computes in a fleet, call ListCompute.
//
// To request information on a specific compute, provide the fleet ID and compute
// name.
//
// If successful, this operation returns details for the requested compute
// resource. Depending on the fleet's compute type, the result includes the
// following information:
//
//   - For EC2 fleets, this operation returns information about the EC2 instance.
//
//   - For ANYWHERE fleets, this operation returns information about the registered
//     compute.
//
//   - For CONTAINER fleets, this operation returns information about the container
//     that's registered as a compute, and the instance it's running on. The compute
//     name is the container name.
func (c *Client) DescribeCompute(ctx context.Context, params *DescribeComputeInput, optFns ...func(*Options)) (*DescribeComputeOutput, error) {
	if params == nil {
		params = &DescribeComputeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCompute", params, optFns, c.addOperationDescribeComputeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeComputeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeComputeInput struct {

	// The unique identifier of the compute resource to retrieve properties for. For
	// an Anywhere fleet compute, use the registered compute name. For an EC2 fleet
	// instance, use the instance ID. For a container fleet, use the compute name (for
	// example, a123b456c789012d3e4567f8a901b23c/1a234b56-7cd8-9e0f-a1b2-c34d567ef8a9 )
	// or the compute ARN.
	//
	// This member is required.
	ComputeName *string

	// A unique identifier for the fleet that the compute belongs to. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	noSmithyDocumentSerde
}

type DescribeComputeOutput struct {

	// The set of properties for the requested compute resource.
	Compute *types.Compute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeComputeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCompute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCompute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCompute"); err != nil {
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
	if err = addOpDescribeComputeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCompute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCompute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCompute",
	}
}
