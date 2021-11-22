package yuv4mpeg2

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseHeader(reader io.Reader) (Yuv, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(scanFrames)

	if scanner.Scan() {
		headerStr := scanner.Text()

		return parseHeader(headerStr)
	} else {
		return Yuv{}, errors.New("unable to find FRAME start")
	}
}

func parseHeader(headerStr string) (yuv Yuv, err error) {
	if !strings.HasPrefix(headerStr, string(Prefix)) {
		return Yuv{}, errors.New("missing YUV4MPEG2 introduction")
	}

	yuv.HeaderSize = len(headerStr)

	for _, keyword := range strings.Split(headerStr, " ") {
		keyword = strings.TrimSpace(keyword)

		switch keyword[0:1] {
		case "W": // Width
			if yuv.Width != 0 {
				return Yuv{}, errors.New("width was already set")
			}

			w, err := strconv.Atoi(keyword[1:])
			if err != nil {
				return Yuv{}, errors.New("bad width")
			}
			yuv.Width = w

		case "H": // Height
			if yuv.Height != 0 {
				return Yuv{}, errors.New("height was already set")
			}

			h, err := strconv.Atoi(keyword[1:])
			if err != nil {
				return Yuv{}, errors.New("bad height")
			}
			yuv.Height = h

		case "F": // Frame rate
			if yuv.FrameRateNum != 0 || yuv.FrameRateDen != 0 {
				return Yuv{}, errors.New("frame rate was already set")
			}

			numerator, denominator, err := frac(keyword, "framerate")
			if err != nil {
				return Yuv{}, err
			}

			yuv.FrameRateNum = numerator
			yuv.FrameRateDen = denominator

		case "A": // Aspect Ratio
			if yuv.AspectRatioNum != 0 || yuv.AspectRatioDen != 0 {
				return Yuv{}, errors.New("aspect was already set")
			}

			numerator, denominator, err := frac(keyword, "aspect ratio")
			if err != nil {
				return Yuv{}, err
			}

			yuv.AspectRatioNum = numerator
			yuv.AspectRatioDen = denominator

		case "I": // Interlaced Mode
			if yuv.InterlacedMode != InterlacedUnset {
				return Yuv{}, errors.New("interlace mode was already set")
			}

			switch keyword {
			case "Ip":
				yuv.InterlacedMode = Progressive
			case "It":
				yuv.InterlacedMode = TopFieldFirst
			case "Ib":
				yuv.InterlacedMode = BottomFieldFirst
			case "Im":
				yuv.InterlacedMode = MixedMode
			}

		case "C": // Color Mode
			if yuv.ColorMode != ColorUnset {
				return Yuv{}, errors.New("color mode was already set")
			}

			switch keyword {
			case "C420jpeg":
				yuv.ColorMode = C420jpeg
			case "C420paldv":
				yuv.ColorMode = C420paldv
			case "C420":
				yuv.ColorMode = C420
			case "C422":
				yuv.ColorMode = C422
			case "C444":
				yuv.ColorMode = C444
			}
		}
	}

	// Let's set some defaults if those are not set previously
	if yuv.ColorMode == ColorUnset {
		yuv.ColorMode = C444
	}

	if yuv.AspectRatioNum == 0 || yuv.AspectRatioDen == 0 {
		yuv.AspectRatioNum = 1
		yuv.AspectRatioDen = 1
	}

	return
}

func frac(keyword string, label string) (int, int, error) {
	frac := strings.Split(keyword[1:], ":")
	if len(frac) != 2 {
		return 0, 0, errors.New(fmt.Sprintf("bad %s fraction", label))
	}

	numerator, err := strconv.Atoi(frac[0])
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("bad %s numerator", label))
	}

	denominator, err := strconv.Atoi(frac[1])
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("bad %s denominator", label))
	}
	return numerator, denominator, nil
}

func scanFrames(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, FrameStart); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
