// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the output destination configuration from your SQL-based Kinesis Data
// Analytics application's configuration. Kinesis Data Analytics will no longer
// write data from the corresponding in-application stream to the external output
// destination.
func (c *Client) DeleteApplicationOutput(ctx context.Context, params *DeleteApplicationOutputInput, optFns ...func(*Options)) (*DeleteApplicationOutputOutput, error) {
	if params == nil {
		params = &DeleteApplicationOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApplicationOutput", params, optFns, c.addOperationDeleteApplicationOutputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApplicationOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationOutputInput struct {

	// The application name.
	//
	// This member is required.
	ApplicationName *string

	// The application version. You can use the DescribeApplication operation to get the current
	// application version. If the version specified is not the current version, the
	// ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The ID of the configuration to delete. Each output configuration that is added
	// to the application (either when the application is created or later) using the AddApplicationOutput
	// operation has a unique ID. You need to provide the ID to uniquely identify the
	// output configuration that you want to delete from the application configuration.
	// You can use the DescribeApplicationoperation to get the specific OutputId .
	//
	// This member is required.
	OutputId *string

	noSmithyDocumentSerde
}

type DeleteApplicationOutputOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string

	// The current application version ID.
	ApplicationVersionId *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteApplicationOutputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationOutput{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteApplicationOutput"); err != nil {
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
	if err = addOpDeleteApplicationOutputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationOutput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteApplicationOutput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteApplicationOutput",
	}
}
