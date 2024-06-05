// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of servers that are one network hop away from a specified
// server.
func (c *Client) ListServerNeighbors(ctx context.Context, params *ListServerNeighborsInput, optFns ...func(*Options)) (*ListServerNeighborsOutput, error) {
	if params == nil {
		params = &ListServerNeighborsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServerNeighbors", params, optFns, c.addOperationListServerNeighborsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServerNeighborsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServerNeighborsInput struct {

	// Configuration ID of the server for which neighbors are being listed.
	//
	// This member is required.
	ConfigurationId *string

	// Maximum number of results to return in a single page of output.
	MaxResults int32

	// List of configuration IDs to test for one-hop-away.
	NeighborConfigurationIds []string

	// Token to retrieve the next set of results. For example, if you previously
	// specified 100 IDs for ListServerNeighborsRequest$neighborConfigurationIds but
	// set ListServerNeighborsRequest$maxResults to 10, you received a set of 10
	// results along with a token. Use that token in this query to get the next set of
	// 10.
	NextToken *string

	// Flag to indicate if port and protocol information is needed as part of the
	// response.
	PortInformationNeeded bool

	noSmithyDocumentSerde
}

type ListServerNeighborsOutput struct {

	// List of distinct servers that are one hop away from the given server.
	//
	// This member is required.
	Neighbors []types.NeighborConnectionDetail

	// Count of distinct servers that are one hop away from the given server.
	KnownDependencyCount int64

	// Token to retrieve the next set of results. For example, if you specified 100
	// IDs for ListServerNeighborsRequest$neighborConfigurationIds but set
	// ListServerNeighborsRequest$maxResults to 10, you received a set of 10 results
	// along with this token. Use this token in the next query to retrieve the next set
	// of 10.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServerNeighborsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServerNeighbors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServerNeighbors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServerNeighbors"); err != nil {
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
	if err = addOpListServerNeighborsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServerNeighbors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListServerNeighbors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServerNeighbors",
	}
}
