package baseha

import (
	"unicode/utf8"
)

// nolint
const (
	encodeStd   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789蛤丝们不要搞个大新闻苟利国家生死以岂因祸福避趋之黄沙百战穿金甲按照香港基本法爱慕安格瑞闷声发财命运靠自我奋斗也考虑历史行程光阴似箭识得唔华莱士螳臂当车鸭嘴笔总书记微小工作很惭愧续一秒走火入膜病中惊坐起二院视察三代表五可奉告六月水柜八门外语荷叶蟾宣传报道有偏差猴蛙另请高明玩具鸡黑框眼镜扬州江边赛艇念句诗夏威夷云腾致雨露结为霜玉出崑冈剑号巨阙珠称夜果珍李柰菜重芥姜容止若思言笃初诚美慎终宜令同"
	StdEncoding = encodeStd
)

// nolint
var STD = NewEncoding(StdEncoding)

// nolint
type Encoding struct {
	encode    []rune
	decodeMap map[int32]byte
}

// nolint
func NewEncoding(encoder string) *Encoding {
	e := new(Encoding)
	e.encode = []rune(encoder)

	if len(e.encode) < 256 {
		panic(len(e.encode))
	}

	i := 0
	e.decodeMap = map[int32]byte{}
	for _, v := range e.encode {
		e.decodeMap[v] = byte(i)
		i++
	}
	return e
}

// nolint
func (e *Encoding) EncodedLen(n int) int {
	return n
}

// nolint
func (e *Encoding) DecodedLen(src string) int {
	return utf8.RuneCountInString(src)
}

// nolint
func (e *Encoding) Encode(src []byte) []byte {
	dst := make([]rune, e.EncodedLen(len(src)))
	encMap := e.encode

	for i, v := range src {
		dst[i] = encMap[v]
	}

	return []byte(string(dst))
}

// nolint
func (e *Encoding) EncodeToString(src []byte) string {
	dst := make([]rune, e.EncodedLen(len(src)))
	encMap := e.encode

	for i, v := range src {
		dst[i] = encMap[v]
	}

	return string(dst)
}

// nolint
func (e *Encoding) Decode(srcStr string) []byte {
	dstSize := e.DecodedLen(srcStr)
	var dst = make([]byte, dstSize)

	var i = 0
	for _, v := range srcStr {
		dst[i] = e.decodeMap[v]
		i++
	}

	return dst
}
