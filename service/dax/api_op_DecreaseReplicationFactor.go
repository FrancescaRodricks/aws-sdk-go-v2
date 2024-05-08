// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes one or more nodes from a DAX cluster.
//
// You cannot use DecreaseReplicationFactor to remove the last node in a DAX
// cluster. If you need to do this, use DeleteCluster instead.
func (c *Client) DecreaseReplicationFactor(ctx context.Context, params *DecreaseReplicationFactorInput, optFns ...func(*Options)) (*DecreaseReplicationFactorOutput, error) {
	if params == nil {
		params = &DecreaseReplicationFactorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DecreaseReplicationFactor", params, optFns, c.addOperationDecreaseReplicationFactorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecreaseReplicationFactorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecreaseReplicationFactorInput struct {

	// The name of the DAX cluster from which you want to remove nodes.
	//
	// This member is required.
	ClusterName *string

	// The new number of nodes for the DAX cluster.
	//
	// This member is required.
	NewReplicationFactor int32

	// The Availability Zone(s) from which to remove nodes.
	AvailabilityZones []string

	// The unique identifiers of the nodes to be removed from the cluster.
	NodeIdsToRemove []string

	noSmithyDocumentSerde
}

type DecreaseReplicationFactorOutput struct {

	// A description of the DAX cluster, after you have decreased its replication
	// factor.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDecreaseReplicationFactorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDecreaseReplicationFactor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDecreaseReplicationFactor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DecreaseReplicationFactor"); err != nil {
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
	if err = addOpDecreaseReplicationFactorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDecreaseReplicationFactor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDecreaseReplicationFactor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DecreaseReplicationFactor",
	}
}
