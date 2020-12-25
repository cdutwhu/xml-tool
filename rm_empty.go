package xmltool

import (
	"github.com/cdutwhu/gotil/queue"
)

// locEmptyEle1 : "<tag/>"
func locEmptyEle1(xml string) (q queue.Queue) {
	remainder := xml
	offset := 0
	for e := sIndex(remainder, "/>"); e >= 0; e = sIndex(remainder, "/>") {
		e += offset + 2
		s := sLastIndex(xml[:e], "<")
		// fPln(s, e, xml[s:e])
		q.Enqueue([2]int{s, e})
		remainder = xml[e:]
		offset = e
	}
	return
}

// locEmptyEle2 : "<tag></tag>"
func locEmptyEle2(xml string) (q queue.Queue) {
	remainder := xml
	offset := 0
	for e := sIndex(remainder, "></"); e >= 0; e = sIndex(remainder, "></") {
		enq := true

		e += offset
		s1 := sLastIndex(xml[:e], "<")
		// exclude "</tag1></tag2>"
		if xml[s1+1] == '/' {
			enq = false
		}

		s2 := sIndex(xml[e+1:], ">") + e + 2
		// fPln(s1, s2)

		if enq {
			q.Enqueue([2]int{s1, s2})
		}

		remainder = xml[s2:]
		offset = s2
	}
	return
}

// locMergeV2 :
func locMergeV2(q0, q1 queue.Queue) (q queue.Queue) {
	qs := [2]*queue.Queue{&q0, &q1}
	for {
		flag := -1
		a0, ok0 := q0.Peek()
		a1, ok1 := q1.Peek()
		if ok0 && ok1 {
			if a0.([2]int)[0] < a1.([2]int)[0] {
				flag = 0
			} else {
				flag = 1
			}
		} else if ok0 && !ok1 {
			flag = 0
		} else if !ok0 && ok1 {
			flag = 1
		} else {
			break
		}

		if a, ok := qs[flag].Dequeue(); ok {
			q.Enqueue(a)
		}
	}
	return
}

// RmEmptyEle :
func RmEmptyEle(xml string, empType int, rmEmptyLine bool) string {
	q := queue.Queue{}
	switch empType {
	case 3:
		q = locMergeV2(locEmptyEle1(xml), locEmptyEle2(xml))
	case 2:
		q = locEmptyEle2(xml)
	case 1:
		q = locEmptyEle1(xml)
	case 0:
		return xml
	default:
		panic("empType: 1 <tag/>, 2 <tag></tag>, 3 both 1&2")
	}

	sb := sBuilder{}
	start, end := 0, 0
	for value, ok := q.Dequeue(); ok; value, ok = q.Dequeue() {
		se := value.([2]int)
		s, e := se[0], se[1]
		// fPln(s, e, xml[s:e])

		end = s
		sb.WriteString(xml[start:end])
		start = e
	}
	// last keep part
	sb.WriteString(xml[start:])
	removed := sb.String()

	if rmEmptyLine {
		r := rxMustCompile(`(\n+\s+\n+)+`)
		removed = r.ReplaceAllString(removed, "\n")
	}

	return removed
}
