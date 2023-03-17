// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the evaluation results for the specified Amazon Web Services resource.
// The results indicate which Config rules were used to evaluate the resource, when
// each rule was last invoked, and whether the resource complies with each rule.
func (c *Client) GetComplianceDetailsByResource(ctx context.Context, params *GetComplianceDetailsByResourceInput, optFns ...func(*Options)) (*GetComplianceDetailsByResourceOutput, error) {
	if params == nil {
		params = &GetComplianceDetailsByResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComplianceDetailsByResource", params, optFns, c.addOperationGetComplianceDetailsByResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComplianceDetailsByResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComplianceDetailsByResourceInput struct {

	// Filters the results by compliance. INSUFFICIENT_DATA is a valid ComplianceType
	// that is returned when an Config rule cannot be evaluated. However,
	// INSUFFICIENT_DATA cannot be used as a ComplianceType for filtering results.
	ComplianceTypes []types.ComplianceType

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The unique ID of Amazon Web Services resource execution for which you want to
	// retrieve evaluation results. You need to only provide either a
	// ResourceEvaluationID or a ResourceID and ResourceType.
	ResourceEvaluationId *string

	// The ID of the Amazon Web Services resource for which you want compliance
	// information.
	ResourceId *string

	// The type of the Amazon Web Services resource for which you want compliance
	// information.
	ResourceType *string

	noSmithyDocumentSerde
}

type GetComplianceDetailsByResourceOutput struct {

	// Indicates whether the specified Amazon Web Services resource complies each
	// Config rule.
	EvaluationResults []types.EvaluationResult

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComplianceDetailsByResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComplianceDetailsByResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComplianceDetailsByResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComplianceDetailsByResource(options.Region), middleware.Before); err != nil {
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
	return nil
}

// GetComplianceDetailsByResourceAPIClient is a client that implements the
// GetComplianceDetailsByResource operation.
type GetComplianceDetailsByResourceAPIClient interface {
	GetComplianceDetailsByResource(context.Context, *GetComplianceDetailsByResourceInput, ...func(*Options)) (*GetComplianceDetailsByResourceOutput, error)
}

var _ GetComplianceDetailsByResourceAPIClient = (*Client)(nil)

// GetComplianceDetailsByResourcePaginatorOptions is the paginator options for
// GetComplianceDetailsByResource
type GetComplianceDetailsByResourcePaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetComplianceDetailsByResourcePaginator is a paginator for
// GetComplianceDetailsByResource
type GetComplianceDetailsByResourcePaginator struct {
	options   GetComplianceDetailsByResourcePaginatorOptions
	client    GetComplianceDetailsByResourceAPIClient
	params    *GetComplianceDetailsByResourceInput
	nextToken *string
	firstPage bool
}

// NewGetComplianceDetailsByResourcePaginator returns a new
// GetComplianceDetailsByResourcePaginator
func NewGetComplianceDetailsByResourcePaginator(client GetComplianceDetailsByResourceAPIClient, params *GetComplianceDetailsByResourceInput, optFns ...func(*GetComplianceDetailsByResourcePaginatorOptions)) *GetComplianceDetailsByResourcePaginator {
	if params == nil {
		params = &GetComplianceDetailsByResourceInput{}
	}

	options := GetComplianceDetailsByResourcePaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetComplianceDetailsByResourcePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetComplianceDetailsByResourcePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetComplianceDetailsByResource page.
func (p *GetComplianceDetailsByResourcePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetComplianceDetailsByResourceOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetComplianceDetailsByResource(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetComplianceDetailsByResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetComplianceDetailsByResource",
	}
}
