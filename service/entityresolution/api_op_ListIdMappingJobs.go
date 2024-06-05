// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all ID mapping jobs for a given workflow.
func (c *Client) ListIdMappingJobs(ctx context.Context, params *ListIdMappingJobsInput, optFns ...func(*Options)) (*ListIdMappingJobsOutput, error) {
	if params == nil {
		params = &ListIdMappingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdMappingJobs", params, optFns, c.addOperationListIdMappingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdMappingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIdMappingJobsInput struct {

	// The name of the workflow to be retrieved.
	//
	// This member is required.
	WorkflowName *string

	// The maximum number of objects returned per page.
	MaxResults *int32

	// The pagination token from the previous API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListIdMappingJobsOutput struct {

	// A list of JobSummary objects.
	Jobs []types.JobSummary

	// The pagination token from the previous API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIdMappingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIdMappingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIdMappingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIdMappingJobs"); err != nil {
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
	if err = addOpListIdMappingJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIdMappingJobs(options.Region), middleware.Before); err != nil {
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

// ListIdMappingJobsAPIClient is a client that implements the ListIdMappingJobs
// operation.
type ListIdMappingJobsAPIClient interface {
	ListIdMappingJobs(context.Context, *ListIdMappingJobsInput, ...func(*Options)) (*ListIdMappingJobsOutput, error)
}

var _ ListIdMappingJobsAPIClient = (*Client)(nil)

// ListIdMappingJobsPaginatorOptions is the paginator options for ListIdMappingJobs
type ListIdMappingJobsPaginatorOptions struct {
	// The maximum number of objects returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIdMappingJobsPaginator is a paginator for ListIdMappingJobs
type ListIdMappingJobsPaginator struct {
	options   ListIdMappingJobsPaginatorOptions
	client    ListIdMappingJobsAPIClient
	params    *ListIdMappingJobsInput
	nextToken *string
	firstPage bool
}

// NewListIdMappingJobsPaginator returns a new ListIdMappingJobsPaginator
func NewListIdMappingJobsPaginator(client ListIdMappingJobsAPIClient, params *ListIdMappingJobsInput, optFns ...func(*ListIdMappingJobsPaginatorOptions)) *ListIdMappingJobsPaginator {
	if params == nil {
		params = &ListIdMappingJobsInput{}
	}

	options := ListIdMappingJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIdMappingJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIdMappingJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIdMappingJobs page.
func (p *ListIdMappingJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIdMappingJobsOutput, error) {
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

	result, err := p.client.ListIdMappingJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIdMappingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIdMappingJobs",
	}
}
