// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DBSecurityGroup descriptions. If a DBSecurityGroupName is
// specified, the list will contain only the descriptions of the specified DB
// security group.
//
// EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see [Migrate from EC2-Classic to a VPC]in the Amazon EC2 User Guide, the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare], and [Moving a DB instance not in a VPC into a VPC] in the
// Amazon RDS User Guide.
//
// [Migrate from EC2-Classic to a VPC]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html
// [EC2-Classic Networking is Retiring – Here’s How to Prepare]: http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/
// [Moving a DB instance not in a VPC into a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html
func (c *Client) DescribeDBSecurityGroups(ctx context.Context, params *DescribeDBSecurityGroupsInput, optFns ...func(*Options)) (*DescribeDBSecurityGroupsOutput, error) {
	if params == nil {
		params = &DescribeDBSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBSecurityGroups", params, optFns, c.addOperationDescribeDBSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBSecurityGroupsInput struct {

	// The name of the DB security group to return details for.
	DBSecurityGroupName *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBSecurityGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the DescribeDBSecurityGroups
// action.
type DescribeDBSecurityGroupsOutput struct {

	// A list of DBSecurityGroup instances.
	DBSecurityGroups []types.DBSecurityGroup

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBSecurityGroups"); err != nil {
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
	if err = addOpDescribeDBSecurityGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBSecurityGroups(options.Region), middleware.Before); err != nil {
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

// DescribeDBSecurityGroupsAPIClient is a client that implements the
// DescribeDBSecurityGroups operation.
type DescribeDBSecurityGroupsAPIClient interface {
	DescribeDBSecurityGroups(context.Context, *DescribeDBSecurityGroupsInput, ...func(*Options)) (*DescribeDBSecurityGroupsOutput, error)
}

var _ DescribeDBSecurityGroupsAPIClient = (*Client)(nil)

// DescribeDBSecurityGroupsPaginatorOptions is the paginator options for
// DescribeDBSecurityGroups
type DescribeDBSecurityGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBSecurityGroupsPaginator is a paginator for DescribeDBSecurityGroups
type DescribeDBSecurityGroupsPaginator struct {
	options   DescribeDBSecurityGroupsPaginatorOptions
	client    DescribeDBSecurityGroupsAPIClient
	params    *DescribeDBSecurityGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBSecurityGroupsPaginator returns a new
// DescribeDBSecurityGroupsPaginator
func NewDescribeDBSecurityGroupsPaginator(client DescribeDBSecurityGroupsAPIClient, params *DescribeDBSecurityGroupsInput, optFns ...func(*DescribeDBSecurityGroupsPaginatorOptions)) *DescribeDBSecurityGroupsPaginator {
	if params == nil {
		params = &DescribeDBSecurityGroupsInput{}
	}

	options := DescribeDBSecurityGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBSecurityGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBSecurityGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBSecurityGroups page.
func (p *DescribeDBSecurityGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBSecurityGroupsOutput, error) {
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

	result, err := p.client.DescribeDBSecurityGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDBSecurityGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBSecurityGroups",
	}
}
