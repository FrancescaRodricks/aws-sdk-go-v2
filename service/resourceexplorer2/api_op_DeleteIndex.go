// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes the specified index and turns off Amazon Web Services Resource Explorer
// in the specified Amazon Web Services Region. When you delete an index, Resource
// Explorer stops discovering and indexing resources in that Region. Resource
// Explorer also deletes all views in that Region. These actions occur as
// asynchronous background tasks. You can check to see when the actions are
// complete by using the GetIndexoperation and checking the Status response value.
//
// If the index you delete is the aggregator index for the Amazon Web Services
// account, you must wait 24 hours before you can promote another local index to be
// the aggregator index for the account. Users can't perform account-wide searches
// using Resource Explorer until another aggregator index is configured.
func (c *Client) DeleteIndex(ctx context.Context, params *DeleteIndexInput, optFns ...func(*Options)) (*DeleteIndexOutput, error) {
	if params == nil {
		params = &DeleteIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIndex", params, optFns, c.addOperationDeleteIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIndexInput struct {

	// The [Amazon resource name (ARN)] of the index that you want to delete.
	//
	// [Amazon resource name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type DeleteIndexOutput struct {

	// The [Amazon resource name (ARN)] of the index that you successfully started the deletion process.
	//
	// This operation is asynchronous. To check its status, call the GetIndex operation.
	//
	// [Amazon resource name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	Arn *string

	// The date and time when you last updated this index.
	LastUpdatedAt *time.Time

	// Indicates the current state of the index.
	State types.IndexState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIndex{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteIndex"); err != nil {
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
	if err = addOpDeleteIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIndex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteIndex",
	}
}
