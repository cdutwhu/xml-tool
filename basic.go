package xmltool

import "encoding/xml"

// IsValid :
func IsValid(xmlstr string) bool {
	return xml.Unmarshal([]byte(xmlstr), new(interface{})) == nil
}

// Root :
func Root(xml string) (tag string) {
	xml = sTrim(xml, " \t\n\r")
	if loc := rxHead.FindStringIndex(xml); loc != nil {
		s, e := loc[0], loc[1]
		sTag := xml[s:e]
		// fPln("sTag:", sTag)
		tag = rxTag.FindString(sTag)
		tag = sTrim(tag, "< \t\r\n>/")
		// fPln("tag:", tag)
	}
	return
}

// TagAttrVal :
func TagAttrVal(xml string) (tag string, attrs []string, mAttrVal map[string]string) {
	xml = sTrim(xml, " \t\n\r")
	if pair := rxHead.FindAllStringIndex(xml, 1); pair != nil {
		s, e := pair[0][0], pair[0][1]
		sTag := xml[s:e]
		// fPln("sTag:", sTag)

		tag = rxTag.FindAllString(sTag, 1)[0]
		tag = sTrim(tag, "< \t\r\n>/")
		// fPln("tag:", tag)

		mAttrVal = make(map[string]string)
		for _, av := range rxAttrPart.FindAllString(sTag, -1) {
			av = sTrim(av, " \t\r\n")
			ss := sSplit(av, "=")
			a, v := sTrim(ss[0], " \t\r\n"), sTrim(sTrim(ss[1], " \t\r\n"), "\"")
			attrs = append(attrs, a)
			mAttrVal[a] = v
		}
	}
	return
}

// Lvl0 : xml should be '<name>***</name>'
func Lvl0(xml string) (string, string, string) {
	sTag, name, eTag := "", "", ""
	end1, end2 := 0, 0

	if loc := rxHead.FindStringIndex(xml); loc != nil {
		s, e := loc[0], loc[1]
		sTag = xml[s:e]
		// fPln("sTag", sTag)
		end1 = e

		if loc := rxTag.FindStringIndex(sTag); loc != nil {
			s, e := loc[0], loc[1]
			name = sTrimRight(sTag[s+1:e], " \t/>")
			// fPln("name", name)
			eTag = fSf(`</%s>`, name)
			// fPln("eTag", eTag)
			end2 = sLastIndex(xml, eTag)
		}
	}

	cont := xml[end1:end2]
	cont = sTrimLeft(cont, "\n\r")
	cont = sTrimRight(cont, " \t\n\r")
	return name, sJoin([]string{sTag, eTag}, "\n"), cont
}

// Cont :
func Cont(xml string) string {
	_, _, cont := Lvl0(xml)
	return cont
}
