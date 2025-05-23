package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/AlwxSin/noinlineerr"
)

func main() {
	singlechecker.Main(noinlineerr.NewAnalyzer())
}
