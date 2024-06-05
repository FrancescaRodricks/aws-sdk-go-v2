// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a sync configuration which allows Amazon Web Services to sync content
// from a Git repository to update a specified Amazon Web Services resource.
// Parameters for the sync configuration are determined by the sync type.
func (c *Client) CreateSyncConfiguration(ctx context.Context, params *CreateSyncConfigurationInput, optFns ...func(*Options)) (*CreateSyncConfigurationOutput, error) {
	if params == nil {
		params = &CreateSyncConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSyncConfiguration", params, optFns, c.addOperationCreateSyncConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSyncConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSyncConfigurationInput struct {

	// The branch in the repository from which changes will be synced.
	//
	// This member is required.
	Branch *string

	// The file name of the configuration file that manages syncing between the
	// connection and the repository. This configuration file is stored in the
	// repository.
	//
	// This member is required.
	ConfigFile *string

	// The ID of the repository link created for the connection. A repository link
	// allows Git sync to monitor and sync changes to files in a specified Git
	// repository.
	//
	// This member is required.
	RepositoryLinkId *string

	// The name of the Amazon Web Services resource (for example, a CloudFormation
	// stack in the case of CFN_STACK_SYNC) that will be synchronized from the linked
	// repository.
	//
	// This member is required.
	ResourceName *string

	// The ARN of the IAM role that grants permission for Amazon Web Services to use
	// Git sync to update a given Amazon Web Services resource on your behalf.
	//
	// This member is required.
	RoleArn *string

	// The type of sync configuration.
	//
	// This member is required.
	SyncType types.SyncConfigurationType

	// Whether to enable or disable publishing of deployment status to source
	// providers.
	PublishDeploymentStatus types.PublishDeploymentStatus

	// When to trigger Git sync to begin the stack update.
	TriggerResourceUpdateOn types.TriggerResourceUpdateOn

	noSmithyDocumentSerde
}

type CreateSyncConfigurationOutput struct {

	// The created sync configuration for the connection. A sync configuration allows
	// Amazon Web Services to sync content from a Git repository to update a specified
	// Amazon Web Services resource.
	//
	// This member is required.
	SyncConfiguration *types.SyncConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSyncConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateSyncConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateSyncConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSyncConfiguration"); err != nil {
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
	if err = addOpCreateSyncConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSyncConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSyncConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSyncConfiguration",
	}
}
