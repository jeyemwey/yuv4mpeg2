package yuv4mpeg2

import (
	"errors"
	"fmt"
)

func (y Yuv) GenerateHeader() ([]byte, error) {
	interlacedMode, err := interlacedMode(y.InterlacedMode)
	if err != nil {
		return nil, err
	}

	color, err := colorMode(y.ColorMode)
	if err != nil {
		return nil, err
	}

	header := []byte(fmt.Sprintf("%sW%d H%d F%d:%d A%d:%d %s %s\n", Prefix, y.Width, y.Height, y.FrameRateNum, y.FrameRateDen, y.AspectRatioNum, y.AspectRatioDen, interlacedMode, color))

	return header, nil
}

func interlacedMode(mode InterlacedMode) (string, error) {
	switch mode {
	case Progressive:
		return "Ip", nil
	case TopFieldFirst:
		return "It", nil
	case BottomFieldFirst:
		return "Ib", nil
	case MixedMode:
		return "Im", nil
	default:
		return "", errors.New("bad interlaced mode")
	}
}

func colorMode(mode ColorMode) (string, error) {
	switch mode {
	case C420jpeg:
		return "C420jpeg", nil
	case C420paldv:
		return "C420paldv", nil
	case C420:
		return "C420", nil
	case C422:
		return "C422", nil
	case C444:
		return "C444", nil
	default:
		return "", errors.New("bad color mode")
	}
}
