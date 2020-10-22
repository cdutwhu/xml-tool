package xmltool

// BreakCont :
func BreakCont(xml string) (names, subs []string) {
	remain := xml

AGAIN:
	if loc := rxTag.FindStringIndex(remain); loc != nil {
		s, e := loc[0], loc[1]
		name := remain[s+1 : e-1]
		names = append(names, name)

		remain = remain[s:] // from first '<tag>'
		// fPln("remain:", remain)

		end1, end2 := -1, -1
		if loc := rxMustCompile(fSf(`</%s\s*>`, name)).FindStringIndex(remain); loc != nil {
			_, end1 = loc[0], loc[1] // update e to '</tag>' end
			// fPln("end:", remain[s:end1]) // end tag
		}
		if i := sIndex(remain, "/>"); i >= 0 {
			end2 = i + 2 // update e to '/>' end
		}

		// if '/>' is found, and before '</tag>', and this part is valid XML
		switch {
		case end1 >= 0 && end2 < 0:
			e = end1
		case end1 < 0 && end2 >= 0:
			e = end2
		case end1 >= 0 && end2 >= 0:
			if end2 < end1 && IsValid(remain[:end2]) {
				e = end2
			} else {
				e = end1
			}
		default:
			panic("invalid sub xml")
		}

		sub := remain[:e]
		// fPln("sub:", sub)
		subs = append(subs, sub)

		remain = remain[e:] // from end of first '</tag>' or '/>'
		goto AGAIN
	}
	return
}
