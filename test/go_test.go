package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	var pwd = "admin123"

	sum := md5.Sum([]byte(pwd))
	md5strPwd := hex.EncodeToString(sum[:])

	fmt.Println(md5strPwd)

}
