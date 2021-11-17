# yuv4mpeg2

[![go-docs](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/jeyemwey/yuv4mpeg2)

Parsing headers of `yuv4mpeg2` files.

```go
file, _ := os.Open(filepath)	
yuv, err := yuv4mpeg2.ParseHeader(file)
size, err := yuv.Size()

fmt.Printf("width=%d height=%d size=%d", yuv.Width, yuv.Height, size)
```

## Further readings

* [YUV4MPEG2 in Multimedia Wiki](https://wiki.multimedia.cx/index.php/YUV4MPEG2)