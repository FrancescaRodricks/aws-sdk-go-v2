// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Export an existing lens.
//
// Only the owner of a lens can export it. Lenses provided by Amazon Web Services
// (Amazon Web Services Official Content) cannot be exported.
//
// Lenses are defined in JSON. For more information, see [JSON format specification] in the Well-Architected
// Tool User Guide.
//
// # Disclaimer
//
// Do not include or gather personal identifiable information (PII) of end users
// or other identifiable individuals in or via your custom lenses. If your custom
// lens or those shared with you and used in your account do include or collect PII
// you are responsible for: ensuring that the included PII is processed in
// accordance with applicable law, providing adequate privacy notices, and
// obtaining necessary consents for processing such data.
//
// [JSON format specification]: https://docs.aws.amazon.com/wellarchitected/latest/userguide/lenses-format-specification.html
func (c *Client) ExportLens(ctx context.Context, params *ExportLensInput, optFns ...func(*Options)) (*ExportLensOutput, error) {
	if params == nil {
		params = &ExportLensInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportLens", params, optFns, c.addOperationExportLensMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportLensOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportLensInput struct {

	// The alias of the lens.
	//
	// For Amazon Web Services official lenses, this is either the lens alias, such as
	// serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses.
	//
	// For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// .
	//
	// Each lens is identified by its LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The lens version to be exported.
	LensVersion *string

	noSmithyDocumentSerde
}

type ExportLensOutput struct {

	// The JSON representation of a lens.
	LensJSON *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportLensMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportLens{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportLens{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportLens"); err != nil {
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
	if err = addOpExportLensValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportLens(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportLens(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportLens",
	}
}
