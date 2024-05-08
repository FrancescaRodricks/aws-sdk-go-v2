// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Starts the specified Amazon Kinesis Analytics application. After creating an
// application, you must exclusively call this operation to start your application.
//
// After the application starts, it begins consuming the input data, processes it,
// and writes the output to the configured destination.
//
// The application status must be READY for you to start an application. You can
// get the application status in the console or using the [DescribeApplication]operation.
//
// After you start the application, you can stop the application from processing
// the input by calling the [StopApplication]operation.
//
// This operation requires permissions to perform the
// kinesisanalytics:StartApplication action.
//
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
// [StopApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_StopApplication.html
func (c *Client) StartApplication(ctx context.Context, params *StartApplicationInput, optFns ...func(*Options)) (*StartApplicationOutput, error) {
	if params == nil {
		params = &StartApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartApplication", params, optFns, c.addOperationStartApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartApplicationInput struct {

	// Name of the application.
	//
	// This member is required.
	ApplicationName *string

	// Identifies the specific input, by ID, that the application starts consuming.
	// Amazon Kinesis Analytics starts reading the streaming source associated with the
	// input. You can also specify where in the streaming source you want Amazon
	// Kinesis Analytics to start reading.
	//
	// This member is required.
	InputConfigurations []types.InputConfiguration

	noSmithyDocumentSerde
}

type StartApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartApplication"); err != nil {
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
	if err = addOpStartApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartApplication",
	}
}
