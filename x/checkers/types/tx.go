package types

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogogrpc "github.com/gogo/protobuf/grpc"
	"regexp"
)

// baseAppSimulateFn is the signature of the Baseapp#Simulate function.
type baseAppSimulateFn func(txBytes []byte) (sdk.GasInfo, *sdk.Result, error)

// msgServer is the server for the protobuf Msg service.
type msgServer struct {
	clientCtx         client.Context
	simulate          baseAppSimulateFn
	interfaceRegistry codectypes.InterfaceRegistry
}

func (m msgServer) CreateGame(ctx context.Context, game *MsgCreateGame) (*MsgCreateGameResponse, error) {
	//TODO implement me
	//panic("implement me")
	return nil, nil
}

func (m msgServer) PlayMove(ctx context.Context, move *MsgPlayMove) (*MsgPlayMoveResponse, error) {
	//TODO implement me
	//panic("implement me")
	return nil, nil
}

// NewMsgServer creates a new Msg service server.
func NewMsgServer(clientCtx client.Context, simulate baseAppSimulateFn, interfaceRegistry codectypes.InterfaceRegistry) MsgServer {
	return msgServer{
		clientCtx:         clientCtx,
		simulate:          simulate,
		interfaceRegistry: interfaceRegistry,
	}
}

var (
	_ MsgServer = msgServer{}

	// EventRegex checks that an event string is formatted with {alphabetic}.{alphabetic}={value}
	EventRegex = regexp.MustCompile(`^[a-zA-Z_]+\.[a-zA-Z_]+=\S+$`)
)

const (
	eventFormat = "{eventType}.{eventAttribute}={value}"
)

// RegisterMsgService registers the tx service on the gRPC router.
func RegisterMsgService(
	qrt gogogrpc.Server,
	clientCtx client.Context,
	simulateFn baseAppSimulateFn,
	interfaceRegistry codectypes.InterfaceRegistry,
) {
	RegisterMsgServer(
		qrt,
		NewMsgServer(clientCtx, simulateFn, interfaceRegistry),
	)
}
