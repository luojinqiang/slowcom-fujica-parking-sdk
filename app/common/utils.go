package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gdes"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"strconv"
	"time"
)

func GetSign() (encryptSign string, timestamp string) {
	timestamp = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	fmt.Println("timestamp=", timestamp)
	paramStr := "appId=" + AppId + "appSecret=" + AppSecret + "mchId=" + MchId + "timestamp=" + timestamp
	fmt.Println("排序后参数=", paramStr)
	sign := CreateSign(PartnerKey, paramStr)
	encryptSign = Encrypt(PartnerKey, sign)
	fmt.Println("签名sign=", encryptSign)
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

func BuildUrl(url string) string {
	return fmt.Sprintf("%s%s", BaseUrl, url)
}
