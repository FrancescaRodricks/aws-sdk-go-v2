// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of job templates.
//
// Requires permission to access the [ListJobTemplates] action.
//
// [ListJobTemplates]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListJobTemplates(ctx context.Context, params *ListJobTemplatesInput, optFns ...func(*Options)) (*ListJobTemplatesOutput, error) {
	if params == nil {
		params = &ListJobTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobTemplates", params, optFns, c.addOperationListJobTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobTemplatesInput struct {

	// The maximum number of results to return in the list.
	MaxResults *int32

	// The token to use to return the next set of results in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListJobTemplatesOutput struct {

	// A list of objects that contain information about the job templates.
	JobTemplates []types.JobTemplateSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJobTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListJobTemplates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobTemplates(options.Region), middleware.Before); err != nil {
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

// ListJobTemplatesAPIClient is a client that implements the ListJobTemplates
// operation.
type ListJobTemplatesAPIClient interface {
	ListJobTemplates(context.Context, *ListJobTemplatesInput, ...func(*Options)) (*ListJobTemplatesOutput, error)
}

var _ ListJobTemplatesAPIClient = (*Client)(nil)

// ListJobTemplatesPaginatorOptions is the paginator options for ListJobTemplates
type ListJobTemplatesPaginatorOptions struct {
	// The maximum number of results to return in the list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobTemplatesPaginator is a paginator for ListJobTemplates
type ListJobTemplatesPaginator struct {
	options   ListJobTemplatesPaginatorOptions
	client    ListJobTemplatesAPIClient
	params    *ListJobTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListJobTemplatesPaginator returns a new ListJobTemplatesPaginator
func NewListJobTemplatesPaginator(client ListJobTemplatesAPIClient, params *ListJobTemplatesInput, optFns ...func(*ListJobTemplatesPaginatorOptions)) *ListJobTemplatesPaginator {
	if params == nil {
		params = &ListJobTemplatesInput{}
	}

	options := ListJobTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJobTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJobTemplates page.
func (p *ListJobTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobTemplatesOutput, error) {
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

	result, err := p.client.ListJobTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListJobTemplates",
	}
}
