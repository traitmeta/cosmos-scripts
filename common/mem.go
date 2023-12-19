package common

import (
	"encoding/hex"
	"fmt"
	"log"

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

func GetPrivKeyFromMnemonic() {

	seed := bip39.NewSeed("blast about old claw current first paste risk involve victory edit current", "")
	fmt.Println("Seed: ", hex.EncodeToString(seed)) // Seed:  dd5ffa7088c0fa4c665085bca7096a61e42ba92e7243a8ad7fbc6975a4aeea1845c6b668ebacd024fd2ca215c6cd510be7a9815528016af3a5e6f47d1cca30dd

	master, ch := hd.ComputeMastersFromSeed(seed)
	path := "m/44'/118'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Derivation Path: ", path)                 // Derivation Path:  m/44'/118'/0'/0/0'
	fmt.Println("Private Key: ", hex.EncodeToString(priv)) // Private Key:  69668f2378b43009b16b5c6eb5e405d9224ca2a326a65a17919e567105fa4e5a

	_, ecPub := btcec.PrivKeyFromBytes(priv[:])
	pubkeyBytes := ecPub.SerializeCompressed()
	fmt.Println("Public Key: ", hex.EncodeToString(pubkeyBytes)) // Public Key:  03de79435cbc8a799efc24cdce7d3b180fb014d5f19949fb8d61de3f21b9f6c1f8

	decodeString, err := hex.DecodeString(fmt.Sprintf("04%x", pubkeyBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Convert test data to base32:
	conv, err := bech32.ConvertBits(decodeString, 8, 5, true)
	if err != nil {
		fmt.Println("Error:", err)
	}
	encoded, err := bech32.Encode("atom", conv)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show the encoded data.
	fmt.Println("Wallet Address:", encoded) // Wallet Address: atom1qspau72rtj7g57v7lsjvmnna8vvqlvq56hcejj0m34sau0eph8mvr7qgl9avu
}
