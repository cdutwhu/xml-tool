package xmltool

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/misc"
)

func TestBreakCont(t *testing.T) {
	defer misc.TrackTime(time.Now())

	// xml := `<Activity RefId="C27E1FCF-C163-485F-BEF0-F36F18A0493A">	<Title>Shakespeare Essay - Much Ado About Nothing</Title>
	// <Preamble>This is a very funny comedy - students should have passing familiarity with Shakespeare</Preamble>
	// <LearningStandards>
	// <LearningStandardItemRefId>9DB15CEA-B2C5-4F66-94C3-7D0A0CAEDDA4</LearningStandardItemRefId>
	// </LearningStandards>
	// <SourceObjects>
	// <SourceObject SIF_RefObject="Lesson">A71ADBD3-D93D-A64B-7166-E420D50EDABC</SourceObject>
	// </SourceObjects>
	// <Points>50</Points>
	// <ActivityTime>
	// <CreationDate>2002-06-15</CreationDate>
	// <Duration Units="minute">30</Duration>
	// <StartDate>2002-09-10</StartDate>
	// <FinishDate>2002-09-12</FinishDate>
	// <DueDate>2002-09-12</DueDate>
	// </ActivityTime>
	// <AssessmentRefId>03EDB29E-8116-B450-0435-FA87E42A0AD2</AssessmentRefId>
	// <MaxAttemptsAllowed>3</MaxAttemptsAllowed>
	// <ActivityWeight>5</ActivityWeight>
	// <Evaluation EvaluationType="Inline">
	// <Description>Students should be able to correctly identify all major characters.</Description>
	// </Evaluation>
	// <LearningResources>
	// <LearningResourceRefId>B7337698-BF6D-B193-7F79-A07B87211B93</LearningResourceRefId>
	// </LearningResources>
	// </Activity>`

	xmlbytes, _ := ioutil.ReadFile("./examples/siftest.xml")
	xml := string(xmlbytes)

	name, out, in := Lvl0(xml)
	fPln(name)
	fPln(" ---------- ")
	fPln(out)
	fPln(" ---------- ")
	fPln(in)
	fPln(" ---------- ")
	names, subContents := BreakCont(in)
	mFNIdx := make(map[string]int)
	for i, name := range names {
		fPln(" *** ", name)
		sub := subContents[i]
		sub = RmComment(sub)
		subFmt := Fmt(sub)

		if !IsValid(sub) || !IsValid(subFmt) {
			io.MustWriteFile("./debug/"+name+".xml", []byte(sub))
			io.MustWriteFile("./debug/"+name+"_fmt.xml", []byte(subFmt))
			panic("Invalid XML @ " + name)
		}
		// fPln(sub)
		// fPln(subFmt)

		// ----- save to file ----- //
		idx := mFNIdx[name]
		filename := fSf("%s_%d.xml", name, idx)
		io.MustWriteFile("./examples348/"+filename, []byte(subFmt))
		mFNIdx[name]++
	}
}
