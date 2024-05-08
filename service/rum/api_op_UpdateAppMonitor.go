// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of an existing app monitor. When you use this
// operation, only the parts of the app monitor configuration that you specify in
// this operation are changed. For any parameters that you omit, the existing
// values are kept.
//
// You can't use this operation to change the tags of an existing app monitor. To
// change the tags of an existing app monitor, use [TagResource].
//
// To create a new app monitor, use [CreateAppMonitor].
//
// After you update an app monitor, sign in to the CloudWatch RUM console to get
// the updated JavaScript code snippet to add to your web application. For more
// information, see [How do I find a code snippet that I've already generated?]
//
// [CreateAppMonitor]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_CreateAppMonitor.html
// [TagResource]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_TagResource.html
// [How do I find a code snippet that I've already generated?]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-find-code-snippet.html
func (c *Client) UpdateAppMonitor(ctx context.Context, params *UpdateAppMonitorInput, optFns ...func(*Options)) (*UpdateAppMonitorOutput, error) {
	if params == nil {
		params = &UpdateAppMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAppMonitor", params, optFns, c.addOperationUpdateAppMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAppMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAppMonitorInput struct {

	// The name of the app monitor to update.
	//
	// This member is required.
	Name *string

	// A structure that contains much of the configuration data for the app monitor.
	// If you are using Amazon Cognito for authorization, you must include this
	// structure in your request, and it must include the ID of the Amazon Cognito
	// identity pool to use for authorization. If you don't include
	// AppMonitorConfiguration , you must set up your own authorization method. For
	// more information, see [Authorize your application to send data to Amazon Web Services].
	//
	// [Authorize your application to send data to Amazon Web Services]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-get-started-authorization.html
	AppMonitorConfiguration *types.AppMonitorConfiguration

	// Specifies whether this app monitor allows the web client to define and send
	// custom events. The default is for custom events to be DISABLED .
	//
	// For more information about custom events, see [Send custom events].
	//
	// [Send custom events]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-custom-events.html
	CustomEvents *types.CustomEvents

	// Data collected by RUM is kept by RUM for 30 days and then deleted. This
	// parameter specifies whether RUM sends a copy of this telemetry data to Amazon
	// CloudWatch Logs in your account. This enables you to keep the telemetry data for
	// more than 30 days, but it does incur Amazon CloudWatch Logs charges.
	CwLogEnabled *bool

	// The top-level internet domain name for which your application has
	// administrative authority.
	Domain *string

	noSmithyDocumentSerde
}

type UpdateAppMonitorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAppMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAppMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAppMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAppMonitor"); err != nil {
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
	if err = addOpUpdateAppMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAppMonitor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAppMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAppMonitor",
	}
}
