// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the RestApi resource in the collection.
func (c *Client) GetRestApi(ctx context.Context, params *GetRestApiInput, optFns ...func(*Options)) (*GetRestApiOutput, error) {
	if params == nil {
		params = &GetRestApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRestApi", params, optFns, c.addOperationGetRestApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRestApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to list an existing RestApi defined for your collection.
type GetRestApiInput struct {

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	noSmithyDocumentSerde
}

// Represents a REST API.
type GetRestApiOutput struct {

	// The source of the API key for metering requests according to a usage plan.
	// Valid values are: > HEADER to read the API key from the X-API-Key header of a
	// request. AUTHORIZER to read the API key from the UsageIdentifierKey from a
	// custom authorizer.
	ApiKeySource types.ApiKeySourceType

	// The list of binary media types supported by the RestApi. By default, the
	// RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []string

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The API's description.
	Description *string

	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint bool

	// The endpoint configuration of this RestApi showing the endpoint types of the
	// API.
	EndpointConfiguration *types.EndpointConfiguration

	// The API's identifier. This identifier is unique across all of your APIs in API
	// Gateway.
	Id *string

	// A nullable integer that is used to enable compression (with non-negative
	// between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a
	// null value) on an API. When compression is enabled, compression or decompression
	// is not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int32

	// The API's name.
	Name *string

	// A stringified JSON policy document that applies to this RestApi regardless of
	// the caller and Method configuration.
	Policy *string

	// The API's root resource ID.
	RootResourceId *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRestApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRestApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRestApi{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRestApi"); err != nil {
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
	if err = addOpGetRestApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRestApi(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetRestApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRestApi",
	}
}
