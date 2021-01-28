package xmltool

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestValStrIdx(t *testing.T) {
	misc.TrackTime(time.Now())
	bytes, _ := ioutil.ReadFile("./examples/StudentPersonal.xml")
	xmlstr := string(bytes)
	if !IsValid(xmlstr) {
		panic("error")
	}
	fPln(Fmt(xmlstr))

	// bLocGrp, types := brktLoc(xmlstr)
	// fPln(bLocGrp)
	// // fPln(types)
	// fPln(" -------------------- ")
	// cLocGrp, types := conTxtLoc(xmlstr, bLocGrp, types)
	// fPln(cLocGrp)

	// for _, loc := range cLocGrp {
	// 	s, e := loc[0], loc[1]
	// 	fPln(xmlstr[s:e])
	// }

	// merged := locMerge(bLocGrp, cLocGrp)
	// fPln(merged)

	// for _, loc := range merged {
	// 	fPln(xmlstr[loc[0]:loc[1]])
	// }

	// aLocGrp, types := attrValLoc(xmlstr, types)
	// fPln(aLocGrp)
	// fPln(types)

}
