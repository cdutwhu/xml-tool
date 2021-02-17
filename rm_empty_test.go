package xmltool

import (
	"os"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
	jt "github.com/cdutwhu/json-tool"
)

func TestRmEmptyEle(T *testing.T) {
	defer misc.TrackTime(time.Now())

	bytes, err := os.ReadFile("./examples/err-json.xml")
	if err != nil {
		panic(err)
	}
	xml := string(bytes)

	if !IsValid(xml) {
		panic("invalid xml")
	}

	remainder := RmEmptyEle(xml, 3, false)
	// fPln(remainder)
	// os.WriteFile("debug.xml", []byte(remainder), 0666)

	if !IsValid(remainder) {
		panic("invalid remainder")
	}

	fPln("-----------------------")

	SetSlim(false)
	SetListPathSuffix("List")
	jsonstr := MkJSON(remainder)
	if !jt.IsValid(jsonstr) {
		panic("invalid json str")
	}
	// fPln(jsonstr)
}

func BenchmarkRmEmptyEle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := os.ReadFile("./examples/n2sif_2.xml")
		if err != nil {
			panic(err)
		}
		xml := string(bytes)
		SetSlim(false)
		SetListPathSuffix("List")

		remainder := RmEmptyEle(xml, 3, true)
		MkJSON(remainder)
	}
}
