// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the Availability Zones for the specified public subnets for the
// specified Application Load Balancer, Network Load Balancer or Gateway Load
// Balancer. The specified subnets replace the previously enabled subnets.
//
// When you specify subnets for a Network Load Balancer, or Gateway Load Balancer
// you must include all subnets that were enabled previously, with their existing
// configurations, plus any additional subnets.
func (c *Client) SetSubnets(ctx context.Context, params *SetSubnetsInput, optFns ...func(*Options)) (*SetSubnetsOutput, error) {
	if params == nil {
		params = &SetSubnetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetSubnets", params, optFns, c.addOperationSetSubnetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetSubnetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetSubnetsInput struct {

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// This member is required.
	LoadBalancerArn *string

	// [Application Load Balancers] The IP address type. The possible values are ipv4
	// (for only IPv4 addresses), dualstack (for IPv4 and IPv6 addresses), and
	// dualstack-without-public-ipv4 (for IPv6 only public addresses, with private IPv4
	// and IPv6 addresses).
	//
	// [Network Load Balancers] The type of IP addresses used by the subnets for your
	// load balancer. The possible values are ipv4 (for IPv4 addresses) and dualstack
	// (for IPv4 and IPv6 addresses). You can’t specify dualstack for a load balancer
	// with a UDP or TCP_UDP listener.
	//
	// [Gateway Load Balancers] The type of IP addresses used by the subnets for your
	// load balancer. The possible values are ipv4 (for IPv4 addresses) and dualstack
	// (for IPv4 and IPv6 addresses).
	IpAddressType types.IpAddressType

	// The IDs of the public subnets. You can specify only one subnet per Availability
	// Zone. You must specify either subnets or subnet mappings.
	//
	// [Application Load Balancers] You must specify subnets from at least two
	// Availability Zones. You cannot specify Elastic IP addresses for your subnets.
	//
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//
	// [Application Load Balancers on Local Zones] You can specify subnets from one or
	// more Local Zones.
	//
	// [Network Load Balancers] You can specify subnets from one or more Availability
	// Zones. You can specify one Elastic IP address per subnet if you need static IP
	// addresses for your internet-facing load balancer. For internal load balancers,
	// you can specify one private IP address per subnet from the IPv4 range of the
	// subnet. For internet-facing load balancer, you can specify one IPv6 address per
	// subnet.
	//
	// [Gateway Load Balancers] You can specify subnets from one or more Availability
	// Zones.
	SubnetMappings []types.SubnetMapping

	// The IDs of the public subnets. You can specify only one subnet per Availability
	// Zone. You must specify either subnets or subnet mappings.
	//
	// [Application Load Balancers] You must specify subnets from at least two
	// Availability Zones.
	//
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//
	// [Application Load Balancers on Local Zones] You can specify subnets from one or
	// more Local Zones.
	//
	// [Network Load Balancers] You can specify subnets from one or more Availability
	// Zones.
	//
	// [Gateway Load Balancers] You can specify subnets from one or more Availability
	// Zones.
	Subnets []string

	noSmithyDocumentSerde
}

type SetSubnetsOutput struct {

	// Information about the subnets.
	AvailabilityZones []types.AvailabilityZone

	// [Application Load Balancers] The IP address type.
	//
	// [Network Load Balancers] The IP address type.
	//
	// [Gateway Load Balancers] The IP address type.
	IpAddressType types.IpAddressType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetSubnetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetSubnets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetSubnets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetSubnets"); err != nil {
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
	if err = addOpSetSubnetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetSubnets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetSubnets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetSubnets",
	}
}
