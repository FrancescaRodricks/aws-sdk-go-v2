// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the maintenance configuration of the Managed Service for Apache Flink
// application.
//
// You can invoke this operation on an application that is in one of the two
// following states: READY or RUNNING . If you invoke it when the application is in
// a state other than these two states, it throws a ResourceInUseException . The
// service makes use of the updated configuration the next time it schedules
// maintenance for the application. If you invoke this operation after the service
// schedules maintenance, the service will apply the configuration update the next
// time it schedules maintenance for the application. This means that you might not
// see the maintenance configuration update applied to the maintenance process that
// follows a successful invocation of this operation, but to the following
// maintenance process instead.
//
// To see the current maintenance configuration of your application, invoke the DescribeApplication
// operation.
//
// For information about application maintenance, see [Managed Service for Apache Flink for Apache Flink Maintenance].
//
// This operation is supported only for Managed Service for Apache Flink.
//
// [Managed Service for Apache Flink for Apache Flink Maintenance]: https://docs.aws.amazon.com/kinesisanalytics/latest/java/maintenance.html
func (c *Client) UpdateApplicationMaintenanceConfiguration(ctx context.Context, params *UpdateApplicationMaintenanceConfigurationInput, optFns ...func(*Options)) (*UpdateApplicationMaintenanceConfigurationOutput, error) {
	if params == nil {
		params = &UpdateApplicationMaintenanceConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApplicationMaintenanceConfiguration", params, optFns, c.addOperationUpdateApplicationMaintenanceConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApplicationMaintenanceConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationMaintenanceConfigurationInput struct {

	// Describes the application maintenance configuration update.
	//
	// This member is required.
	ApplicationMaintenanceConfigurationUpdate *types.ApplicationMaintenanceConfigurationUpdate

	// The name of the application for which you want to update the maintenance
	// configuration.
	//
	// This member is required.
	ApplicationName *string

	noSmithyDocumentSerde
}

type UpdateApplicationMaintenanceConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the application.
	ApplicationARN *string

	// The application maintenance configuration description after the update.
	ApplicationMaintenanceConfigurationDescription *types.ApplicationMaintenanceConfigurationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateApplicationMaintenanceConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApplicationMaintenanceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApplicationMaintenanceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateApplicationMaintenanceConfiguration"); err != nil {
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
	if err = addOpUpdateApplicationMaintenanceConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplicationMaintenanceConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApplicationMaintenanceConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateApplicationMaintenanceConfiguration",
	}
}
