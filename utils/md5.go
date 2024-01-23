package utils

import (
	"crypto/md5"
	_ "crypto/md5"
	"fmt"
	_ "fmt"
	"strings"
	_ "strings"
)

func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)

	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
