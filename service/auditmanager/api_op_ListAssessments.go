// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of current and past assessments from Audit Manager.
func (c *Client) ListAssessments(ctx context.Context, params *ListAssessmentsInput, optFns ...func(*Options)) (*ListAssessmentsOutput, error) {
	if params == nil {
		params = &ListAssessmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessments", params, optFns, c.addOperationListAssessmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentsInput struct {

	//  Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	//  The pagination token that's used to fetch the next set of results.
	NextToken *string

	//  The current status of the assessment.
	Status types.AssessmentStatus

	noSmithyDocumentSerde
}

type ListAssessmentsOutput struct {

	// The metadata that the ListAssessments API returns for each assessment.
	AssessmentMetadata []types.AssessmentMetadataItem

	//  The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssessmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssessments"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessments(options.Region), middleware.Before); err != nil {
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

// ListAssessmentsAPIClient is a client that implements the ListAssessments
// operation.
type ListAssessmentsAPIClient interface {
	ListAssessments(context.Context, *ListAssessmentsInput, ...func(*Options)) (*ListAssessmentsOutput, error)
}

var _ ListAssessmentsAPIClient = (*Client)(nil)

// ListAssessmentsPaginatorOptions is the paginator options for ListAssessments
type ListAssessmentsPaginatorOptions struct {
	//  Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssessmentsPaginator is a paginator for ListAssessments
type ListAssessmentsPaginator struct {
	options   ListAssessmentsPaginatorOptions
	client    ListAssessmentsAPIClient
	params    *ListAssessmentsInput
	nextToken *string
	firstPage bool
}

// NewListAssessmentsPaginator returns a new ListAssessmentsPaginator
func NewListAssessmentsPaginator(client ListAssessmentsAPIClient, params *ListAssessmentsInput, optFns ...func(*ListAssessmentsPaginatorOptions)) *ListAssessmentsPaginator {
	if params == nil {
		params = &ListAssessmentsInput{}
	}

	options := ListAssessmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssessmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssessmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssessments page.
func (p *ListAssessmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssessmentsOutput, error) {
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

	result, err := p.client.ListAssessments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssessments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssessments",
	}
}
