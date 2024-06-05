// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing custom verification email template.
//
// For more information about custom verification email templates, see [Using custom verification email templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using custom verification email templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func (c *Client) UpdateCustomVerificationEmailTemplate(ctx context.Context, params *UpdateCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*UpdateCustomVerificationEmailTemplateOutput, error) {
	if params == nil {
		params = &UpdateCustomVerificationEmailTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCustomVerificationEmailTemplate", params, optFns, c.addOperationUpdateCustomVerificationEmailTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCustomVerificationEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to update an existing custom verification email template.
type UpdateCustomVerificationEmailTemplateInput struct {

	// The URL that the recipient of the verification email is sent to if his or her
	// address is not successfully verified.
	//
	// This member is required.
	FailureRedirectionURL *string

	// The email address that the custom verification email is sent from.
	//
	// This member is required.
	FromEmailAddress *string

	// The URL that the recipient of the verification email is sent to if his or her
	// address is successfully verified.
	//
	// This member is required.
	SuccessRedirectionURL *string

	// The content of the custom verification email. The total size of the email must
	// be less than 10 MB. The message body may contain HTML, with some limitations.
	// For more information, see [Custom verification email frequently asked questions]in the Amazon SES Developer Guide.
	//
	// [Custom verification email frequently asked questions]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom-faq
	//
	// This member is required.
	TemplateContent *string

	// The name of the custom verification email template that you want to update.
	//
	// This member is required.
	TemplateName *string

	// The subject line of the custom verification email.
	//
	// This member is required.
	TemplateSubject *string

	noSmithyDocumentSerde
}

// If the action is successful, the service sends back an HTTP 200 response with
// an empty HTTP body.
type UpdateCustomVerificationEmailTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCustomVerificationEmailTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCustomVerificationEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCustomVerificationEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCustomVerificationEmailTemplate"); err != nil {
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
	if err = addOpUpdateCustomVerificationEmailTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCustomVerificationEmailTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCustomVerificationEmailTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCustomVerificationEmailTemplate",
	}
}
