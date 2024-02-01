// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Places a request for a new game session in a queue. When processing a placement
// request, Amazon GameLift searches for available resources on the queue's
// destinations, scanning each until it finds resources or the placement request
// times out. A game session placement request can also request player sessions.
// When a new game session is successfully created, Amazon GameLift creates a
// player session for each player included in the request. When placing a game
// session, by default Amazon GameLift tries each fleet in the order they are
// listed in the queue configuration. Ideally, a queue's destinations are listed in
// preference order. Alternatively, when requesting a game session with players,
// you can also provide latency data for each player in relevant Regions. Latency
// data indicates the performance lag a player experiences when connected to a
// fleet in the Region. Amazon GameLift uses latency data to reorder the list of
// destinations to place the game session in a Region with minimal lag. If latency
// data is provided for multiple players, Amazon GameLift calculates each Region's
// average lag for all players and reorders to get the best game play across all
// players. To place a new game session request, specify the following:
//   - The queue name and a set of game session properties and settings
//   - A unique ID (such as a UUID) for the placement. You use this ID to track
//     the status of the placement request
//   - (Optional) A set of player data and a unique player ID for each player that
//     you are joining to the new game session (player data is optional, but if you
//     include it, you must also provide a unique ID for each player)
//   - Latency data for all players (if you want to optimize game play for the
//     players)
//
// If successful, a new game session placement is created. To track the status of
// a placement request, call DescribeGameSessionPlacement (https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeGameSessionPlacement.html)
// and check the request's status. If the status is FULFILLED , a new game session
// has been created and a game session ARN and Region are referenced. If the
// placement request times out, you can resubmit the request or retry it with a
// different queue.
func (c *Client) StartGameSessionPlacement(ctx context.Context, params *StartGameSessionPlacementInput, optFns ...func(*Options)) (*StartGameSessionPlacementOutput, error) {
	if params == nil {
		params = &StartGameSessionPlacementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartGameSessionPlacement", params, optFns, c.addOperationStartGameSessionPlacementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartGameSessionPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartGameSessionPlacementInput struct {

	// Name of the queue to use to place the new game session. You can use either the
	// queue name or ARN value.
	//
	// This member is required.
	GameSessionQueueName *string

	// The maximum number of players that can be connected simultaneously to the game
	// session.
	//
	// This member is required.
	MaximumPlayerSessionCount *int32

	// A unique identifier to assign to the new game session placement. This value is
	// developer-defined. The value must be unique across all Regions and cannot be
	// reused.
	//
	// This member is required.
	PlacementId *string

	// Set of information on each player to create a player session for.
	DesiredPlayerSessions []types.DesiredPlayerSession

	// A set of key-value pairs that can store custom data in a game session. For
	// example: {"Key": "difficulty", "Value": "novice"} .
	GameProperties []types.GameProperty

	// A set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process in the GameSession object with a
	// request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)
	// ).
	GameSessionData *string

	// A descriptive label that is associated with a game session. Session names do
	// not need to be unique.
	GameSessionName *string

	// A set of values, expressed in milliseconds, that indicates the amount of
	// latency that a player experiences when connected to Amazon Web Services Regions.
	// This information is used to try to place the new game session where it can offer
	// the best possible gameplay experience for the players.
	PlayerLatencies []types.PlayerLatency

	noSmithyDocumentSerde
}

type StartGameSessionPlacementOutput struct {

	// Object that describes the newly created game session placement. This object
	// includes all the information provided in the request, as well as start/end time
	// stamps and placement status.
	GameSessionPlacement *types.GameSessionPlacement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartGameSessionPlacementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartGameSessionPlacement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartGameSessionPlacement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartGameSessionPlacement"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpStartGameSessionPlacementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartGameSessionPlacement(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opStartGameSessionPlacement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartGameSessionPlacement",
	}
}
