// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates replication info of a replicator.
func (c *Client) UpdateReplicationInfo(ctx context.Context, params *UpdateReplicationInfoInput, optFns ...func(*Options)) (*UpdateReplicationInfoOutput, error) {
	if params == nil {
		params = &UpdateReplicationInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateReplicationInfo", params, optFns, c.addOperationUpdateReplicationInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateReplicationInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Update information relating to replication between a given source and target
// Kafka cluster.
type UpdateReplicationInfoInput struct {

	// Current replicator version.
	//
	// This member is required.
	CurrentVersion *string

	// The Amazon Resource Name (ARN) of the replicator to be updated.
	//
	// This member is required.
	ReplicatorArn *string

	// The ARN of the source Kafka cluster.
	//
	// This member is required.
	SourceKafkaClusterArn *string

	// The ARN of the target Kafka cluster.
	//
	// This member is required.
	TargetKafkaClusterArn *string

	// Updated consumer group replication information.
	ConsumerGroupReplication *types.ConsumerGroupReplicationUpdate

	// Updated topic replication information.
	TopicReplication *types.TopicReplicationUpdate

	noSmithyDocumentSerde
}

type UpdateReplicationInfoOutput struct {

	// The Amazon Resource Name (ARN) of the replicator.
	ReplicatorArn *string

	// State of the replicator.
	ReplicatorState types.ReplicatorState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReplicationInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateReplicationInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateReplicationInfo{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateReplicationInfo"); err != nil {
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
	if err = addOpUpdateReplicationInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReplicationInfo(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateReplicationInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateReplicationInfo",
	}
}
