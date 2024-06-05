// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifyuibuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports form configurations to code that is ready to integrate into an Amplify
// app.
func (c *Client) ExportForms(ctx context.Context, params *ExportFormsInput, optFns ...func(*Options)) (*ExportFormsOutput, error) {
	if params == nil {
		params = &ExportFormsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportForms", params, optFns, c.addOperationExportFormsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportFormsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportFormsInput struct {

	// The unique ID of the Amplify app to export forms to.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ExportFormsOutput struct {

	// Represents the configuration of the exported forms.
	//
	// This member is required.
	Entities []types.Form

	// The pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportFormsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportForms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportForms{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportForms"); err != nil {
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
	if err = addOpExportFormsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportForms(options.Region), middleware.Before); err != nil {
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

// ExportFormsAPIClient is a client that implements the ExportForms operation.
type ExportFormsAPIClient interface {
	ExportForms(context.Context, *ExportFormsInput, ...func(*Options)) (*ExportFormsOutput, error)
}

var _ ExportFormsAPIClient = (*Client)(nil)

// ExportFormsPaginatorOptions is the paginator options for ExportForms
type ExportFormsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ExportFormsPaginator is a paginator for ExportForms
type ExportFormsPaginator struct {
	options   ExportFormsPaginatorOptions
	client    ExportFormsAPIClient
	params    *ExportFormsInput
	nextToken *string
	firstPage bool
}

// NewExportFormsPaginator returns a new ExportFormsPaginator
func NewExportFormsPaginator(client ExportFormsAPIClient, params *ExportFormsInput, optFns ...func(*ExportFormsPaginatorOptions)) *ExportFormsPaginator {
	if params == nil {
		params = &ExportFormsInput{}
	}

	options := ExportFormsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ExportFormsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ExportFormsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ExportForms page.
func (p *ExportFormsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ExportFormsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ExportForms(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opExportForms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportForms",
	}
}
