// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Neptune global database spread across multiple Amazon Regions. The
// global database contains a single primary cluster with read-write capability,
// and read-only secondary clusters that receive data from the primary cluster
// through high-speed replication performed by the Neptune storage subsystem.
//
// You can create a global database that is initially empty, and then add a
// primary cluster and secondary clusters to it, or you can specify an existing
// Neptune cluster during the create operation to become the primary cluster of the
// global database.
func (c *Client) CreateGlobalCluster(ctx context.Context, params *CreateGlobalClusterInput, optFns ...func(*Options)) (*CreateGlobalClusterOutput, error) {
	if params == nil {
		params = &CreateGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGlobalCluster", params, optFns, c.addOperationCreateGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGlobalClusterInput struct {

	// The cluster identifier of the new global database cluster.
	//
	// This member is required.
	GlobalClusterIdentifier *string

	// The deletion protection setting for the new global database. The global
	// database can't be deleted when deletion protection is enabled.
	DeletionProtection *bool

	// The name of the database engine to be used in the global database.
	//
	// Valid values: neptune
	Engine *string

	// The Neptune engine version to be used by the global database.
	//
	// Valid values: 1.2.0.0 or above.
	EngineVersion *string

	// (Optional) The Amazon Resource Name (ARN) of an existing Neptune DB cluster to
	// use as the primary cluster of the new global database.
	SourceDBClusterIdentifier *string

	// The storage encryption setting for the new global database cluster.
	StorageEncrypted *bool

	noSmithyDocumentSerde
}

type CreateGlobalClusterOutput struct {

	// Contains the details of an Amazon Neptune global database.
	//
	// This data type is used as a response element for the CreateGlobalCluster, DescribeGlobalClusters, ModifyGlobalCluster, DeleteGlobalCluster, FailoverGlobalCluster, and RemoveFromGlobalCluster actions.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateGlobalCluster"); err != nil {
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
	if err = addOpCreateGlobalClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateGlobalCluster",
	}
}
