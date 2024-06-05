// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this action to inspect your lineage and discover relationships between
// entities. For more information, see [Querying Lineage Entities]in the Amazon SageMaker Developer Guide.
//
// [Querying Lineage Entities]: https://docs.aws.amazon.com/sagemaker/latest/dg/querying-lineage-entities.html
func (c *Client) QueryLineage(ctx context.Context, params *QueryLineageInput, optFns ...func(*Options)) (*QueryLineageOutput, error) {
	if params == nil {
		params = &QueryLineageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryLineage", params, optFns, c.addOperationQueryLineageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryLineageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryLineageInput struct {

	// Associations between lineage entities have a direction. This parameter
	// determines the direction from the StartArn(s) that the query traverses.
	Direction types.Direction

	// A set of filtering parameters that allow you to specify which entities should
	// be returned.
	//
	//   - Properties - Key-value pairs to match on the lineage entities' properties.
	//
	//   - LineageTypes - A set of lineage entity types to match on. For example:
	//   TrialComponent , Artifact , or Context .
	//
	//   - CreatedBefore - Filter entities created before this date.
	//
	//   - ModifiedBefore - Filter entities modified before this date.
	//
	//   - ModifiedAfter - Filter entities modified after this date.
	Filters *types.QueryFilters

	//  Setting this value to True retrieves not only the entities of interest but
	// also the [Associations]and lineage entities on the path. Set to False to only return lineage
	// entities that match your query.
	//
	// [Associations]: https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking-entities.html
	IncludeEdges *bool

	// The maximum depth in lineage relationships from the StartArns that are
	// traversed. Depth is a measure of the number of Associations from the StartArn
	// entity to the matched results.
	MaxDepth *int32

	// Limits the number of vertices in the results. Use the NextToken in a response
	// to to retrieve the next page of results.
	MaxResults *int32

	// Limits the number of vertices in the request. Use the NextToken in a response
	// to to retrieve the next page of results.
	NextToken *string

	// A list of resource Amazon Resource Name (ARN) that represent the starting point
	// for your lineage query.
	StartArns []string

	noSmithyDocumentSerde
}

type QueryLineageOutput struct {

	// A list of edges that connect vertices in the response.
	Edges []types.Edge

	// Limits the number of vertices in the response. Use the NextToken in a response
	// to to retrieve the next page of results.
	NextToken *string

	// A list of vertices connected to the start entity(ies) in the lineage graph.
	Vertices []types.Vertex

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryLineageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpQueryLineage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpQueryLineage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "QueryLineage"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQueryLineage(options.Region), middleware.Before); err != nil {
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

// QueryLineageAPIClient is a client that implements the QueryLineage operation.
type QueryLineageAPIClient interface {
	QueryLineage(context.Context, *QueryLineageInput, ...func(*Options)) (*QueryLineageOutput, error)
}

var _ QueryLineageAPIClient = (*Client)(nil)

// QueryLineagePaginatorOptions is the paginator options for QueryLineage
type QueryLineagePaginatorOptions struct {
	// Limits the number of vertices in the results. Use the NextToken in a response
	// to to retrieve the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// QueryLineagePaginator is a paginator for QueryLineage
type QueryLineagePaginator struct {
	options   QueryLineagePaginatorOptions
	client    QueryLineageAPIClient
	params    *QueryLineageInput
	nextToken *string
	firstPage bool
}

// NewQueryLineagePaginator returns a new QueryLineagePaginator
func NewQueryLineagePaginator(client QueryLineageAPIClient, params *QueryLineageInput, optFns ...func(*QueryLineagePaginatorOptions)) *QueryLineagePaginator {
	if params == nil {
		params = &QueryLineageInput{}
	}

	options := QueryLineagePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &QueryLineagePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *QueryLineagePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next QueryLineage page.
func (p *QueryLineagePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*QueryLineageOutput, error) {
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

	result, err := p.client.QueryLineage(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opQueryLineage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "QueryLineage",
	}
}
