// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the inbound cross-cluster search connections for a destination domain.
func (c *Client) DescribeInboundCrossClusterSearchConnections(ctx context.Context, params *DescribeInboundCrossClusterSearchConnectionsInput, optFns ...func(*Options)) (*DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	if params == nil {
		params = &DescribeInboundCrossClusterSearchConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInboundCrossClusterSearchConnections", params, optFns, c.addOperationDescribeInboundCrossClusterSearchConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInboundCrossClusterSearchConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeInboundCrossClusterSearchConnections operation.
type DescribeInboundCrossClusterSearchConnectionsInput struct {

	//  A list of filters used to match properties for inbound cross-cluster search
	// connection. Available Filternames for this operation are:
	//
	//   - cross-cluster-search-connection-id
	//   - source-domain-info.domain-name
	//   - source-domain-info.owner-id
	//   - source-domain-info.region
	//   - destination-domain-info.domain-name
	Filters []types.Filter

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults int32

	//  NextToken is sent in case the earlier API call results contain the NextToken.
	// It is used for pagination.
	NextToken *string

	noSmithyDocumentSerde
}

// The result of a DescribeInboundCrossClusterSearchConnections request. Contains the list of connections matching the filter
// criteria.
type DescribeInboundCrossClusterSearchConnectionsOutput struct {

	// Consists of list of InboundCrossClusterSearchConnection matching the specified filter criteria.
	CrossClusterSearchConnections []types.InboundCrossClusterSearchConnection

	// If more results are available and NextToken is present, make the next request
	// to the same API with the received NextToken to paginate the remaining results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInboundCrossClusterSearchConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInboundCrossClusterSearchConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInboundCrossClusterSearchConnections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInboundCrossClusterSearchConnections"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInboundCrossClusterSearchConnections(options.Region), middleware.Before); err != nil {
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

// DescribeInboundCrossClusterSearchConnectionsAPIClient is a client that
// implements the DescribeInboundCrossClusterSearchConnections operation.
type DescribeInboundCrossClusterSearchConnectionsAPIClient interface {
	DescribeInboundCrossClusterSearchConnections(context.Context, *DescribeInboundCrossClusterSearchConnectionsInput, ...func(*Options)) (*DescribeInboundCrossClusterSearchConnectionsOutput, error)
}

var _ DescribeInboundCrossClusterSearchConnectionsAPIClient = (*Client)(nil)

// DescribeInboundCrossClusterSearchConnectionsPaginatorOptions is the paginator
// options for DescribeInboundCrossClusterSearchConnections
type DescribeInboundCrossClusterSearchConnectionsPaginatorOptions struct {
	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInboundCrossClusterSearchConnectionsPaginator is a paginator for
// DescribeInboundCrossClusterSearchConnections
type DescribeInboundCrossClusterSearchConnectionsPaginator struct {
	options   DescribeInboundCrossClusterSearchConnectionsPaginatorOptions
	client    DescribeInboundCrossClusterSearchConnectionsAPIClient
	params    *DescribeInboundCrossClusterSearchConnectionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeInboundCrossClusterSearchConnectionsPaginator returns a new
// DescribeInboundCrossClusterSearchConnectionsPaginator
func NewDescribeInboundCrossClusterSearchConnectionsPaginator(client DescribeInboundCrossClusterSearchConnectionsAPIClient, params *DescribeInboundCrossClusterSearchConnectionsInput, optFns ...func(*DescribeInboundCrossClusterSearchConnectionsPaginatorOptions)) *DescribeInboundCrossClusterSearchConnectionsPaginator {
	if params == nil {
		params = &DescribeInboundCrossClusterSearchConnectionsInput{}
	}

	options := DescribeInboundCrossClusterSearchConnectionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInboundCrossClusterSearchConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInboundCrossClusterSearchConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInboundCrossClusterSearchConnections page.
func (p *DescribeInboundCrossClusterSearchConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeInboundCrossClusterSearchConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInboundCrossClusterSearchConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInboundCrossClusterSearchConnections",
	}
}
