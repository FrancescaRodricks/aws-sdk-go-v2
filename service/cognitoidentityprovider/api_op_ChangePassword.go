// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the password for a specified user in a user pool.
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
func (c *Client) ChangePassword(ctx context.Context, params *ChangePasswordInput, optFns ...func(*Options)) (*ChangePasswordOutput, error) {
	if params == nil {
		params = &ChangePasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangePassword", params, optFns, c.addOperationChangePasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangePasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to change a user password.
type ChangePasswordInput struct {

	// A valid access token that Amazon Cognito issued to the user whose password you
	// want to change.
	//
	// This member is required.
	AccessToken *string

	// The old password.
	//
	// This member is required.
	PreviousPassword *string

	// The new password.
	//
	// This member is required.
	ProposedPassword *string

	noSmithyDocumentSerde
}

// The response from the server to the change password request.
type ChangePasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationChangePasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpChangePassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpChangePassword{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ChangePassword"); err != nil {
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
	if err = addOpChangePasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChangePassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opChangePassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ChangePassword",
	}
}
