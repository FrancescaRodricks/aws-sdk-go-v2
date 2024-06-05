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

// Returns a list of orderable cluster options. Before you create a new cluster
// you can use this operation to find what options are available, such as the EC2
// Availability Zones (AZ) in the specific Amazon Web Services Region that you can
// specify, and the node types you can request. The node types differ by available
// storage, memory, CPU and price. With the cost involved you might want to obtain
// a list of cluster options in the specific region and specify values when
// creating a cluster. For more information about managing clusters, go to [Amazon Redshift Clusters]in the
// Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
func (c *Client) DescribeOrderableClusterOptions(ctx context.Context, params *DescribeOrderableClusterOptionsInput, optFns ...func(*Options)) (*DescribeOrderableClusterOptionsOutput, error) {
	if params == nil {
		params = &DescribeOrderableClusterOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrderableClusterOptions", params, optFns, c.addOperationDescribeOrderableClusterOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrderableClusterOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrderableClusterOptionsInput struct {

	// The version filter value. Specify this parameter to show only the available
	// offerings matching the specified version.
	//
	// Default: All versions.
	//
	// Constraints: Must be one of the version returned from DescribeClusterVersions.
	ClusterVersion *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeOrderableClusterOptionsrequest exceed the value specified in
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

	// The node type filter value. Specify this parameter to show only the available
	// offerings matching the specified node type.
	NodeType *string

	noSmithyDocumentSerde
}

// Contains the output from the DescribeOrderableClusterOptions action.
type DescribeOrderableClusterOptionsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// An OrderableClusterOption structure containing information about orderable
	// options for the cluster.
	OrderableClusterOptions []types.OrderableClusterOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrderableClusterOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOrderableClusterOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOrderableClusterOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeOrderableClusterOptions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrderableClusterOptions(options.Region), middleware.Before); err != nil {
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

// DescribeOrderableClusterOptionsAPIClient is a client that implements the
// DescribeOrderableClusterOptions operation.
type DescribeOrderableClusterOptionsAPIClient interface {
	DescribeOrderableClusterOptions(context.Context, *DescribeOrderableClusterOptionsInput, ...func(*Options)) (*DescribeOrderableClusterOptionsOutput, error)
}

var _ DescribeOrderableClusterOptionsAPIClient = (*Client)(nil)

// DescribeOrderableClusterOptionsPaginatorOptions is the paginator options for
// DescribeOrderableClusterOptions
type DescribeOrderableClusterOptionsPaginatorOptions struct {
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

// DescribeOrderableClusterOptionsPaginator is a paginator for
// DescribeOrderableClusterOptions
type DescribeOrderableClusterOptionsPaginator struct {
	options   DescribeOrderableClusterOptionsPaginatorOptions
	client    DescribeOrderableClusterOptionsAPIClient
	params    *DescribeOrderableClusterOptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrderableClusterOptionsPaginator returns a new
// DescribeOrderableClusterOptionsPaginator
func NewDescribeOrderableClusterOptionsPaginator(client DescribeOrderableClusterOptionsAPIClient, params *DescribeOrderableClusterOptionsInput, optFns ...func(*DescribeOrderableClusterOptionsPaginatorOptions)) *DescribeOrderableClusterOptionsPaginator {
	if params == nil {
		params = &DescribeOrderableClusterOptionsInput{}
	}

	options := DescribeOrderableClusterOptionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrderableClusterOptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrderableClusterOptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrderableClusterOptions page.
func (p *DescribeOrderableClusterOptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrderableClusterOptionsOutput, error) {
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

	result, err := p.client.DescribeOrderableClusterOptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOrderableClusterOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeOrderableClusterOptions",
	}
}
