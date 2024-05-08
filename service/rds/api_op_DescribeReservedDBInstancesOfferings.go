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

// Lists available reserved DB instance offerings.
func (c *Client) DescribeReservedDBInstancesOfferings(ctx context.Context, params *DescribeReservedDBInstancesOfferingsInput, optFns ...func(*Options)) (*DescribeReservedDBInstancesOfferingsOutput, error) {
	if params == nil {
		params = &DescribeReservedDBInstancesOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedDBInstancesOfferings", params, optFns, c.addOperationDescribeReservedDBInstancesOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedDBInstancesOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReservedDBInstancesOfferingsInput struct {

	// The DB instance class filter value. Specify this parameter to show only the
	// available offerings matching the specified DB instance class.
	DBInstanceClass *string

	// Duration filter value, specified in years or seconds. Specify this parameter to
	// show only reservations for this duration.
	//
	// Valid Values: 1 | 3 | 31536000 | 94608000
	Duration *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// Specifies whether to show only those reservations that support Multi-AZ.
	MultiAZ *bool

	// The offering type filter value. Specify this parameter to show only the
	// available offerings matching the specified offering type.
	//
	// Valid Values: "Partial Upfront" | "All Upfront" | "No Upfront"
	OfferingType *string

	// Product description filter value. Specify this parameter to show only the
	// available offerings that contain the specified product description.
	//
	// The results show offerings that partially match the filter value.
	ProductDescription *string

	// The offering identifier filter value. Specify this parameter to show only the
	// available offering that matches the specified reservation identifier.
	//
	// Example: 438012d3-4052-4cc7-b2e3-8d3372e0e706
	ReservedDBInstancesOfferingId *string

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the
// DescribeReservedDBInstancesOfferings action.
type DescribeReservedDBInstancesOfferingsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// A list of reserved DB instance offerings.
	ReservedDBInstancesOfferings []types.ReservedDBInstancesOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedDBInstancesOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReservedDBInstancesOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReservedDBInstancesOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReservedDBInstancesOfferings"); err != nil {
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
	if err = addOpDescribeReservedDBInstancesOfferingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedDBInstancesOfferings(options.Region), middleware.Before); err != nil {
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

// DescribeReservedDBInstancesOfferingsAPIClient is a client that implements the
// DescribeReservedDBInstancesOfferings operation.
type DescribeReservedDBInstancesOfferingsAPIClient interface {
	DescribeReservedDBInstancesOfferings(context.Context, *DescribeReservedDBInstancesOfferingsInput, ...func(*Options)) (*DescribeReservedDBInstancesOfferingsOutput, error)
}

var _ DescribeReservedDBInstancesOfferingsAPIClient = (*Client)(nil)

// DescribeReservedDBInstancesOfferingsPaginatorOptions is the paginator options
// for DescribeReservedDBInstancesOfferings
type DescribeReservedDBInstancesOfferingsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedDBInstancesOfferingsPaginator is a paginator for
// DescribeReservedDBInstancesOfferings
type DescribeReservedDBInstancesOfferingsPaginator struct {
	options   DescribeReservedDBInstancesOfferingsPaginatorOptions
	client    DescribeReservedDBInstancesOfferingsAPIClient
	params    *DescribeReservedDBInstancesOfferingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedDBInstancesOfferingsPaginator returns a new
// DescribeReservedDBInstancesOfferingsPaginator
func NewDescribeReservedDBInstancesOfferingsPaginator(client DescribeReservedDBInstancesOfferingsAPIClient, params *DescribeReservedDBInstancesOfferingsInput, optFns ...func(*DescribeReservedDBInstancesOfferingsPaginatorOptions)) *DescribeReservedDBInstancesOfferingsPaginator {
	if params == nil {
		params = &DescribeReservedDBInstancesOfferingsInput{}
	}

	options := DescribeReservedDBInstancesOfferingsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedDBInstancesOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedDBInstancesOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedDBInstancesOfferings page.
func (p *DescribeReservedDBInstancesOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedDBInstancesOfferingsOutput, error) {
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

	result, err := p.client.DescribeReservedDBInstancesOfferings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReservedDBInstancesOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReservedDBInstancesOfferings",
	}
}
