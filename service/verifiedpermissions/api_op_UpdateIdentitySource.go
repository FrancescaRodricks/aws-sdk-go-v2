// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified identity source to use a new identity provider (IdP)
// source, or to change the mapping of identities from the IdP to a different
// principal entity type.
func (c *Client) UpdateIdentitySource(ctx context.Context, params *UpdateIdentitySourceInput, optFns ...func(*Options)) (*UpdateIdentitySourceOutput, error) {
	if params == nil {
		params = &UpdateIdentitySourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIdentitySource", params, optFns, c.addOperationUpdateIdentitySourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIdentitySourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIdentitySourceInput struct {

	// Specifies the ID of the identity source that you want to update.
	//
	// This member is required.
	IdentitySourceId *string

	// Specifies the ID of the policy store that contains the identity source that you
	// want to update.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies the details required to communicate with the identity provider (IdP)
	// associated with this identity source. At this time, the only valid member of
	// this structure is a Amazon Cognito user pool configuration. You must specify a
	// userPoolArn , and optionally, a ClientId .
	//
	// This member is required.
	UpdateConfiguration types.UpdateConfiguration

	// Specifies the data type of principals generated for identities authenticated by
	// the identity source.
	PrincipalEntityType *string

	noSmithyDocumentSerde
}

type UpdateIdentitySourceOutput struct {

	// The date and time that the updated identity source was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The ID of the updated identity source.
	//
	// This member is required.
	IdentitySourceId *string

	// The date and time that the identity source was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The ID of the policy store that contains the updated identity source.
	//
	// This member is required.
	PolicyStoreId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIdentitySourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateIdentitySource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateIdentitySource{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateIdentitySourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIdentitySource(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opUpdateIdentitySource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "verifiedpermissions",
		OperationName: "UpdateIdentitySource",
	}
}
