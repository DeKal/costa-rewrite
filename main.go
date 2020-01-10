package main

import (
	ModeCompare "github.com/DeKal/costa-rewrite/mode/compare"
	ModeNormal "github.com/DeKal/costa-rewrite/mode/normal"
	Parser "github.com/DeKal/costa-rewrite/parser"
)

func main() {
	params := Parser.ParseCommandLineParams()
	if Parser.IsNormalMode(params.Mode) {
		ModeNormal.RunRewriteAndProduceReports(params)
	} else {
		ModeCompare.Compare2Report(params)
	}
}
