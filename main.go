package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
)
func main {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	
	md := []byte("markdown text")
	html := markdown.ToHTML(md, nil, renderer)
}
