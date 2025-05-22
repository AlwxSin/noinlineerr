package main

import (
	"github.com/AlwxSin/noinlineerr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(noinlineerr.Analyzer)
}
