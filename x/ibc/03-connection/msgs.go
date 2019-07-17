package connection

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/ibc/23-commitment"
)

type MsgOpenInit struct {
	ConnectionID       string
	Connection         Connection
	CounterpartyClient string
	NextTimeout        uint64
	Signer             sdk.AccAddress
}

var _ sdk.Msg = MsgOpenInit{}

func (msg MsgOpenInit) Route() string {
	return "ibc"
}

func (msg MsgOpenInit) Type() string {
	return "open-init"
}

func (msg MsgOpenInit) ValidateBasic() sdk.Error {
	return nil // TODO
}

func (msg MsgOpenInit) GetSignBytes() []byte {
	return nil // TODO
}

func (msg MsgOpenInit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

type MsgOpenTry struct {
	ConnectionID       string
	Connection         Connection
	CounterpartyClient string
	Timeout            uint64
	NextTimeout        uint64
	Proofs             []commitment.Proof
	Signer             sdk.AccAddress
}

var _ sdk.Msg = MsgOpenTry{}

func (msg MsgOpenTry) Route() string {
	return "ibc"
}

func (msg MsgOpenTry) Type() string {
	return "open-init"
}

func (msg MsgOpenTry) ValidateBasic() sdk.Error {
	return nil // TODO
}

func (msg MsgOpenTry) GetSignBytes() []byte {
	return nil // TODO
}

func (msg MsgOpenTry) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

type MsgOpenAck struct {
	ConnectionID string
	Timeout      uint64
	NextTimeout  uint64
	Proofs       []commitment.Proof
	Signer       sdk.AccAddress
}

var _ sdk.Msg = MsgOpenAck{}

func (msg MsgOpenAck) Route() string {
	return "ibc"
}

func (msg MsgOpenAck) Type() string {
	return "open-init"
}

func (msg MsgOpenAck) ValidateBasic() sdk.Error {
	return nil // TODO
}

func (msg MsgOpenAck) GetSignBytes() []byte {
	return nil // TODO
}

func (msg MsgOpenAck) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

type MsgOpenConfirm struct {
	ConnectionID string
	Timeout      uint64
	Proofs       []commitment.Proof
	Signer       sdk.AccAddress
}

var _ sdk.Msg = MsgOpenConfirm{}

func (msg MsgOpenConfirm) Route() string {
	return "ibc"
}

func (msg MsgOpenConfirm) Type() string {
	return "open-init"
}

func (msg MsgOpenConfirm) ValidateBasic() sdk.Error {
	return nil // TODO
}

func (msg MsgOpenConfirm) GetSignBytes() []byte {
	return nil // TODO
}

func (msg MsgOpenConfirm) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}