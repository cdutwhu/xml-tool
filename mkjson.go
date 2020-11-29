package xmltool

var (
	contAttrName = "#content"
	attrPrefix   = "@"
	ignoreAttr   = []string{}
	suf4LsEleGrp = []string{}
)

// SetContAttrName :
func SetContAttrName(name string) {
	if name != "" {
		contAttrName = name
	}
}

// SetAttrPrefix :
func SetAttrPrefix(prefix string) {
	if prefix != "" {
		attrPrefix = prefix
	}
}

// SetIgnrAttr :
func SetIgnrAttr(attrGrp ...string) {
	ignoreAttr = append(ignoreAttr, attrGrp...)
}

// SetSuf4LsEle :
func SetSuf4LsEle(sufGrp ...string) {
	suf4LsEleGrp = append(suf4LsEleGrp, sufGrp...)
}

// if incoming-part is neither <start> nor <whole/> type, return both nil
func attrInfo(swBrktPart string, ignore ...string) (attrs []string, mav map[string]string) {
	mav = make(map[string]string)
NEXT:
	for _, attrstr := range rxAttrPart.FindAllString(swBrktPart, -1) {
		eqLoc := sIndex(attrstr, "=")
		name := sTrim(attrstr[:eqLoc], " \t\r\n")
		for _, ign := range ignore {
			if ign == name {
				continue NEXT
			}
		}
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

// var listFmt = true

func cat4json(
	sb *sBuilder,
	part string,
	partType int8,
	mLvlEle *map[int]string,
	stk *stack,
	singleCont *bool,
	plainList *bool,
	stk4lslvl *stack,
	mLslvlMark *map[int]struct{},
) *sBuilder {

	ele := ""
	lvl := stk.Len()
	ind := xIndent[lvl]
	appInd := func() string {
		return xIndent[len(*mLslvlMark)]
	}
	lastChar := func() byte {
		buf := sb.String()
		return buf[len(buf)-1]
	}
	tailStr := func(n int) string {
		buf := sb.String()
		if len(buf) < n {
			return buf
		}
		return buf[len(buf)-n:]
	}
	tracePrt := func(s string) {
		sb.WriteString(s)
		if false {
			fPln("---", tailStr(80))
		}
	}

	// ------------------------------ //

	defer func() {
		switch partType {
		case sBrkt:
			(*mLvlEle)[lvl] = ele // root is lvl0
			stk.Push(ele)
		case eBrkt:
			(*mLvlEle)[lvl] = ""
			if top, ok := stk.Peek(); ok && top == part[2:len(part)-1] {
				stk.Pop()
			}
		}
	}()

	// ------------------------------ //

	switch partType {
	case sBrkt, wBrkt: // push

		ele = rxTag.FindString(part)
		ele = ele[1 : len(ele)-1]

		if *plainList {
			if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl {
				break
			}
		}

		switch lastChar() {
		case '}', '"', 'l': // if this time element is Not the 1st one, append a Comma to buf.
			tracePrt(",")

		case ' ': // step into a sub object(s)

			// list-elements bunch checking ...
			if buf := sb.String(); sHasAnySuffix(buf, suf4LsEleGrp...) {
				stk4lslvl.Push(lvl)
			}

			// list CPLX element begin, 1st list element
			if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl-1 {
				if _, ok := (*mLslvlMark)[lslvl.(int)]; !ok {
					(*mLslvlMark)[lslvl.(int)] = struct{}{}
					tracePrt("[\n" + ind + appInd() + "{")
				}
			} else { // non-list element begin
				tracePrt("{")
			}
		}

		// pack indents
		tracePrt("\n" + ind + appInd() + "\t")

		// 2nd... list element content
		if ele == (*mLvlEle)[lvl] {
			tracePrt("{")
		} else { // non-list element begin, prepare for attrs & subs
			tracePrt(fSf(`"%s": `, ele)) // write "***":
		}

		attrs, mav := attrInfo(part, ignoreAttr...)
		for i, attr := range attrs {
			if i == 0 {
				if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl {
					if _, ok := (*mLslvlMark)[lslvl.(int)]; !ok {
						(*mLslvlMark)[lslvl.(int)] = struct{}{}
						tracePrt("[\n" + ind + appInd() + "\t{")
					}
				} else {
					tracePrt("{")
				}
			}

			tracePrt("\n" + ind + appInd() + "\t\t") // supplement two '\t' to attribute indent
			tracePrt(fSf("\"%s%s\": %s", attrPrefix, sTrimLeft(attr, " \t\r\n"), mav[attr]))

			if i != len(attrs)-1 {
				tracePrt(",") // if Not the last attr, append a Comma to existing buf.
			}
		}

		// end up the single whole element
		if partType == wBrkt {
			if len(attrs) > 0 {
				tracePrt("\n" + ind + appInd() + "\t}") // supplement one '\t' to root indent
			} else {
				tracePrt("null") // pure empty whole element (type <ele />)
			}
		}

	case cText: // push
		part = sTrim(part, " \t\n\r")
		part = sReplaceAll(part, "\"", "\\\"")
		part = sReplaceAll(part, "\n", "\\n")

		// if Not the first position for text content, append a Comma to existing buf.
		switch lastChar() {
		case '"', 'l': // here text is not the first sub, above are attributes subs
			tracePrt(",")

			if *plainList {
				tracePrt("\n" + ind + appInd()) // supplement one '\t' to text content indent
				tracePrt(fSf("\"%s\"", part))
			} else {
				tracePrt("\n" + ind + appInd() + "\t")              // supplement one '\t' to text content indent
				tracePrt(fSf("\"%s\": \"%s\"", contAttrName, part)) // remove tail blank or line-feed
			}

		case ' ': // here text is the first & only sub
			*singleCont = true

			if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl-1 {
				if _, ok := (*mLslvlMark)[lslvl.(int)]; !ok {
					(*mLslvlMark)[lslvl.(int)] = struct{}{}
					tracePrt("[\n" + ind + appInd())
					*plainList = true
				}
			}

			tracePrt(fSf("\"%s\"", part))
		}

	case eBrkt: // pop

		// end list-element bunch
		if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl {
			if _, ok := (*mLslvlMark)[lslvl.(int)]; ok {
				tracePrt("\n" + ind + appInd() + "]")
				delete(*mLslvlMark, lslvl.(int))
				stk4lslvl.Pop()
				*singleCont = false
				*plainList = false
			}
		}

		if *plainList {
			break
		}

		// singleCont does NOT need '}' at jumping out
		if !*singleCont {
			if lastChar() == ' ' {
				tracePrt("null") // empty element (type <ele></ele>)
			} else {
				tracePrt("\n" + ind + appInd() + "}")
			}
		}
		*singleCont = false
	}
	return sb
}

// MkJSON :
func MkJSON(xstr string) string {

	stk, stk4lslvl := stack{}, stack{}
	mLvlEle := make(map[int]string)
	mLslvlMark := make(map[int]struct{})

	bLocGrp, types := brktLoc(xstr)
	cLocGrp, types := conTxtLoc(xstr, bLocGrp, types)
	bcLocGrp := locMerge(bLocGrp, cLocGrp) // bracket & content

	sb := &sBuilder{}
	singleCont, plainList := false, false

	sb.Grow(len(xstr) * 2)
	sb.WriteString("{")
	for _, loc := range bcLocGrp {
		s, e := loc[0], loc[1]
		sb = cat4json(sb, xstr[s:e], types[s], &mLvlEle, &stk, &singleCont, &plainList, &stk4lslvl, &mLslvlMark)
	}
	sb.WriteString("\n}")

	return sb.String()
}
