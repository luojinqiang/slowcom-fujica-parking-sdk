package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gdes"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"sort"
	"strings"
)

func GetSign(data map[string]interface{}) (encryptSign string) {
	param := JoinStringsInASCII(data)
	fmt.Println("排序后参数：", param)
	partnerKey := "4B180E05FFF1CF42EB73D07AEE1E0BE61EBE7D4481634E5BC568D8683660ED0E"
	sign := CreateSign(partnerKey, param)
	encryptSign = Encrypt(partnerKey, sign)
	return
}

func CreateSign(partnerKey string, paramStr string) string {
	hash := hmac.New(sha256.New, []byte(partnerKey))
	hash.Write([]byte(paramStr))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func Encrypt(partnerKey string, sign string) string {
	partnerKey = partnerKey[0:8]
	bytes, err := gdes.EncryptCBC([]byte(sign), []byte(partnerKey), []byte("12345678"), 1)
	fmt.Println("err", err)
	return string(gbase64.Encode(bytes))
}

// JoinStringsInASCII 按照规则，参数名ASCII码从小到大排序后拼接
func JoinStringsInASCII(data map[string]interface{}) string {
	var keyValList []string
	var keyList []string
	for k := range data {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		keyValList = append(keyValList, fmt.Sprintf("%s=%v", k, data[k]))
	}
	return strings.Join(keyValList, "")
}

func BuildUrl(baseUrl string, url string) string {
	return fmt.Sprintf("%s%s", baseUrl, url)
}
