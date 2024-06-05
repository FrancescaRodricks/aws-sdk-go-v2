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

// Creates a new user in the specified user pool.
//
// If MessageAction isn't set, the default is to send a welcome message via email
// or phone (SMS).
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
// This message is based on a template that you configured in your call to create
// or update a user pool. This template includes your custom sign-up instructions
// and placeholders for user name and temporary password.
//
// Alternatively, you can call AdminCreateUser with SUPPRESS for the MessageAction
// parameter, and Amazon Cognito won't send any email.
//
// In either case, the user will be in the FORCE_CHANGE_PASSWORD state until they
// sign in and change their password.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func (c *Client) AdminCreateUser(ctx context.Context, params *AdminCreateUserInput, optFns ...func(*Options)) (*AdminCreateUserOutput, error) {
	if params == nil {
		params = &AdminCreateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminCreateUser", params, optFns, c.addOperationAdminCreateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminCreateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to create a user in the specified user pool.
type AdminCreateUserInput struct {

	// The user pool ID for the user pool where the user will be created.
	//
	// This member is required.
	UserPoolId *string

	// The value that you want to set as the username sign-in attribute. The following
	// conditions apply to the username parameter.
	//
	//   - The username can't be a duplicate of another username in the same user pool.
	//
	//   - You can't change the value of a username after you create it.
	//
	//   - You can only provide a value if usernames are a valid sign-in attribute for
	//   your user pool. If your user pool only supports phone numbers or email addresses
	//   as sign-in attributes, Amazon Cognito automatically generates a username value.
	//   For more information, see [Customizing sign-in attributes].
	//
	// [Customizing sign-in attributes]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-aliases
	//
	// This member is required.
	Username *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers.
	//
	// You create custom workflows by assigning Lambda functions to user pool
	// triggers. When you use the AdminCreateUser API action, Amazon Cognito invokes
	// the function that is assigned to the pre sign-up trigger. When Amazon Cognito
	// invokes this function, it passes a JSON payload, which the function receives as
	// input. This payload contains a clientMetadata attribute, which provides the
	// data that you assigned to the ClientMetadata parameter in your AdminCreateUser
	// request. In your function code in Lambda, you can process the clientMetadata
	// value to enhance your workflow for your specific needs.
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

	// Specify "EMAIL" if email will be used to send the welcome message. Specify "SMS"
	// if the phone number will be used. The default value is "SMS" . You can specify
	// more than one value.
	DesiredDeliveryMediums []types.DeliveryMediumType

	// This parameter is used only if the phone_number_verified or email_verified
	// attribute is set to True . Otherwise, it is ignored.
	//
	// If this parameter is set to True and the phone number or email address
	// specified in the UserAttributes parameter already exists as an alias with a
	// different user, the API call will migrate the alias from the previous user to
	// the newly created user. The previous user will no longer be able to log in using
	// that alias.
	//
	// If this parameter is set to False , the API throws an AliasExistsException
	// error if the alias already exists. The default value is False .
	ForceAliasCreation bool

	// Set to RESEND to resend the invitation message to a user that already exists
	// and reset the expiration limit on the user's account. Set to SUPPRESS to
	// suppress sending the message. You can specify only one value.
	MessageAction types.MessageActionType

	// The user's temporary password. This password must conform to the password
	// policy that you specified when you created the user pool.
	//
	// The temporary password is valid only once. To complete the Admin Create User
	// flow, the user must enter the temporary password in the sign-in page, along with
	// a new password to be used in all future sign-ins.
	//
	// This parameter isn't required. If you don't specify a value, Amazon Cognito
	// generates one for you.
	//
	// The temporary password can only be used until the user account expiration limit
	// that you set for your user pool. To reset the account after that time limit, you
	// must call AdminCreateUser again and specify RESEND for the MessageAction
	// parameter.
	TemporaryPassword *string

	// An array of name-value pairs that contain user attributes and attribute values
	// to be set for the user to be created. You can create a user without specifying
	// any attributes other than Username . However, any attributes that you specify as
	// required (when creating a user pool or in the Attributes tab of the console)
	// either you should supply (in your call to AdminCreateUser ) or the user should
	// supply (when they sign up in response to your welcome message).
	//
	// For custom attributes, you must prepend the custom: prefix to the attribute
	// name.
	//
	// To send a message inviting the user to sign up, you must specify the user's
	// email address or phone number. You can do this in your call to AdminCreateUser
	// or in the Users tab of the Amazon Cognito console for managing your user pools.
	//
	// In your call to AdminCreateUser , you can set the email_verified attribute to
	// True , and you can set the phone_number_verified attribute to True . You can
	// also do this by calling [AdminUpdateUserAttributes].
	//
	//   - email: The email address of the user to whom the message that contains the
	//   code and username will be sent. Required if the email_verified attribute is
	//   set to True , or if "EMAIL" is specified in the DesiredDeliveryMediums
	//   parameter.
	//
	//   - phone_number: The phone number of the user to whom the message that
	//   contains the code and username will be sent. Required if the
	//   phone_number_verified attribute is set to True , or if "SMS" is specified in
	//   the DesiredDeliveryMediums parameter.
	//
	// [AdminUpdateUserAttributes]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminUpdateUserAttributes.html
	UserAttributes []types.AttributeType

	// Temporary user attributes that contribute to the outcomes of your pre sign-up
	// Lambda trigger. This set of key-value pairs are for custom validation of
	// information that you collect from your users but don't need to retain.
	//
	// Your Lambda function can analyze this additional data and act on it. Your
	// function might perform external API operations like logging user attributes and
	// validation data to Amazon CloudWatch Logs. Validation data might also affect the
	// response that your function returns to Amazon Cognito, like automatically
	// confirming the user if they sign up from within your network.
	//
	// For more information about the pre sign-up Lambda trigger, see [Pre sign-up Lambda trigger].
	//
	// [Pre sign-up Lambda trigger]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html
	ValidationData []types.AttributeType

	noSmithyDocumentSerde
}

// Represents the response from the server to the request to create the user.
type AdminCreateUserOutput struct {

	// The newly created user.
	User *types.UserType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminCreateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminCreateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminCreateUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdminCreateUser"); err != nil {
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
	if err = addOpAdminCreateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminCreateUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminCreateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdminCreateUser",
	}
}
