package xmltool

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestValStrIdx(t *testing.T) {
	misc.TrackTime(time.Now())
	bytes, _ := ioutil.ReadFile("./one.xml")
	xstr := string(bytes)

	// bLocGrp, types := brktLoc(xstr)
	// fPln(bLocGrp)
	// // fPln(types)
	// fPln(" -------------------- ")
	// cLocGrp, types := conTxtLoc(xstr, bLocGrp, types)
	// fPln(cLocGrp)

	// for _, loc := range cLocGrp {
	// 	s, e := loc[0], loc[1]
	// 	fPln(xstr[s:e])
	// }

	// merged := locMerge(bLocGrp, cLocGrp)
	// fPln(merged)

	// for _, loc := range merged {
	// 	fPln(xstr[loc[0]:loc[1]])
	// }

	// aLocGrp, types = attrValLoc(xstr, types)
	// fPln(aLocGrp)
	// fPln(types)

	fPln(Fmt(xstr))
}
