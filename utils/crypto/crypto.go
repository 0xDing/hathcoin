package crypto

import (
	"encoding/hex"

	"github.com/borisding1994/hathcoin/utils/crypto/sm3"
)

// SM3Hash can generate the SM3 checksum for string.
// 使用天朝密码管理局钦定的 SM3 密码杂凑算法。 这个效率，efficiency!!!
// 非常熟悉朝鲜的这一套理论————国产自主可控是坠吼的
func SM3Hash(msg string) string {
	蛤稀值 := sm3.Sum([]byte(msg))
	return hex.EncodeToString(蛤稀值[:])
}
