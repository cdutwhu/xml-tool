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

func cat4json(sb *sBuilder, part string, partType int8, mLvlEle *map[int8]string, stk *stack) *sBuilder {
	ele := ""
	switch partType {
	case sBrkt, wBrkt: // push
		ele = rxTag.FindString(part)
		ele = ele[1 : len(ele)-1]
		ind := xIndent[stk.len()]
		sb.WriteString("\n")
		sb.WriteString(ind)  // '\t' as const indent
		sb.WriteString("\t") // supplement one '\t' to root indent
		sb.WriteString(fSf(`"%s": {`, ele))

		attrs, mav := attrinfo(part)
		for _, attr := range attrs {
			sb.WriteString("\n")
			sb.WriteString(ind)
			sb.WriteString("\t\t") // supplement two '\t' to attribute indent
			sb.WriteString(fSf("\"@%s\": %s,", attr, mav[attr]))
		}

		if partType == wBrkt {
			sb.WriteString("\n")
			sb.WriteString(ind)  // '\t' as const indent
			sb.WriteString("\t") // supplement one '\t' to root indent
			sb.WriteString("},")
		}

		// ----------------------------- //
		if partType == sBrkt {
			stk.push(ele)
		}

	case cText: // push
		ind := xIndent[stk.len()]
		sb.WriteString("\n")
		sb.WriteString(ind)
		sb.WriteString("\t")               // supplement one '\t' to text content indent
		txt := sTrimRight(part, " \t\n\r") // remove tail blank or line-feed
		sb.WriteString(fSf("\"#content\": \"%s\"", txt))

	case eBrkt: // pop
		ind := xIndent[stk.len()]
		sb.WriteString("\n")
		sb.WriteString(ind) // '\t' as const indent
		sb.WriteString("},")
		// ----------------------------- //
		ele = part[2 : len(part)-1]
		if top, ok := stk.peek(); ok && top == ele {
			stk.pop()
		}

	case aQuot:
	}
	// sb.WriteString(part)
	return sb
}

// MkJSON :
func MkJSON(xstr string) string {
	// misc.TrackTime(time.Now())

	stk := stack{}
	mLvlEle := make(map[int8]string)

	bLocGrp, types := brktLoc(xstr)
	cLocGrp, types := conTxtLoc(xstr, bLocGrp, types)
	bcLocGrp := locMerge(bLocGrp, cLocGrp) // bracket & content

	sb := &sBuilder{}
	sb.WriteString("{")
	for _, loc := range bcLocGrp {
		s, e := loc[0], loc[1]
		sb = cat4json(sb, xstr[s:e], types[s], &mLvlEle, &stk)
	}
	sb.WriteString("\n}")

	jstr := sb.String()

	// -- remove last or redundant comma -- //
	jstr = rxExtComma.ReplaceAllStringFunc(jstr, func(m string) string {
		return sReplace(m, ",", "", 1)
	})

	// -- reform no-attributes & text content only element -- //
	jstr = rxContNoAttr.ReplaceAllStringFunc(jstr, func(m string) string {
		m = sTrimRight(m, "}")
		txt := sSplit(m, `"#content":`)[1]
		txt = sTrim(txt, " \t\n\r")
		return txt
	})

	return jstr
}
