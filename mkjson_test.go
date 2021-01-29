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

	SetSlim(false)
	SetAttrPrefix("")
	SetContAttrName("value")
	SetIgnrAttr(
		"xsi:nil",
		"xmlns:xsd",
		"xmlns:xsi",
		"xmlns",
	)
	// SetListPathSuf(
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
	SetListPath(
		'/',
		"NAPCodeFrame/TestletList/Testlet/TestItemList/TestItem/TestItemContent/NAPWritingRubricList",
		"NAPCodeFrame/TestletList/Testlet/TestItemList/TestItem/TestItemContent/NAPWritingRubricList/NAPWritingRubric/ScoreList/Score/ScoreDescriptionList",
		"TimeTable/TimeTableDayList",
		"TimeTable/TimeTableDayList/TimeTableDay/TimeTablePeriodList",
		"TimeTableCell/TeacherList",
		"TeachingGroup/TeachingGroupPeriodList",
		"TeachingGroup/TeacherList",
		"TeachingGroup/StudentList",
		"SystemRole/SystemContextList/SystemContext/RoleList/Role/RoleScopeList",
		"SystemRole/SystemContextList",
		"StudentSchoolEnrollment/PublishingPermissionList",
		"StudentPersonal/LocalCodeList",
		"StudentPersonal/PersonInfo/PhoneNumberList",
		"StudentPersonal/PersonInfo/EmailList",
		"StudentPersonal/PersonInfo/AddressList",
		"StudentPersonal/PersonInfo/Demographics/ReligiousEventList",
		"StudentPersonal/PersonInfo/Demographics/CountriesOfResidency",
		"StudentPersonal/PersonInfo/Demographics/CountriesOfCitizenship",
		"StudentPersonal/PersonInfo/OtherNames",
		"StudentPersonal/ElectronicIdList",
		"StudentPersonal/MedicalAlertMessages",
		"StudentDailyAttendance/AttendanceCode/OtherCodeList",
		"StudentContactRelationship/HouseholdList",
		"StudentContactPersonal/PersonInfo/EmailList",
		"StudentContactPersonal/PersonInfo/PhoneNumberList",
		"StudentContactPersonal/PersonInfo/AddressList",
		"StudentContactPersonal/PersonInfo/Demographics/ReligiousEventList",
		"StudentContactPersonal/PersonInfo/Demographics/CountriesOfResidency",
		"StudentContactPersonal/PersonInfo/Demographics/CountriesOfCitizenship",
		"StudentContactPersonal/PersonInfo/OtherNames",
		"StudentContactPersonal/OtherIdList",
		"StudentAttendanceTimeList/PeriodAttendances/PeriodAttendance/TeacherList",
		"StudentAttendanceTimeList/PeriodAttendances",
		"StudentAttendanceTimeList/AttendanceTimes/AttendanceTime/AttendanceCode/OtherCodeList",
		"StudentAttendanceTimeList/AttendanceTimes",
		"StudentAttendanceCollection/StudentAttendanceCollectionReportingList/StudentAttendanceCollectionReporting/StatsCohortYearLevelList",
		"StudentAttendanceCollection/StudentAttendanceCollectionReportingList/StudentAttendanceCollectionReporting/StatsCohortYearLevelList/StatsCohortYearLevel/StatsCohortList",
		"StudentActivityInfo/YearLevels",
		"Activity/NumList",
		"Activity/Num1List",
		"Activity/SoftwareRequirementList",
		"AddressCollection/AddressCollectionReportingList/AddressCollectionReporting/AGContextualQuestionList",
		"CalendarSummary/YearLevels",
		"CensusCollection/CensusReportingList/CensusReporting/CensusStaffList",
		"CensusCollection/CensusReportingList/CensusReporting/CensusStudentList",
		"CollectionRound/AGRoundList",
		"CollectionStatus/AGReportingObjectResponseList",
		"StudentPersonal/OtherIdList",
		"NAPStudentResponseSet/TestletList",
		"NAPStudentResponseSet/TestletList/Testlet/ItemResponseList",
		"NAPStudentResponseSet/DomainScore/PlausibleScaledValueList",
		"NAPStudentResponseSet/TestletList/Testlet/ItemResponseList/ItemResponse/SubscoreList",
		"FinancialQuestionnaireCollection/FQReportingList",
		"FinancialQuestionnaireCollection/FQReportingList/FQReporting/FQContextualQuestionList",
		"FinancialQuestionnaireCollection/FQReportingList/FQReporting/FQItemList",
		"Identity/IdentityAssertions",
		"LearningResource/YearLevels",
		"LearningResource/LearningStandards",
		"LearningStandardDocument/YearLevels",
		"LearningStandardItem/YearLevels",
		"LearningStandardItem/StandardIdentifier/YearLevels",
		"LearningStandardItem/RelatedLearningStandardItems",
		"LibraryPatronStatus/ElectronicIdList",
		"LibraryPatronStatus/TransactionList",
		"LibraryPatronStatus/MessageList",
		"MarkValueInfo/ValidLetterMarkList",
		"NAPCodeFrame/TestletList/Testlet/TestItemList/TestItem/TestItemContent/ContentDescriptionList",
		"NAPCodeFrame/TestletList/Testlet/TestItemList",
		"NAPTestItem/TestItemContent/ContentDescriptionList",
		"NAPTestlet/TestItemList",
		"PersonPicture/PublishingPermissionList",
		"ResourceUsage/ResourceReportColumnList",
		"ResourceUsage/ResourceReportLineList",
		"ScheduledActivity/TeacherList",
		"ScheduledActivity/TeachingGroupList",
		"SchoolInfo/SchoolFocusList",
		"SchoolInfo/YearLevels",
		"SchoolInfo/SchoolGroupList",
		"SchoolInfo/SchoolContactList/SchoolContact/ContactInfo/EmailList",
		"SchoolPrograms/SchoolProgramList",
		"StaffAssignment/StaffSubjectList",
		"StaffAssignment/YearLevels",
		"StaffPersonal/PersonInfo/OtherNames",
		"StaffPersonal/PersonInfo/Demographics/CountriesOfCitizenship",
		"StaffPersonal/PersonInfo/Demographics/CountriesOfResidency",
		"StaffPersonal/PersonInfo/Demographics/ReligiousEventList",
		"StaffPersonal/PersonInfo/AddressList",
		"StaffPersonal/PersonInfo/PhoneNumberList",
		"StaffPersonal/PersonInfo/EmailList",
		"StaffPersonal/MostRecent/NAPLANClassList",
	)
	SetNonStrPath(
		'/',
		"Activity/Title/Bool",
		"Activity/Num",
		"Activity/NumList/Num/Bool",
		"Activity/Points",
		"Activity/NumList/Num",
		"Activity/Num1List/Num",
		"Activity/ActivityWeight",
		"Activity/MaxAttemptsAllowed",
		"Activity/ActivityTime/Duration/value",
		"MarkValueInfo/ValidLetterMarkList/ValidLetterMark/NumericEquivalent",
		"NAPCodeFrame/TestletList/Testlet/TestItemList/TestItem/TestItemContent/MultipleChoiceOptionCount",
		"NAPCodeFrame/TestletList/Testlet/TestItemList/TestItem/TestItemContent/ItemDifficultyLogit5",
		"NAPTestItem/TestItemContent/ItemDifficultyLogit5",
	)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if xmlfile := info.Name(); sHasSuffix(xmlfile, ".xml") {

			// if xmlfile != "StudentContactPersonal_0.xml" {
			// 	return nil
			// }

			if xmlfile == "n2sif.xml" || xmlfile == "siftest.xml" {
				return nil
			}

			fPln("--->", xmlfile)

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
	SetListPathSuf(
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
