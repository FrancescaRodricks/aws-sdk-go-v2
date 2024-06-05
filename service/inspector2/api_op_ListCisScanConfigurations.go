// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists CIS scan configurations.
func (c *Client) ListCisScanConfigurations(ctx context.Context, params *ListCisScanConfigurationsInput, optFns ...func(*Options)) (*ListCisScanConfigurationsOutput, error) {
	if params == nil {
		params = &ListCisScanConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCisScanConfigurations", params, optFns, c.addOperationListCisScanConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCisScanConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCisScanConfigurationsInput struct {

	// The CIS scan configuration filter criteria.
	FilterCriteria *types.ListCisScanConfigurationsFilterCriteria

	// The maximum number of CIS scan configurations to be returned in a single page
	// of results.
	MaxResults *int32

	// The pagination token from a previous request that's used to retrieve the next
	// page of results.
	NextToken *string

	// The CIS scan configuration sort by order.
	SortBy types.CisScanConfigurationsSortBy

	// The CIS scan configuration sort order order.
	SortOrder types.CisSortOrder

	noSmithyDocumentSerde
}

type ListCisScanConfigurationsOutput struct {

	// The pagination token from a previous request that's used to retrieve the next
	// page of results.
	NextToken *string

	// The CIS scan configuration scan configurations.
	ScanConfigurations []types.CisScanConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCisScanConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCisScanConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCisScanConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCisScanConfigurations"); err != nil {
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
	if err = addOpListCisScanConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCisScanConfigurations(options.Region), middleware.Before); err != nil {
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

// ListCisScanConfigurationsAPIClient is a client that implements the
// ListCisScanConfigurations operation.
type ListCisScanConfigurationsAPIClient interface {
	ListCisScanConfigurations(context.Context, *ListCisScanConfigurationsInput, ...func(*Options)) (*ListCisScanConfigurationsOutput, error)
}

var _ ListCisScanConfigurationsAPIClient = (*Client)(nil)

// ListCisScanConfigurationsPaginatorOptions is the paginator options for
// ListCisScanConfigurations
type ListCisScanConfigurationsPaginatorOptions struct {
	// The maximum number of CIS scan configurations to be returned in a single page
	// of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCisScanConfigurationsPaginator is a paginator for ListCisScanConfigurations
type ListCisScanConfigurationsPaginator struct {
	options   ListCisScanConfigurationsPaginatorOptions
	client    ListCisScanConfigurationsAPIClient
	params    *ListCisScanConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListCisScanConfigurationsPaginator returns a new
// ListCisScanConfigurationsPaginator
func NewListCisScanConfigurationsPaginator(client ListCisScanConfigurationsAPIClient, params *ListCisScanConfigurationsInput, optFns ...func(*ListCisScanConfigurationsPaginatorOptions)) *ListCisScanConfigurationsPaginator {
	if params == nil {
		params = &ListCisScanConfigurationsInput{}
	}

	options := ListCisScanConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCisScanConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCisScanConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCisScanConfigurations page.
func (p *ListCisScanConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCisScanConfigurationsOutput, error) {
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

	result, err := p.client.ListCisScanConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCisScanConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCisScanConfigurations",
	}
}
