// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a user's profile in AWS CodeStar. The user profile is not
// project-specific. Information in the user profile is displayed wherever the
// user's information appears to other users in AWS CodeStar.
func (c *Client) UpdateUserProfile(ctx context.Context, params *UpdateUserProfileInput, optFns ...func(*Options)) (*UpdateUserProfileOutput, error) {
	if params == nil {
		params = &UpdateUserProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserProfile", params, optFns, c.addOperationUpdateUserProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserProfileInput struct {

	// The name that will be displayed as the friendly name for the user in AWS
	// CodeStar.
	//
	// This member is required.
	UserArn *string

	// The name that is displayed as the friendly name for the user in AWS CodeStar.
	DisplayName *string

	// The email address that is displayed as part of the user's profile in AWS
	// CodeStar.
	EmailAddress *string

	// The SSH public key associated with the user in AWS CodeStar. If a project owner
	// allows the user remote access to project resources, this public key will be used
	// along with the user's private key for SSH access.
	SshPublicKey *string

	noSmithyDocumentSerde
}

type UpdateUserProfileOutput struct {

	// The Amazon Resource Name (ARN) of the user in IAM.
	//
	// This member is required.
	UserArn *string

	// The date the user profile was created, in timestamp format.
	CreatedTimestamp *time.Time

	// The name that is displayed as the friendly name for the user in AWS CodeStar.
	DisplayName *string

	// The email address that is displayed as part of the user's profile in AWS
	// CodeStar.
	EmailAddress *string

	// The date the user profile was last modified, in timestamp format.
	LastModifiedTimestamp *time.Time

	// The SSH public key associated with the user in AWS CodeStar. This is the public
	// portion of the public/private keypair the user can use to access project
	// resources if a project owner allows the user remote access to those resources.
	SshPublicKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateUserProfile"); err != nil {
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
	if err = addOpUpdateUserProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUserProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateUserProfile",
	}
}
