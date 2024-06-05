// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
	"time"
)

// Retrieves the details of a particular location registered in your S3 Access
// Grants instance.
//
// Permissions You must have the s3:GetAccessGrantsLocation permission to use this
// operation.
func (c *Client) GetAccessGrantsLocation(ctx context.Context, params *GetAccessGrantsLocationInput, optFns ...func(*Options)) (*GetAccessGrantsLocationOutput, error) {
	if params == nil {
		params = &GetAccessGrantsLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessGrantsLocation", params, optFns, c.addOperationGetAccessGrantsLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessGrantsLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessGrantsLocationInput struct {

	// The ID of the registered location that you are retrieving. S3 Access Grants
	// assigns this ID when you register the location. S3 Access Grants assigns the ID
	// default to the default location s3:// and assigns an auto-generated ID to other
	// locations that you register.
	//
	// This member is required.
	AccessGrantsLocationId *string

	// The ID of the Amazon Web Services account that is making this request.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

func (in *GetAccessGrantsLocationInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type GetAccessGrantsLocationOutput struct {

	// The Amazon Resource Name (ARN) of the registered location.
	AccessGrantsLocationArn *string

	// The ID of the registered location to which you are granting access. S3 Access
	// Grants assigns this ID when you register the location. S3 Access Grants assigns
	// the ID default to the default location s3:// and assigns an auto-generated ID
	// to other locations that you register.
	AccessGrantsLocationId *string

	// The date and time when you registered the location.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the IAM role for the registered location. S3
	// Access Grants assumes this role to manage access to the registered location.
	IAMRoleArn *string

	// The S3 URI path to the registered location. The location scope can be the
	// default S3 location s3:// , the S3 path to a bucket, or the S3 path to a bucket
	// and prefix. A prefix in S3 is a string of characters at the beginning of an
	// object key name used to organize the objects that you store in your S3 buckets.
	// For example, object key names that start with the engineering/ prefix or object
	// key names that start with the marketing/campaigns/ prefix.
	LocationScope *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessGrantsLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessGrantsLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessGrantsLocation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessGrantsLocation"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetAccessGrantsLocationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAccessGrantsLocationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessGrantsLocation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetAccessGrantsLocationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetAccessGrantsLocationMiddleware struct {
}

func (*endpointPrefix_opGetAccessGrantsLocationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAccessGrantsLocationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*GetAccessGrantsLocationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetAccessGrantsLocationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetAccessGrantsLocationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessGrantsLocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessGrantsLocation",
	}
}

func copyGetAccessGrantsLocationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetAccessGrantsLocationInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetAccessGrantsLocationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetAccessGrantsLocationInput) copy() interface{} {
	v := *in
	return &v
}
func backFillGetAccessGrantsLocationAccountID(input interface{}, v string) error {
	in := input.(*GetAccessGrantsLocationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetAccessGrantsLocationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetAccessGrantsLocationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
