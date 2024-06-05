// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Verifies the hash-based message authentication code (HMAC) for a specified
// message, HMAC KMS key, and MAC algorithm. To verify the HMAC, VerifyMac
// computes an HMAC using the message, HMAC KMS key, and MAC algorithm that you
// specify, and compares the computed HMAC to the HMAC that you specify. If the
// HMACs are identical, the verification succeeds; otherwise, it fails.
// Verification indicates that the message hasn't changed since the HMAC was
// calculated, and the specified key was used to generate and verify the HMAC.
//
// HMAC KMS keys and the HMAC algorithms that KMS uses conform to industry
// standards defined in [RFC 2104].
//
// This operation is part of KMS support for HMAC KMS keys. For details, see [HMAC keys in KMS] in
// the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:VerifyMac] (key policy)
//
// Related operations: GenerateMac
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [RFC 2104]: https://datatracker.ietf.org/doc/html/rfc2104
// [kms:VerifyMac]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
// [HMAC keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
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

	// The KMS key that will be used in the verification.
	//
	// Enter a key ID of the KMS key that was used to generate the HMAC. If you
	// identify a different KMS key, the VerifyMac operation fails.
	//
	// This member is required.
	KeyId *string

	// The HMAC to verify. Enter the HMAC that was generated by the GenerateMac operation when
	// you specified the same message, HMAC KMS key, and MAC algorithm as the values
	// specified in this request.
	//
	// This member is required.
	Mac []byte

	// The MAC algorithm that will be used in the verification. Enter the same MAC
	// algorithm that was used to compute the HMAC. This algorithm must be supported by
	// the HMAC KMS key identified by the KeyId parameter.
	//
	// This member is required.
	MacAlgorithm types.MacAlgorithmSpec

	// The message that will be used in the verification. Enter the same message that
	// was used to generate the HMAC.
	//
	// GenerateMacand VerifyMac do not provide special handling for message digests. If you
	// generated an HMAC for a hash digest of a message, you must verify the HMAC for
	// the same hash digest.
	//
	// This member is required.
	Message []byte

	// Checks if your request will succeed. DryRun is an optional parameter.
	//
	// To learn more about how to use this parameter, see [Testing your KMS API calls] in the Key Management
	// Service Developer Guide.
	//
	// [Testing your KMS API calls]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-dryrun.html
	DryRun *bool

	// A list of grant tokens.
	//
	// Use a grant token when your permission to call this operation comes from a new
	// grant that has not yet achieved eventual consistency. For more information, see [Grant token]
	// and [Using a grant token]in the Key Management Service Developer Guide.
	//
	// [Grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token
	// [Using a grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token
	GrantTokens []string

	noSmithyDocumentSerde
}

type VerifyMacOutput struct {

	// The HMAC KMS key used in the verification.
	KeyId *string

	// The MAC algorithm used in the verification.
	MacAlgorithm types.MacAlgorithmSpec

	// A Boolean value that indicates whether the HMAC was verified. A value of True
	// indicates that the HMAC ( Mac ) was generated with the specified Message , HMAC
	// KMS key ( KeyID ) and MacAlgorithm. .
	//
	// If the HMAC is not verified, the VerifyMac operation fails with a
	// KMSInvalidMacException exception. This exception indicates that one or more of
	// the inputs changed since the HMAC was computed.
	MacValid bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyMacMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpVerifyMac{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerifyMac{}, middleware.After)
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
