package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"jsonfmt/decode"
	"jsonfmt/indent"
	"log"
	"os"
	"regexp"
)

const READBYTES int = 1024
const JSONP_RE string = "^([\n]?[A-Za-z_0-9.]+[(]{1})(.*)([)]|[)][\n]+)$"

func main() {

	var (
		head bytes.Buffer
		body *bytes.Buffer
		tail bytes.Buffer
	)

	var opts struct {
		Sort bool `short:"s" long:"sort" description:"Sort keys alphabetically"`
	}

	args, _ := flags.Parse(&opts)

	// Parse args.
	if len(args) < 1 {
		fmt.Println("Usage: jsonfmt [file]")
		os.Exit(1)
	}
	filename := args[0]

	body = loadFile(filename)

	// Try parsing JSONP.
	if parts, err := ParseJSONP(body.Bytes()); err == nil {
		head.Write(parts[0])
		body.Reset()
		body.Write(parts[1])
		tail.Write(parts[2])
	}

	// Make a new buffer of indented JSON.
	// TODO: need to initialize like this?
	indentedBody := bytes.NewBufferString("")
	i, err := decode.RawInterfaceMap(body.Bytes())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	indent.Indent(indentedBody, i, "    ", opts.Sort)

	// Write the buffer into the same file.
	fo, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	fo.Write(head.Bytes())
	fo.Write(indentedBody.Bytes())
	fo.Write(tail.Bytes())
	fo.Close()
}

func ParseJSONP(contents []byte) ([][]byte, error) {
	re, _ := regexp.Compile(JSONP_RE)
	matches := re.FindAllSubmatch(contents, -1)
	if len(matches) == 0 {
		return nil, errors.New("Could not parse into JSONP")
	}
	parts := matches[0]
	if len(parts) < 3 {
		return nil, errors.New("Could not parse into JSONP")
	}
	return parts[1:], nil
}

func loadFile(filename string) *bytes.Buffer {
	fi, err := os.Open(filename)
	if err != nil {
        fmt.Println(err)
		os.Exit(1)
	}
    var buf bytes.Buffer
	data := make([]byte, 1024)
	for {
		n, err := fi.Read(data)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		buf.Write(data[:n])
	}
	fi.Close()
    return &buf
}
