// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation gets an Amazon S3 on Outposts bucket's replication
// configuration. To get an S3 bucket's replication configuration, see [GetBucketReplication]in the
// Amazon S3 API Reference.
//
// Returns the replication configuration of an S3 on Outposts bucket. For more
// information about S3 on Outposts, see [Using Amazon S3 on Outposts]in the Amazon S3 User Guide. For
// information about S3 replication on Outposts configuration, see [Replicating objects for S3 on Outposts]in the Amazon
// S3 User Guide.
//
// It can take a while to propagate PUT or DELETE requests for a replication
// configuration to all S3 on Outposts systems. Therefore, the replication
// configuration that's returned by a GET request soon after a PUT or DELETE
// request might return a more recent result than what's on the Outpost. If an
// Outpost is offline, the delay in updating the replication configuration on that
// Outpost can be significant.
//
// This action requires permissions for the s3-outposts:GetReplicationConfiguration
// action. The Outposts bucket owner has this permission by default and can grant
// it to others. For more information about permissions, see [Setting up IAM with S3 on Outposts]and [Managing access to S3 on Outposts bucket] in the Amazon S3
// User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// If you include the Filter element in a replication configuration, you must also
// include the DeleteMarkerReplication , Status , and Priority elements. The
// response also returns those elements.
//
// For information about S3 on Outposts replication failure reasons, see [Replication failure reasons] in the
// Amazon S3 User Guide.
//
// The following operations are related to GetBucketReplication :
//
// [PutBucketReplication]
//
// [DeleteBucketReplication]
//
// [Replicating objects for S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html
// [Replication failure reasons]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/outposts-replication-eventbridge.html#outposts-replication-failure-codes
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html
// [Setting up IAM with S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html
// [Managing access to S3 on Outposts bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html#API_control_GetBucketReplication_Examples
func (c *Client) GetBucketReplication(ctx context.Context, params *GetBucketReplicationInput, optFns ...func(*Options)) (*GetBucketReplicationOutput, error) {
	if params == nil {
		params = &GetBucketReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketReplication", params, optFns, c.addOperationGetBucketReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketReplicationInput struct {

	// The Amazon Web Services account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// Specifies the bucket to get the replication information for.
	//
	// For using this parameter with Amazon S3 on Outposts with the REST API, you must
	// specify the name and the x-amz-outpost-id as well.
	//
	// For using this parameter with S3 on Outposts with the Amazon Web Services SDK
	// and CLI, you must specify the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/ . For example, to access the bucket
	// reports through Outpost my-outpost owned by account 123456789012 in Region
	// us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	noSmithyDocumentSerde
}

func (in *GetBucketReplicationInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type GetBucketReplicationOutput struct {

	// A container for one or more replication rules. A replication configuration must
	// have at least one rule and you can add up to 100 rules. The maximum size of a
	// replication configuration is 128 KB.
	ReplicationConfiguration *types.ReplicationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBucketReplication"); err != nil {
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
	if err = addEndpointPrefix_opGetBucketReplicationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetBucketReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketReplication(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetBucketReplicationUpdateEndpoint(stack, options); err != nil {
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

func (m *GetBucketReplicationInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *GetBucketReplicationInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opGetBucketReplicationMiddleware struct {
}

func (*endpointPrefix_opGetBucketReplicationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetBucketReplicationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
	input, ok := opaqueInput.(*GetBucketReplicationInput)
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
func addEndpointPrefix_opGetBucketReplicationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetBucketReplicationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBucketReplication",
	}
}

func copyGetBucketReplicationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetBucketReplicationInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetBucketReplicationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetBucketReplicationInput) copy() interface{} {
	v := *in
	return &v
}
func getGetBucketReplicationARNMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketReplicationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setGetBucketReplicationARNMember(input interface{}, v string) error {
	in := input.(*GetBucketReplicationInput)
	in.Bucket = &v
	return nil
}
func backFillGetBucketReplicationAccountID(input interface{}, v string) error {
	in := input.(*GetBucketReplicationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetBucketReplicationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getGetBucketReplicationARNMember,
			BackfillAccountID: backFillGetBucketReplicationAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setGetBucketReplicationARNMember,
			CopyInput:         copyGetBucketReplicationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
