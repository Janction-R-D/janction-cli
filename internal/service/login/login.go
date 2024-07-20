package login

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spruceid/siwe-go"
	"jct/common/config"
	"jct/types"
	"jct/utils"
	"strings"
	"time"
)

func FetchNonce() (*types.NonceRes, error) {
	url := config.TestnetUrl + "/api/v1/auth/nonce"
	resp, err := utils.GetWithTimeout(url, nil, nil, 30*time.Second)
	if err != nil {
		return nil, err
	}
	var nonceResp types.NonceRes
	err = json.Unmarshal(resp, &nonceResp)
	if err != nil {
		return nil, err
	}
	return &nonceResp, nil
}

// TODO handle token expire
func Login(nonce string, serialNumber string) (*types.LoginResp, error) {
	fmt.Println("[serialNumber] ", serialNumber)
	url := config.TestnetUrl + "/api/v1/auth/login"
	privateKey, err := crypto.HexToECDSA(config.PrivateKey[2:])
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	message, err := siwe.InitMessage(
		"janction.io",
		address,
		"https://janction.io",
		nonce,
		map[string]interface{}{
			"statement": serialNumber,
		})
	if err != nil {
		return nil, err
	}
	prepare := message.String()
	// LINUX BUG
	prepare = strings.Replace(prepare, "\\nURI", "URI", 1)
	//messageHash := crypto.Keccak256Hash([]byte(prepare))
	//signature, err := crypto.Sign(messageHash.Bytes(), privateKey)
	fmt.Println("---------msg----------")
	fmt.Println(prepare)
	signature, err := signMessage(prepare, privateKey)
	if err != nil {
		return nil, err
	}
	reqLogin := types.LoginReq{
		Message:   prepare,
		Signature: hexutil.Encode(signature),
		IsNode:    true,
	}
	body, err := json.Marshal(reqLogin)
	if err != nil {
		return nil, err
	}
	fmt.Println("---------req----------")
	fmt.Println(string(body))
	resp, err := utils.PostWithTimeout(url, body, nil, 30*time.Second)
	if err != nil {
		return nil, err
	}
	var loginResp types.LoginResp
	fmt.Println("---------resp----------")
	fmt.Println(string(resp))
	err = json.Unmarshal(resp, &loginResp)
	if err != nil {
		return nil, err
	}
	return &loginResp, nil
}

func signHash(data []byte) common.Hash {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}

func signMessage(message string, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	sign := signHash([]byte(message))
	signature, err := crypto.Sign(sign.Bytes(), privateKey)

	if err != nil {
		return nil, err
	}

	signature[64] += 27
	return signature, nil
}
