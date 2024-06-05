// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Creates a new Amazon Web Services secret access key and corresponding Amazon
//
// Web Services access key ID for the specified user. The default status for new
// keys is Active .
//
// If you do not specify a user name, IAM determines the user name implicitly
// based on the Amazon Web Services access key ID signing the request. This
// operation works for access keys under the Amazon Web Services account.
// Consequently, you can use this operation to manage Amazon Web Services account
// root user credentials. This is true even if the Amazon Web Services account has
// no associated users.
//
// For information about quotas on the number of keys you can create, see [IAM and STS quotas] in the
// IAM User Guide.
//
// To ensure the security of your Amazon Web Services account, the secret access
// key is accessible only during key and user creation. You must save the key (for
// example, in a text file) if you want to be able to access it again. If a secret
// key is lost, you can delete the access keys for the associated user and then
// create new keys.
//
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
func (c *Client) CreateAccessKey(ctx context.Context, params *CreateAccessKeyInput, optFns ...func(*Options)) (*CreateAccessKeyOutput, error) {
	if params == nil {
		params = &CreateAccessKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessKey", params, optFns, c.addOperationCreateAccessKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessKeyInput struct {

	// The name of the IAM user that the new key will belong to.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	UserName *string

	noSmithyDocumentSerde
}

// Contains the response to a successful CreateAccessKey request.
type CreateAccessKeyOutput struct {

	// A structure with details about the access key.
	//
	// This member is required.
	AccessKey *types.AccessKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAccessKey"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAccessKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAccessKey",
	}
}
