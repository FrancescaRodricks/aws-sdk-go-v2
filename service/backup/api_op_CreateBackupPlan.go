// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a backup plan using a backup plan name and backup rules. A backup plan
// is a document that contains information that Backup uses to schedule tasks that
// create recovery points for resources.
//
// If you call CreateBackupPlan with a plan that already exists, you receive an
// AlreadyExistsException exception.
func (c *Client) CreateBackupPlan(ctx context.Context, params *CreateBackupPlanInput, optFns ...func(*Options)) (*CreateBackupPlanOutput, error) {
	if params == nil {
		params = &CreateBackupPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBackupPlan", params, optFns, c.addOperationCreateBackupPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBackupPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBackupPlanInput struct {

	// Specifies the body of a backup plan. Includes a BackupPlanName and one or more
	// sets of Rules .
	//
	// This member is required.
	BackupPlan *types.BackupPlanInput

	// To help organize your resources, you can assign your own metadata to the
	// resources that you create. Each tag is a key-value pair. The specified tags are
	// assigned to all backups created with this plan.
	BackupPlanTags map[string]string

	// Identifies the request and allows failed requests to be retried without the
	// risk of running the operation twice. If the request includes a CreatorRequestId
	// that matches an existing backup plan, that plan is returned. This parameter is
	// optional.
	//
	// If used, this parameter must contain 1 to 50 alphanumeric or '-_.' characters.
	CreatorRequestId *string

	noSmithyDocumentSerde
}

type CreateBackupPlanOutput struct {

	// A list of BackupOptions settings for a resource type. This option is only
	// available for Windows Volume Shadow Copy Service (VSS) backup jobs.
	AdvancedBackupSettings []types.AdvancedBackupSetting

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50 .
	BackupPlanArn *string

	// Uniquely identifies a backup plan.
	BackupPlanId *string

	// The date and time that a backup plan is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. They cannot be edited.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBackupPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBackupPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBackupPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBackupPlan"); err != nil {
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
	if err = addOpCreateBackupPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackupPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBackupPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBackupPlan",
	}
}
