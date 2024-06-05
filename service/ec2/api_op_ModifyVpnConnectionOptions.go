// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the connection options for your Site-to-Site VPN connection.
//
// When you modify the VPN connection options, the VPN endpoint IP addresses on
// the Amazon Web Services side do not change, and the tunnel options do not
// change. Your VPN connection will be temporarily unavailable for a brief period
// while the VPN connection is updated.
func (c *Client) ModifyVpnConnectionOptions(ctx context.Context, params *ModifyVpnConnectionOptionsInput, optFns ...func(*Options)) (*ModifyVpnConnectionOptionsOutput, error) {
	if params == nil {
		params = &ModifyVpnConnectionOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpnConnectionOptions", params, optFns, c.addOperationModifyVpnConnectionOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpnConnectionOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpnConnectionOptionsInput struct {

	// The ID of the Site-to-Site VPN connection.
	//
	// This member is required.
	VpnConnectionId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.
	//
	// Default: 0.0.0.0/0
	LocalIpv4NetworkCidr *string

	// The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.
	//
	// Default: ::/0
	LocalIpv6NetworkCidr *string

	// The IPv4 CIDR on the Amazon Web Services side of the VPN connection.
	//
	// Default: 0.0.0.0/0
	RemoteIpv4NetworkCidr *string

	// The IPv6 CIDR on the Amazon Web Services side of the VPN connection.
	//
	// Default: ::/0
	RemoteIpv6NetworkCidr *string

	noSmithyDocumentSerde
}

type ModifyVpnConnectionOptionsOutput struct {

	// Information about the VPN connection.
	VpnConnection *types.VpnConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVpnConnectionOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpnConnectionOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpnConnectionOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyVpnConnectionOptions"); err != nil {
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
	if err = addOpModifyVpnConnectionOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpnConnectionOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpnConnectionOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyVpnConnectionOptions",
	}
}
