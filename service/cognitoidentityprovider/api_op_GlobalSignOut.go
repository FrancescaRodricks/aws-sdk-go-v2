// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Invalidates the identity, access, and refresh tokens that Amazon Cognito issued
// to a user. Call this operation when your user signs out of your app. This
// results in the following behavior.
//
//   - Amazon Cognito no longer accepts token-authorized user operations that you
//     authorize with a signed-out user's access tokens. For more information, see [Using the Amazon Cognito user pools API and user pool endpoints].
//
// Amazon Cognito returns an Access Token has been revoked error when your app
//
//	attempts to authorize a user pools API request with a revoked access token that
//	contains the scope aws.cognito.signin.user.admin .
//
//	- Amazon Cognito no longer accepts a signed-out user's ID token in a [GetId]request
//	to an identity pool with ServerSideTokenCheck enabled for its user pool IdP
//	configuration in [CognitoIdentityProvider].
//
//	- Amazon Cognito no longer accepts a signed-out user's refresh tokens in
//	refresh requests.
//
// Other requests might be valid until your user's token expires.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
//
// [CognitoIdentityProvider]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_CognitoIdentityProvider.html
// [GetId]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetId.html
func (c *Client) GlobalSignOut(ctx context.Context, params *GlobalSignOutInput, optFns ...func(*Options)) (*GlobalSignOutOutput, error) {
	if params == nil {
		params = &GlobalSignOutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GlobalSignOut", params, optFns, c.addOperationGlobalSignOutMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GlobalSignOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to sign out all devices.
type GlobalSignOutInput struct {

	// A valid access token that Amazon Cognito issued to the user who you want to
	// sign out.
	//
	// This member is required.
	AccessToken *string

	noSmithyDocumentSerde
}

// The response to the request to sign out all devices.
type GlobalSignOutOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGlobalSignOutMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGlobalSignOut{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGlobalSignOut{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GlobalSignOut"); err != nil {
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
	if err = addOpGlobalSignOutValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGlobalSignOut(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGlobalSignOut(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GlobalSignOut",
	}
}
