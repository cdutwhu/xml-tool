package xmltool

// RmComment : remove <!-- *** -->
func RmComment(xml string) string {
	sb := sBuilder{}
	remainder := xml
AGAIN:
	if s := sIndex(remainder, "<!--"); s >= 0 {
		sb.WriteString(remainder[:s])
		remainder = remainder[s:]
		if re := sIndex(remainder, "-->"); re >= 0 {
			remainder = remainder[re+3:]
			goto AGAIN
		}
	} else {
		sb.WriteString(remainder)
		goto END
	}
END:
	return sb.String()
}
