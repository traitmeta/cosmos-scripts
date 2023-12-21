package common

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/go-bip39"
)

func getPrivKeyFromMnemonic(mnemonic string) ([]byte, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}

	masterpriv, ch := hd.ComputeMastersFromSeed(seed)
	path := "m/44'/118'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(masterpriv, ch, path)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func getSeedFromMnemonic(mnemonic string) (string, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(seed), nil
}

// func getAddressFromPrivKey(privKey secp256k1.PrivKey) (sdk.AccAddress, error) {
// 	return sdk.AccAddress(privKey.PubKey().Address()), nil
// }

func GetPubKeyFromPriv(priv []byte) []byte {
	_, ecPub := btcec.PrivKeyFromBytes(priv[:])
	pubkeyBytes := ecPub.SerializeCompressed()

	return pubkeyBytes
}

func GetAddrFromPriv(priv []byte) (string, error) {
	pubkeyBytes := GetPubKeyFromPriv(priv)
	decodeString, err := hex.DecodeString(fmt.Sprintf("04%x", pubkeyBytes))
	if err != nil {
		return "", err
	}

	// Convert test data to base32:
	conv, err := bech32.ConvertBits(decodeString, 8, 5, true)
	if err != nil {
		return "", err
	}
	encoded, err := bech32.Encode("atom", conv)
	if err != nil {
		return "", err
	}

	return encoded, nil
}
