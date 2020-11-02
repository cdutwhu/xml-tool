package xmltool

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestMkJSON(t *testing.T) {
	misc.TrackTime(time.Now())
	bytes, _ := ioutil.ReadFile("./one.xml")
	xstr := string(bytes)
	fPln(MkJSON(xstr))
}
