package xmltool

import (
	"fmt"
	"regexp"
	"strings"

	stk "github.com/cdutwhu/gotil/stack"
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
	sTrimSuffix   = strings.TrimSuffix
	sTrim         = strings.Trim
	sHasPrefix    = strings.HasPrefix
	sHasSuffix    = strings.HasSuffix
	sContains     = strings.Contains
	sReplace      = strings.Replace
	sReplaceAll   = strings.ReplaceAll
	rxMustCompile = regexp.MustCompile

	sHasAnySuffix = func(s string, suffix ...string) bool {
		for _, suf := range suffix {
			if sHasSuffix(s, suf) {
				return true
			}
		}
		return false
	}
)

type (
	sBuilder = strings.Builder
	stack    = stk.Stack
)

var (
	rxTag        = rxMustCompile(`<\w+[\s/>]`)
	rxHead       = rxMustCompile(`<\w+(\s+[\w:]+\s*=\s*"[^"]*"\s*)*\s*/?>`)
	rxAttrPart   = rxMustCompile(`\s+[\w:]+\s*=\s*(("[^"]*")|('[^']*'))`)
	rxLF1more    = rxMustCompile(`\n+`)
	rxSpace2more = rxMustCompile(` {2,}`)
)
