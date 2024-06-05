// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a job to create a camera stream node.
func (c *Client) DescribeNodeFromTemplateJob(ctx context.Context, params *DescribeNodeFromTemplateJobInput, optFns ...func(*Options)) (*DescribeNodeFromTemplateJobOutput, error) {
	if params == nil {
		params = &DescribeNodeFromTemplateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNodeFromTemplateJob", params, optFns, c.addOperationDescribeNodeFromTemplateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNodeFromTemplateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNodeFromTemplateJobInput struct {

	// The job's ID.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeNodeFromTemplateJobOutput struct {

	// When the job was created.
	//
	// This member is required.
	CreatedTime *time.Time

	// The job's ID.
	//
	// This member is required.
	JobId *string

	// When the job was updated.
	//
	// This member is required.
	LastUpdatedTime *time.Time

	// The node's name.
	//
	// This member is required.
	NodeName *string

	// The job's output package name.
	//
	// This member is required.
	OutputPackageName *string

	// The job's output package version.
	//
	// This member is required.
	OutputPackageVersion *string

	// The job's status.
	//
	// This member is required.
	Status types.NodeFromTemplateJobStatus

	// The job's status message.
	//
	// This member is required.
	StatusMessage *string

	// The job's template parameters.
	//
	// This member is required.
	TemplateParameters map[string]string

	// The job's template type.
	//
	// This member is required.
	TemplateType types.TemplateType

	// The job's tags.
	JobTags []types.JobResourceTags

	// The node's description.
	NodeDescription *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNodeFromTemplateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeNodeFromTemplateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeNodeFromTemplateJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeNodeFromTemplateJob"); err != nil {
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
	if err = addOpDescribeNodeFromTemplateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNodeFromTemplateJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeNodeFromTemplateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeNodeFromTemplateJob",
	}
}
