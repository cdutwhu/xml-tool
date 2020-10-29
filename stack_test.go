package xmltool

import "testing"

func TestStack(t *testing.T) {
	var s4s stack
	s4s.push("abc")
	s4s.push("def")
	s4s.push("ghi")

	val, ok := s4s.pop()
	fPln(val, ok, s4s)

	val, ok = s4s.pop()
	s4s.push("jkl")
	fPln(val, ok, s4s)

	val, ok = s4s.pop()
	fPln(val, ok, s4s)

	val, ok = s4s.pop()
	fPln(val, ok, s4s)
}
