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

// Describes a root volume replacement task. For more information, see [Replace a root volume] in the
// Amazon Elastic Compute Cloud User Guide.
//
// [Replace a root volume]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/replace-root.html
func (c *Client) DescribeReplaceRootVolumeTasks(ctx context.Context, params *DescribeReplaceRootVolumeTasksInput, optFns ...func(*Options)) (*DescribeReplaceRootVolumeTasksOutput, error) {
	if params == nil {
		params = &DescribeReplaceRootVolumeTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplaceRootVolumeTasks", params, optFns, c.addOperationDescribeReplaceRootVolumeTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplaceRootVolumeTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplaceRootVolumeTasksInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Filter to use:
	//
	//   - instance-id - The ID of the instance for which the root volume replacement
	//   task was created.
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

	// The ID of the root volume replacement task to view.
	ReplaceRootVolumeTaskIds []string

	noSmithyDocumentSerde
}

type DescribeReplaceRootVolumeTasksOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Information about the root volume replacement task.
	ReplaceRootVolumeTasks []types.ReplaceRootVolumeTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplaceRootVolumeTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeReplaceRootVolumeTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeReplaceRootVolumeTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReplaceRootVolumeTasks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplaceRootVolumeTasks(options.Region), middleware.Before); err != nil {
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

// DescribeReplaceRootVolumeTasksAPIClient is a client that implements the
// DescribeReplaceRootVolumeTasks operation.
type DescribeReplaceRootVolumeTasksAPIClient interface {
	DescribeReplaceRootVolumeTasks(context.Context, *DescribeReplaceRootVolumeTasksInput, ...func(*Options)) (*DescribeReplaceRootVolumeTasksOutput, error)
}

var _ DescribeReplaceRootVolumeTasksAPIClient = (*Client)(nil)

// DescribeReplaceRootVolumeTasksPaginatorOptions is the paginator options for
// DescribeReplaceRootVolumeTasks
type DescribeReplaceRootVolumeTasksPaginatorOptions struct {
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

// DescribeReplaceRootVolumeTasksPaginator is a paginator for
// DescribeReplaceRootVolumeTasks
type DescribeReplaceRootVolumeTasksPaginator struct {
	options   DescribeReplaceRootVolumeTasksPaginatorOptions
	client    DescribeReplaceRootVolumeTasksAPIClient
	params    *DescribeReplaceRootVolumeTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplaceRootVolumeTasksPaginator returns a new
// DescribeReplaceRootVolumeTasksPaginator
func NewDescribeReplaceRootVolumeTasksPaginator(client DescribeReplaceRootVolumeTasksAPIClient, params *DescribeReplaceRootVolumeTasksInput, optFns ...func(*DescribeReplaceRootVolumeTasksPaginatorOptions)) *DescribeReplaceRootVolumeTasksPaginator {
	if params == nil {
		params = &DescribeReplaceRootVolumeTasksInput{}
	}

	options := DescribeReplaceRootVolumeTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReplaceRootVolumeTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplaceRootVolumeTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReplaceRootVolumeTasks page.
func (p *DescribeReplaceRootVolumeTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplaceRootVolumeTasksOutput, error) {
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

	result, err := p.client.DescribeReplaceRootVolumeTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReplaceRootVolumeTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReplaceRootVolumeTasks",
	}
}
