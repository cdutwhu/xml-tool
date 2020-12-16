package xmltool

import (
	"bufio"
	"os"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestNewXMLEleV2(t *testing.T) {
	defer misc.TrackTime(time.Now())
	file, err := os.Open("./examples/n2sif.xml")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	br := bufio.NewReader(file)
	count := 0

	var dataTypes = []string{
		"NAPStudentResponseSet",
		"NAPEventStudentLink",
		"StudentPersonal",
		"NAPTestlet",
		"NAPTestItem",
		"NAPTest",
		"NAPCodeFrame",
		"SchoolInfo",
		"NAPTestScoreSummary",
	}

	for ele := range StreamEle(br, dataTypes...) {
		count++
		if count < 100 {
			fPln(count, len(ele))
		}
	}
	fPln("total", count)
}
