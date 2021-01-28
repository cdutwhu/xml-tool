package xmltool

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
	jt "github.com/cdutwhu/json-tool"
)

func TestRmEmptyEle(T *testing.T) {
	defer misc.TrackTime(time.Now())

	bytes, err := ioutil.ReadFile("./examples/err-json.xml")
	if err != nil {
		panic(err)
	}
	xml := string(bytes)

	if !IsValid(xml) {
		panic("invalid xml")
	}

	remainder := RmEmptyEle(xml, 3, false)
	// fPln(remainder)
	// ioutil.WriteFile("debug.xml", []byte(remainder), 0666)

	if !IsValid(remainder) {
		panic("invalid remainder")
	}

	fPln("-----------------------")

	SetSlim(false)
	SetListPathSuf("List")
	jsonstr := MkJSON(remainder)
	if !jt.IsValid(jsonstr) {
		panic("invalid json str")
	}
	// fPln(jsonstr)
}

func BenchmarkRmEmptyEle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := ioutil.ReadFile("./examples/n2sif_2.xml")
		if err != nil {
			panic(err)
		}
		xml := string(bytes)
		SetSlim(false)
		SetListPathSuf("List")

		remainder := RmEmptyEle(xml, 3, true)
		MkJSON(remainder)
	}
}
