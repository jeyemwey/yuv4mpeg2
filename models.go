package yuv4mpeg2

type Yuv struct {
	Width          int
	Height         int
	FrameRateNum   int
	FrameRateDen   int
	AspectRatioNum int
	AspectRatioDen int

	InterlacedMode InterlacedMode
	ColorMode      ColorMode
	HeaderSize     int
}

// InterlacedMode type hint
type InterlacedMode int

const (
	InterlacedUnset InterlacedMode = iota

	Progressive      InterlacedMode = iota
	TopFieldFirst    InterlacedMode = iota
	BottomFieldFirst InterlacedMode = iota
	MixedMode        InterlacedMode = iota
)

// ColorMode type hint
type ColorMode int

const (
	ColorUnset ColorMode = iota

	C420jpeg  ColorMode = iota // 4:2:0 with biaxially-displaced chroma planes
	C420paldv ColorMode = iota // 4:2:0 with vertically-displaced chroma planes
	C420      ColorMode = iota // 4:2:0 with coincident chroma planes
	C422      ColorMode = iota // 4:2:2
	C444      ColorMode = iota // 4:4:4
)

var FrameStart []byte
var Prefix []byte

func init() {
	FrameStart = []byte("FRAME")
	Prefix = []byte("YUV4MPEG2 ")
}
