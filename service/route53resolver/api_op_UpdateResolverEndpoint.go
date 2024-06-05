// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the name, or endpoint type for an inbound or an outbound Resolver
// endpoint. You can only update between IPV4 and DUALSTACK, IPV6 endpoint type
// can't be updated to other type.
func (c *Client) UpdateResolverEndpoint(ctx context.Context, params *UpdateResolverEndpointInput, optFns ...func(*Options)) (*UpdateResolverEndpointOutput, error) {
	if params == nil {
		params = &UpdateResolverEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResolverEndpoint", params, optFns, c.addOperationUpdateResolverEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResolverEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResolverEndpointInput struct {

	// The ID of the Resolver endpoint that you want to update.
	//
	// This member is required.
	ResolverEndpointId *string

	// The name of the Resolver endpoint that you want to update.
	Name *string

	//  The protocols you want to use for the endpoint. DoH-FIPS is applicable for
	// inbound endpoints only.
	//
	// For an inbound endpoint you can apply the protocols as follows:
	//
	//   - Do53 and DoH in combination.
	//
	//   - Do53 and DoH-FIPS in combination.
	//
	//   - Do53 alone.
	//
	//   - DoH alone.
	//
	//   - DoH-FIPS alone.
	//
	//   - None, which is treated as Do53.
	//
	// For an outbound endpoint you can apply the protocols as follows:
	//
	//   - Do53 and DoH in combination.
	//
	//   - Do53 alone.
	//
	//   - DoH alone.
	//
	//   - None, which is treated as Do53.
	//
	// You can't change the protocol of an inbound endpoint directly from only Do53 to
	// only DoH, or DoH-FIPS. This is to prevent a sudden disruption to incoming
	// traffic that relies on Do53. To change the protocol from Do53 to DoH, or
	// DoH-FIPS, you must first enable both Do53 and DoH, or Do53 and DoH-FIPS, to make
	// sure that all incoming traffic has transferred to using the DoH protocol, or
	// DoH-FIPS, and then remove the Do53.
	Protocols []types.Protocol

	//  Specifies the endpoint type for what type of IP address the endpoint uses to
	// forward DNS queries.
	//
	// Updating to IPV6 type isn't currently supported.
	ResolverEndpointType types.ResolverEndpointType

	//  Specifies the IPv6 address when you update the Resolver endpoint from IPv4 to
	// dual-stack. If you don't specify an IPv6 address, one will be automatically
	// chosen from your subnet.
	UpdateIpAddresses []types.UpdateIpAddress

	noSmithyDocumentSerde
}

type UpdateResolverEndpointOutput struct {

	// The response to an UpdateResolverEndpoint request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResolverEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateResolverEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateResolverEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateResolverEndpoint"); err != nil {
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
	if err = addOpUpdateResolverEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResolverEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResolverEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateResolverEndpoint",
	}
}
