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
// Requests authorization to remotely connect to a hosting resource in a Amazon
// GameLift managed fleet. This operation is not used with Amazon GameLift Anywhere
// fleets
//
// To request access, specify the compute name and the fleet ID. If successful,
// this operation returns a set of temporary Amazon Web Services credentials,
// including a two-part access key and a session token.
//
// # EC2 fleets
//
// With an EC2 fleet (where compute type is EC2 ), use these credentials with
// Amazon EC2 Systems Manager (SSM) to start a session with the compute. For more
// details, see [Starting a session (CLI)]in the Amazon EC2 Systems Manager User Guide.
//
// # Container fleets
//
// With a container fleet (where compute type is CONTAINER ), use these credentials
// and the target value with SSM to connect to the fleet instance where the
// container is running. After you're connected to the instance, use Docker
// commands to interact with the container.
//
// # Learn more
//
// [Remotely connect to fleet instances]
//
// [Debug fleet issues]
//
// [Remotely connect to a container fleet]
//
// [Remotely connect to fleet instances]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-remote-access.html
// [Remotely connect to a container fleet]: https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-remote-access.html
// [Debug fleet issues]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html
// [Starting a session (CLI)]: https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-sessions-start.html#sessions-start-cli
func (c *Client) GetComputeAccess(ctx context.Context, params *GetComputeAccessInput, optFns ...func(*Options)) (*GetComputeAccessOutput, error) {
	if params == nil {
		params = &GetComputeAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComputeAccess", params, optFns, c.addOperationGetComputeAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComputeAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComputeAccessInput struct {

	// A unique identifier for the compute resource that you want to connect to. For
	// an EC2 fleet compute, use the instance ID. For a container fleet, use the
	// compute name (for example,
	// a123b456c789012d3e4567f8a901b23c/1a234b56-7cd8-9e0f-a1b2-c34d567ef8a9 ) or the
	// compute ARN.
	//
	// This member is required.
	ComputeName *string

	// A unique identifier for the fleet that holds the compute resource that you want
	// to connect to. You can use either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	noSmithyDocumentSerde
}

type GetComputeAccessOutput struct {

	// The Amazon Resource Name ([ARN] ) that is assigned to an Amazon GameLift compute
	// resource and uniquely identifies it. ARNs are unique across all Regions. Format
	// is arn:aws:gamelift:::compute/compute-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	//
	// [ARN]: https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html
	ComputeArn *string

	// The identifier of the compute resource to be accessed. This value might be
	// either a compute name or an instance ID.
	ComputeName *string

	// A set of temporary Amazon Web Services credentials for use when connecting to
	// the compute resource with Amazon EC2 Systems Manager (SSM).
	Credentials *types.AwsCredentials

	// The Amazon Resource Name ([ARN] ) that is assigned to a Amazon GameLift fleet
	// resource and uniquely identifies it. ARNs are unique across all Regions. Format
	// is arn:aws:gamelift:::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912 .
	//
	// [ARN]: https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html
	FleetArn *string

	// The ID of the fleet that holds the compute resource to be accessed.
	FleetId *string

	// (For container fleets only) The instance ID where the compute resource is
	// running.
	Target *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComputeAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComputeAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComputeAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetComputeAccess"); err != nil {
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
	if err = addOpGetComputeAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComputeAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetComputeAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetComputeAccess",
	}
}
