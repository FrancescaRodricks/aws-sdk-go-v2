// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Calling this API causes a message to be sent to the end user with a
// confirmation code that is required to change the user's password. For the
// Username parameter, you can use the username or user alias. The method used to
// send the confirmation code is sent according to the specified
// AccountRecoverySetting. For more information, see [Recovering User Accounts]in the Amazon Cognito
// Developer Guide. To use the confirmation code for resetting the password, call [ConfirmForgotPassword]
// .
//
// If neither a verified phone number nor a verified email exists, this API
// returns InvalidParameterException . If your app client has a client secret and
// you don't provide a SECRET_HASH parameter, this API returns
// NotAuthorizedException .
//
// To use this API operation, your user pool must have self-service account
// recovery configured. Use [AdminSetUserPassword]if you manage passwords as an administrator.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Service, Amazon Simple Notification Service might place your account
// in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone numbers.
// After you test your app while in the sandbox environment, you can move out of
// the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [ConfirmForgotPassword]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_ConfirmForgotPassword.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [AdminSetUserPassword]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminSetUserPassword.html
// [Recovering User Accounts]: https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-recover-a-user-account.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func (c *Client) ForgotPassword(ctx context.Context, params *ForgotPasswordInput, optFns ...func(*Options)) (*ForgotPasswordOutput, error) {
	if params == nil {
		params = &ForgotPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ForgotPassword", params, optFns, c.addOperationForgotPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ForgotPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to reset a user's password.
type ForgotPasswordInput struct {

	// The ID of the client associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The username of the user that you want to query or modify. The value of this
	// parameter is typically your user's username, but it can be any of their alias
	// attributes. If username isn't an alias attribute in your user pool, this value
	// must be the sub of a local user or the username of a user from a third-party
	// IdP.
	//
	// This member is required.
	Username *string

	// The Amazon Pinpoint analytics metadata that contributes to your metrics for
	// ForgotPassword calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers.
	//
	// You create custom workflows by assigning Lambda functions to user pool
	// triggers. When you use the ForgotPassword API action, Amazon Cognito invokes any
	// functions that are assigned to the following triggers: pre sign-up, custom
	// message, and user migration. When Amazon Cognito invokes any of these functions,
	// it passes a JSON payload, which the function receives as input. This payload
	// contains a clientMetadata attribute, which provides the data that you assigned
	// to the ClientMetadata parameter in your ForgotPassword request. In your function
	// code in Lambda, you can process the clientMetadata value to enhance your
	// workflow for your specific needs.
	//
	// For more information, see [Customizing user pool Workflows with Lambda Triggers] in the Amazon Cognito Developer Guide.
	//
	// When you use the ClientMetadata parameter, remember that Amazon Cognito won't
	// do the following:
	//
	//   - Store the ClientMetadata value. This data is available only to Lambda
	//   triggers that are assigned to a user pool to support custom workflows. If your
	//   user pool configuration doesn't include triggers, the ClientMetadata parameter
	//   serves no purpose.
	//
	//   - Validate the ClientMetadata value.
	//
	//   - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide
	//   sensitive information.
	//
	// [Customizing user pool Workflows with Lambda Triggers]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
	ClientMetadata map[string]string

	// A keyed-hash message authentication code (HMAC) calculated using the secret key
	// of a user pool client and username plus the client ID in the message.
	SecretHash *string

	// Contextual data about your user session, such as the device fingerprint, IP
	// address, or location. Amazon Cognito advanced security evaluates the risk of an
	// authentication event based on the context that your app generates and passes to
	// Amazon Cognito when it makes API requests.
	UserContextData *types.UserContextDataType

	noSmithyDocumentSerde
}

// The response from Amazon Cognito to a request to reset a password.
type ForgotPasswordOutput struct {

	// The code delivery details returned by the server in response to the request to
	// reset a password.
	CodeDeliveryDetails *types.CodeDeliveryDetailsType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationForgotPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpForgotPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpForgotPassword{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ForgotPassword"); err != nil {
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
	if err = addOpForgotPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opForgotPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opForgotPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ForgotPassword",
	}
}
