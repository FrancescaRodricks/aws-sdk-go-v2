// Code generated by smithy-go-codegen DO NOT EDIT.

package medicalimaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medicalimaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List import jobs created for a specific data store.
func (c *Client) ListDICOMImportJobs(ctx context.Context, params *ListDICOMImportJobsInput, optFns ...func(*Options)) (*ListDICOMImportJobsOutput, error) {
	if params == nil {
		params = &ListDICOMImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDICOMImportJobs", params, optFns, c.addOperationListDICOMImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDICOMImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDICOMImportJobsInput struct {

	// The data store identifier.
	//
	// This member is required.
	DatastoreId *string

	// The filters for listing import jobs based on status.
	JobStatus types.JobStatus

	// The max results count. The upper bound is determined by load testing.
	MaxResults *int32

	// The pagination token used to request the list of import jobs on the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDICOMImportJobsOutput struct {

	// A list of job summaries.
	//
	// This member is required.
	JobSummaries []types.DICOMImportJobSummary

	// The pagination token used to retrieve the list of import jobs on the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDICOMImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDICOMImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDICOMImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDICOMImportJobs"); err != nil {
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
	if err = addOpListDICOMImportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDICOMImportJobs(options.Region), middleware.Before); err != nil {
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

// ListDICOMImportJobsAPIClient is a client that implements the
// ListDICOMImportJobs operation.
type ListDICOMImportJobsAPIClient interface {
	ListDICOMImportJobs(context.Context, *ListDICOMImportJobsInput, ...func(*Options)) (*ListDICOMImportJobsOutput, error)
}

var _ ListDICOMImportJobsAPIClient = (*Client)(nil)

// ListDICOMImportJobsPaginatorOptions is the paginator options for
// ListDICOMImportJobs
type ListDICOMImportJobsPaginatorOptions struct {
	// The max results count. The upper bound is determined by load testing.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDICOMImportJobsPaginator is a paginator for ListDICOMImportJobs
type ListDICOMImportJobsPaginator struct {
	options   ListDICOMImportJobsPaginatorOptions
	client    ListDICOMImportJobsAPIClient
	params    *ListDICOMImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListDICOMImportJobsPaginator returns a new ListDICOMImportJobsPaginator
func NewListDICOMImportJobsPaginator(client ListDICOMImportJobsAPIClient, params *ListDICOMImportJobsInput, optFns ...func(*ListDICOMImportJobsPaginatorOptions)) *ListDICOMImportJobsPaginator {
	if params == nil {
		params = &ListDICOMImportJobsInput{}
	}

	options := ListDICOMImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDICOMImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDICOMImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDICOMImportJobs page.
func (p *ListDICOMImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDICOMImportJobsOutput, error) {
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

	result, err := p.client.ListDICOMImportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDICOMImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDICOMImportJobs",
	}
}
