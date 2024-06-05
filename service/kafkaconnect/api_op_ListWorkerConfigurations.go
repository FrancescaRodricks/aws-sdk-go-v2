// Code generated by smithy-go-codegen DO NOT EDIT.

package kafkaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all of the worker configurations in this account and Region.
func (c *Client) ListWorkerConfigurations(ctx context.Context, params *ListWorkerConfigurationsInput, optFns ...func(*Options)) (*ListWorkerConfigurationsOutput, error) {
	if params == nil {
		params = &ListWorkerConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkerConfigurations", params, optFns, c.addOperationListWorkerConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkerConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkerConfigurationsInput struct {

	// The maximum number of worker configurations to list in one response.
	MaxResults int32

	// Lists worker configuration names that start with the specified text string.
	NamePrefix *string

	// If the response of a ListWorkerConfigurations operation is truncated, it will
	// include a NextToken. Send this NextToken in a subsequent request to continue
	// listing from where the previous operation left off.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkerConfigurationsOutput struct {

	// If the response of a ListWorkerConfigurations operation is truncated, it will
	// include a NextToken. Send this NextToken in a subsequent request to continue
	// listing from where the previous operation left off.
	NextToken *string

	// An array of worker configuration descriptions.
	WorkerConfigurations []types.WorkerConfigurationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkerConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkerConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkerConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkerConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkerConfigurations(options.Region), middleware.Before); err != nil {
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

// ListWorkerConfigurationsAPIClient is a client that implements the
// ListWorkerConfigurations operation.
type ListWorkerConfigurationsAPIClient interface {
	ListWorkerConfigurations(context.Context, *ListWorkerConfigurationsInput, ...func(*Options)) (*ListWorkerConfigurationsOutput, error)
}

var _ ListWorkerConfigurationsAPIClient = (*Client)(nil)

// ListWorkerConfigurationsPaginatorOptions is the paginator options for
// ListWorkerConfigurations
type ListWorkerConfigurationsPaginatorOptions struct {
	// The maximum number of worker configurations to list in one response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkerConfigurationsPaginator is a paginator for ListWorkerConfigurations
type ListWorkerConfigurationsPaginator struct {
	options   ListWorkerConfigurationsPaginatorOptions
	client    ListWorkerConfigurationsAPIClient
	params    *ListWorkerConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListWorkerConfigurationsPaginator returns a new
// ListWorkerConfigurationsPaginator
func NewListWorkerConfigurationsPaginator(client ListWorkerConfigurationsAPIClient, params *ListWorkerConfigurationsInput, optFns ...func(*ListWorkerConfigurationsPaginatorOptions)) *ListWorkerConfigurationsPaginator {
	if params == nil {
		params = &ListWorkerConfigurationsInput{}
	}

	options := ListWorkerConfigurationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkerConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkerConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkerConfigurations page.
func (p *ListWorkerConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkerConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListWorkerConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkerConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkerConfigurations",
	}
}
