package login

import (
	"fmt"
	"strings"
	"testing"
)

var sss = `{"message":"janction.io wants you to sign in with your Ethereum account:\n0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266\n\nid00test00linux00amd64\n\n\nURI: https://janction.io\nVersion: 1\nChain ID: 1\nNonce: WYirkJelHtKPqtWN\nIssued At: 2024-07-20T11:07:17Z","signature":"0x4728e0ece9bc29504208f4509dbe6c8d6d3894f28734c389e3b4a8316c56d3ae4a87e46128c2c9648f60fbf24a96b4cf4e8b9c6cc191907b88ff76383551c8e71b","is_node":true}`

func TestMessage(*testing.T) {
	fmt.Println(sss)
	ss2 := strings.Replace(sss, "\\nURI", "URI", 1)
	fmt.Println(ss2)
}
