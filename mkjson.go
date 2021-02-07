package xmltool

import (
	"bufio"
	"os"
)

var (
	setSlim      = false
	slim         = false
	contAttrName = "#content"
	attrPrefix   = "@"
	ignoreAttr   = []string{}
	suf4LsEleGrp = []string{}
	nspSep       string
	mNonStrPath  = make(map[string]struct{})
	mRegMapNSP   = make(map[string]map[string]struct{})
	lspSep       string
	mListPath    = make(map[string]struct{})
	mRegMapLP    = make(map[string]map[string]struct{})
)

// SetSlim :
func SetSlim(beSlim bool) {
	slim = beSlim
	setSlim = true
}

// SetContAttrName :
func SetContAttrName(name string) {
	contAttrName = name
}

// SetAttrPrefix :
func SetAttrPrefix(prefix string) {
	attrPrefix = prefix
}

// SetIgnrAttr :
func SetIgnrAttr(attrGrp ...string) {
	ignoreAttr = append(ignoreAttr, attrGrp...)
}

// ---------------------------------------------------------------- //

// SetNonStrPath :
func SetNonStrPath(id string, add bool, sep rune, pathGrp ...string) {
	if mRegMapNSP[id] == nil || !add {
		mRegMapNSP[id] = make(map[string]struct{})
	}
	nspSep = string(sep)
	for _, nsPath := range pathGrp {
		nsPath = sTrimSuffix(nsPath, nspSep)
		mRegMapNSP[id][nsPath] = struct{}{}
		mRegMapNSP[id][nsPath+nspSep+contAttrName] = struct{}{}
	}
}

// EnableNonStrPath :
func EnableNonStrPath(id string) error {
	if m, ok := mRegMapNSP[id]; ok {
		mNonStrPath = m
		return nil
	}
	return fEf("[%s] is not set for NonStrPath", id)
}

// ----------------------------- //

// SetListPath :
func SetListPath(id string, add bool, sep rune, pathGrp ...string) {
	if mRegMapLP[id] == nil || !add {
		mRegMapLP[id] = make(map[string]struct{})
	}
	lspSep = string(sep)
	for _, lsPath := range pathGrp {
		lsPath = sTrimSuffix(lsPath, lspSep)
		mRegMapLP[id][lsPath] = struct{}{}
	}
}

// EnableListPath :
func EnableListPath(id string) error {
	if m, ok := mRegMapLP[id]; ok {
		mListPath = m
		return nil
	}
	return fEf("[%s] is not set for ListPath", id)
}

// SetPathByFile :
func SetPathByFile(pathtype, filepath, id string, add bool, sep rune) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	paths := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		paths = append(paths, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	switch pathtype {
	case "LIST", "List", "list":
		SetListPath(id, add, sep, paths...)
	case "Type", "TYPE", "Non-Str", "NON-STR", "NonStr", "NONSTR", "Bool", "BOOL", "Num", "NUM":
		SetNonStrPath(id, add, sep, paths...)
	default:
		panic(fSf("[%s] is not supported @SetPathByFile", pathtype))
	}
	return nil
}

// SetListPathSuffix :
func SetListPathSuffix(sufGrp ...string) {
	if !setSlim {
		panic("MUST explicitly 'SetSlim' before setting List")
	}

	space := " "
	if slim {
		space = ""
	}
	for i, suf := range sufGrp {
		sufGrp[i] = fSf(`%s":%s`, suf, space)
	}
	suf4LsEleGrp = append(suf4LsEleGrp, sufGrp...)
}

// ---------------------------------------------------------------- //

// if incoming-part is neither <start> nor <whole/> type, return both nil
func attrInfo(swBrktPart string, ignore ...string) (tag string, attrs []string, mav map[string]string) {
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
		// value := sTrim(attrstr[iVal:], "\"'") // keep attribute's quotes, so comment this line out
		value := attrstr[iVal:]
		// if 'single quotes', to "double quotes"
		if sHasPrefix(value, "'") && sHasSuffix(value, "'") {
			value = "\"" + value[1:len(value)-1] + "\""
		}
		mav[name] = value
	}
	// --------------------------------------- //
	for i, c := range swBrktPart {
		switch c {
		case ' ', '>', '/':
			tag = swBrktPart[1:i]
			return
		}
	}
	panic("error@swBrktPart")
}

// .123 => 0.123
func legalizeFloat(float string) string {
	if float[0] != '.' {
		return float
	}
	for _, v := range float[1:] {
		if v < '0' || v > '9' {
			return float
		}
	}
	return "0" + float
}

func legalize(s string) string {
	return legalizeFloat(s)
}

// ---------------------------------------------------------------- //

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
	newLineBlank := func() string {
		return "\n" + xIndent[lvl] + xIndent[len(*mLslvlMark)]
	}
	lastChar := func() byte {
		buf := sb.String()
		return buf[len(buf)-1]
	}
	// tailStr := func(n int) string {
	// 	buf := sb.String()
	// 	if len(buf) < n {
	// 		return buf
	// 	}
	// 	return buf[len(buf)-n:]
	// }
	tracePrt := func(s string, chkslim ...bool) {
		if len(chkslim) > 0 && chkslim[0] == true {
			if slim {
				return
			}
		}
		sb.WriteString(s)
		// if false {
		// 	fPln("---", tailStr(80))
		// }
	}
	space := " "
	if slim {
		space = ""
	}
	var spacecolon byte = ' '
	if slim {
		spacecolon = ':'
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

		eChar := lastChar()
		switch {
		case eChar == '}' || eChar == '"' || eChar == 'l' || eChar == 'e' || (eChar >= '0' && eChar <= '9'): // if this time element is Not the 1st one, append a Comma to buf.
			tracePrt(",")

		case eChar == spacecolon: // step into a sub object(s)

			if _, ok := mListPath[stk.Sprint(lspSep)]; ok { // 1) list-elements bunch checking ...
				stk4lslvl.Push(lvl)
			} else if sHasAnySuffix(sb.String(), suf4LsEleGrp...) { // 2) list-elements bunch checking by suffix ...
				stk4lslvl.Push(lvl)
			}

			// list CPLX element begin, 1st list element
			if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl-1 {
				if _, ok := (*mLslvlMark)[lslvl.(int)]; !ok {
					(*mLslvlMark)[lslvl.(int)] = struct{}{}
					tracePrt("[")
					tracePrt(newLineBlank(), true)
					tracePrt("{")
				}
			} else { // non-list element begin
				tracePrt("{")
			}
		}

		tracePrt(newLineBlank()+"\t", true) // pack indents

		// 2nd... list element content
		if ele == (*mLvlEle)[lvl] {
			tracePrt("{")
		} else { // non-list element begin, prepare for attrs & subs
			tracePrt(fSf(`"%s":%s`, ele, space)) // write "***":
		}

		tag, attrs, mav := attrInfo(part, ignoreAttr...)
		for i, attr := range attrs {
			if i == 0 {
				if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl {
					if _, ok := (*mLslvlMark)[lslvl.(int)]; !ok {
						(*mLslvlMark)[lslvl.(int)] = struct{}{}
						tracePrt("[")
						tracePrt(newLineBlank()+"\t", true)
						tracePrt("{")
					}
				} else {
					tracePrt("{")
				}
			}

			str2type := false
			if _, ok := mNonStrPath[sTrimLeft(stk.Sprint(nspSep)+nspSep+tag+nspSep+attr, nspSep)]; ok {
				str2type = true
			}

			tracePrt(newLineBlank()+"\t\t", true) // supplement two '\t' to attribute indent
			val := mav[attr]
			if str2type {
				val = sTrim(mav[attr], "\"")
				val = legalize(val)
			}
			tracePrt(fSf("\"%s%s\":%s%s", attrPrefix, sTrimLeft(attr, " \t\r\n"), space, val))

			if i != len(attrs)-1 {
				tracePrt(",") // if Not the last attr, append a Comma to existing buf.
			}
		}

		// end up the single whole element
		if partType == wBrkt {
			if len(attrs) > 0 {
				tracePrt(newLineBlank()+"\t", true) // supplement one '\t' to root indent
				tracePrt("}")
			} else {
				tracePrt("null") // pure empty whole element (type <ele />)
			}
		}

	case cText: // push
		part = xmlTxtEscCharProc(part)

		str2type := false
		if _, ok := mNonStrPath[stk.Sprint(nspSep)]; ok {
			str2type = true
		}

		// if Not the first position for text content, append a Comma to existing buf.
		eChar := lastChar()
		switch {
		case eChar == '"' || eChar == 'l' || eChar == 'e' || (eChar >= '0' && eChar <= '9'): // here text is not the first sub, above are attributes subs
			tracePrt(",")

			val := part
			if !str2type {
				val = fSf("\"%s\"", part)
			} else {
				val = legalize(val)
			}

			if *plainList {
				tracePrt(newLineBlank(), true)
				tracePrt(val)
			} else {
				tracePrt(newLineBlank()+"\t", true)                    // supplement one '\t' to text content indent
				tracePrt(fSf("\"%s\":%s%s", contAttrName, space, val)) // remove tail blank or line-feed
			}

		case eChar == spacecolon: // here text is the first & only sub
			*singleCont = true

			if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl-1 {
				if _, ok := (*mLslvlMark)[lslvl.(int)]; !ok {
					(*mLslvlMark)[lslvl.(int)] = struct{}{}
					tracePrt("[")
					tracePrt(newLineBlank(), true)
					*plainList = true
				}
			}

			val := part
			if !str2type {
				val = fSf("\"%s\"", part)
			} else {
				val = legalize(val)
			}
			tracePrt(val)
		}

	case eBrkt: // pop

		// end list-element bunch
		if lslvl, ok := stk4lslvl.Peek(); ok && lslvl == lvl {
			if _, ok := (*mLslvlMark)[lslvl.(int)]; ok {
				tracePrt(newLineBlank(), true)
				tracePrt("]")
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
			if lastChar() == spacecolon {
				tracePrt("null") // empty element (type <ele></ele>)
			} else {
				tracePrt(newLineBlank(), true)
				tracePrt("}")
			}
		}
		*singleCont = false
	}
	return sb
}

// MkJSON :
func MkJSON(xmlstr string) string {

	stk, stk4lslvl := stack{}, stack{}
	mLvlEle := make(map[int]string)
	mLslvlMark := make(map[int]struct{})

	bLocGrp, types := brktLoc(xmlstr)
	cLocGrp, types := conTxtLoc(xmlstr, bLocGrp, types)
	bcLocGrp := locMerge(bLocGrp, cLocGrp) // bracket & content

	singleCont, plainList := false, false
	sb := &sBuilder{}
	sb.Grow(len(xmlstr) * 2)
	sb.WriteString("{")
	for _, loc := range bcLocGrp {
		s, e := loc[0], loc[1]
		sb = cat4json(sb, xmlstr[s:e], types[s], &mLvlEle, &stk, &singleCont, &plainList, &stk4lslvl, &mLslvlMark)
	}
	if !slim {
		sb.WriteString("\n")
	}
	sb.WriteString("}")

	return sb.String()
}

// --------------------------------- //

func xmlTxtEscCharProc(txt string) string {
	sb4TxtEscChar := sBuilder{}
	for _, c := range txt {
		switch c {
		case '"':
			sb4TxtEscChar.WriteString("\\\"")
			continue
		case '\n':
			sb4TxtEscChar.WriteString("\\n")
			continue
		case '\t':
			sb4TxtEscChar.WriteString("\\t")
			continue
		default:
			sb4TxtEscChar.WriteRune(c)
		}
	}
	return sTrim(sb4TxtEscChar.String(), " \t\n\r")
}
