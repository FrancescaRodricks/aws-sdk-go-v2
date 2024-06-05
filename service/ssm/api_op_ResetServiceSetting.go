// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// ServiceSetting is an account-level setting for an Amazon Web Services service.
// This setting defines how a user interacts with or uses a service or a feature of
// a service. For example, if an Amazon Web Services service charges money to the
// account based on feature or service usage, then the Amazon Web Services service
// team might create a default setting of "false". This means the user can't use
// this feature unless they change the setting to "true" and intentionally opt in
// for a paid feature.
//
// Services map a SettingId object to a setting value. Amazon Web Services
// services teams define the default value for a SettingId . You can't create a new
// SettingId , but you can overwrite the default value if you have the
// ssm:UpdateServiceSetting permission for the setting. Use the GetServiceSetting API operation to
// view the current value. Use the UpdateServiceSettingAPI operation to change the default setting.
//
// Reset the service setting for the account to the default value as provisioned
// by the Amazon Web Services service team.
func (c *Client) ResetServiceSetting(ctx context.Context, params *ResetServiceSettingInput, optFns ...func(*Options)) (*ResetServiceSettingOutput, error) {
	if params == nil {
		params = &ResetServiceSettingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetServiceSetting", params, optFns, c.addOperationResetServiceSettingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetServiceSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body of the ResetServiceSetting API operation.
type ResetServiceSettingInput struct {

	// The Amazon Resource Name (ARN) of the service setting to reset. The setting ID
	// can be one of the following.
	//
	//   - /ssm/managed-instance/default-ec2-instance-management-role
	//
	//   - /ssm/automation/customer-script-log-destination
	//
	//   - /ssm/automation/customer-script-log-group-name
	//
	//   - /ssm/documents/console/public-sharing-permission
	//
	//   - /ssm/managed-instance/activation-tier
	//
	//   - /ssm/opsinsights/opscenter
	//
	//   - /ssm/parameter-store/default-parameter-tier
	//
	//   - /ssm/parameter-store/high-throughput-enabled
	//
	// This member is required.
	SettingId *string

	noSmithyDocumentSerde
}

// The result body of the ResetServiceSetting API operation.
type ResetServiceSettingOutput struct {

	// The current, effective service setting after calling the ResetServiceSetting
	// API operation.
	ServiceSetting *types.ServiceSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetServiceSettingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResetServiceSetting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetServiceSetting{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ResetServiceSetting"); err != nil {
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
	if err = addOpResetServiceSettingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetServiceSetting(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetServiceSetting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ResetServiceSetting",
	}
}
