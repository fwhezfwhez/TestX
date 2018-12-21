package test_umeng

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func YoumengOpenApiSignature(signParam1 string, signParam2 map[string]string, apiSecurity string) string {

	keyMapSli := make([]string, 0)

	for k := range signParam2{
		keyMapSli = append(keyMapSli, k)
	}

	sort.Strings(keyMapSli)

	param2 := ""
	for _, v := range keyMapSli{
		param2 += v
		param2 += signParam2[v]
	}

	s := signParam1 + param2

	fmt.Println("s:",s)


	mac := hmac.New(sha1.New, []byte(apiSecurity))

	mac.Write([]byte(s))

	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}

