// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new health check.
//
// For information about adding health checks to resource record sets, see [HealthCheckId] in [ChangeResourceRecordSets].
//
// # ELB Load Balancers
//
// If you're registering EC2 instances with an Elastic Load Balancing (ELB) load
// balancer, do not create Amazon Route 53 health checks for the EC2 instances.
// When you register an EC2 instance with a load balancer, you configure settings
// for an ELB health check, which performs a similar function to a Route 53 health
// check.
//
// # Private Hosted Zones
//
// You can associate health checks with failover resource record sets in a private
// hosted zone. Note the following:
//
//   - Route 53 health checkers are outside the VPC. To check the health of an
//     endpoint within a VPC by IP address, you must assign a public IP address to the
//     instance in the VPC.
//
//   - You can configure a health checker to check the health of an external
//     resource that the instance relies on, such as a database server.
//
//   - You can create a CloudWatch metric, associate an alarm with the metric, and
//     then create a health check that is based on the state of the alarm. For example,
//     you might create a CloudWatch metric that checks the status of the Amazon EC2
//     StatusCheckFailed metric, add an alarm to the metric, and then create a health
//     check that is based on the state of the alarm. For information about creating
//     CloudWatch metrics and alarms by using the CloudWatch console, see the [Amazon CloudWatch User Guide].
//
// [HealthCheckId]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html#Route53-Type-ResourceRecordSet-HealthCheckId
// [ChangeResourceRecordSets]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html
// [Amazon CloudWatch User Guide]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/WhatIsCloudWatch.html
func (c *Client) CreateHealthCheck(ctx context.Context, params *CreateHealthCheckInput, optFns ...func(*Options)) (*CreateHealthCheckOutput, error) {
	if params == nil {
		params = &CreateHealthCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHealthCheck", params, optFns, c.addOperationCreateHealthCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains the health check request information.
type CreateHealthCheckInput struct {

	// A unique string that identifies the request and that allows you to retry a
	// failed CreateHealthCheck request without the risk of creating two identical
	// health checks:
	//
	//   - If you send a CreateHealthCheck request with the same CallerReference and
	//   settings as a previous request, and if the health check doesn't exist, Amazon
	//   Route 53 creates the health check. If the health check does exist, Route 53
	//   returns the settings for the existing health check.
	//
	//   - If you send a CreateHealthCheck request with the same CallerReference as a
	//   deleted health check, regardless of the settings, Route 53 returns a
	//   HealthCheckAlreadyExists error.
	//
	//   - If you send a CreateHealthCheck request with the same CallerReference as an
	//   existing health check but with different settings, Route 53 returns a
	//   HealthCheckAlreadyExists error.
	//
	//   - If you send a CreateHealthCheck request with a unique CallerReference but
	//   settings identical to an existing health check, Route 53 creates the health
	//   check.
	//
	// Route 53 does not store the CallerReference for a deleted health check
	// indefinitely. The CallerReference for a deleted health check will be deleted
	// after a number of days.
	//
	// This member is required.
	CallerReference *string

	// A complex type that contains settings for a new health check.
	//
	// This member is required.
	HealthCheckConfig *types.HealthCheckConfig

	noSmithyDocumentSerde
}

// A complex type containing the response information for the new health check.
type CreateHealthCheckOutput struct {

	// A complex type that contains identifying information about the health check.
	//
	// This member is required.
	HealthCheck *types.HealthCheck

	// The unique URL representing the new health check.
	//
	// This member is required.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHealthCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHealthCheck"); err != nil {
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
	if err = addOpCreateHealthCheckValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHealthCheck(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHealthCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHealthCheck",
	}
}
