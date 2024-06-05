// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the export token and the signing key certificate to initiate a TR-34 key
// export from Amazon Web Services Payment Cryptography.
//
// The signing key certificate signs the wrapped key under export within the TR-34
// key payload. The export token and signing key certificate must be in place and
// operational before calling [ExportKey]. The export token expires in 7 days. You can use
// the same export token to export multiple keys from your service account.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [ExportKey]
//
// [GetParametersForImport]
//
// [ExportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ExportKey.html
// [GetParametersForImport]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetParametersForImport.html
func (c *Client) GetParametersForExport(ctx context.Context, params *GetParametersForExportInput, optFns ...func(*Options)) (*GetParametersForExportOutput, error) {
	if params == nil {
		params = &GetParametersForExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetParametersForExport", params, optFns, c.addOperationGetParametersForExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetParametersForExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetParametersForExportInput struct {

	// The key block format type (for example, TR-34 or TR-31) to use during key
	// material export. Export token is only required for a TR-34 key export,
	// TR34_KEY_BLOCK . Export token is not required for TR-31 key export.
	//
	// This member is required.
	KeyMaterialType types.KeyMaterialType

	// The signing key algorithm to generate a signing key certificate. This
	// certificate signs the wrapped key under export within the TR-34 key block.
	// RSA_2048 is the only signing key algorithm allowed.
	//
	// This member is required.
	SigningKeyAlgorithm types.KeyAlgorithm

	noSmithyDocumentSerde
}

type GetParametersForExportOutput struct {

	// The export token to initiate key export from Amazon Web Services Payment
	// Cryptography. The export token expires after 7 days. You can use the same export
	// token to export multiple keys from the same service account.
	//
	// This member is required.
	ExportToken *string

	// The validity period of the export token.
	//
	// This member is required.
	ParametersValidUntilTimestamp *time.Time

	// The algorithm of the signing key certificate for use in TR-34 key block
	// generation. RSA_2048 is the only signing key algorithm allowed.
	//
	// This member is required.
	SigningKeyAlgorithm types.KeyAlgorithm

	// The signing key certificate in PEM format (base64 encoded) of the public key
	// for signature within the TR-34 key block. The certificate expires after 7 days.
	//
	// This member is required.
	SigningKeyCertificate *string

	// The root certificate authority (CA) that signed the signing key certificate in
	// PEM format (base64 encoded).
	//
	// This member is required.
	SigningKeyCertificateChain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetParametersForExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetParametersForExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetParametersForExport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetParametersForExport"); err != nil {
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
	if err = addOpGetParametersForExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetParametersForExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetParametersForExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetParametersForExport",
	}
}
