// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Amazon Web Services account users that are assigned to an application.
func (c *Client) ListApplicationAssignments(ctx context.Context, params *ListApplicationAssignmentsInput, optFns ...func(*Options)) (*ListApplicationAssignmentsOutput, error) {
	if params == nil {
		params = &ListApplicationAssignmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationAssignments", params, optFns, c.addOperationListApplicationAssignmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationAssignmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationAssignmentsInput struct {

	// Specifies the ARN of the application.
	//
	// This member is required.
	ApplicationArn *string

	// Specifies the total number of results that you want included in each response.
	// If additional items exist beyond the number you specify, the NextToken response
	// element is returned with a value (not null). Include the specified value as the
	// NextToken request parameter in the next call to the operation to get the next
	// set of results. Note that the service might return fewer results than the
	// maximum even when there are more results available. You should check NextToken
	// after every operation to ensure that you receive all of the results.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListApplicationAssignmentsOutput struct {

	// The list of users assigned to an application.
	ApplicationAssignments []types.ApplicationAssignment

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationAssignmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApplicationAssignments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApplicationAssignments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListApplicationAssignments"); err != nil {
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
	if err = addOpListApplicationAssignmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationAssignments(options.Region), middleware.Before); err != nil {
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

// ListApplicationAssignmentsAPIClient is a client that implements the
// ListApplicationAssignments operation.
type ListApplicationAssignmentsAPIClient interface {
	ListApplicationAssignments(context.Context, *ListApplicationAssignmentsInput, ...func(*Options)) (*ListApplicationAssignmentsOutput, error)
}

var _ ListApplicationAssignmentsAPIClient = (*Client)(nil)

// ListApplicationAssignmentsPaginatorOptions is the paginator options for
// ListApplicationAssignments
type ListApplicationAssignmentsPaginatorOptions struct {
	// Specifies the total number of results that you want included in each response.
	// If additional items exist beyond the number you specify, the NextToken response
	// element is returned with a value (not null). Include the specified value as the
	// NextToken request parameter in the next call to the operation to get the next
	// set of results. Note that the service might return fewer results than the
	// maximum even when there are more results available. You should check NextToken
	// after every operation to ensure that you receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationAssignmentsPaginator is a paginator for
// ListApplicationAssignments
type ListApplicationAssignmentsPaginator struct {
	options   ListApplicationAssignmentsPaginatorOptions
	client    ListApplicationAssignmentsAPIClient
	params    *ListApplicationAssignmentsInput
	nextToken *string
	firstPage bool
}

// NewListApplicationAssignmentsPaginator returns a new
// ListApplicationAssignmentsPaginator
func NewListApplicationAssignmentsPaginator(client ListApplicationAssignmentsAPIClient, params *ListApplicationAssignmentsInput, optFns ...func(*ListApplicationAssignmentsPaginatorOptions)) *ListApplicationAssignmentsPaginator {
	if params == nil {
		params = &ListApplicationAssignmentsInput{}
	}

	options := ListApplicationAssignmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationAssignmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationAssignmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationAssignments page.
func (p *ListApplicationAssignmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationAssignmentsOutput, error) {
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

	result, err := p.client.ListApplicationAssignments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationAssignments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListApplicationAssignments",
	}
}
