// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Releases the specified address range that you provisioned to use with your
// Amazon Web Services resources through bring your own IP addresses (BYOIP) and
// deletes the corresponding address pool.
//
// Before you can release an address range, you must stop advertising it by using [WithdrawByoipCidr]
// and you must not have any accelerators that are using static IP addresses
// allocated from its address range.
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
// [WithdrawByoipCidr]: https://docs.aws.amazon.com/global-accelerator/latest/api/WithdrawByoipCidr.html
func (c *Client) DeprovisionByoipCidr(ctx context.Context, params *DeprovisionByoipCidrInput, optFns ...func(*Options)) (*DeprovisionByoipCidrOutput, error) {
	if params == nil {
		params = &DeprovisionByoipCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeprovisionByoipCidr", params, optFns, c.addOperationDeprovisionByoipCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeprovisionByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeprovisionByoipCidrInput struct {

	// The address range, in CIDR notation. The prefix must be the same prefix that
	// you specified when you provisioned the address range.
	//
	// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
	//
	// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
	//
	// This member is required.
	Cidr *string

	noSmithyDocumentSerde
}

type DeprovisionByoipCidrOutput struct {

	// Information about the address range.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeprovisionByoipCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeprovisionByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeprovisionByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeprovisionByoipCidr"); err != nil {
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
	if err = addOpDeprovisionByoipCidrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeprovisionByoipCidr(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeprovisionByoipCidr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeprovisionByoipCidr",
	}
}
