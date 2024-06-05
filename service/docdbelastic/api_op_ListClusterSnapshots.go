// Code generated by smithy-go-codegen DO NOT EDIT.

package docdbelastic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdbelastic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about snapshots for a specified elastic cluster.
func (c *Client) ListClusterSnapshots(ctx context.Context, params *ListClusterSnapshotsInput, optFns ...func(*Options)) (*ListClusterSnapshotsOutput, error) {
	if params == nil {
		params = &ListClusterSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusterSnapshots", params, optFns, c.addOperationListClusterSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClusterSnapshotsInput struct {

	// The ARN identifier of the elastic cluster.
	ClusterArn *string

	// The maximum number of elastic cluster snapshot results to receive in the
	// response.
	MaxResults *int32

	// A pagination token provided by a previous request. If this parameter is
	// specified, the response includes only records beyond this token, up to the value
	// specified by max-results .
	//
	// If there is no more data in the responce, the nextToken will not be returned.
	NextToken *string

	// The type of cluster snapshots to be returned. You can specify one of the
	// following values:
	//
	//   - automated - Return all cluster snapshots that Amazon DocumentDB has
	//   automatically created for your Amazon Web Services account.
	//
	//   - manual - Return all cluster snapshots that you have manually created for
	//   your Amazon Web Services account.
	SnapshotType *string

	noSmithyDocumentSerde
}

type ListClusterSnapshotsOutput struct {

	// A pagination token provided by a previous request. If this parameter is
	// specified, the response includes only records beyond this token, up to the value
	// specified by max-results .
	//
	// If there is no more data in the responce, the nextToken will not be returned.
	NextToken *string

	// A list of snapshots for a specified elastic cluster.
	Snapshots []types.ClusterSnapshotInList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClusterSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListClusterSnapshots"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusterSnapshots(options.Region), middleware.Before); err != nil {
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

// ListClusterSnapshotsAPIClient is a client that implements the
// ListClusterSnapshots operation.
type ListClusterSnapshotsAPIClient interface {
	ListClusterSnapshots(context.Context, *ListClusterSnapshotsInput, ...func(*Options)) (*ListClusterSnapshotsOutput, error)
}

var _ ListClusterSnapshotsAPIClient = (*Client)(nil)

// ListClusterSnapshotsPaginatorOptions is the paginator options for
// ListClusterSnapshots
type ListClusterSnapshotsPaginatorOptions struct {
	// The maximum number of elastic cluster snapshot results to receive in the
	// response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClusterSnapshotsPaginator is a paginator for ListClusterSnapshots
type ListClusterSnapshotsPaginator struct {
	options   ListClusterSnapshotsPaginatorOptions
	client    ListClusterSnapshotsAPIClient
	params    *ListClusterSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewListClusterSnapshotsPaginator returns a new ListClusterSnapshotsPaginator
func NewListClusterSnapshotsPaginator(client ListClusterSnapshotsAPIClient, params *ListClusterSnapshotsInput, optFns ...func(*ListClusterSnapshotsPaginatorOptions)) *ListClusterSnapshotsPaginator {
	if params == nil {
		params = &ListClusterSnapshotsInput{}
	}

	options := ListClusterSnapshotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClusterSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClusterSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListClusterSnapshots page.
func (p *ListClusterSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClusterSnapshotsOutput, error) {
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

	result, err := p.client.ListClusterSnapshots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListClusterSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListClusterSnapshots",
	}
}
