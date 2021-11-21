package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"githun.com/heuristicwave/buildingBlockchain/utils"
)

const (
	signature     string = "998b1664d8728ea6d30735619621be430045fdc8d51ad9a367ce0b53456026ed919c7eed20fba9c7936facac5221df78eadc1a56ecd0079a337ff7ff16d68dae"
	privateKey    string = "30770201010420c20e276559d369f54ea3166c7b720bb28b84386993d55f555730af14a4a532faa00a06082a8648ce3d030107a144034200042db60a2051617d3f417a63c6b330c9bcbfeef42be2ece421f8ee6931eb2cc6152b26b78e86121d7092b8772b0bed90c4e77d34b25ad3eca5d41e31e02029b040"
	hashedMessage string = "ac84723f3a759cbef710c5937dce777ac1b17d492b62a988cf701666efd4495b"
)

func Start() {
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
	// utils.HandleErr(err)

	hashBytes, err := hex.DecodeString(hashedMessage)

	utils.HandleErr(err)

	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)

	fmt.Println(ok)
}
