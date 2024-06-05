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

// Exports theme configurations to code that is ready to integrate into an Amplify
// app.
func (c *Client) ExportThemes(ctx context.Context, params *ExportThemesInput, optFns ...func(*Options)) (*ExportThemesOutput, error) {
	if params == nil {
		params = &ExportThemesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportThemes", params, optFns, c.addOperationExportThemesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportThemesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportThemesInput struct {

	// The unique ID of the Amplify app to export the themes to.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment that is part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ExportThemesOutput struct {

	// Represents the configuration of the exported themes.
	//
	// This member is required.
	Entities []types.Theme

	// The pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportThemesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportThemes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportThemes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportThemes"); err != nil {
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
	if err = addOpExportThemesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportThemes(options.Region), middleware.Before); err != nil {
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

// ExportThemesAPIClient is a client that implements the ExportThemes operation.
type ExportThemesAPIClient interface {
	ExportThemes(context.Context, *ExportThemesInput, ...func(*Options)) (*ExportThemesOutput, error)
}

var _ ExportThemesAPIClient = (*Client)(nil)

// ExportThemesPaginatorOptions is the paginator options for ExportThemes
type ExportThemesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ExportThemesPaginator is a paginator for ExportThemes
type ExportThemesPaginator struct {
	options   ExportThemesPaginatorOptions
	client    ExportThemesAPIClient
	params    *ExportThemesInput
	nextToken *string
	firstPage bool
}

// NewExportThemesPaginator returns a new ExportThemesPaginator
func NewExportThemesPaginator(client ExportThemesAPIClient, params *ExportThemesInput, optFns ...func(*ExportThemesPaginatorOptions)) *ExportThemesPaginator {
	if params == nil {
		params = &ExportThemesInput{}
	}

	options := ExportThemesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ExportThemesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ExportThemesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ExportThemes page.
func (p *ExportThemesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ExportThemesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ExportThemes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opExportThemes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportThemes",
	}
}
