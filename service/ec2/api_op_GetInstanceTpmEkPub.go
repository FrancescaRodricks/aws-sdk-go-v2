// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the public endorsement key associated with the Nitro Trusted Platform
// Module (NitroTPM) for the specified instance.
func (c *Client) GetInstanceTpmEkPub(ctx context.Context, params *GetInstanceTpmEkPubInput, optFns ...func(*Options)) (*GetInstanceTpmEkPubOutput, error) {
	if params == nil {
		params = &GetInstanceTpmEkPubInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstanceTpmEkPub", params, optFns, c.addOperationGetInstanceTpmEkPubMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstanceTpmEkPubOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstanceTpmEkPubInput struct {

	// The ID of the instance for which to get the public endorsement key.
	//
	// This member is required.
	InstanceId *string

	// The required public endorsement key format. Specify der for a DER-encoded
	// public key that is compatible with OpenSSL. Specify tpmt for a TPM 2.0 format
	// that is compatible with tpm2-tools. The returned key is base64 encoded.
	//
	// This member is required.
	KeyFormat types.EkPubKeyFormat

	// The required public endorsement key type.
	//
	// This member is required.
	KeyType types.EkPubKeyType

	// Specify this parameter to verify whether the request will succeed, without
	// actually making the request. If the request will succeed, the response is
	// DryRunOperation . Otherwise, the response is UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type GetInstanceTpmEkPubOutput struct {

	// The ID of the instance.
	InstanceId *string

	// The public endorsement key format.
	KeyFormat types.EkPubKeyFormat

	// The public endorsement key type.
	KeyType types.EkPubKeyType

	// The public endorsement key material.
	KeyValue *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInstanceTpmEkPubMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetInstanceTpmEkPub{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetInstanceTpmEkPub{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetInstanceTpmEkPub"); err != nil {
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
	if err = addOpGetInstanceTpmEkPubValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceTpmEkPub(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInstanceTpmEkPub(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetInstanceTpmEkPub",
	}
}
