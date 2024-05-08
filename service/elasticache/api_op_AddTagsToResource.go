// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A tag is a key-value pair where the key and value are case-sensitive. You can
// use tags to categorize and track all your ElastiCache resources, with the
// exception of global replication group. When you add or remove tags on
// replication groups, those actions will be replicated to all nodes in the
// replication group. For more information, see [Resource-level permissions].
//
// For example, you can use cost-allocation tags to your ElastiCache resources,
// Amazon generates a cost allocation report as a comma-separated value (CSV) file
// with your usage and costs aggregated by your tags. You can apply tags that
// represent business categories (such as cost centers, application names, or
// owners) to organize your costs across multiple services.
//
// For more information, see [Using Cost Allocation Tags in Amazon ElastiCache] in the ElastiCache User Guide.
//
// [Using Cost Allocation Tags in Amazon ElastiCache]: https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Tagging.html
// [Resource-level permissions]: http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/IAM.ResourceLevelPermissions.html
func (c *Client) AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) {
	if params == nil {
		params = &AddTagsToResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddTagsToResource", params, optFns, c.addOperationAddTagsToResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddTagsToResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an AddTagsToResource operation.
type AddTagsToResourceInput struct {

	// The Amazon Resource Name (ARN) of the resource to which the tags are to be
	// added, for example arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster
	// or arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot . ElastiCache
	// resources are cluster and snapshot.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs) and Amazon Service Namespaces].
	//
	// [Amazon Resource Names (ARNs) and Amazon Service Namespaces]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	ResourceName *string

	// A list of tags to be added to this resource. A tag is a key-value pair. A tag
	// key must be accompanied by a tag value, although null is accepted.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Represents the output from the AddTagsToResource , ListTagsForResource , and
// RemoveTagsFromResource operations.
type AddTagsToResourceOutput struct {

	// A list of tags as key-value pairs.
	TagList []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddTagsToResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddTagsToResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddTagsToResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddTagsToResource"); err != nil {
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
	if err = addOpAddTagsToResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddTagsToResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddTagsToResource",
	}
}
