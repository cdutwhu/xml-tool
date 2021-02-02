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
	dir := "./examples348/"
	if !dirExists(dir) {
		fPln(dir, " does not exist")
		return
	}

	SetSlim(false)
	SetAttrPrefix("")
	SetContAttrName("value")
	SetIgnrAttr(
		"xsi:nil",
		"xmlns:xsd",
		"xmlns:xsi",
		"xmlns",
	)
	// SetListPathSuffix(
	// 	`List`,
	// 	`MedicalAlertMessages`,
	// 	`OtherNames`,
	// 	`CountriesOfCitizenship`,
	// 	`CountriesOfResidency`,
	// 	`YearLevels`,
	// 	`IdentityAssertions`,
	// 	`LearningStandards`,
	// 	`RelatedLearningStandardItems`,
	// 	`AttendanceTimes`,
	// 	`PeriodAttendances`,
	// )

	SetPathByFile("LIST", "./mkjson/LIST.txt", "listSIF347", true, '/')
	SetPathByFile("LIST", "./mkjson/LIST1.txt", "listSIF347", true, '/')
	if err := EnableListPath("listSIF347"); err != nil {
		panic(err)
	}

	SetPathByFile("TYPE", "./mkjson/TYPE.txt", "typeSIF347", true, '/')
	if err := EnableNonStrPath("typeSIF347"); err != nil {
		panic(err)
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if xmlfile := info.Name(); sHasSuffix(xmlfile, ".xml") {

			// if xmlfile != "StudentContactPersonal_0.xml" {
			// 	return nil
			// }

			if xmlfile == "n2sif.xml" || xmlfile == "siftest.xml" {
				return nil
			}

			fPln("processing...", xmlfile)

			bytes, _ := ioutil.ReadFile(dir + xmlfile)
			jsonstr := MkJSON(string(bytes))

			if !jt.IsValid(jsonstr) {
				io.MustWriteFile(fSf("debug_%s.json", sTrimSuffix(xmlfile, ".xml")), []byte(jsonstr))
				panic("error on MkJSON")
			}

			//if xmlfile == "CensusCollection_0.xml" {
			io.MustWriteFile(fSf("./out/record_%s.json", sTrimSuffix(xmlfile, ".xml")), []byte(jsonstr))
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
	SetListPathSuffix(
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
				// fPln("processing...", xmlfile)
				if xmlfile == "n2sif.xml" {
					return nil
				}
				bytes, _ := ioutil.ReadFile(dir + xmlfile)
				MkJSON(string(bytes))

				// jsonstr := MkJSON(string(bytes))
				// if !jt.IsValid(jsonstr) {
				// 	fPln(jsonstr)
				// 	panic("error on MkJSON")
				// }
			}
			return nil
		})
	}
}
