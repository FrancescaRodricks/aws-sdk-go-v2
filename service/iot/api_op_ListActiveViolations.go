// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the active violations for a given Device Defender security profile.
//
// Requires permission to access the [ListActiveViolations] action.
//
// [ListActiveViolations]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListActiveViolations(ctx context.Context, params *ListActiveViolationsInput, optFns ...func(*Options)) (*ListActiveViolationsOutput, error) {
	if params == nil {
		params = &ListActiveViolationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListActiveViolations", params, optFns, c.addOperationListActiveViolationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListActiveViolationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListActiveViolationsInput struct {

	//  The criteria for a behavior.
	BehaviorCriteriaType types.BehaviorCriteriaType

	//  A list of all suppressed alerts.
	ListSuppressedAlerts *bool

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// The name of the Device Defender security profile for which violations are
	// listed.
	SecurityProfileName *string

	// The name of the thing whose active violations are listed.
	ThingName *string

	// The verification state of the violation (detect alarm).
	VerificationState types.VerificationState

	noSmithyDocumentSerde
}

type ListActiveViolationsOutput struct {

	// The list of active violations.
	ActiveViolations []types.ActiveViolation

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListActiveViolationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListActiveViolations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListActiveViolations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListActiveViolations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListActiveViolations(options.Region), middleware.Before); err != nil {
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

// ListActiveViolationsAPIClient is a client that implements the
// ListActiveViolations operation.
type ListActiveViolationsAPIClient interface {
	ListActiveViolations(context.Context, *ListActiveViolationsInput, ...func(*Options)) (*ListActiveViolationsOutput, error)
}

var _ ListActiveViolationsAPIClient = (*Client)(nil)

// ListActiveViolationsPaginatorOptions is the paginator options for
// ListActiveViolations
type ListActiveViolationsPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListActiveViolationsPaginator is a paginator for ListActiveViolations
type ListActiveViolationsPaginator struct {
	options   ListActiveViolationsPaginatorOptions
	client    ListActiveViolationsAPIClient
	params    *ListActiveViolationsInput
	nextToken *string
	firstPage bool
}

// NewListActiveViolationsPaginator returns a new ListActiveViolationsPaginator
func NewListActiveViolationsPaginator(client ListActiveViolationsAPIClient, params *ListActiveViolationsInput, optFns ...func(*ListActiveViolationsPaginatorOptions)) *ListActiveViolationsPaginator {
	if params == nil {
		params = &ListActiveViolationsInput{}
	}

	options := ListActiveViolationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListActiveViolationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListActiveViolationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListActiveViolations page.
func (p *ListActiveViolationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListActiveViolationsOutput, error) {
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

	result, err := p.client.ListActiveViolations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListActiveViolations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListActiveViolations",
	}
}
