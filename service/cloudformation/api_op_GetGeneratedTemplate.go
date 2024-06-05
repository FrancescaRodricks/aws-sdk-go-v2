// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a generated template. If the template is in an InProgress or Pending
// status then the template returned will be the template when the template was
// last in a Complete status. If the template has not yet been in a Complete
// status then an empty template will be returned.
func (c *Client) GetGeneratedTemplate(ctx context.Context, params *GetGeneratedTemplateInput, optFns ...func(*Options)) (*GetGeneratedTemplateOutput, error) {
	if params == nil {
		params = &GetGeneratedTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGeneratedTemplate", params, optFns, c.addOperationGetGeneratedTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGeneratedTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGeneratedTemplateInput struct {

	// The name or Amazon Resource Name (ARN) of the generated template. The format is
	// arn:${Partition}:cloudformation:${Region}:${Account}:generatedtemplate/${Id} .
	// For example,
	// arn:aws:cloudformation:us-east-1:123456789012:generatedtemplate/2e8465c1-9a80-43ea-a3a3-4f2d692fe6dc
	// .
	//
	// This member is required.
	GeneratedTemplateName *string

	// The language to use to retrieve for the generated template. Supported values
	// are:
	//
	//   - JSON
	//
	//   - YAML
	Format types.TemplateFormat

	noSmithyDocumentSerde
}

type GetGeneratedTemplateOutput struct {

	// The status of the template generation. Supported values are:
	//
	//   - CreatePending - the creation of the template is pending.
	//
	//   - CreateInProgress - the creation of the template is in progress.
	//
	//   - DeletePending - the deletion of the template is pending.
	//
	//   - DeleteInProgress - the deletion of the template is in progress.
	//
	//   - UpdatePending - the update of the template is pending.
	//
	//   - UpdateInProgress - the update of the template is in progress.
	//
	//   - Failed - the template operation failed.
	//
	//   - Complete - the template operation is complete.
	Status types.GeneratedTemplateStatus

	// The template body of the generated template, in the language specified by the
	// Language parameter.
	TemplateBody *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGeneratedTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetGeneratedTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetGeneratedTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGeneratedTemplate"); err != nil {
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
	if err = addOpGetGeneratedTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGeneratedTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGeneratedTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGeneratedTemplate",
	}
}
