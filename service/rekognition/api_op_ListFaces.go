// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns metadata for faces in the specified collection. This metadata includes
// information such as the bounding box coordinates, the confidence (that the
// bounding box contains a face), and face ID. For an example, see Listing Faces in
// a Collection in the Amazon Rekognition Developer Guide.
//
// This operation requires permissions to perform the rekognition:ListFaces action.
func (c *Client) ListFaces(ctx context.Context, params *ListFacesInput, optFns ...func(*Options)) (*ListFacesOutput, error) {
	if params == nil {
		params = &ListFacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFaces", params, optFns, c.addOperationListFacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFacesInput struct {

	// ID of the collection from which to list the faces.
	//
	// This member is required.
	CollectionId *string

	// An array of face IDs to filter results with when listing faces in a collection.
	FaceIds []string

	// Maximum number of faces to return.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Rekognition returns a pagination token in the response. You
	// can use this pagination token to retrieve the next set of faces.
	NextToken *string

	// An array of user IDs to filter results with when listing faces in a collection.
	UserId *string

	noSmithyDocumentSerde
}

type ListFacesOutput struct {

	// Version number of the face detection model associated with the input collection
	// ( CollectionId ).
	FaceModelVersion *string

	// An array of Face objects.
	Faces []types.Face

	// If the response is truncated, Amazon Rekognition returns this token that you
	// can use in the subsequent request to retrieve the next set of faces.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFaces"); err != nil {
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
	if err = addOpListFacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFaces(options.Region), middleware.Before); err != nil {
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

// ListFacesAPIClient is a client that implements the ListFaces operation.
type ListFacesAPIClient interface {
	ListFaces(context.Context, *ListFacesInput, ...func(*Options)) (*ListFacesOutput, error)
}

var _ ListFacesAPIClient = (*Client)(nil)

// ListFacesPaginatorOptions is the paginator options for ListFaces
type ListFacesPaginatorOptions struct {
	// Maximum number of faces to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFacesPaginator is a paginator for ListFaces
type ListFacesPaginator struct {
	options   ListFacesPaginatorOptions
	client    ListFacesAPIClient
	params    *ListFacesInput
	nextToken *string
	firstPage bool
}

// NewListFacesPaginator returns a new ListFacesPaginator
func NewListFacesPaginator(client ListFacesAPIClient, params *ListFacesInput, optFns ...func(*ListFacesPaginatorOptions)) *ListFacesPaginator {
	if params == nil {
		params = &ListFacesInput{}
	}

	options := ListFacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFaces page.
func (p *ListFacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFacesOutput, error) {
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

	result, err := p.client.ListFaces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFaces",
	}
}
