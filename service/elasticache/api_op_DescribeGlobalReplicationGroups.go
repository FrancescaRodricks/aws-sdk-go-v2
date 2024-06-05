// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a particular global replication group. If no
// identifier is specified, returns information about all Global datastores.
func (c *Client) DescribeGlobalReplicationGroups(ctx context.Context, params *DescribeGlobalReplicationGroupsInput, optFns ...func(*Options)) (*DescribeGlobalReplicationGroupsOutput, error) {
	if params == nil {
		params = &DescribeGlobalReplicationGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGlobalReplicationGroups", params, optFns, c.addOperationDescribeGlobalReplicationGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGlobalReplicationGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalReplicationGroupsInput struct {

	// The name of the Global datastore
	GlobalReplicationGroupId *string

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	MaxRecords *int32

	// Returns the list of members that comprise the Global datastore.
	ShowMemberInfo *bool

	noSmithyDocumentSerde
}

type DescribeGlobalReplicationGroupsOutput struct {

	// Indicates the slot configuration and global identifier for each slice group.
	GlobalReplicationGroups []types.GlobalReplicationGroup

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords. >
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGlobalReplicationGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeGlobalReplicationGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeGlobalReplicationGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeGlobalReplicationGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalReplicationGroups(options.Region), middleware.Before); err != nil {
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

// DescribeGlobalReplicationGroupsAPIClient is a client that implements the
// DescribeGlobalReplicationGroups operation.
type DescribeGlobalReplicationGroupsAPIClient interface {
	DescribeGlobalReplicationGroups(context.Context, *DescribeGlobalReplicationGroupsInput, ...func(*Options)) (*DescribeGlobalReplicationGroupsOutput, error)
}

var _ DescribeGlobalReplicationGroupsAPIClient = (*Client)(nil)

// DescribeGlobalReplicationGroupsPaginatorOptions is the paginator options for
// DescribeGlobalReplicationGroups
type DescribeGlobalReplicationGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeGlobalReplicationGroupsPaginator is a paginator for
// DescribeGlobalReplicationGroups
type DescribeGlobalReplicationGroupsPaginator struct {
	options   DescribeGlobalReplicationGroupsPaginatorOptions
	client    DescribeGlobalReplicationGroupsAPIClient
	params    *DescribeGlobalReplicationGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeGlobalReplicationGroupsPaginator returns a new
// DescribeGlobalReplicationGroupsPaginator
func NewDescribeGlobalReplicationGroupsPaginator(client DescribeGlobalReplicationGroupsAPIClient, params *DescribeGlobalReplicationGroupsInput, optFns ...func(*DescribeGlobalReplicationGroupsPaginatorOptions)) *DescribeGlobalReplicationGroupsPaginator {
	if params == nil {
		params = &DescribeGlobalReplicationGroupsInput{}
	}

	options := DescribeGlobalReplicationGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeGlobalReplicationGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeGlobalReplicationGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeGlobalReplicationGroups page.
func (p *DescribeGlobalReplicationGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeGlobalReplicationGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeGlobalReplicationGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeGlobalReplicationGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeGlobalReplicationGroups",
	}
}
