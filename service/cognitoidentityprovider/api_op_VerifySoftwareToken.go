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

// Use this API to register a user's entered time-based one-time password (TOTP)
// code and mark the user's software token MFA status as "verified" if successful.
// The request takes an access token or a session string, but not both.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func (c *Client) VerifySoftwareToken(ctx context.Context, params *VerifySoftwareTokenInput, optFns ...func(*Options)) (*VerifySoftwareTokenOutput, error) {
	if params == nil {
		params = &VerifySoftwareTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifySoftwareToken", params, optFns, c.addOperationVerifySoftwareTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifySoftwareTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VerifySoftwareTokenInput struct {

	// The one- time password computed using the secret code returned by [AssociateSoftwareToken].
	//
	// [AssociateSoftwareToken]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AssociateSoftwareToken.html
	//
	// This member is required.
	UserCode *string

	// A valid access token that Amazon Cognito issued to the user whose software
	// token you want to verify.
	AccessToken *string

	// The friendly device name.
	FriendlyDeviceName *string

	// The session that should be passed both ways in challenge-response calls to the
	// service.
	Session *string

	noSmithyDocumentSerde
}

type VerifySoftwareTokenOutput struct {

	// The session that should be passed both ways in challenge-response calls to the
	// service.
	Session *string

	// The status of the verify software token.
	Status types.VerifySoftwareTokenResponseType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifySoftwareTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpVerifySoftwareToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerifySoftwareToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "VerifySoftwareToken"); err != nil {
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
	if err = addOpVerifySoftwareTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifySoftwareToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifySoftwareToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifySoftwareToken",
	}
}
