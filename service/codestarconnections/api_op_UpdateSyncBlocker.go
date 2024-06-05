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

// Allows you to update the status of a sync blocker, resolving the blocker and
// allowing syncing to continue.
func (c *Client) UpdateSyncBlocker(ctx context.Context, params *UpdateSyncBlockerInput, optFns ...func(*Options)) (*UpdateSyncBlockerOutput, error) {
	if params == nil {
		params = &UpdateSyncBlockerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSyncBlocker", params, optFns, c.addOperationUpdateSyncBlockerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSyncBlockerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSyncBlockerInput struct {

	// The ID of the sync blocker to be updated.
	//
	// This member is required.
	Id *string

	// The reason for resolving the sync blocker.
	//
	// This member is required.
	ResolvedReason *string

	// The name of the resource for the sync blocker to be updated.
	//
	// This member is required.
	ResourceName *string

	// The sync type of the sync blocker to be updated.
	//
	// This member is required.
	SyncType types.SyncConfigurationType

	noSmithyDocumentSerde
}

type UpdateSyncBlockerOutput struct {

	// The resource name for the sync blocker.
	//
	// This member is required.
	ResourceName *string

	// Information about the sync blocker to be updated.
	//
	// This member is required.
	SyncBlocker *types.SyncBlocker

	// The parent resource name for the sync blocker.
	ParentResourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSyncBlockerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateSyncBlocker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateSyncBlocker{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSyncBlocker"); err != nil {
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
	if err = addOpUpdateSyncBlockerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSyncBlocker(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSyncBlocker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSyncBlocker",
	}
}
