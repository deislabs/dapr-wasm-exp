# ptr

<p align="center">
    <a href="https://travis-ci.org/agrea/ptr"><img src="https://travis-ci.org/agrea/ptr.svg?branch=master"></a>
    <a href="https://godoc.org/github.com/agrea/ptr"><img src="https://img.shields.io/badge/godoc-documentation-blue.svg"></a>
    <a href="https://codeclimate.com/github/agrea/ptr/maintainability"><img src="https://api.codeclimate.com/v1/badges/6220f7d67a8a2332b7df/maintainability" /></a>
</p>

The `ptr` package gives you some basic helpers for working with pointers in Go.
The package is simply intended to make it easy to create pointers to things.
E.g. instead of writing:

    s := "some string"
    b := &s

You'd just write

    b := ptr.String("some string")

## Usage

    ptr.Bool(true)        // Returns *bool
    ptr.Byte(byte('a'))   // Returns *byte
    ptr.Float32(123.3)    // Returns *float32
    ptr.Float64(123.3)    // Returns *float64
    ptr.Int(123)          // Returns *int
    ptr.Int8(123)         // Returns *int8
    ptr.Int16(123)        // Returns *int16
    ptr.Int32(123)        // Returns *int32
    ptr.Int64(123)        // Returns *int64
    ptr.Rune(123)         // Returns *rune
    ptr.String("string")  // Returns *string
    ptr.Uint(123)         // Returns *uint
    ptr.Uint8(123)        // Returns *uint8
    ptr.Uint16(123)       // Returns *uint16
    ptr.Uint32(123)       // Returns *uint32
    ptr.Uint64(123)       // Returns *uint64

## Running the tests

    go test -v ./...

## License

MIT
