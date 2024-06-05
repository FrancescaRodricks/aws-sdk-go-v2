// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the record of an existing introspection. If the retrieval is
// successful, the result of the instrospection will also be returned. If the
// retrieval fails the operation, an error message will be returned instead.
func (c *Client) GetDataSourceIntrospection(ctx context.Context, params *GetDataSourceIntrospectionInput, optFns ...func(*Options)) (*GetDataSourceIntrospectionOutput, error) {
	if params == nil {
		params = &GetDataSourceIntrospectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataSourceIntrospection", params, optFns, c.addOperationGetDataSourceIntrospectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataSourceIntrospectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataSourceIntrospectionInput struct {

	// The introspection ID. Each introspection contains a unique ID that can be used
	// to reference the instrospection record.
	//
	// This member is required.
	IntrospectionId *string

	// A boolean flag that determines whether SDL should be generated for introspected
	// types or not. If set to true , each model will contain an sdl property that
	// contains the SDL for that type. The SDL only contains the type data and no
	// additional metadata or directives.
	IncludeModelsSDL bool

	// The maximum number of introspected types that will be returned in a single
	// response.
	MaxResults int32

	// Determines the number of types to be returned in a single response before
	// paginating. This value is typically taken from nextToken value from the
	// previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetDataSourceIntrospectionOutput struct {

	// The introspection ID. Each introspection contains a unique ID that can be used
	// to reference the instrospection record.
	IntrospectionId *string

	// The DataSourceIntrospectionResult object data.
	IntrospectionResult *types.DataSourceIntrospectionResult

	// The status of the introspection during retrieval. By default, when a new
	// instrospection is being retrieved, the status will be set to PROCESSING . Once
	// the operation has been completed, the status will change to SUCCESS or FAILED
	// depending on how the data was parsed. A FAILED operation will return an error
	// and its details as an introspectionStatusDetail .
	IntrospectionStatus types.DataSourceIntrospectionStatus

	// The error detail field. When a FAILED introspectionStatus is returned, the
	// introspectionStatusDetail will also return the exact error that was generated
	// during the operation.
	IntrospectionStatusDetail *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataSourceIntrospectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataSourceIntrospection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataSourceIntrospection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataSourceIntrospection"); err != nil {
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
	if err = addOpGetDataSourceIntrospectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataSourceIntrospection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataSourceIntrospection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataSourceIntrospection",
	}
}
