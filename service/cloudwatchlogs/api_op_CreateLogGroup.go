// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a log group with the specified name. You can create up to 1,000,000 log
// groups per Region per account.
//
// You must use the following guidelines when naming a log group:
//
//   - Log group names must be unique within a Region for an Amazon Web Services
//     account.
//
//   - Log group names can be between 1 and 512 characters long.
//
//   - Log group names consist of the following characters: a-z, A-Z, 0-9, '_'
//     (underscore), '-' (hyphen), '/' (forward slash), '.' (period), and '#' (number
//     sign)
//
//   - Log group names can't start with the string aws/
//
// When you create a log group, by default the log events in the log group do not
// expire. To set a retention policy so that events expire and are deleted after a
// specified time, use [PutRetentionPolicy].
//
// If you associate an KMS key with the log group, ingested data is encrypted
// using the KMS key. This association is stored as long as the data encrypted with
// the KMS key is still within CloudWatch Logs. This enables CloudWatch Logs to
// decrypt this data whenever it is requested.
//
// If you attempt to associate a KMS key with the log group but the KMS key does
// not exist or the KMS key is disabled, you receive an InvalidParameterException
// error.
//
// CloudWatch Logs supports only symmetric KMS keys. Do not associate an
// asymmetric KMS key with your log group. For more information, see [Using Symmetric and Asymmetric Keys].
//
// [Using Symmetric and Asymmetric Keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [PutRetentionPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutRetentionPolicy.html
func (c *Client) CreateLogGroup(ctx context.Context, params *CreateLogGroupInput, optFns ...func(*Options)) (*CreateLogGroupOutput, error) {
	if params == nil {
		params = &CreateLogGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLogGroup", params, optFns, c.addOperationCreateLogGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLogGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogGroupInput struct {

	// A name for the log group.
	//
	// This member is required.
	LogGroupName *string

	// The Amazon Resource Name (ARN) of the KMS key to use when encrypting log data.
	// For more information, see [Amazon Resource Names].
	//
	// [Amazon Resource Names]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms
	KmsKeyId *string

	// Use this parameter to specify the log group class for this log group. There are
	// two classes:
	//
	//   - The Standard log class supports all CloudWatch Logs features.
	//
	//   - The Infrequent Access log class supports a subset of CloudWatch Logs
	//   features and incurs lower costs.
	//
	// If you omit this parameter, the default of STANDARD is used.
	//
	// The value of logGroupClass can't be changed after a log group is created.
	//
	// For details about the features supported by each class, see [Log classes]
	//
	// [Log classes]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch_Logs_Log_Classes.html
	LogGroupClass types.LogGroupClass

	// The key-value pairs to use for the tags.
	//
	// You can grant users access to certain log groups while preventing them from
	// accessing other log groups. To do so, tag your groups and use IAM policies that
	// refer to those tags. To assign tags when you create a log group, you must have
	// either the logs:TagResource or logs:TagLogGroup permission. For more
	// information about tagging, see [Tagging Amazon Web Services resources]. For more information about using tags to
	// control access, see [Controlling access to Amazon Web Services resources using tags].
	//
	// [Controlling access to Amazon Web Services resources using tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
	// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLogGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLogGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLogGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLogGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLogGroup"); err != nil {
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
	if err = addOpCreateLogGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLogGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLogGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLogGroup",
	}
}
