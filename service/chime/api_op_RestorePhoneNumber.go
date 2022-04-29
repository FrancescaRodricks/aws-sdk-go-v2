// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Moves a phone number from the Deletion queue back into the phone number
// Inventory.
func (c *Client) RestorePhoneNumber(ctx context.Context, params *RestorePhoneNumberInput, optFns ...func(*Options)) (*RestorePhoneNumberOutput, error) {
	if params == nil {
		params = &RestorePhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestorePhoneNumber", params, optFns, c.addOperationRestorePhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestorePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestorePhoneNumberInput struct {

	// The phone number.
	//
	// This member is required.
	PhoneNumberId *string

	noSmithyDocumentSerde
}

type RestorePhoneNumberOutput struct {

	// The phone number details.
	PhoneNumber *types.PhoneNumber

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestorePhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRestorePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRestorePhoneNumber{}, middleware.After)
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
	if err = addOpRestorePhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestorePhoneNumber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestorePhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "RestorePhoneNumber",
	}
}