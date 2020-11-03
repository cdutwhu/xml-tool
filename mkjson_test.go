package xmltool

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
	jt "github.com/cdutwhu/json-tool"
)

func TestMkJSON(t *testing.T) {
	misc.TrackTime(time.Now())
	bytes, _ := ioutil.ReadFile("./one.xml")
	xstr := string(bytes)
	jstr := MkJSON(xstr)
	fPln(jt.IsValid(jstr))
	fPln(jstr)
}
