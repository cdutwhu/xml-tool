package xmltool

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestRmEmptyEle(T *testing.T) {
	defer misc.TrackTime(time.Now())

	bytes, err := ioutil.ReadFile("./examples/n2sif_2.xml")
	if err != nil {
		panic(err)
	}
	xml := string(bytes)

	SetSlim(false)
	SetSuffix4List("List")
	
	remainder := RmEmptyEle(xml, 3, true)
	fPln(remainder)

	fPln("-----------------------")

	jstr := MkJSON(remainder)
	fPln(jstr)
}

func BenchmarkRmEmptyEle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := ioutil.ReadFile("./examples/n2sif_2.xml")
		if err != nil {
			panic(err)
		}
		xml := string(bytes)
		SetSlim(false)
		SetSuffix4List("List")
		
		remainder := RmEmptyEle(xml, 3, true)
		MkJSON(remainder)
	}
}
