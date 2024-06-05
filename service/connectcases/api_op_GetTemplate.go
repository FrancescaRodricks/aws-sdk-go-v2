// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the details for the requested template.
func (c *Client) GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) {
	if params == nil {
		params = &GetTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTemplate", params, optFns, c.addOperationGetTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTemplateInput struct {

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// A unique identifier of a template.
	//
	// This member is required.
	TemplateId *string

	noSmithyDocumentSerde
}

type GetTemplateOutput struct {

	// The name of the template.
	//
	// This member is required.
	Name *string

	// The status of the template.
	//
	// This member is required.
	Status types.TemplateStatus

	// The Amazon Resource Name (ARN) of the template.
	//
	// This member is required.
	TemplateArn *string

	// A unique identifier of a template.
	//
	// This member is required.
	TemplateId *string

	// Timestamp at which the resource was created.
	CreatedTime *time.Time

	// Denotes whether or not the resource has been deleted.
	Deleted bool

	// A brief description of the template.
	Description *string

	// Timestamp at which the resource was created or last modified.
	LastModifiedTime *time.Time

	// Configuration of layouts associated to the template.
	LayoutConfiguration *types.LayoutConfiguration

	// A list of fields that must contain a value for a case to be successfully
	// created with this template.
	RequiredFields []types.RequiredField

	// A map of of key-value pairs that represent tags on a resource. Tags are used to
	// organize, track, or control access for this resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTemplate"); err != nil {
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
	if err = addOpGetTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTemplate",
	}
}
