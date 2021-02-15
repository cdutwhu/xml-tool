package xmltool

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/embres"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/misc"
	jt "github.com/cdutwhu/json-tool"
	// m "github.com/cdutwhu/xml-tool/mkjson"
)

func TestPrepareLineKeyMap(t *testing.T) {
	io.EditFileByLine("./mkjson/348d.txt", func(ln string) (bool, string) {
		if strings.HasPrefix(ln, "NUMERIC:") {
			return true, ln[len("NUMERIC: "):] + "\n" + ln[len("NUMERIC: "):] + "/value"
		}
		return false, ""
	}, "./mkjson/NUMERIC.txt")
	embres.PrintLineKeyMap("mkjson", "MNumPath348D", "./mkjson/numpath348draft", "./mkjson/NUMERIC.txt")

	io.EditFileByLine("./mkjson/348d.txt", func(ln string) (bool, string) {
		if strings.HasPrefix(ln, "BOOLEAN:") {
			return true, ln[len("BOOLEAN: "):] + "\n" + ln[len("BOOLEAN: "):] + "/value"
		}
		return false, ""
	}, "./mkjson/BOOLEAN.txt")
	embres.PrintLineKeyMap("mkjson", "MBoolPath348D", "./mkjson/boolpath348draft", "./mkjson/BOOLEAN.txt")

	io.EditFileByLine("./mkjson/348d.txt", func(ln string) (bool, string) {
		if strings.HasPrefix(ln, "LIST:") {
			part := ln[len("LIST: "):]
			ss := strings.Split(part, "/")
			return true, strings.Join(ss[:len(ss)-1], "/")
		}
		return false, ""
	}, "./mkjson/LIST.txt")
	embres.PrintLineKeyMap("mkjson", "MListPath348D", "./mkjson/listpath348draft", "./mkjson/LIST.txt")
}

func TestXmlTxtEscCharProc(t *testing.T) {
	part := ` he"llo
	world  `
	part = xmlTxtEscCharProc(part)
	fPln(part)
}

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

	if err := SetPathByFile("LIST", "./mkjson/LIST.txt", "listSIF347", true, '/'); err != nil {
		panic(err)
	}
	if err := EnableListPath("listSIF347"); err != nil {
		panic(err)
	}
	// DirectListPath('/', m.MListPath348D)

	if err := SetPathByFile("TYPE", "./mkjson/BOOLEAN.txt", "typeSIF347", true, '/'); err != nil {
		panic(err)
	}
	if err := SetPathByFile("TYPE", "./mkjson/NUMERIC.txt", "typeSIF347", true, '/'); err != nil {
		panic(err)
	}
	if err := EnableNonStrPath("typeSIF347"); err != nil {
		panic(err)
	}
	// DirectNonStrPath('/', m.MBoolPath348D, m.MNumPath348D)

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
