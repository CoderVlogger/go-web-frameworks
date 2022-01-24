module demoapp

go 1.17

replace github.com/CoderVlogger/go-web-frameworks/pkg => ../pkg

require (
	github.com/CoderVlogger/go-web-frameworks/pkg v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.6.3
	github.com/mattn/go-colorable v0.1.12 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/net v0.0.0-20220121210141-e204ce36a2ba // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
)

