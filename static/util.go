package static

import (
	"crypto/md5"
	"encoding/hex"
)

/**
  *2018/11/16
  *author:xiaoC
  *func:生成md5加盐字符串
  *param:string string
 */
func MakeMD5(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}