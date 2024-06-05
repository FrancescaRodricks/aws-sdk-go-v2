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

// Describes the stale security group rules for security groups in a specified
// VPC. Rules are stale when they reference a deleted security group in the same
// VPC or peered VPC. Rules can also be stale if they reference a security group in
// a peer VPC for which the VPC peering connection has been deleted.
func (c *Client) DescribeStaleSecurityGroups(ctx context.Context, params *DescribeStaleSecurityGroupsInput, optFns ...func(*Options)) (*DescribeStaleSecurityGroupsOutput, error) {
	if params == nil {
		params = &DescribeStaleSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStaleSecurityGroups", params, optFns, c.addOperationDescribeStaleSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStaleSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStaleSecurityGroupsInput struct {

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

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

type DescribeStaleSecurityGroupsOutput struct {

	// The token to include in another request to get the next page of items. If there
	// are no additional items to return, the string is empty.
	NextToken *string

	// Information about the stale security groups.
	StaleSecurityGroupSet []types.StaleSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStaleSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeStaleSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeStaleSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStaleSecurityGroups"); err != nil {
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
	if err = addOpDescribeStaleSecurityGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStaleSecurityGroups(options.Region), middleware.Before); err != nil {
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

// DescribeStaleSecurityGroupsAPIClient is a client that implements the
// DescribeStaleSecurityGroups operation.
type DescribeStaleSecurityGroupsAPIClient interface {
	DescribeStaleSecurityGroups(context.Context, *DescribeStaleSecurityGroupsInput, ...func(*Options)) (*DescribeStaleSecurityGroupsOutput, error)
}

var _ DescribeStaleSecurityGroupsAPIClient = (*Client)(nil)

// DescribeStaleSecurityGroupsPaginatorOptions is the paginator options for
// DescribeStaleSecurityGroups
type DescribeStaleSecurityGroupsPaginatorOptions struct {
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

// DescribeStaleSecurityGroupsPaginator is a paginator for
// DescribeStaleSecurityGroups
type DescribeStaleSecurityGroupsPaginator struct {
	options   DescribeStaleSecurityGroupsPaginatorOptions
	client    DescribeStaleSecurityGroupsAPIClient
	params    *DescribeStaleSecurityGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeStaleSecurityGroupsPaginator returns a new
// DescribeStaleSecurityGroupsPaginator
func NewDescribeStaleSecurityGroupsPaginator(client DescribeStaleSecurityGroupsAPIClient, params *DescribeStaleSecurityGroupsInput, optFns ...func(*DescribeStaleSecurityGroupsPaginatorOptions)) *DescribeStaleSecurityGroupsPaginator {
	if params == nil {
		params = &DescribeStaleSecurityGroupsInput{}
	}

	options := DescribeStaleSecurityGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeStaleSecurityGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeStaleSecurityGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeStaleSecurityGroups page.
func (p *DescribeStaleSecurityGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeStaleSecurityGroupsOutput, error) {
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

	result, err := p.client.DescribeStaleSecurityGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeStaleSecurityGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStaleSecurityGroups",
	}
}
