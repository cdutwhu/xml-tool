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

			switch xmlfile {
			case "CalendarSummary_0.xml", "Identity_1.xml", "Identity_2.xml", "LearningResource_0.xml":
				return nil
			case "LearningStandardDocument_0.xml", "LearningStandardItem_0.xml":
				return nil
			case "SchoolInfo_0.xml", "StaffAssignment_0.xml":
				return nil
			case "StaffPersonal_0.xml", "StudentActivityInfo_0.xml":
				return nil
			case "StudentAttendanceTimeList_0.xml", "StudentAttendanceTimeList_1.xml":
				return nil
			case "StudentContactPersonal_0.xml":
				return nil
			case "StudentPersonal_0.xml":
				return nil
			}

			// if xmlfile != "Debug3.xml" {
			// 	return nil
			// }

			bytes, _ := ioutil.ReadFile(dir + xmlfile)
			jstr := MkJSON(string(bytes), "", "")

			if !jt.IsValid(jstr) {
				ioutil.WriteFile(fSf("debug_%s.json", xmlfile), []byte(jstr), 0666)
				panic("error on MkJSON")
			}
			//if xmlfile == "CensusCollection_0.xml" {
			ioutil.WriteFile(fSf("record_%s.json", xmlfile), []byte(jstr), 0666)
			//}
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
				jstr := MkJSON(string(bytes), "", "")
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
