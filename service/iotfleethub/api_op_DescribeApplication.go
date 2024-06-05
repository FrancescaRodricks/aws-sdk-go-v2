// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleethub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotfleethub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a Fleet Hub for IoT Device Management web application.
func (c *Client) DescribeApplication(ctx context.Context, params *DescribeApplicationInput, optFns ...func(*Options)) (*DescribeApplicationOutput, error) {
	if params == nil {
		params = &DescribeApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplication", params, optFns, c.addOperationDescribeApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeApplicationInput struct {

	// The unique Id of the web application.
	//
	// This member is required.
	ApplicationId *string

	noSmithyDocumentSerde
}

type DescribeApplicationOutput struct {

	// The ARN of the web application.
	//
	// This member is required.
	ApplicationArn *string

	// The date (in Unix epoch time) when the application was created.
	//
	// This member is required.
	ApplicationCreationDate int64

	// The unique Id of the web application.
	//
	// This member is required.
	ApplicationId *string

	// The date (in Unix epoch time) when the application was last updated.
	//
	// This member is required.
	ApplicationLastUpdateDate int64

	// The name of the web application.
	//
	// This member is required.
	ApplicationName *string

	// The current state of the web application.
	//
	// This member is required.
	ApplicationState types.ApplicationState

	// The URL of the web application.
	//
	// This member is required.
	ApplicationUrl *string

	// The ARN of the role that the web application assumes when it interacts with
	// Amazon Web Services IoT Core.
	//
	// This member is required.
	RoleArn *string

	// An optional description of the web application.
	ApplicationDescription *string

	// A message that explains any failures included in the applicationState response
	// field. This message explains failures in the CreateApplication and
	// DeleteApplication actions.
	ErrorMessage *string

	// The Id of the single sign-on client that you use to authenticate and authorize
	// users on the web application.
	SsoClientId *string

	// A set of key/value pairs that you can use to manage the web application
	// resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeApplication"); err != nil {
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
	if err = addOpDescribeApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeApplication",
	}
}
