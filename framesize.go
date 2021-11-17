package yuv4mpeg2

import "errors"

// Size needs to be multiplied with the bit depth of the image.
func (yuv Yuv) Size() (int, error) {
	// frame length = width * height * 3 / 2 (4:2:0)
	if yuv.ColorMode == C420 || yuv.ColorMode == C420jpeg || yuv.ColorMode == C420paldv {
		return yuv.Width * yuv.Height * 3 / 2, nil
	}

	// frame length = width * height * 2 (4:2:2)
	if yuv.ColorMode == C422 {
		return yuv.Width * yuv.Height * 2, nil
	}

	// frame length = width * height * 3 (4:4:4)
	if yuv.ColorMode == C444 {
		return yuv.Width * yuv.Height * 3, nil
	}

	return 0, errors.New("bad color mode")
}