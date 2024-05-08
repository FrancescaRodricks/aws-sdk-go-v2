// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of limitations for recommendations of target Amazon
// Web Services engines.
func (c *Client) DescribeRecommendationLimitations(ctx context.Context, params *DescribeRecommendationLimitationsInput, optFns ...func(*Options)) (*DescribeRecommendationLimitationsOutput, error) {
	if params == nil {
		params = &DescribeRecommendationLimitationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecommendationLimitations", params, optFns, c.addOperationDescribeRecommendationLimitationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecommendationLimitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecommendationLimitationsInput struct {

	// Filters applied to the limitations described in the form of key-value pairs.
	Filters []types.Filter

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, Fleet Advisor includes a pagination token
	// in the response so that you can retrieve the remaining results.
	MaxRecords *int32

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If NextToken is returned by a previous response, there are more results
	// available. The value of NextToken is a unique pagination token for each page.
	// Make the call again using the returned token to retrieve the next page. Keep all
	// other arguments unchanged.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeRecommendationLimitationsOutput struct {

	// The list of limitations for recommendations of target Amazon Web Services
	// engines.
	Limitations []types.Limitation

	// The unique pagination token returned for you to pass to a subsequent request.
	// Fleet Advisor returns this token when the number of records in the response is
	// greater than the MaxRecords value. To retrieve the next page, make the call
	// again using the returned token and keeping all other arguments unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecommendationLimitationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRecommendationLimitations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRecommendationLimitations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRecommendationLimitations"); err != nil {
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
	if err = addOpDescribeRecommendationLimitationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecommendationLimitations(options.Region), middleware.Before); err != nil {
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

// DescribeRecommendationLimitationsAPIClient is a client that implements the
// DescribeRecommendationLimitations operation.
type DescribeRecommendationLimitationsAPIClient interface {
	DescribeRecommendationLimitations(context.Context, *DescribeRecommendationLimitationsInput, ...func(*Options)) (*DescribeRecommendationLimitationsOutput, error)
}

var _ DescribeRecommendationLimitationsAPIClient = (*Client)(nil)

// DescribeRecommendationLimitationsPaginatorOptions is the paginator options for
// DescribeRecommendationLimitations
type DescribeRecommendationLimitationsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, Fleet Advisor includes a pagination token
	// in the response so that you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRecommendationLimitationsPaginator is a paginator for
// DescribeRecommendationLimitations
type DescribeRecommendationLimitationsPaginator struct {
	options   DescribeRecommendationLimitationsPaginatorOptions
	client    DescribeRecommendationLimitationsAPIClient
	params    *DescribeRecommendationLimitationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRecommendationLimitationsPaginator returns a new
// DescribeRecommendationLimitationsPaginator
func NewDescribeRecommendationLimitationsPaginator(client DescribeRecommendationLimitationsAPIClient, params *DescribeRecommendationLimitationsInput, optFns ...func(*DescribeRecommendationLimitationsPaginatorOptions)) *DescribeRecommendationLimitationsPaginator {
	if params == nil {
		params = &DescribeRecommendationLimitationsInput{}
	}

	options := DescribeRecommendationLimitationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRecommendationLimitationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRecommendationLimitationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRecommendationLimitations page.
func (p *DescribeRecommendationLimitationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRecommendationLimitationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeRecommendationLimitations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRecommendationLimitations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRecommendationLimitations",
	}
}
