package bitcoins

import (
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

const derivationPath = "m/44'/0'/0'/0/0"

// NewApplication creates a new application
func NewApplication() Application {
	return createApplication()
}

// Application represents the bitcoin application
type Application interface {
	MasterKey(
		words []string,
	) (*hdkeychain.ExtendedKey, error) // generate a bitcoin master key
	Transact(
		masterKey *hdkeychain.ExtendedKey,
		utxoHash string,
		utxoIndex uint32,
		toAddress string,
		amountInSatoshi int64,
		customData []byte,
	) (*wire.MsgTx, error) // creates a bitcoin transaction
	Broadcast(
		trx *wire.MsgTx,
		rpcUser string,
		rpcPassword string,
		rpcPort string,
		rpcHost string,
	) (*chainhash.Hash, error)
}
