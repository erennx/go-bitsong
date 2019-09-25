package track

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultCodespace codespace for the module
	DefaultCodespace sdk.CodespaceType = ModuleName

	// CodeInvalidGenesis error code
	CodeInvalidGenesis sdk.CodeType = 101

	// CodeEmptyInput error code
	CodeEmptyInput sdk.CodeType = 102

	// CodeSongNotExist error code
	CodeSongNotExist sdk.CodeType = 103

	// CodePlayNotExist error code
	CodePlayNotExist sdk.CodeType = 104
)

// ErrInvalidGenesis Error constructor
func ErrInvalidGenesis(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidGenesis, fmt.Sprintf("InitialTrackID never set."))
}

// ErrInitialTrackIDAlreadySet Error constructor
func ErrInitialTrackIDAlreadySet(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidGenesis, fmt.Sprintf("InitialTrackID already set."))
}

// ErrEmptyInput Error constructor
func ErrEmptyInput(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeEmptyInput, fmt.Sprintf("Input must not be empty."))
}

// ErrSongNotExist Error constructor
func ErrSongNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeEmptyInput, fmt.Sprintf("Track not exist."))
}
