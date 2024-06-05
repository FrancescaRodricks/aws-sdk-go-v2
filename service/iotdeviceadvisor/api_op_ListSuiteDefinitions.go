// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdeviceadvisor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Device Advisor test suites you have created.
//
// Requires permission to access the [ListSuiteDefinitions] action.
//
// [ListSuiteDefinitions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListSuiteDefinitions(ctx context.Context, params *ListSuiteDefinitionsInput, optFns ...func(*Options)) (*ListSuiteDefinitionsOutput, error) {
	if params == nil {
		params = &ListSuiteDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSuiteDefinitions", params, optFns, c.addOperationListSuiteDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSuiteDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSuiteDefinitionsInput struct {

	// The maximum number of results to return at once.
	MaxResults *int32

	// A token used to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSuiteDefinitionsOutput struct {

	// A token used to get the next set of results.
	NextToken *string

	// An array of objects that provide summaries of information about the suite
	// definitions in the list.
	SuiteDefinitionInformationList []types.SuiteDefinitionInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSuiteDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSuiteDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSuiteDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSuiteDefinitions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSuiteDefinitions(options.Region), middleware.Before); err != nil {
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

// ListSuiteDefinitionsAPIClient is a client that implements the
// ListSuiteDefinitions operation.
type ListSuiteDefinitionsAPIClient interface {
	ListSuiteDefinitions(context.Context, *ListSuiteDefinitionsInput, ...func(*Options)) (*ListSuiteDefinitionsOutput, error)
}

var _ ListSuiteDefinitionsAPIClient = (*Client)(nil)

// ListSuiteDefinitionsPaginatorOptions is the paginator options for
// ListSuiteDefinitions
type ListSuiteDefinitionsPaginatorOptions struct {
	// The maximum number of results to return at once.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSuiteDefinitionsPaginator is a paginator for ListSuiteDefinitions
type ListSuiteDefinitionsPaginator struct {
	options   ListSuiteDefinitionsPaginatorOptions
	client    ListSuiteDefinitionsAPIClient
	params    *ListSuiteDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListSuiteDefinitionsPaginator returns a new ListSuiteDefinitionsPaginator
func NewListSuiteDefinitionsPaginator(client ListSuiteDefinitionsAPIClient, params *ListSuiteDefinitionsInput, optFns ...func(*ListSuiteDefinitionsPaginatorOptions)) *ListSuiteDefinitionsPaginator {
	if params == nil {
		params = &ListSuiteDefinitionsInput{}
	}

	options := ListSuiteDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSuiteDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSuiteDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSuiteDefinitions page.
func (p *ListSuiteDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSuiteDefinitionsOutput, error) {
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

	result, err := p.client.ListSuiteDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSuiteDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSuiteDefinitions",
	}
}
