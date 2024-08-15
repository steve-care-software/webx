package bitcoins

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/tyler-smith/go-bip39"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// MasterKey creates a master key
func (app *application) MasterKey(words []string) (*hdkeychain.ExtendedKey, error) {
	// mnemonic creates the mnemonic
	mnemonic := strings.Join(words, " ")

	// Generate a seed from the mnemonic phrase
	seed := bip39.NewSeed(mnemonic, "")

	// Generate a master private key from the seed
	return hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
}

// Transact creates a bitcoin transaction
func (app *application) Transact(
	masterKey *hdkeychain.ExtendedKey,
	utxoHash string,
	utxoIndex uint32,
	toAddress string,
	amountInSatoshi int64,
	customData []byte,
) (*wire.MsgTx, error) {
	// Derive child key from master key using the derivation path
	childKey, err := deriveKeyFromPath(masterKey, derivationPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the private key
	privateKey, err := childKey.ECPrivKey()
	if err != nil {
		log.Fatal(err)
	}

	// create the amount:
	amount := btcutil.Amount(amountInSatoshi)

	// Create a new transaction
	tx := wire.NewMsgTx(wire.TxVersion)

	// Add the UTXO as input
	hash, err := chainhash.NewHashFromStr(utxoHash)
	if err != nil {
		return nil, err
	}

	outPoint := wire.NewOutPoint(hash, utxoIndex)
	txIn := wire.NewTxIn(outPoint, nil, nil)
	tx.AddTxIn(txIn)

	// Add the recipient output
	address, err := btcutil.DecodeAddress(toAddress, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	pkScript, err := txscript.PayToAddrScript(address)
	if err != nil {
		return nil, err
	}

	txOut := wire.NewTxOut(int64(amount), pkScript)
	tx.AddTxOut(txOut)

	// Create an OP_RETURN output with custom bytes, if any:
	if len(customData) > 0 {
		opReturnScript, err := txscript.NewScriptBuilder().
			AddOp(txscript.OP_RETURN).
			AddData(customData).
			Script()

		if err != nil {
			return nil, err
		}

		// Value is 0 as we are just embedding data
		opReturnOutput := wire.NewTxOut(0, opReturnScript)
		tx.AddTxOut(opReturnOutput)
	}

	// Sign the transaction
	sigScript, err := txscript.SignatureScript(tx, 0, pkScript, txscript.SigHashAll, privateKey, true)
	if err != nil {
		return nil, err
	}

	txIn.SignatureScript = sigScript
	return tx, nil
}

// Broadcast broadcasts a bitcoin transaction
func (app *application) Broadcast(
	trx *wire.MsgTx,
	rpcUser string,
	rpcPassword string,
	rpcPort string,
	rpcHost string,
) (*chainhash.Hash, error) {
	// serialize the signed trx:
	buf := new(bytes.Buffer)
	if err := trx.Serialize(buf); err != nil {
		str := fmt.Sprintf("Could not serialize the input transaction: %v", err)
		return nil, errors.New(str)
	}

	// raw trx:
	rawTxHex := hex.EncodeToString(buf.Bytes())

	// Create the RPC request
	rpcRequest := RpcRequest{
		Jsonrpc: "1.0",
		Method:  "sendrawtransaction",
		Params:  []interface{}{rawTxHex},
		Id:      1,
	}

	requestBody, err := json.Marshal(rpcRequest)
	if err != nil {
		log.Fatalf("Error marshaling RPC request: %v", err)
	}

	// Create the HTTP request
	url := fmt.Sprintf("http://%s:%s@%s:%s/", rpcUser, rpcPassword, rpcHost, rpcPort)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		str := fmt.Sprintf("Error creating HTTP request: %v", err)
		return nil, errors.New(str)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		str := fmt.Sprintf("Error sending HTTP request: %v", err)
		return nil, errors.New(str)
	}
	defer resp.Body.Close()

	// Parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		str := fmt.Sprintf("Error reading HTTP response: %v", err)
		return nil, errors.New(str)
	}

	var rpcResponse RpcResponse
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		str := fmt.Sprintf("Error unmarshaling RPC response: %v", err)
		return nil, errors.New(str)
	}

	if rpcResponse.Error != nil {
		str := fmt.Sprintf("RPC Error: %s", rpcResponse.Error.Message)
		return nil, errors.New(str)
	}

	// Parse the transaction ID into a chainhash.Hash object
	txHash, err := chainhash.NewHashFromStr(rpcResponse.Result)
	if err != nil {
		str := fmt.Sprintf("Failed to parse transaction ID: %v", err)
		return nil, errors.New(str)
	}

	return txHash, nil
}
