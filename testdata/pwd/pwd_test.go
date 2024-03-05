package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}

func TestCheckPwd(t *testing.T) {
	hash := HashPwd("123456")
	fmt.Println(CheckPwd(hash, "123456"))
}
