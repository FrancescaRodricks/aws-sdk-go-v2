// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns one or more cluster subnet group objects, which contain metadata about
// your cluster subnet groups. By default, this operation returns information about
// all cluster subnet groups that are defined in your Amazon Web Services account.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all subnet groups that match any combination of the specified
// keys and values. For example, if you have owner and environment for tag keys,
// and admin and test for tag values, all subnet groups that have any combination
// of those values are returned.
//
// If both tag keys and values are omitted from the request, subnet groups are
// returned regardless of whether they have tag keys or values associated with
// them.
func (c *Client) DescribeClusterSubnetGroups(ctx context.Context, params *DescribeClusterSubnetGroupsInput, optFns ...func(*Options)) (*DescribeClusterSubnetGroupsOutput, error) {
	if params == nil {
		params = &DescribeClusterSubnetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusterSubnetGroups", params, optFns, c.addOperationDescribeClusterSubnetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterSubnetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterSubnetGroupsInput struct {

	// The name of the cluster subnet group for which information is requested.
	ClusterSubnetGroupName *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterSubnetGroupsrequest exceed the value specified in
	// MaxRecords , Amazon Web Services returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// A tag key or keys for which you want to return all matching cluster subnet
	// groups that are associated with the specified key or keys. For example, suppose
	// that you have subnet groups that are tagged with keys called owner and
	// environment . If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the subnet groups that have either or both of
	// these tag keys associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching cluster subnet
	// groups that are associated with the specified tag value or values. For example,
	// suppose that you have subnet groups that are tagged with values called admin
	// and test . If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the subnet groups that have either or both of
	// these tag values associated with them.
	TagValues []string

	noSmithyDocumentSerde
}

// Contains the output from the DescribeClusterSubnetGroups action.
type DescribeClusterSubnetGroupsOutput struct {

	// A list of ClusterSubnetGroup instances.
	ClusterSubnetGroups []types.ClusterSubnetGroup

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterSubnetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterSubnetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterSubnetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeClusterSubnetGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterSubnetGroups(options.Region), middleware.Before); err != nil {
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

// DescribeClusterSubnetGroupsAPIClient is a client that implements the
// DescribeClusterSubnetGroups operation.
type DescribeClusterSubnetGroupsAPIClient interface {
	DescribeClusterSubnetGroups(context.Context, *DescribeClusterSubnetGroupsInput, ...func(*Options)) (*DescribeClusterSubnetGroupsOutput, error)
}

var _ DescribeClusterSubnetGroupsAPIClient = (*Client)(nil)

// DescribeClusterSubnetGroupsPaginatorOptions is the paginator options for
// DescribeClusterSubnetGroups
type DescribeClusterSubnetGroupsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClusterSubnetGroupsPaginator is a paginator for
// DescribeClusterSubnetGroups
type DescribeClusterSubnetGroupsPaginator struct {
	options   DescribeClusterSubnetGroupsPaginatorOptions
	client    DescribeClusterSubnetGroupsAPIClient
	params    *DescribeClusterSubnetGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeClusterSubnetGroupsPaginator returns a new
// DescribeClusterSubnetGroupsPaginator
func NewDescribeClusterSubnetGroupsPaginator(client DescribeClusterSubnetGroupsAPIClient, params *DescribeClusterSubnetGroupsInput, optFns ...func(*DescribeClusterSubnetGroupsPaginatorOptions)) *DescribeClusterSubnetGroupsPaginator {
	if params == nil {
		params = &DescribeClusterSubnetGroupsInput{}
	}

	options := DescribeClusterSubnetGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClusterSubnetGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClusterSubnetGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClusterSubnetGroups page.
func (p *DescribeClusterSubnetGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClusterSubnetGroupsOutput, error) {
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

	result, err := p.client.DescribeClusterSubnetGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeClusterSubnetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeClusterSubnetGroups",
	}
}
