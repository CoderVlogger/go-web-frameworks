module demoapp

go 1.17

replace github.com/CoderVlogger/go-web-frameworks/pkg => ../pkg

require (
	github.com/CoderVlogger/go-web-frameworks/pkg v0.0.0-20220219210813-22c124f9e92d
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/gofiber/fiber/v2 v2.27.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/klauspost/compress v1.14.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.33.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
)
