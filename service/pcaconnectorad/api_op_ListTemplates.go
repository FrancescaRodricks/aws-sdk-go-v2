// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorad/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the templates, if any, that are associated with a connector.
func (c *Client) ListTemplates(ctx context.Context, params *ListTemplatesInput, optFns ...func(*Options)) (*ListTemplatesOutput, error) {
	if params == nil {
		params = &ListTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTemplates", params, optFns, c.addOperationListTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTemplatesInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called [CreateConnector].
	//
	// [CreateConnector]: https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html
	//
	// This member is required.
	ConnectorArn *string

	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	MaxResults *int32

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of the NextToken
	// parameter from the response you just received.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTemplatesOutput struct {

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of the NextToken
	// parameter from the response you just received.
	NextToken *string

	// Custom configuration templates used when issuing a certificate.
	Templates []types.TemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTemplates"); err != nil {
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
	if err = addOpListTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTemplates(options.Region), middleware.Before); err != nil {
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

// ListTemplatesAPIClient is a client that implements the ListTemplates operation.
type ListTemplatesAPIClient interface {
	ListTemplates(context.Context, *ListTemplatesInput, ...func(*Options)) (*ListTemplatesOutput, error)
}

var _ ListTemplatesAPIClient = (*Client)(nil)

// ListTemplatesPaginatorOptions is the paginator options for ListTemplates
type ListTemplatesPaginatorOptions struct {
	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTemplatesPaginator is a paginator for ListTemplates
type ListTemplatesPaginator struct {
	options   ListTemplatesPaginatorOptions
	client    ListTemplatesAPIClient
	params    *ListTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListTemplatesPaginator returns a new ListTemplatesPaginator
func NewListTemplatesPaginator(client ListTemplatesAPIClient, params *ListTemplatesInput, optFns ...func(*ListTemplatesPaginatorOptions)) *ListTemplatesPaginator {
	if params == nil {
		params = &ListTemplatesInput{}
	}

	options := ListTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTemplates page.
func (p *ListTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTemplatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTemplates",
	}
}
