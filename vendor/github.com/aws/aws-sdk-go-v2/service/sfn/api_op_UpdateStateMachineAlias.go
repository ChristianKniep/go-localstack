// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the configuration of an existing state machine [alias] by modifying its
// description or routingConfiguration .
//
// You must specify at least one of the description or routingConfiguration
// parameters to update a state machine alias.
//
// UpdateStateMachineAlias is an idempotent API. Step Functions bases the
// idempotency check on the stateMachineAliasArn , description , and
// routingConfiguration parameters. Requests with the same parameters return an
// idempotent response.
//
// This operation is eventually consistent. All StartExecution requests made within a few
// seconds use the latest alias configuration. Executions started immediately after
// calling UpdateStateMachineAlias may use the previous routing configuration.
//
// Related operations:
//
// # CreateStateMachineAlias
//
// # DescribeStateMachineAlias
//
// # ListStateMachineAliases
//
// # DeleteStateMachineAlias
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
func (c *Client) UpdateStateMachineAlias(ctx context.Context, params *UpdateStateMachineAliasInput, optFns ...func(*Options)) (*UpdateStateMachineAliasOutput, error) {
	if params == nil {
		params = &UpdateStateMachineAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStateMachineAlias", params, optFns, c.addOperationUpdateStateMachineAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStateMachineAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStateMachineAliasInput struct {

	// The Amazon Resource Name (ARN) of the state machine alias.
	//
	// This member is required.
	StateMachineAliasArn *string

	// A description of the state machine alias.
	Description *string

	// The routing configuration of the state machine alias.
	//
	// An array of RoutingConfig objects that specifies up to two state machine
	// versions that the alias starts executions for.
	RoutingConfiguration []types.RoutingConfigurationListItem

	noSmithyDocumentSerde
}

type UpdateStateMachineAliasOutput struct {

	// The date and time the state machine alias was updated.
	//
	// This member is required.
	UpdateDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStateMachineAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateStateMachineAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateStateMachineAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateStateMachineAlias"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateStateMachineAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStateMachineAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStateMachineAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateStateMachineAlias",
	}
}