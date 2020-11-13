package xmltool

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
	jt "github.com/cdutwhu/json-tool"
)

func TestMkJSON(t *testing.T) {
	misc.TrackTime(time.Now())
	dir := "./examples/"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if xmlfile := info.Name(); sHasSuffix(xmlfile, ".xml") {
			fPln("--->", xmlfile)
			bytes, _ := ioutil.ReadFile(dir + xmlfile)
			xstr := string(bytes)
			jstr := MkJSON(xstr, "", "")
			if !jt.IsValid(jstr) {
				// fPln(jstr)
				ioutil.WriteFile("debug.json", []byte(jstr), 0666)
				panic("error on MkJSON")
			}
			// if xmlfile == "Debug3.xml" {
			// 	fPln(jstr)
			// }
		}
		return nil
	})
}

func BenchmarkMkJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// misc.TrackTime(time.Now())
		dir := "./examples/"
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if xmlfile := info.Name(); sHasSuffix(xmlfile, ".xml") {
				// fPln("--->", xmlfile)
				bytes, _ := ioutil.ReadFile(dir + xmlfile)
				xstr := string(bytes)
				jstr := MkJSON(xstr, "", "")
				if !jt.IsValid(jstr) {
					fPln(jstr)
					panic("error on MkJSON")
				}
				// if xmlfile == "Activity_2.xml" {
				// 	fPln(jstr)
				// }
			}
			return nil
		})
	}
}
