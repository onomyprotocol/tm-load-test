package types

import (
	"github.com/onomyprotocol/tm-load-test/customclient/codec"
	"github.com/onomyprotocol/tm-load-test/customclient/codec/types"
	cryptocodec "github.com/onomyprotocol/tm-load-test/customclient/crypto/codec"
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	"github.com/onomyprotocol/tm-load-test/customclient/types/msgservice"
	"github.com/onomyprotocol/tm-load-test/customclient/x/authz"
)

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSend{}, "cosmos-sdk/MsgSend", nil)
	cdc.RegisterConcrete(&MsgMultiSend{}, "cosmos-sdk/MsgMultiSend", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
	)
	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&SendAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/bank module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
