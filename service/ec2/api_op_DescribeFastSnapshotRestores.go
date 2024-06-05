// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the state of fast snapshot restores for your snapshots.
func (c *Client) DescribeFastSnapshotRestores(ctx context.Context, params *DescribeFastSnapshotRestoresInput, optFns ...func(*Options)) (*DescribeFastSnapshotRestoresOutput, error) {
	if params == nil {
		params = &DescribeFastSnapshotRestoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFastSnapshotRestores", params, optFns, c.addOperationDescribeFastSnapshotRestoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFastSnapshotRestoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFastSnapshotRestoresInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters. The possible values are:
	//
	//   - availability-zone : The Availability Zone of the snapshot.
	//
	//   - owner-id : The ID of the Amazon Web Services account that enabled fast
	//   snapshot restore on the snapshot.
	//
	//   - snapshot-id : The ID of the snapshot.
	//
	//   - state : The state of fast snapshot restores for the snapshot ( enabling |
	//   optimizing | enabled | disabling | disabled ).
	Filters []types.Filter

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFastSnapshotRestoresOutput struct {

	// Information about the state of fast snapshot restores.
	FastSnapshotRestores []types.DescribeFastSnapshotRestoreSuccessItem

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFastSnapshotRestoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeFastSnapshotRestores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFastSnapshotRestores{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFastSnapshotRestores"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFastSnapshotRestores(options.Region), middleware.Before); err != nil {
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

// DescribeFastSnapshotRestoresAPIClient is a client that implements the
// DescribeFastSnapshotRestores operation.
type DescribeFastSnapshotRestoresAPIClient interface {
	DescribeFastSnapshotRestores(context.Context, *DescribeFastSnapshotRestoresInput, ...func(*Options)) (*DescribeFastSnapshotRestoresOutput, error)
}

var _ DescribeFastSnapshotRestoresAPIClient = (*Client)(nil)

// DescribeFastSnapshotRestoresPaginatorOptions is the paginator options for
// DescribeFastSnapshotRestores
type DescribeFastSnapshotRestoresPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFastSnapshotRestoresPaginator is a paginator for
// DescribeFastSnapshotRestores
type DescribeFastSnapshotRestoresPaginator struct {
	options   DescribeFastSnapshotRestoresPaginatorOptions
	client    DescribeFastSnapshotRestoresAPIClient
	params    *DescribeFastSnapshotRestoresInput
	nextToken *string
	firstPage bool
}

// NewDescribeFastSnapshotRestoresPaginator returns a new
// DescribeFastSnapshotRestoresPaginator
func NewDescribeFastSnapshotRestoresPaginator(client DescribeFastSnapshotRestoresAPIClient, params *DescribeFastSnapshotRestoresInput, optFns ...func(*DescribeFastSnapshotRestoresPaginatorOptions)) *DescribeFastSnapshotRestoresPaginator {
	if params == nil {
		params = &DescribeFastSnapshotRestoresInput{}
	}

	options := DescribeFastSnapshotRestoresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFastSnapshotRestoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFastSnapshotRestoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFastSnapshotRestores page.
func (p *DescribeFastSnapshotRestoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFastSnapshotRestoresOutput, error) {
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

	result, err := p.client.DescribeFastSnapshotRestores(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFastSnapshotRestores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFastSnapshotRestores",
	}
}
