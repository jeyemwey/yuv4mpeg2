package yuv4mpeg2

import (
	"reflect"
	"testing"
)

func Test_frac(t *testing.T) {
	type args struct {
		keyword string
		label   string
	}
	tests := []struct {
		name        string
		args        args
		numerator   int
		denominator int
		wantErr     bool
	}{
		{
			name: "Unknown frame rate",
			args: args{
				keyword: "F0:0",
				label:   "framerate",
			},
			numerator: 0, denominator: 0,
			wantErr: false,
		},
		{
			name: "Square",
			args: args{
				keyword: "A1:1",
				label:   "aspect ratio",
			},
			numerator: 1, denominator: 1,
			wantErr: false,
		},
		{
			name: "Four by three",
			args: args{
				keyword: "A4:3",
				label:   "aspect ratio",
			},
			numerator: 4, denominator: 3,
			wantErr: false,
		},
		{
			name: "Missing Colon",
			args: args{
				keyword: "F1_1",
				label:   "framerate",
			},
			numerator: 0, denominator: 0,
			wantErr: true,
		},
		{
			name: "Bad number of colons",
			args: args{
				keyword: "F1:1:1",
				label:   "framerate",
			},
			numerator: 0, denominator: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumerator, gotDenominator, err := frac(tt.args.keyword, tt.args.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("frac() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNumerator != tt.numerator {
				t.Errorf("frac() gotNumerator = %v, numerator %v", gotNumerator, tt.numerator)
			}
			if gotDenominator != tt.denominator {
				t.Errorf("frac() gotDenominator = %v, numerator %v", gotDenominator, tt.denominator)
			}
		})
	}
}

func Test_parseHeader(t *testing.T) {
	tests := []struct {
		name      string
		headerStr string
		wantYuv   Yuv
		wantErr   bool
	}{
		{name: "Bad start", headerStr: "XYZ_NOT_WHAT_I_WANT", wantErr: true, wantYuv: Yuv{}},
		{name: "Interlaced mode set multiple times", headerStr: "YUV4MPEG2 W1920 H1080 F50:1 Ip Ip", wantErr: true, wantYuv: Yuv{}},
		{name: "Good one with induced defaults", headerStr: "YUV4MPEG2 W1920 H1080 F50:1 Ip", wantErr: false, wantYuv: Yuv{
			Width:          1920,
			Height:         1080,
			FrameRateNum:   50,
			FrameRateDen:   1,
			AspectRatioNum: 1,
			AspectRatioDen: 1,
			InterlacedMode: Progressive,
			ColorMode:      C444,
			HeaderSize:     30,
		}},
		{name: "Good one", headerStr: "YUV4MPEG2 W1920 H1080 F50:1 Ip C420 A4:3", wantErr: false, wantYuv: Yuv{
			Width:          1920,
			Height:         1080,
			FrameRateNum:   50,
			FrameRateDen:   1,
			AspectRatioNum: 4,
			AspectRatioDen: 3,
			InterlacedMode: Progressive,
			ColorMode:      C420,
			HeaderSize:     40,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYuv, err := parseHeader(tt.headerStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotYuv, tt.wantYuv) {
				t.Errorf("parseHeader() gotYuv = %v, size %v", gotYuv, tt.wantYuv)
			}
		})
	}
}
