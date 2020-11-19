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

	SetIgnrAttr(
		"xsi:nil",
		"xmlns:xsd",
		"xmlns:xsi",
		"xmlns",
	)
	SetSuf4LsEle(
		`List": `,
		`MedicalAlertMessages": `,
		`OtherNames": `,
		`CountriesOfCitizenship": `,
		`CountriesOfResidency": `,
		`YearLevels": `,
		`IdentityAssertions": `,
		`LearningStandards": `,
		`RelatedLearningStandardItems": `,
		`AttendanceTimes": `,
		`PeriodAttendances": `,
	)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if xmlfile := info.Name(); sHasSuffix(xmlfile, ".xml") {
			fPln("--->", xmlfile)

			// if xmlfile != "StudentPersonal_0.xml" {
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

	SetIgnrAttr(
		"xsi:nil",
		"xmlns:xsd",
		"xmlns:xsi",
		"xmlns",
	)
	SetSuf4LsEle(
		`List": `,
		`MedicalAlertMessages": `,
		`OtherNames": `,
		`CountriesOfCitizenship": `,
		`CountriesOfResidency": `,
		`YearLevels": `,
		`IdentityAssertions": `,
		`LearningStandards": `,
		`RelatedLearningStandardItems": `,
		`AttendanceTimes": `,
		`PeriodAttendances": `,
	)

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
			}
			return nil
		})
	}
}
