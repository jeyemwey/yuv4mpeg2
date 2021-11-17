package yuv4mpeg2

import "testing"

func TestYuv_Size(t *testing.T) {
	tests := []struct {
		name      string
		yuvHeader string
		size      int
		wantErr   bool
	}{
		{name: "1080p 4:4:4", yuvHeader: "YUV4MPEG2 W1920 H1080 C444", size: 6220800, wantErr: false},
		{name: "1080p 4:2:0", yuvHeader: "YUV4MPEG2 W1920 H1080 C420", size: 3110400, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yuv, err := parseHeader(tt.yuvHeader)
			if err != nil {
				t.Errorf("yuvHeader is bad")
				return
			}

			got, err := yuv.Size()
			if (err != nil) != tt.wantErr {
				t.Errorf("Size() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.size {
				t.Errorf("Size() got = %v, size %v", got, tt.size)
			}
		})
	}
}
