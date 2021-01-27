package xmltool

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/misc"
	jt "github.com/cdutwhu/json-tool"
)

func TestMkJSON(t *testing.T) {
	defer misc.TrackTime(time.Now())
	dir := "./examples/"

	SetSlim(false)
	SetAttrPrefix("")
	SetContAttrName("value")
	SetIgnrAttr(
		"xsi:nil",
		"xmlns:xsd",
		"xmlns:xsi",
		"xmlns",
	)
	SetSuffix4List(
		`List`,
		`MedicalAlertMessages`,
		`OtherNames`,
		`CountriesOfCitizenship`,
		`CountriesOfResidency`,
		`YearLevels`,
		`IdentityAssertions`,
		`LearningStandards`,
		`RelatedLearningStandardItems`,
		`AttendanceTimes`,
		`PeriodAttendances`,
	)
	SetNonstrPath(
		'/',
		"Activity/Title/Bool",
		"Activity/Num",
		"Activity/NumList/Num/Bool",
		"Activity/Points",
		"Activity/NumList/Num",
		"Activity/ActivityWeight",
		"Activity/MaxAttemptsAllowed",
		"Activity/ActivityTime/Duration/value",
		"MarkValueInfo/ValidLetterMarkList/ValidLetterMark/NumericEquivalent",
		"NAPCodeFrame/TestletList/Testlet/TestItemList/TestItem/TestItemContent/MultipleChoiceOptionCount",
		// "NAPCodeFrame/TestletList/Testlet/TestItemList/TestItem/TestItemContent/ItemDifficultyLogit5",
	)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if xmlfile := info.Name(); sHasSuffix(xmlfile, ".xml") {

			// if xmlfile != "Activity_0.xml" {
			// 	return nil
			// }

			if xmlfile == "n2sif.xml" {
				return nil
			}

			fPln("--->", xmlfile)

			bytes, _ := ioutil.ReadFile(dir + xmlfile)
			jstr := MkJSON(string(bytes))

			if !jt.IsValid(jstr) {
				io.MustWriteFile(fSf("debug_%s.json", sTrimSuffix(xmlfile, ".xml")), []byte(jstr))
				panic("error on MkJSON")
			}

			//if xmlfile == "CensusCollection_0.xml" {
			io.MustWriteFile(fSf("./out/record_%s.json", sTrimSuffix(xmlfile, ".xml")), []byte(jstr))
			//}
		}
		return nil
	})
}

func BenchmarkMkJSON(b *testing.B) {
	defer misc.TrackTime(time.Now())

	SetSlim(true)
	SetIgnrAttr(
		"xsi:nil",
		"xmlns:xsd",
		"xmlns:xsi",
		"xmlns",
	)
	SetSuffix4List(
		`List`,
		`MedicalAlertMessages`,
		`OtherNames`,
		`CountriesOfCitizenship`,
		`CountriesOfResidency`,
		`YearLevels`,
		`IdentityAssertions`,
		`LearningStandards`,
		`RelatedLearningStandardItems`,
		`AttendanceTimes`,
		`PeriodAttendances`,
	)

	for n := 0; n < b.N; n++ {
		dir := "./examples/"
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if xmlfile := info.Name(); sHasSuffix(xmlfile, ".xml") {
				// fPln("--->", xmlfile)
				if xmlfile == "n2sif.xml" {
					return nil
				}
				bytes, _ := ioutil.ReadFile(dir + xmlfile)
				MkJSON(string(bytes))

				// jstr := MkJSON(string(bytes))
				// if !jt.IsValid(jstr) {
				// 	fPln(jstr)
				// 	panic("error on MkJSON")
				// }
			}
			return nil
		})
	}
}
