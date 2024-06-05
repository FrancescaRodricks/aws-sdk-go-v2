// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns configuration details about all the experiments in the specified
// project.
func (c *Client) ListExperiments(ctx context.Context, params *ListExperimentsInput, optFns ...func(*Options)) (*ListExperimentsOutput, error) {
	if params == nil {
		params = &ListExperimentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExperiments", params, optFns, c.addOperationListExperimentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExperimentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExperimentsInput struct {

	// The name or ARN of the project to return the experiment list from.
	//
	// This member is required.
	Project *string

	// The maximum number of results to include in the response.
	MaxResults *int32

	// The token to use when requesting the next set of results. You received this
	// token from a previous ListExperiments operation.
	NextToken *string

	// Use this optional parameter to limit the returned results to only the
	// experiments with the status that you specify here.
	Status types.ExperimentStatus

	noSmithyDocumentSerde
}

type ListExperimentsOutput struct {

	// An array of structures that contain the configuration details of the
	// experiments in the specified project.
	Experiments []types.Experiment

	// The token to use in a subsequent ListExperiments operation to return the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExperimentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExperiments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExperiments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListExperiments"); err != nil {
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
	if err = addOpListExperimentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExperiments(options.Region), middleware.Before); err != nil {
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

// ListExperimentsAPIClient is a client that implements the ListExperiments
// operation.
type ListExperimentsAPIClient interface {
	ListExperiments(context.Context, *ListExperimentsInput, ...func(*Options)) (*ListExperimentsOutput, error)
}

var _ ListExperimentsAPIClient = (*Client)(nil)

// ListExperimentsPaginatorOptions is the paginator options for ListExperiments
type ListExperimentsPaginatorOptions struct {
	// The maximum number of results to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExperimentsPaginator is a paginator for ListExperiments
type ListExperimentsPaginator struct {
	options   ListExperimentsPaginatorOptions
	client    ListExperimentsAPIClient
	params    *ListExperimentsInput
	nextToken *string
	firstPage bool
}

// NewListExperimentsPaginator returns a new ListExperimentsPaginator
func NewListExperimentsPaginator(client ListExperimentsAPIClient, params *ListExperimentsInput, optFns ...func(*ListExperimentsPaginatorOptions)) *ListExperimentsPaginator {
	if params == nil {
		params = &ListExperimentsInput{}
	}

	options := ListExperimentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExperimentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExperimentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExperiments page.
func (p *ListExperimentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExperimentsOutput, error) {
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

	result, err := p.client.ListExperiments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExperiments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListExperiments",
	}
}
