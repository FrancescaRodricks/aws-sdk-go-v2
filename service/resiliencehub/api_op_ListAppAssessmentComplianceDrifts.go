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

// List of compliance drifts that were detected while running an assessment.
func (c *Client) ListAppAssessmentComplianceDrifts(ctx context.Context, params *ListAppAssessmentComplianceDriftsInput, optFns ...func(*Options)) (*ListAppAssessmentComplianceDriftsOutput, error) {
	if params == nil {
		params = &ListAppAssessmentComplianceDriftsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppAssessmentComplianceDrifts", params, optFns, c.addOperationListAppAssessmentComplianceDriftsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppAssessmentComplianceDriftsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppAssessmentComplianceDriftsInput struct {

	// Amazon Resource Name (ARN) of the assessment. The format for this ARN is: arn:
	// partition :resiliencehub: region : account :app-assessment/ app-id . For more
	// information about ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference guide.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	AssessmentArn *string

	// Indicates the maximum number of applications requested.
	MaxResults *int32

	// Indicates the unique token number of the next application to be checked for
	// compliance and regulatory requirements from the list of applications.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppAssessmentComplianceDriftsOutput struct {

	// Indicates compliance drifts (recovery time objective (RTO) and recovery point
	// objective (RPO)) detected for an assessed entity.
	//
	// This member is required.
	ComplianceDrifts []types.ComplianceDrift

	// Token number of the next application to be checked for compliance and
	// regulatory requirements from the list of applications.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppAssessmentComplianceDriftsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppAssessmentComplianceDrifts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppAssessmentComplianceDrifts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppAssessmentComplianceDrifts"); err != nil {
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
	if err = addOpListAppAssessmentComplianceDriftsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppAssessmentComplianceDrifts(options.Region), middleware.Before); err != nil {
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

// ListAppAssessmentComplianceDriftsAPIClient is a client that implements the
// ListAppAssessmentComplianceDrifts operation.
type ListAppAssessmentComplianceDriftsAPIClient interface {
	ListAppAssessmentComplianceDrifts(context.Context, *ListAppAssessmentComplianceDriftsInput, ...func(*Options)) (*ListAppAssessmentComplianceDriftsOutput, error)
}

var _ ListAppAssessmentComplianceDriftsAPIClient = (*Client)(nil)

// ListAppAssessmentComplianceDriftsPaginatorOptions is the paginator options for
// ListAppAssessmentComplianceDrifts
type ListAppAssessmentComplianceDriftsPaginatorOptions struct {
	// Indicates the maximum number of applications requested.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppAssessmentComplianceDriftsPaginator is a paginator for
// ListAppAssessmentComplianceDrifts
type ListAppAssessmentComplianceDriftsPaginator struct {
	options   ListAppAssessmentComplianceDriftsPaginatorOptions
	client    ListAppAssessmentComplianceDriftsAPIClient
	params    *ListAppAssessmentComplianceDriftsInput
	nextToken *string
	firstPage bool
}

// NewListAppAssessmentComplianceDriftsPaginator returns a new
// ListAppAssessmentComplianceDriftsPaginator
func NewListAppAssessmentComplianceDriftsPaginator(client ListAppAssessmentComplianceDriftsAPIClient, params *ListAppAssessmentComplianceDriftsInput, optFns ...func(*ListAppAssessmentComplianceDriftsPaginatorOptions)) *ListAppAssessmentComplianceDriftsPaginator {
	if params == nil {
		params = &ListAppAssessmentComplianceDriftsInput{}
	}

	options := ListAppAssessmentComplianceDriftsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppAssessmentComplianceDriftsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppAssessmentComplianceDriftsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppAssessmentComplianceDrifts page.
func (p *ListAppAssessmentComplianceDriftsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppAssessmentComplianceDriftsOutput, error) {
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

	result, err := p.client.ListAppAssessmentComplianceDrifts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppAssessmentComplianceDrifts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppAssessmentComplianceDrifts",
	}
}
