// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all metadata generation runs.
func (c *Client) ListMetadataGenerationRuns(ctx context.Context, params *ListMetadataGenerationRunsInput, optFns ...func(*Options)) (*ListMetadataGenerationRunsOutput, error) {
	if params == nil {
		params = &ListMetadataGenerationRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMetadataGenerationRuns", params, optFns, c.addOperationListMetadataGenerationRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMetadataGenerationRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMetadataGenerationRunsInput struct {

	// The ID of the Amazon DataZone domain where you want to list metadata generation
	// runs.
	//
	// This member is required.
	DomainIdentifier *string

	// The maximum number of metadata generation runs to return in a single call to
	// ListMetadataGenerationRuns. When the number of metadata generation runs to be
	// listed is greater than the value of MaxResults, the response contains a
	// NextToken value that you can use in a subsequent call to
	// ListMetadataGenerationRuns to list the next set of revisions.
	MaxResults *int32

	// When the number of metadata generation runs is greater than the default value
	// for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of metadata generation runs, the
	// response includes a pagination token named NextToken. You can specify this
	// NextToken value in a subsequent call to ListMetadataGenerationRuns to list the
	// next set of revisions.
	NextToken *string

	// The status of the metadata generation runs.
	Status types.MetadataGenerationRunStatus

	// The type of the metadata generation runs.
	Type types.MetadataGenerationRunType

	noSmithyDocumentSerde
}

type ListMetadataGenerationRunsOutput struct {

	// The results of the ListMetadataGenerationRuns action.
	Items []types.MetadataGenerationRunItem

	// When the number of metadata generation runs is greater than the default value
	// for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of metadata generation runs, the
	// response includes a pagination token named NextToken. You can specify this
	// NextToken value in a subsequent call to ListMetadataGenerationRuns to list the
	// next set of revisions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMetadataGenerationRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMetadataGenerationRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMetadataGenerationRuns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMetadataGenerationRuns"); err != nil {
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
	if err = addOpListMetadataGenerationRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMetadataGenerationRuns(options.Region), middleware.Before); err != nil {
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

// ListMetadataGenerationRunsAPIClient is a client that implements the
// ListMetadataGenerationRuns operation.
type ListMetadataGenerationRunsAPIClient interface {
	ListMetadataGenerationRuns(context.Context, *ListMetadataGenerationRunsInput, ...func(*Options)) (*ListMetadataGenerationRunsOutput, error)
}

var _ ListMetadataGenerationRunsAPIClient = (*Client)(nil)

// ListMetadataGenerationRunsPaginatorOptions is the paginator options for
// ListMetadataGenerationRuns
type ListMetadataGenerationRunsPaginatorOptions struct {
	// The maximum number of metadata generation runs to return in a single call to
	// ListMetadataGenerationRuns. When the number of metadata generation runs to be
	// listed is greater than the value of MaxResults, the response contains a
	// NextToken value that you can use in a subsequent call to
	// ListMetadataGenerationRuns to list the next set of revisions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMetadataGenerationRunsPaginator is a paginator for
// ListMetadataGenerationRuns
type ListMetadataGenerationRunsPaginator struct {
	options   ListMetadataGenerationRunsPaginatorOptions
	client    ListMetadataGenerationRunsAPIClient
	params    *ListMetadataGenerationRunsInput
	nextToken *string
	firstPage bool
}

// NewListMetadataGenerationRunsPaginator returns a new
// ListMetadataGenerationRunsPaginator
func NewListMetadataGenerationRunsPaginator(client ListMetadataGenerationRunsAPIClient, params *ListMetadataGenerationRunsInput, optFns ...func(*ListMetadataGenerationRunsPaginatorOptions)) *ListMetadataGenerationRunsPaginator {
	if params == nil {
		params = &ListMetadataGenerationRunsInput{}
	}

	options := ListMetadataGenerationRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMetadataGenerationRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMetadataGenerationRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMetadataGenerationRuns page.
func (p *ListMetadataGenerationRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMetadataGenerationRunsOutput, error) {
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

	result, err := p.client.ListMetadataGenerationRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMetadataGenerationRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMetadataGenerationRuns",
	}
}
