// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists notebook instance lifestyle configurations created with the [CreateNotebookInstanceLifecycleConfig] API.
//
// [CreateNotebookInstanceLifecycleConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateNotebookInstanceLifecycleConfig.html
func (c *Client) ListNotebookInstanceLifecycleConfigs(ctx context.Context, params *ListNotebookInstanceLifecycleConfigsInput, optFns ...func(*Options)) (*ListNotebookInstanceLifecycleConfigsOutput, error) {
	if params == nil {
		params = &ListNotebookInstanceLifecycleConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotebookInstanceLifecycleConfigs", params, optFns, c.addOperationListNotebookInstanceLifecycleConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotebookInstanceLifecycleConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotebookInstanceLifecycleConfigsInput struct {

	// A filter that returns only lifecycle configurations that were created after the
	// specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only lifecycle configurations that were created before
	// the specified time (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only lifecycle configurations that were modified after
	// the specified time (timestamp).
	LastModifiedTimeAfter *time.Time

	// A filter that returns only lifecycle configurations that were modified before
	// the specified time (timestamp).
	LastModifiedTimeBefore *time.Time

	// The maximum number of lifecycle configurations to return in the response.
	MaxResults *int32

	// A string in the lifecycle configuration name. This filter returns only
	// lifecycle configurations whose name contains the specified string.
	NameContains *string

	// If the result of a ListNotebookInstanceLifecycleConfigs request was truncated,
	// the response includes a NextToken . To get the next set of lifecycle
	// configurations, use the token in the next request.
	NextToken *string

	// Sorts the list of results. The default is CreationTime .
	SortBy types.NotebookInstanceLifecycleConfigSortKey

	// The sort order for results.
	SortOrder types.NotebookInstanceLifecycleConfigSortOrder

	noSmithyDocumentSerde
}

type ListNotebookInstanceLifecycleConfigsOutput struct {

	// If the response is truncated, SageMaker returns this token. To get the next set
	// of lifecycle configurations, use it in the next request.
	NextToken *string

	// An array of NotebookInstanceLifecycleConfiguration objects, each listing a
	// lifecycle configuration.
	NotebookInstanceLifecycleConfigs []types.NotebookInstanceLifecycleConfigSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotebookInstanceLifecycleConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNotebookInstanceLifecycleConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNotebookInstanceLifecycleConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNotebookInstanceLifecycleConfigs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotebookInstanceLifecycleConfigs(options.Region), middleware.Before); err != nil {
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

// ListNotebookInstanceLifecycleConfigsAPIClient is a client that implements the
// ListNotebookInstanceLifecycleConfigs operation.
type ListNotebookInstanceLifecycleConfigsAPIClient interface {
	ListNotebookInstanceLifecycleConfigs(context.Context, *ListNotebookInstanceLifecycleConfigsInput, ...func(*Options)) (*ListNotebookInstanceLifecycleConfigsOutput, error)
}

var _ ListNotebookInstanceLifecycleConfigsAPIClient = (*Client)(nil)

// ListNotebookInstanceLifecycleConfigsPaginatorOptions is the paginator options
// for ListNotebookInstanceLifecycleConfigs
type ListNotebookInstanceLifecycleConfigsPaginatorOptions struct {
	// The maximum number of lifecycle configurations to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNotebookInstanceLifecycleConfigsPaginator is a paginator for
// ListNotebookInstanceLifecycleConfigs
type ListNotebookInstanceLifecycleConfigsPaginator struct {
	options   ListNotebookInstanceLifecycleConfigsPaginatorOptions
	client    ListNotebookInstanceLifecycleConfigsAPIClient
	params    *ListNotebookInstanceLifecycleConfigsInput
	nextToken *string
	firstPage bool
}

// NewListNotebookInstanceLifecycleConfigsPaginator returns a new
// ListNotebookInstanceLifecycleConfigsPaginator
func NewListNotebookInstanceLifecycleConfigsPaginator(client ListNotebookInstanceLifecycleConfigsAPIClient, params *ListNotebookInstanceLifecycleConfigsInput, optFns ...func(*ListNotebookInstanceLifecycleConfigsPaginatorOptions)) *ListNotebookInstanceLifecycleConfigsPaginator {
	if params == nil {
		params = &ListNotebookInstanceLifecycleConfigsInput{}
	}

	options := ListNotebookInstanceLifecycleConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNotebookInstanceLifecycleConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNotebookInstanceLifecycleConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNotebookInstanceLifecycleConfigs page.
func (p *ListNotebookInstanceLifecycleConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNotebookInstanceLifecycleConfigsOutput, error) {
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

	result, err := p.client.ListNotebookInstanceLifecycleConfigs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNotebookInstanceLifecycleConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNotebookInstanceLifecycleConfigs",
	}
}
