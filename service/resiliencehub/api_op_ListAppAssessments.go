// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the assessments for an Resilience Hub application. You can use request
// parameters to refine the results for the response object.
func (c *Client) ListAppAssessments(ctx context.Context, params *ListAppAssessmentsInput, optFns ...func(*Options)) (*ListAppAssessmentsOutput, error) {
	if params == nil {
		params = &ListAppAssessmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppAssessments", params, optFns, c.addOperationListAppAssessmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppAssessmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppAssessmentsInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference
	// guide.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	AppArn *string

	// The name for the assessment.
	AssessmentName *string

	// The current status of the assessment for the resiliency policy.
	AssessmentStatus []types.AssessmentStatus

	// The current status of compliance for the resiliency policy.
	ComplianceStatus types.ComplianceStatus

	// Specifies the entity that invoked a specific assessment, either a User or the
	// System .
	Invoker types.AssessmentInvoker

	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	// The default is to sort by ascending startTime. To sort by descending startTime,
	// set reverseOrder to true .
	ReverseOrder *bool

	noSmithyDocumentSerde
}

type ListAppAssessmentsOutput struct {

	// The summaries for the specified assessments, returned as an object. This object
	// includes application versions, associated Amazon Resource Numbers (ARNs), cost,
	// messages, resiliency scores, and more.
	//
	// This member is required.
	AssessmentSummaries []types.AppAssessmentSummary

	// Token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppAssessmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppAssessments"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppAssessments(options.Region), middleware.Before); err != nil {
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

// ListAppAssessmentsAPIClient is a client that implements the ListAppAssessments
// operation.
type ListAppAssessmentsAPIClient interface {
	ListAppAssessments(context.Context, *ListAppAssessmentsInput, ...func(*Options)) (*ListAppAssessmentsOutput, error)
}

var _ ListAppAssessmentsAPIClient = (*Client)(nil)

// ListAppAssessmentsPaginatorOptions is the paginator options for
// ListAppAssessments
type ListAppAssessmentsPaginatorOptions struct {
	// Maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so
	// that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppAssessmentsPaginator is a paginator for ListAppAssessments
type ListAppAssessmentsPaginator struct {
	options   ListAppAssessmentsPaginatorOptions
	client    ListAppAssessmentsAPIClient
	params    *ListAppAssessmentsInput
	nextToken *string
	firstPage bool
}

// NewListAppAssessmentsPaginator returns a new ListAppAssessmentsPaginator
func NewListAppAssessmentsPaginator(client ListAppAssessmentsAPIClient, params *ListAppAssessmentsInput, optFns ...func(*ListAppAssessmentsPaginatorOptions)) *ListAppAssessmentsPaginator {
	if params == nil {
		params = &ListAppAssessmentsInput{}
	}

	options := ListAppAssessmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppAssessmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppAssessmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppAssessments page.
func (p *ListAppAssessmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppAssessmentsOutput, error) {
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

	result, err := p.client.ListAppAssessments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppAssessments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppAssessments",
	}
}
