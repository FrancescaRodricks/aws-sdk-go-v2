// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Changes the state of a signing profile to REVOKED. This indicates that
// signatures generated using the signing profile after an effective start date are
// no longer valid.
func (c *Client) RevokeSigningProfile(ctx context.Context, params *RevokeSigningProfileInput, optFns ...func(*Options)) (*RevokeSigningProfileOutput, error) {
	if params == nil {
		params = &RevokeSigningProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeSigningProfile", params, optFns, c.addOperationRevokeSigningProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeSigningProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeSigningProfileInput struct {

	// A timestamp for when revocation of a Signing Profile should become effective.
	// Signatures generated using the signing profile after this timestamp are not
	// trusted.
	//
	// This member is required.
	EffectiveTime *time.Time

	// The name of the signing profile to be revoked.
	//
	// This member is required.
	ProfileName *string

	// The version of the signing profile to be revoked.
	//
	// This member is required.
	ProfileVersion *string

	// The reason for revoking a signing profile.
	//
	// This member is required.
	Reason *string

	noSmithyDocumentSerde
}

type RevokeSigningProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeSigningProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRevokeSigningProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRevokeSigningProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RevokeSigningProfile"); err != nil {
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
	if err = addOpRevokeSigningProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeSigningProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeSigningProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RevokeSigningProfile",
	}
}
