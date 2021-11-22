package yuv4mpeg2

import (
	"reflect"
	"testing"
)

func TestYuv_GenerateHeader(t *testing.T) {
	tests := []struct {
		name    string
		yuv     Yuv
		want    []byte
		wantErr bool
	}{
		{
			name: "good",
			yuv: Yuv{
				Width:          1920,
				Height:         1080,
				FrameRateDen:   1,
				FrameRateNum:   50,
				AspectRatioNum: 1,
				AspectRatioDen: 1,
				InterlacedMode: Progressive,
				ColorMode:      C422,
			},
			want:    []byte("YUV4MPEG2 W1920 H1080 F50:1 A1:1 Ip C422\n"),
			wantErr: false,
		},
		{
			name: "bad colors",
			yuv: Yuv{
				Width:          1920,
				Height:         1080,
				FrameRateDen:   1,
				FrameRateNum:   50,
				AspectRatioNum: 1,
				AspectRatioDen: 1,
				InterlacedMode: Progressive,
				ColorMode:      ColorUnset,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "bad interlaced mode",
			yuv: Yuv{
				Width:          1920,
				Height:         1080,
				FrameRateDen:   1,
				FrameRateNum:   50,
				AspectRatioNum: 1,
				AspectRatioDen: 1,
				InterlacedMode: InterlacedUnset,
				ColorMode:      C422,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.yuv.GenerateHeader()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHeader() got = %v, want %v", got, tt.want)
			}
		})
	}
}
