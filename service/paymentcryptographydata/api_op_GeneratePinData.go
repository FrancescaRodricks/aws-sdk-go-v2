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

// Generates pin-related data such as PIN, PIN Verification Value (PVV), PIN
// Block, and PIN Offset during new card issuance or reissuance. For more
// information, see [Generate PIN data]in the Amazon Web Services Payment Cryptography User Guide.
//
// PIN data is never transmitted in clear to or from Amazon Web Services Payment
// Cryptography. This operation generates PIN, PVV, or PIN Offset and then encrypts
// it using Pin Encryption Key (PEK) to create an EncryptedPinBlock for
// transmission from Amazon Web Services Payment Cryptography. This operation uses
// a separate Pin Verification Key (PVK) for VISA PVV generation.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # GenerateCardValidationData
//
// # TranslatePinData
//
// # VerifyPinData
//
// [Generate PIN data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/generate-pin-data.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
func (c *Client) GeneratePinData(ctx context.Context, params *GeneratePinDataInput, optFns ...func(*Options)) (*GeneratePinDataOutput, error) {
	if params == nil {
		params = &GeneratePinDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GeneratePinData", params, optFns, c.addOperationGeneratePinDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GeneratePinDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GeneratePinDataInput struct {

	// The keyARN of the PEK that Amazon Web Services Payment Cryptography uses to
	// encrypt the PIN Block.
	//
	// This member is required.
	EncryptionKeyIdentifier *string

	// The attributes and values to use for PIN, PVV, or PIN Offset generation.
	//
	// This member is required.
	GenerationAttributes types.PinGenerationAttributes

	// The keyARN of the PEK that Amazon Web Services Payment Cryptography uses for
	// pin data generation.
	//
	// This member is required.
	GenerationKeyIdentifier *string

	// The PIN encoding format for pin data generation as specified in ISO 9564.
	// Amazon Web Services Payment Cryptography supports ISO_Format_0 and ISO_Format_3 .
	//
	// The ISO_Format_0 PIN block format is equivalent to the ANSI X9.8, VISA-1, and
	// ECI-1 PIN block formats. It is similar to a VISA-4 PIN block format. It supports
	// a PIN from 4 to 12 digits in length.
	//
	// The ISO_Format_3 PIN block format is the same as ISO_Format_0 except that the
	// fill digits are random values from 10 to 15.
	//
	// This member is required.
	PinBlockFormat types.PinBlockFormatForPinData

	// The Primary Account Number (PAN), a unique identifier for a payment credit or
	// debit card that associates the card with a specific account holder.
	//
	// This member is required.
	PrimaryAccountNumber *string

	// The length of PIN under generation.
	PinDataLength *int32

	noSmithyDocumentSerde
}

type GeneratePinDataOutput struct {

	// The PIN block encrypted under PEK from Amazon Web Services Payment
	// Cryptography. The encrypted PIN block is a composite of PAN (Primary Account
	// Number) and PIN (Personal Identification Number), generated in accordance with
	// ISO 9564 standard.
	//
	// This member is required.
	EncryptedPinBlock *string

	// The keyARN of the PEK that Amazon Web Services Payment Cryptography uses for
	// encrypted pin block generation.
	//
	// This member is required.
	EncryptionKeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed.
	//
	// Amazon Web Services Payment Cryptography computes the KCV according to the CMAC
	// specification.
	//
	// This member is required.
	EncryptionKeyCheckValue *string

	// The keyARN of the pin data generation key that Amazon Web Services Payment
	// Cryptography uses for PIN, PVV or PIN Offset generation.
	//
	// This member is required.
	GenerationKeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed.
	//
	// Amazon Web Services Payment Cryptography computes the KCV according to the CMAC
	// specification.
	//
	// This member is required.
	GenerationKeyCheckValue *string

	// The attributes and values Amazon Web Services Payment Cryptography uses for pin
	// data generation.
	//
	// This member is required.
	PinData types.PinData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGeneratePinDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGeneratePinData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGeneratePinData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GeneratePinData"); err != nil {
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
	if err = addOpGeneratePinDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGeneratePinData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGeneratePinData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GeneratePinData",
	}
}
