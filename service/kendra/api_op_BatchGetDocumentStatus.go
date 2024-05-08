// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the indexing status for one or more documents submitted with the [BatchPutDocument] API.
//
// When you use the BatchPutDocument API, documents are indexed asynchronously.
// You can use the BatchGetDocumentStatus API to get the current status of a list
// of documents so that you can determine if they have been successfully indexed.
//
// You can also use the BatchGetDocumentStatus API to check the status of the [BatchDeleteDocument]
// API. When a document is deleted from the index, Amazon Kendra returns NOT_FOUND
// as the status.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html
// [BatchDeleteDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchDeleteDocument.html
func (c *Client) BatchGetDocumentStatus(ctx context.Context, params *BatchGetDocumentStatusInput, optFns ...func(*Options)) (*BatchGetDocumentStatusOutput, error) {
	if params == nil {
		params = &BatchGetDocumentStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetDocumentStatus", params, optFns, c.addOperationBatchGetDocumentStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetDocumentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetDocumentStatusInput struct {

	// A list of DocumentInfo objects that identify the documents for which to get the
	// status. You identify the documents by their document ID and optional attributes.
	//
	// This member is required.
	DocumentInfoList []types.DocumentInfo

	// The identifier of the index to add documents to. The index ID is returned by
	// the [CreateIndex]API.
	//
	// [CreateIndex]: https://docs.aws.amazon.com/kendra/latest/dg/API_CreateIndex.html
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type BatchGetDocumentStatusOutput struct {

	// The status of documents. The status indicates if the document is waiting to be
	// indexed, is in the process of indexing, has completed indexing, or failed
	// indexing. If a document failed indexing, the status provides the reason why.
	DocumentStatusList []types.Status

	// A list of documents that Amazon Kendra couldn't get the status for. The list
	// includes the ID of the document and the reason that the status couldn't be
	// found.
	Errors []types.BatchGetDocumentStatusResponseError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetDocumentStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetDocumentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetDocumentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetDocumentStatus"); err != nil {
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
	if err = addOpBatchGetDocumentStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetDocumentStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetDocumentStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetDocumentStatus",
	}
}
