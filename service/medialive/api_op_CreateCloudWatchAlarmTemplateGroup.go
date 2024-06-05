// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a cloudwatch alarm template group to group your cloudwatch alarm
// templates and to attach to signal maps for dynamically creating alarms.
func (c *Client) CreateCloudWatchAlarmTemplateGroup(ctx context.Context, params *CreateCloudWatchAlarmTemplateGroupInput, optFns ...func(*Options)) (*CreateCloudWatchAlarmTemplateGroupOutput, error) {
	if params == nil {
		params = &CreateCloudWatchAlarmTemplateGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCloudWatchAlarmTemplateGroup", params, optFns, c.addOperationCreateCloudWatchAlarmTemplateGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCloudWatchAlarmTemplateGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for CreateCloudWatchAlarmTemplateGroupRequest
type CreateCloudWatchAlarmTemplateGroupInput struct {

	// A resource's name. Names must be unique within the scope of a resource type in
	// a specific region.
	//
	// This member is required.
	Name *string

	// A resource's optional description.
	Description *string

	// Represents the tags associated with a resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Placeholder documentation for CreateCloudWatchAlarmTemplateGroupResponse
type CreateCloudWatchAlarmTemplateGroupOutput struct {

	// A cloudwatch alarm template group's ARN (Amazon Resource Name)
	Arn *string

	// Placeholder documentation for __timestampIso8601
	CreatedAt *time.Time

	// A resource's optional description.
	Description *string

	// A cloudwatch alarm template group's id. AWS provided template groups have ids
	// that start with aws-
	Id *string

	// Placeholder documentation for __timestampIso8601
	ModifiedAt *time.Time

	// A resource's name. Names must be unique within the scope of a resource type in
	// a specific region.
	Name *string

	// Represents the tags associated with a resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCloudWatchAlarmTemplateGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCloudWatchAlarmTemplateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCloudWatchAlarmTemplateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCloudWatchAlarmTemplateGroup"); err != nil {
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
	if err = addOpCreateCloudWatchAlarmTemplateGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCloudWatchAlarmTemplateGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCloudWatchAlarmTemplateGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCloudWatchAlarmTemplateGroup",
	}
}
