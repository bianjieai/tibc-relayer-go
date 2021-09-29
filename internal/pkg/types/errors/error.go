package errors

import "fmt"

const RootCodeSpace = "relayer"

var (
	ErrInternal            = Register(RootCodeSpace, 1, "internal")
	ErrChainConn           = Register(RootCodeSpace, 2, "connection chain failed")
	ErrGetLightClientState = Register(RootCodeSpace, 3, "failed to get light client state")
	ErrGetBlockHeader      = Register(RootCodeSpace, 4, "failed to get block header")
	ErrUpdateClient        = Register(RootCodeSpace, 5, "failed to update client")
	ErrGetPackets          = Register(RootCodeSpace, 6, "failed to get packets")
	ErrGetCommitmentPacket = Register(RootCodeSpace, 7, "failed to get commitment packet")
	ErrGetAckPacket        = Register(RootCodeSpace, 8, "failed to get ack packet")
	ErrGetReceiptPacket    = Register(RootCodeSpace, 9, "failed to get receipt packet")
	ErrGetProof            = Register(RootCodeSpace, 10, "failed to get proof")
	ErrGetLatestHeight     = Register(RootCodeSpace, 11, "failed to get latest height")
	ErrRecvPacket          = Register(RootCodeSpace, 12, "failed to recv packet")
	ErrNotProduced         = Register(RootCodeSpace, 13, "failed to not produced")
	ErrDelayTime           = Register(RootCodeSpace, 14, "failed to get delay time")
	ErrDelayHeight         = Register(RootCodeSpace, 15, "failed to get delay height")
	ErrCurBlockTime        = Register(RootCodeSpace, 16, "failed to get current block time")
	ErrUnknownMsg          = Register(RootCodeSpace, 17, "failed to unknown msg type")
)

var usedCodes = map[string]*Error{}

func getUsed(codespace string, code uint32) *Error {
	return usedCodes[errorID(codespace, code)]
}

func setUsed(err *Error) {
	usedCodes[errorID(err.codeSpace, err.code)] = err
}

func errorID(codespace string, code uint32) string {
	return fmt.Sprintf("%s:%d", codespace, code)
}

type IError interface {
	error
	Code() uint32
	CodeSpace() string
}

type Error struct {
	codeSpace string
	code      uint32
	desc      string
}

func New(codeSpace string, code uint32, desc string) *Error {
	return &Error{codeSpace: codeSpace, code: code, desc: desc}
}

func (e Error) Error() string {
	return e.desc
}

func (e Error) Code() uint32 {
	return e.code
}

func (e Error) CodeSpace() string {
	return e.codeSpace
}

func Register(codespace string, code uint32, description string) *Error {
	if e := getUsed(codespace, code); e != nil {
		panic(fmt.Sprintf("error with code %d is already registered: %q", code, e.desc))
	}

	err := New(codespace, code, description)
	setUsed(err)

	return err
}
