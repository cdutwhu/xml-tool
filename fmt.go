package xmltool

// func mkIndent(n int) string {
// 	var sb strings.Builder
// 	for i := 0; i < n; i++ {
// 		sb.WriteRune('\t')
// 	}
// 	return "\n" + sb.String()
// }

var (
	xIndent = [33]string{
		"",                                       // 0
		"\t",                                     // 1
		"\t\t",                                   // 2
		"\t\t\t",                                 // 3
		"\t\t\t\t",                               // 4
		"\t\t\t\t\t",                             // 5
		"\t\t\t\t\t\t",                           // 6
		"\t\t\t\t\t\t\t",                         // 7
		"\t\t\t\t\t\t\t\t",                       // 8
		"\t\t\t\t\t\t\t\t\t",                     // 9
		"\t\t\t\t\t\t\t\t\t\t",                   // 10
		"\t\t\t\t\t\t\t\t\t\t\t",                 // 11
		"\t\t\t\t\t\t\t\t\t\t\t\t",               // 12
		"\t\t\t\t\t\t\t\t\t\t\t\t\t",             // 13
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t",           // 14
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",         // 15
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",       // 16
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",     // 17
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",   // 18
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t", // 19
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",                         // 20
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",                       // 21
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",                     // 22
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",                   // 23
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",                 // 24
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",               // 25
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",             // 26
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",           // 27
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",         // 28
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",       // 29
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",     // 30
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t",   // 31
		"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t", // 32
	}
)

const (
	sBrkt = 1 // start angle bracket
	eBrkt = 2 // end angle bracket
	wBrkt = 3 // empty whole angle bracket
	cText = 4 // content text
	aQuot = 5 // attribute double/single quotes
)

func attrValLoc(xstr string, partType map[int]int8) ([][2]int, map[int]int8) {
	locGrp := [][2]int{}
	loc := [2]int{}
	for i := 0; i < len(xstr); i++ {
		c := xstr[i]
		if c == '"' || c == '\'' {
			if loc[0] == 0 {
				loc[0] = i
				// type
				partType[i] = aQuot
				continue
			}
			if loc[0] > 0 {
				loc[1] = i + 1
				locGrp = append(locGrp, loc)
				loc = [2]int{}
			}
		}
	}
	return locGrp, partType
}

func brktLoc(xstr string) (locGrp [][2]int, partType map[int]int8) {
	partType = make(map[int]int8)
	loc := [2]int{}
	for i := 0; i < len(xstr); i++ {
		switch c := xstr[i]; c {
		case '<':
			loc[0] = i
			// type
			if xstr[i+1] == '/' {
				partType[i] = eBrkt
			} else {
				partType[i] = sBrkt
			}
		case '>':
			loc[1] = i + 1
			locGrp = append(locGrp, loc)
			// type
			if xstr[i-1] == '/' {
				partType[loc[0]] = wBrkt
			}
		}
	}
	return
}

func conTxtLoc(xstr string, brktLocGrp [][2]int, partType map[int]int8) ([][2]int, map[int]int8) {
	locGrp := [][2]int{}
	for iBrkt, loc := range brktLocGrp {
		s, e := loc[0], loc[1]
		t := partType[s]
		if t == eBrkt { //
			eTagTxt := xstr[s+2 : e-1]
			pBrkt := brktLocGrp[iBrkt-1]
			ps, pe := pBrkt[0], pBrkt[1]
			pTagTxt := xstr[ps+1 : pe-1]
			if pTagTxt == eTagTxt || sHasPrefix(pTagTxt, eTagTxt+" ") {
				// exclude empty text content, let this pos be end tag pos
				if pe != s { // && sTrim(xstr[pe:s], " \t\r\n") != "" {
					locGrp = append(locGrp, [2]int{pe, s})
					// type
					partType[pe] = cText
				}
			}
		}
	}
	return locGrp, partType
}

func locMerge(locGrp1, locGrp2 [][2]int) (merged [][2]int) {
	lG, lGcmp := locGrp1, locGrp2
	if lG[0][0] > lGcmp[0][0] {
		lG, lGcmp = lGcmp, lG
	}

	cmpIdx := 0
NEXT_LG:
	for _, loc := range lG {
		for _, locCmp := range lGcmp[cmpIdx:] {
			if loc[1] <= locCmp[0] {
				merged = append(merged, loc)
				continue NEXT_LG
			} else {
				merged = append(merged, locCmp)
				cmpIdx++
			}
		}
		merged = append(merged, loc)
	}

	return
}

func cat(sb *sBuilder, part string, partType int8, mLvlEle *map[int8]string, stk *stack) *sBuilder {
	ele := ""
	switch partType {
	case sBrkt, wBrkt: // push
		ele = rxTag.FindString(part)
		ele = ele[1 : len(ele)-1]
		if stk.Len() > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(xIndent[stk.Len()])
		if partType == sBrkt {
			stk.Push(ele)
		}
	case cText: // push
	case eBrkt: // pop
		ele = part[2 : len(part)-1]
		if top, ok := stk.Peek(); ok && top == ele {
			stk.Pop()
		}
	case aQuot:
	}
	sb.WriteString(part)
	return sb
}

// Fmt :
func Fmt(xstr string) string {
	// misc.TrackTime(time.Now())

	stk := stack{}
	mLvlEle := make(map[int8]string)

	bLocGrp, types := brktLoc(xstr)
	cLocGrp, types := conTxtLoc(xstr, bLocGrp, types)
	bcLocGrp := locMerge(bLocGrp, cLocGrp)

	sb := &sBuilder{}
	sb.Grow(len(xstr) * 2)
	for _, loc := range bcLocGrp {
		s, e := loc[0], loc[1]
		sb = cat(sb, xstr[s:e], types[s], &mLvlEle, &stk)
	}

	return sb.String()
}
