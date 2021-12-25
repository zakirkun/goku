module goku

go 1.17

// Direct - A direct dependency is a dependency which the module directly imports.
require (
	github.com/gofiber/fiber/v2 v2.23.0
	github.com/joho/godotenv v1.4.0
	github.com/spf13/cobra v1.3.0
)

// Indirect – It is the dependency that is imported by the module’s direct dependencies. Also,
// any dependency that is mentioned in the go.mod file but not imported in any of the source files
// of the module is also treated as an indirect dependency.
require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.31.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)
