// Code generated by smithy-go-codegen DO NOT EDIT.

package fis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resolved targets information of the specified experiment.
func (c *Client) ListExperimentResolvedTargets(ctx context.Context, params *ListExperimentResolvedTargetsInput, optFns ...func(*Options)) (*ListExperimentResolvedTargetsOutput, error) {
	if params == nil {
		params = &ListExperimentResolvedTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExperimentResolvedTargets", params, optFns, c.addOperationListExperimentResolvedTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExperimentResolvedTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExperimentResolvedTargetsInput struct {

	// The ID of the experiment.
	//
	// This member is required.
	ExperimentId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The name of the target.
	TargetName *string

	noSmithyDocumentSerde
}

type ListExperimentResolvedTargetsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The resolved targets.
	ResolvedTargets []types.ResolvedTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExperimentResolvedTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExperimentResolvedTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExperimentResolvedTargets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListExperimentResolvedTargets"); err != nil {
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
	if err = addOpListExperimentResolvedTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExperimentResolvedTargets(options.Region), middleware.Before); err != nil {
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

// ListExperimentResolvedTargetsAPIClient is a client that implements the
// ListExperimentResolvedTargets operation.
type ListExperimentResolvedTargetsAPIClient interface {
	ListExperimentResolvedTargets(context.Context, *ListExperimentResolvedTargetsInput, ...func(*Options)) (*ListExperimentResolvedTargetsOutput, error)
}

var _ ListExperimentResolvedTargetsAPIClient = (*Client)(nil)

// ListExperimentResolvedTargetsPaginatorOptions is the paginator options for
// ListExperimentResolvedTargets
type ListExperimentResolvedTargetsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExperimentResolvedTargetsPaginator is a paginator for
// ListExperimentResolvedTargets
type ListExperimentResolvedTargetsPaginator struct {
	options   ListExperimentResolvedTargetsPaginatorOptions
	client    ListExperimentResolvedTargetsAPIClient
	params    *ListExperimentResolvedTargetsInput
	nextToken *string
	firstPage bool
}

// NewListExperimentResolvedTargetsPaginator returns a new
// ListExperimentResolvedTargetsPaginator
func NewListExperimentResolvedTargetsPaginator(client ListExperimentResolvedTargetsAPIClient, params *ListExperimentResolvedTargetsInput, optFns ...func(*ListExperimentResolvedTargetsPaginatorOptions)) *ListExperimentResolvedTargetsPaginator {
	if params == nil {
		params = &ListExperimentResolvedTargetsInput{}
	}

	options := ListExperimentResolvedTargetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExperimentResolvedTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExperimentResolvedTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExperimentResolvedTargets page.
func (p *ListExperimentResolvedTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExperimentResolvedTargetsOutput, error) {
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

	result, err := p.client.ListExperimentResolvedTargets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExperimentResolvedTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListExperimentResolvedTargets",
	}
}
