// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specify the Amazon EC2 instance (destination) IP addresses and ports for a VPC
// subnet endpoint that cannot receive traffic for a custom routing accelerator.
// You can deny traffic to all destinations in the VPC endpoint, or deny traffic to
// a specified list of destination IP addresses and ports. Note that you cannot
// specify IP addresses or ports outside of the range that you configured for the
// endpoint group.
//
// After you make changes, you can verify that the updates are complete by
// checking the status of your accelerator: the status changes from IN_PROGRESS to
// DEPLOYED.
func (c *Client) DenyCustomRoutingTraffic(ctx context.Context, params *DenyCustomRoutingTrafficInput, optFns ...func(*Options)) (*DenyCustomRoutingTrafficOutput, error) {
	if params == nil {
		params = &DenyCustomRoutingTrafficInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DenyCustomRoutingTraffic", params, optFns, c.addOperationDenyCustomRoutingTrafficMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DenyCustomRoutingTrafficOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DenyCustomRoutingTrafficInput struct {

	// The Amazon Resource Name (ARN) of the endpoint group.
	//
	// This member is required.
	EndpointGroupArn *string

	// An ID for the endpoint. For custom routing accelerators, this is the virtual
	// private cloud (VPC) subnet ID.
	//
	// This member is required.
	EndpointId *string

	// Indicates whether all destination IP addresses and ports for a specified VPC
	// subnet endpoint cannot receive traffic from a custom routing accelerator. The
	// value is TRUE or FALSE.
	//
	// When set to TRUE, no destinations in the custom routing VPC subnet can receive
	// traffic. Note that you cannot specify destination IP addresses and ports when
	// the value is set to TRUE.
	//
	// When set to FALSE (or not specified), you must specify a list of destination IP
	// addresses that cannot receive traffic. A list of ports is optional. If you don't
	// specify a list of ports, the ports that can accept traffic is the same as the
	// ports configured for the endpoint group.
	//
	// The default value is FALSE.
	DenyAllTrafficToEndpoint *bool

	// A list of specific Amazon EC2 instance IP addresses (destination addresses) in
	// a subnet that you want to prevent from receiving traffic. The IP addresses must
	// be a subset of the IP addresses allowed for the VPC subnet associated with the
	// endpoint group.
	DestinationAddresses []string

	// A list of specific Amazon EC2 instance ports (destination ports) in a subnet
	// endpoint that you want to prevent from receiving traffic.
	DestinationPorts []int32

	noSmithyDocumentSerde
}

type DenyCustomRoutingTrafficOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDenyCustomRoutingTrafficMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDenyCustomRoutingTraffic{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDenyCustomRoutingTraffic{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DenyCustomRoutingTraffic"); err != nil {
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
	if err = addOpDenyCustomRoutingTrafficValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDenyCustomRoutingTraffic(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDenyCustomRoutingTraffic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DenyCustomRoutingTraffic",
	}
}
