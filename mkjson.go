package xmltool

// if incoming-part is neither <start> nor <whole/> type, return both nil
func attrinfo(swBrktPart string) (attrs []string, mav map[string]string) {
	mav = make(map[string]string)
	attrstrs := rxAttrPart.FindAllString(swBrktPart, -1)
	for _, attrstr := range attrstrs {
		eqLoc := sIndex(attrstr, "=")
		name := sTrim(attrstr[:eqLoc], " \t")
		attrs = append(attrs, name)

		iVal := 0
	FINDAV:
		for i := eqLoc + 1; i < len(attrstr); i++ {
			switch attrstr[i] {
			case '"', '\'':
				iVal = i
				break FINDAV
			}
		}
		// value := sTrim(attrstr[iVal:], "\"'")
		value := attrstr[iVal:]
		mav[name] = value
	}
	return
}

var (
	contAttrName = "#content"
	attrPrefix   = "@"
	prevLvl      = 0
	onlyCont     = false
)

func cat4json(sb *sBuilder, part string, partType int8, mLvlEle *map[int8]string, stk *stack) *sBuilder {
	ele := ""

	defer func() {
		prevLvl = stk.len()
		switch partType {
		case sBrkt:
			stk.push(ele)
		case eBrkt:
			if top, ok := stk.peek(); ok && top == part[2:len(part)-1] {
				stk.pop()
			}
		}
	}()

	switch partType {
	case sBrkt, wBrkt: // push

		switch buf := sb.String(); buf[len(buf)-1] {
		case '}', '"', 'l': // if this time element is Not the first one, append a Comma to existing buf.
			sb.WriteString(",")
		case ' ': // step into a deeper object immediately
			if stk.len() == prevLvl+1 {
				sb.WriteString("{")
			}
		}

		ele = rxTag.FindString(part)
		ele = ele[1 : len(ele)-1]
		ind := xIndent[stk.len()]
		sb.WriteString("\n")
		sb.WriteString(ind)  // '\t' as const indent
		sb.WriteString("\t") // supplement one '\t' to root indent
		sb.WriteString(fSf(`"%s": `, ele))

		attrs, mav := attrinfo(part)
		for i, attr := range attrs {
			if i == 0 {
				sb.WriteString("{")
			}
			sb.WriteString("\n")
			sb.WriteString(ind)
			sb.WriteString("\t\t") // supplement two '\t' to attribute indent
			sb.WriteString(fSf("\"%s%s\": %s", attrPrefix, sTrimLeft(attr, " \t\r\n"), mav[attr]))
			if i != len(attrs)-1 {
				sb.WriteString(",") // if Not the last attr, append a Comma to existing buf.
			}
		}

		// end up the single whole element
		if partType == wBrkt {
			if len(attrs) > 0 {
				sb.WriteString("\n")
				sb.WriteString(ind)  // '\t' as const indent
				sb.WriteString("\t") // supplement one '\t' to root indent
				sb.WriteString("}")
			} else {
				sb.WriteString("null") // pure empty whole element
			}
		}

	case cText: // push
		part := sTrimRight(part, " \t\n\r")
		// if Not the first position for text content, append a Comma to existing buf.
		switch buf := sb.String(); buf[len(buf)-1] {
		case '"', 'l': // here text is not the first sub, above are attributes subs
			sb.WriteString(",")
			sb.WriteString("\n")
			sb.WriteString(xIndent[stk.len()])
			sb.WriteString("\t")                                      // supplement one '\t' to text content indent
			sb.WriteString(fSf("\"%s\": \"%s\"", contAttrName, part)) // remove tail blank or line-feed
		case ' ': // here text is the first & only sub
			sb.WriteString(fSf("\"%s\"", part))
			onlyCont = true
		}

	case eBrkt: // pop
		if !onlyCont {
			sb.WriteString("\n")
			sb.WriteString(xIndent[stk.len()]) // '\t' as const indent
			sb.WriteString("}")
		}
		onlyCont = false

	case aQuot:
	}
	// sb.WriteString(part) // DEBUG
	return sb
}

// MkJSON :
func MkJSON(xstr, nameOfContAttr, prefixOfAttr string) string {
	// misc.TrackTime(time.Now())

	if nameOfContAttr != "" {
		contAttrName = nameOfContAttr
	}
	if prefixOfAttr != "" {
		attrPrefix = prefixOfAttr
	}

	stk := stack{}
	mLvlEle := make(map[int8]string)

	bLocGrp, types := brktLoc(xstr)
	cLocGrp, types := conTxtLoc(xstr, bLocGrp, types)
	bcLocGrp := locMerge(bLocGrp, cLocGrp) // bracket & content

	sb := &sBuilder{}
	sb.Grow(len(xstr) * 2)
	sb.WriteString("{")
	for _, loc := range bcLocGrp {
		s, e := loc[0], loc[1]
		sb = cat4json(sb, xstr[s:e], types[s], &mLvlEle, &stk)
	}
	sb.WriteString("\n}")

	return sb.String()
}
