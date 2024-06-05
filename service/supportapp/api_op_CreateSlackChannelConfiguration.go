// Code generated by smithy-go-codegen DO NOT EDIT.

package supportapp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/supportapp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Slack channel configuration for your Amazon Web Services account.
//
//   - You can add up to 5 Slack workspaces for your account.
//
//   - You can add up to 20 Slack channels for your account.
//
// A Slack channel can have up to 100 Amazon Web Services accounts. This means
// that only 100 accounts can add the same Slack channel to the Amazon Web Services
// Support App. We recommend that you only add the accounts that you need to manage
// support cases for your organization. This can reduce the notifications about
// case updates that you receive in the Slack channel.
//
// We recommend that you choose a private Slack channel so that only members in
// that channel have read and write access to your support cases. Anyone in your
// Slack channel can create, update, or resolve support cases for your account.
// Users require an invitation to join private channels.
func (c *Client) CreateSlackChannelConfiguration(ctx context.Context, params *CreateSlackChannelConfigurationInput, optFns ...func(*Options)) (*CreateSlackChannelConfigurationOutput, error) {
	if params == nil {
		params = &CreateSlackChannelConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSlackChannelConfiguration", params, optFns, c.addOperationCreateSlackChannelConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSlackChannelConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSlackChannelConfigurationInput struct {

	// The channel ID in Slack. This ID identifies a channel within a Slack workspace.
	//
	// This member is required.
	ChannelId *string

	// The Amazon Resource Name (ARN) of an IAM role that you want to use to perform
	// operations on Amazon Web Services. For more information, see [Managing access to the Amazon Web Services Support App]in the Amazon Web
	// Services Support User Guide.
	//
	// [Managing access to the Amazon Web Services Support App]: https://docs.aws.amazon.com/awssupport/latest/user/support-app-permissions.html
	//
	// This member is required.
	ChannelRoleArn *string

	// The case severity for a support case that you want to receive notifications.
	//
	// If you specify high or all , you must specify true for at least one of the
	// following parameters:
	//
	//   - notifyOnAddCorrespondenceToCase
	//
	//   - notifyOnCreateOrReopenCase
	//
	//   - notifyOnResolveCase
	//
	// If you specify none , the following parameters must be null or false :
	//
	//   - notifyOnAddCorrespondenceToCase
	//
	//   - notifyOnCreateOrReopenCase
	//
	//   - notifyOnResolveCase
	//
	// If you don't specify these parameters in your request, they default to false .
	//
	// This member is required.
	NotifyOnCaseSeverity types.NotificationSeverityLevel

	// The team ID in Slack. This ID uniquely identifies a Slack workspace, such as
	// T012ABCDEFG .
	//
	// This member is required.
	TeamId *string

	// The name of the Slack channel that you configure for the Amazon Web Services
	// Support App.
	ChannelName *string

	// Whether you want to get notified when a support case has a new correspondence.
	NotifyOnAddCorrespondenceToCase *bool

	// Whether you want to get notified when a support case is created or reopened.
	NotifyOnCreateOrReopenCase *bool

	// Whether you want to get notified when a support case is resolved.
	NotifyOnResolveCase *bool

	noSmithyDocumentSerde
}

type CreateSlackChannelConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSlackChannelConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSlackChannelConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSlackChannelConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSlackChannelConfiguration"); err != nil {
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
	if err = addOpCreateSlackChannelConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSlackChannelConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSlackChannelConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSlackChannelConfiguration",
	}
}
