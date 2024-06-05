// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modify a setting for an Amazon DocumentDB global cluster. You can change one or
// more configuration parameters (for example: deletion protection), or the global
// cluster identifier by specifying these parameters and the new values in the
// request.
//
// This action only applies to Amazon DocumentDB clusters.
func (c *Client) ModifyGlobalCluster(ctx context.Context, params *ModifyGlobalClusterInput, optFns ...func(*Options)) (*ModifyGlobalClusterOutput, error) {
	if params == nil {
		params = &ModifyGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyGlobalCluster", params, optFns, c.addOperationModifyGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyGlobalCluster.
type ModifyGlobalClusterInput struct {

	// The identifier for the global cluster being modified. This parameter isn't
	// case-sensitive.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing global cluster.
	//
	// This member is required.
	GlobalClusterIdentifier *string

	// Indicates if the global cluster has deletion protection enabled. The global
	// cluster can't be deleted when deletion protection is enabled.
	DeletionProtection *bool

	// The new identifier for a global cluster when you modify a global cluster. This
	// value is stored as a lowercase string.
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens
	//
	// The first character must be a letter
	//
	// Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-cluster2
	NewGlobalClusterIdentifier *string

	noSmithyDocumentSerde
}

type ModifyGlobalClusterOutput struct {

	// A data type representing an Amazon DocumentDB global cluster.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyGlobalCluster"); err != nil {
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
	if err = addOpModifyGlobalClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyGlobalCluster",
	}
}
