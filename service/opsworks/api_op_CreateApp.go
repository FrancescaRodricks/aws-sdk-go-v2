// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an app for a specified stack. For more information, see [Creating Apps].
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see [Managing User Permissions].
//
// [Creating Apps]: https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html
// [Managing User Permissions]: https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
func (c *Client) CreateApp(ctx context.Context, params *CreateAppInput, optFns ...func(*Options)) (*CreateAppOutput, error) {
	if params == nil {
		params = &CreateAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApp", params, optFns, c.addOperationCreateAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppInput struct {

	// The app name.
	//
	// This member is required.
	Name *string

	// The stack ID.
	//
	// This member is required.
	StackId *string

	// The app type. Each supported type is associated with a particular layer. For
	// example, PHP applications are associated with a PHP layer. AWS OpsWorks Stacks
	// deploys an application to those instances that are members of the corresponding
	// layer. If your app isn't one of the standard types, or you prefer to implement
	// your own Deploy recipes, specify other .
	//
	// This member is required.
	Type types.AppType

	// A Source object that specifies the app repository.
	AppSource *types.Source

	// One or more user-defined key/value pairs to be added to the stack attributes.
	Attributes map[string]string

	// The app's data source.
	DataSources []types.DataSource

	// A description of the app.
	Description *string

	// The app virtual host settings, with multiple domains separated by commas. For
	// example: 'www.example.com, example.com'
	Domains []string

	// Whether to enable SSL for the app.
	EnableSsl *bool

	// An array of EnvironmentVariable objects that specify environment variables to
	// be associated with the app. After you deploy the app, these variables are
	// defined on the associated app server instance. For more information, see [Environment Variables].
	//
	// There is no specific limit on the number of environment variables. However, the
	// size of the associated data structure - which includes the variables' names,
	// values, and protected flag values - cannot exceed 20 KB. This limit should
	// accommodate most if not all use cases. Exceeding it will cause an exception with
	// the message, "Environment: is too large (maximum is 20KB)."
	//
	// If you have specified one or more environment variables, you cannot modify the
	// stack's Chef version.
	//
	// [Environment Variables]: https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html#workingapps-creating-environment
	Environment []types.EnvironmentVariable

	// The app's short name.
	Shortname *string

	// An SslConfiguration object with the SSL configuration.
	SslConfiguration *types.SslConfiguration

	noSmithyDocumentSerde
}

// Contains the response to a CreateApp request.
type CreateAppOutput struct {

	// The app ID.
	AppId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateApp{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateApp"); err != nil {
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
	if err = addOpCreateAppValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateApp",
	}
}
