// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptographydata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Verifies a Message Authentication Code (MAC).
//
// You can use this operation to verify MAC for message data authentication such
// as . In this operation, you must use the same message data, secret encryption
// key and MAC algorithm that was used to generate MAC. You can use this operation
// to verify a DUPKT, CMAC, HMAC or EMV MAC by setting generation attributes and
// algorithm to the associated values.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # GenerateMac
//
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
func (c *Client) VerifyMac(ctx context.Context, params *VerifyMacInput, optFns ...func(*Options)) (*VerifyMacOutput, error) {
	if params == nil {
		params = &VerifyMacInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyMac", params, optFns, c.addOperationVerifyMacMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyMacOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VerifyMacInput struct {

	// The keyARN of the encryption key that Amazon Web Services Payment Cryptography
	// uses to verify MAC data.
	//
	// This member is required.
	KeyIdentifier *string

	// The MAC being verified.
	//
	// This member is required.
	Mac *string

	// The data on for which MAC is under verification. This value must be hexBinary.
	//
	// This member is required.
	MessageData *string

	// The attributes and data values to use for MAC verification within Amazon Web
	// Services Payment Cryptography.
	//
	// This member is required.
	VerificationAttributes types.MacAttributes

	// The length of the MAC.
	MacLength *int32

	noSmithyDocumentSerde
}

type VerifyMacOutput struct {

	// The keyARN of the encryption key that Amazon Web Services Payment Cryptography
	// uses for MAC verification.
	//
	// This member is required.
	KeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed.
	//
	// Amazon Web Services Payment Cryptography computes the KCV according to the CMAC
	// specification.
	//
	// This member is required.
	KeyCheckValue *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyMacMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpVerifyMac{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpVerifyMac{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "VerifyMac"); err != nil {
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
	if err = addOpVerifyMacValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyMac(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifyMac(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifyMac",
	}
}
