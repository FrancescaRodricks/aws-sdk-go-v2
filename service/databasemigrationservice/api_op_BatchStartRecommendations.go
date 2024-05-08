// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the analysis of up to 20 source databases to recommend target engines
// for each source database. This is a batch version of [StartRecommendations].
//
// The result of analysis of each source database is reported individually in the
// response. Because the batch request can result in a combination of successful
// and unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
//
// [StartRecommendations]: https://docs.aws.amazon.com/dms/latest/APIReference/API_StartRecommendations.html
func (c *Client) BatchStartRecommendations(ctx context.Context, params *BatchStartRecommendationsInput, optFns ...func(*Options)) (*BatchStartRecommendationsOutput, error) {
	if params == nil {
		params = &BatchStartRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchStartRecommendations", params, optFns, c.addOperationBatchStartRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchStartRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchStartRecommendationsInput struct {

	// Provides information about source databases to analyze. After this analysis,
	// Fleet Advisor recommends target engines for each source database.
	Data []types.StartRecommendationsRequestEntry

	noSmithyDocumentSerde
}

type BatchStartRecommendationsOutput struct {

	// A list with error details about the analysis of each source database.
	ErrorEntries []types.BatchStartRecommendationsErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchStartRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchStartRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchStartRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchStartRecommendations"); err != nil {
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
	if err = addOpBatchStartRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchStartRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchStartRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchStartRecommendations",
	}
}
