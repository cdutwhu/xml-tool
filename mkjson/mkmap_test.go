package mkjson

import (
	"testing"

	"github.com/cdutwhu/gotil/embres"
)

func TestGenMap(t *testing.T) {
	embres.PrintLineKeyMap("mkjson", "MListPath348D", "./listpath348draft", "./LIST.txt")
	embres.PrintLineKeyMap("mkjson", "MBoolPath348D", "./boolpath348draft", "./BOOLEAN.txt")
	embres.PrintLineKeyMap("mkjson", "MNumPath348D", "./numpath348draft", "./NUMERIC.txt")
}
