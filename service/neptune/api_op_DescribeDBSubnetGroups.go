// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DBSubnetGroup descriptions. If a DBSubnetGroupName is
// specified, the list will contain only the descriptions of the specified
// DBSubnetGroup.
//
// For an overview of CIDR ranges, go to the [Wikipedia Tutorial].
//
// [Wikipedia Tutorial]: http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
func (c *Client) DescribeDBSubnetGroups(ctx context.Context, params *DescribeDBSubnetGroupsInput, optFns ...func(*Options)) (*DescribeDBSubnetGroupsOutput, error) {
	if params == nil {
		params = &DescribeDBSubnetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBSubnetGroups", params, optFns, c.addOperationDescribeDBSubnetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBSubnetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBSubnetGroupsInput struct {

	// The name of the DB subnet group to return details for.
	DBSubnetGroupName *string

	// This parameter is not currently supported.
	Filters []types.Filter

	//  An optional pagination token provided by a previous DescribeDBSubnetGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	//  The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeDBSubnetGroupsOutput struct {

	//  A list of DBSubnetGroup instances.
	DBSubnetGroups []types.DBSubnetGroup

	//  An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBSubnetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBSubnetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBSubnetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBSubnetGroups"); err != nil {
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
	if err = addOpDescribeDBSubnetGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBSubnetGroups(options.Region), middleware.Before); err != nil {
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

// DescribeDBSubnetGroupsAPIClient is a client that implements the
// DescribeDBSubnetGroups operation.
type DescribeDBSubnetGroupsAPIClient interface {
	DescribeDBSubnetGroups(context.Context, *DescribeDBSubnetGroupsInput, ...func(*Options)) (*DescribeDBSubnetGroupsOutput, error)
}

var _ DescribeDBSubnetGroupsAPIClient = (*Client)(nil)

// DescribeDBSubnetGroupsPaginatorOptions is the paginator options for
// DescribeDBSubnetGroups
type DescribeDBSubnetGroupsPaginatorOptions struct {
	//  The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBSubnetGroupsPaginator is a paginator for DescribeDBSubnetGroups
type DescribeDBSubnetGroupsPaginator struct {
	options   DescribeDBSubnetGroupsPaginatorOptions
	client    DescribeDBSubnetGroupsAPIClient
	params    *DescribeDBSubnetGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBSubnetGroupsPaginator returns a new DescribeDBSubnetGroupsPaginator
func NewDescribeDBSubnetGroupsPaginator(client DescribeDBSubnetGroupsAPIClient, params *DescribeDBSubnetGroupsInput, optFns ...func(*DescribeDBSubnetGroupsPaginatorOptions)) *DescribeDBSubnetGroupsPaginator {
	if params == nil {
		params = &DescribeDBSubnetGroupsInput{}
	}

	options := DescribeDBSubnetGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBSubnetGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBSubnetGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBSubnetGroups page.
func (p *DescribeDBSubnetGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBSubnetGroupsOutput, error) {
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

	result, err := p.client.DescribeDBSubnetGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDBSubnetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBSubnetGroups",
	}
}
