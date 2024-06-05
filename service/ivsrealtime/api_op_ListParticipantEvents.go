// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists events for a specified participant that occurred during a specified stage
// session.
func (c *Client) ListParticipantEvents(ctx context.Context, params *ListParticipantEventsInput, optFns ...func(*Options)) (*ListParticipantEventsOutput, error) {
	if params == nil {
		params = &ListParticipantEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListParticipantEvents", params, optFns, c.addOperationListParticipantEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListParticipantEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListParticipantEventsInput struct {

	// Unique identifier for this participant. This is assigned by IVS and returned by CreateParticipantToken
	// .
	//
	// This member is required.
	ParticipantId *string

	// ID of a session within the stage.
	//
	// This member is required.
	SessionId *string

	// Stage ARN.
	//
	// This member is required.
	StageArn *string

	// Maximum number of results to return. Default: 50.
	MaxResults *int32

	// The first participant event to retrieve. This is used for pagination; see the
	// nextToken response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListParticipantEventsOutput struct {

	// List of the matching events.
	//
	// This member is required.
	Events []types.Event

	// If there are more events than maxResults , use nextToken in the request to get
	// the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListParticipantEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListParticipantEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListParticipantEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListParticipantEvents"); err != nil {
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
	if err = addOpListParticipantEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListParticipantEvents(options.Region), middleware.Before); err != nil {
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

// ListParticipantEventsAPIClient is a client that implements the
// ListParticipantEvents operation.
type ListParticipantEventsAPIClient interface {
	ListParticipantEvents(context.Context, *ListParticipantEventsInput, ...func(*Options)) (*ListParticipantEventsOutput, error)
}

var _ ListParticipantEventsAPIClient = (*Client)(nil)

// ListParticipantEventsPaginatorOptions is the paginator options for
// ListParticipantEvents
type ListParticipantEventsPaginatorOptions struct {
	// Maximum number of results to return. Default: 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListParticipantEventsPaginator is a paginator for ListParticipantEvents
type ListParticipantEventsPaginator struct {
	options   ListParticipantEventsPaginatorOptions
	client    ListParticipantEventsAPIClient
	params    *ListParticipantEventsInput
	nextToken *string
	firstPage bool
}

// NewListParticipantEventsPaginator returns a new ListParticipantEventsPaginator
func NewListParticipantEventsPaginator(client ListParticipantEventsAPIClient, params *ListParticipantEventsInput, optFns ...func(*ListParticipantEventsPaginatorOptions)) *ListParticipantEventsPaginator {
	if params == nil {
		params = &ListParticipantEventsInput{}
	}

	options := ListParticipantEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListParticipantEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListParticipantEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListParticipantEvents page.
func (p *ListParticipantEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListParticipantEventsOutput, error) {
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

	result, err := p.client.ListParticipantEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListParticipantEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListParticipantEvents",
	}
}
