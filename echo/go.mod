module demoapp

go 1.17

replace github.com/CoderVlogger/go-web-frameworks/pkg => ../pkg

require (
	github.com/CoderVlogger/go-web-frameworks/pkg v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.6.3
)

require (
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27 // indirect
	golang.org/x/text v0.3.7 // indirect
)
