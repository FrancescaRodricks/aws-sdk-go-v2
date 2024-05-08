// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a review template.
//
// # Disclaimer
//
// Do not include or gather personal identifiable information (PII) of end users
// or other identifiable individuals in or via your review templates. If your
// review template or those shared with you and used in your account do include or
// collect PII you are responsible for: ensuring that the included PII is processed
// in accordance with applicable law, providing adequate privacy notices, and
// obtaining necessary consents for processing such data.
func (c *Client) CreateReviewTemplate(ctx context.Context, params *CreateReviewTemplateInput, optFns ...func(*Options)) (*CreateReviewTemplateOutput, error) {
	if params == nil {
		params = &CreateReviewTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReviewTemplate", params, optFns, c.addOperationCreateReviewTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReviewTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReviewTemplateInput struct {

	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once).
	//
	// You should not reuse the same token for other requests. If you retry a request
	// with the same client request token and the same parameters after the original
	// request has completed successfully, the result of the original request is
	// returned.
	//
	// This token is listed as required, however, if you do not specify it, the Amazon
	// Web Services SDKs automatically generate one for you. If you are not using the
	// Amazon Web Services SDK or the CLI, you must provide this token or the request
	// will fail.
	//
	// This member is required.
	ClientRequestToken *string

	// The review template description.
	//
	// This member is required.
	Description *string

	// Lenses applied to the review template.
	//
	// This member is required.
	Lenses []string

	// Name of the review template.
	//
	// This member is required.
	TemplateName *string

	// The notes associated with the workload.
	//
	// For a review template, these are the notes that will be associated with the
	// workload when the template is applied.
	Notes *string

	// The tags assigned to the review template.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateReviewTemplateOutput struct {

	// The review template ARN.
	TemplateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReviewTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateReviewTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateReviewTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReviewTemplate"); err != nil {
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
	if err = addIdempotencyToken_opCreateReviewTemplateMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateReviewTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReviewTemplate(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateReviewTemplate struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateReviewTemplate) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateReviewTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateReviewTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateReviewTemplateInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateReviewTemplateMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateReviewTemplate{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateReviewTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReviewTemplate",
	}
}
