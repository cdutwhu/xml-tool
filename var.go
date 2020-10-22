package xmltool

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	fPln          = fmt.Println
	fSf           = fmt.Sprintf
	fPf           = fmt.Printf
	fEf           = fmt.Errorf
	sIndex        = strings.Index
	sLastIndex    = strings.LastIndex
	sJoin         = strings.Join
	sSplit        = strings.Split
	sTrimRight    = strings.TrimRight
	sTrimLeft     = strings.TrimLeft
	sTrim         = strings.Trim
	rxMustCompile = regexp.MustCompile
)

var (
	rxTag      = rxMustCompile(`<\w+[\s/>]`)
	rxHead     = rxMustCompile(`<\w+(\s+[\w:]+\s*=\s*"[^"]*"\s*)*\s*/?>`)
	rxAttrPart = rxMustCompile(`\s+[\w:]+\s*=\s*"[^"]*"`)
)
