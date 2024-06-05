// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Manages the generation and use of property graph statistics.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ManageStatistics]IAM action in that cluster.
//
// [neptune-db:ManageStatistics]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#managestatistics
func (c *Client) ManagePropertygraphStatistics(ctx context.Context, params *ManagePropertygraphStatisticsInput, optFns ...func(*Options)) (*ManagePropertygraphStatisticsOutput, error) {
	if params == nil {
		params = &ManagePropertygraphStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ManagePropertygraphStatistics", params, optFns, c.addOperationManagePropertygraphStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ManagePropertygraphStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ManagePropertygraphStatisticsInput struct {

	// The statistics generation mode. One of: DISABLE_AUTOCOMPUTE , ENABLE_AUTOCOMPUTE
	// , or REFRESH , the last of which manually triggers DFE statistics generation.
	Mode types.StatisticsAutoGenerationMode

	noSmithyDocumentSerde
}

type ManagePropertygraphStatisticsOutput struct {

	// The HTTP return code of the request. If the request succeeded, the code is 200.
	//
	// This member is required.
	Status *string

	// This is only returned for refresh mode.
	Payload *types.RefreshStatisticsIdMap

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationManagePropertygraphStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpManagePropertygraphStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpManagePropertygraphStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ManagePropertygraphStatistics"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opManagePropertygraphStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opManagePropertygraphStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ManagePropertygraphStatistics",
	}
}
