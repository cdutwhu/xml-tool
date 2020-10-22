package xmltool

import "testing"

func TestBasic(t *testing.T) {
	xml := `<Activity RefId="C27E1FCF-C163-485F-BEF0-F36F18A0493A" AMD="YES">
	<Title>Shakespeare Essay - Much Ado About Nothing</Title>
	<Preamble>This is a very funny comedy - students should have passing familiarity with Shakespeare</Preamble>
	<LearningStandards>
	<LearningStandardItemRefId>9DB15CEA-B2C5-4F66-94C3-7D0A0CAEDDA4</LearningStandardItemRefId>
	</LearningStandards>
	<SourceObjects>
	<SourceObject SIF_RefObject="Lesson">A71ADBD3-D93D-A64B-7166-E420D50EDABC</SourceObject>
	</SourceObjects>
	<Points>50</Points>
	<ActivityTime>
	<CreationDate>2002-06-15</CreationDate>
	<Duration Units="minute">30</Duration>
	<StartDate>2002-09-10</StartDate>
	<FinishDate>2002-09-12</FinishDate>
	<DueDate>2002-09-12</DueDate>
	</ActivityTime>
	<AssessmentRefId>03EDB29E-8116-B450-0435-FA87E42A0AD2</AssessmentRefId>
	<MaxAttemptsAllowed>3</MaxAttemptsAllowed>
	<ActivityWeight>5</ActivityWeight>
	<Evaluation EvaluationType="Inline">
	<Description>Students should be able to correctly identify all major characters.</Description>
	</Evaluation>
	<LearningResources>
	<LearningResourceRefId>B7337698-BF6D-B193-7F79-A07B87211B93</LearningResourceRefId>
	</LearningResources>
	</Activity>`

	name, out, in := Lvl0(xml)
	fPln(name)
	fPln(" ---------- ")
	fPln(out)
	fPln(" ---------- ")
	fPln(in)

	fPln("valid xml?", IsValid(xml))
	fPln("root:", Root(xml))
	tag, attrs, mVal := TagAttrVal(xml)
	fPln("tag:", tag)
	fPln("attrs:", attrs)
	fPln("av:", mVal)
	fPln("cont:", Cont(xml))
}
