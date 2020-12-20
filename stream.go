package xmltool

import (
	"bufio"
	"io"
	"regexp"
)

// StreamEle :
func StreamEle(reader *bufio.Reader, elements ...string) chan string {
	eleChan := make(chan string, 4096)

	go func() {

		restr := `<[\w\d-_]+[ >]` // if no needed elements provided, get all top elements
		// else, select we wanted
		if len(elements) > 0 {
			restr = "<("
			for _, ele := range elements {
				restr += fSf("(%s)|", ele)
			}
			restr = sTrimRight(restr, "|") + ")[> ]"
		}

		rOpenTag := regexp.MustCompile(restr)
		closeTag := ""
		jump4close := false

		split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF {
				return 0, nil, io.EOF
			}

			str := string(data)

			if jump4close {

				// find closeTag
				if e := sIndex(str, closeTag); e >= 0 {
					jump4close = false

					advance = e + len(closeTag)
					token = []byte(str[:advance])
					return advance, token, nil
				}

				// still looking for closeTag
				// bigger than max buf,
				if len(data) >= bufio.MaxScanTokenSize {
					jump4close = true

					advance, token = bufio.MaxScanTokenSize, data
					return advance, token, nil
				}
				return 0, nil, nil
			}

			if loc := rOpenTag.FindStringIndex(str); loc != nil {
				closeTag = fSf("</%s>", str[loc[0]+1:loc[1]-1])
				s := loc[0]

				// advance buf right to openTag
				if s > 0 {
					advance, token = s, []byte{0}
					return advance, token, nil
				}

				// find closeTag
				if se := sIndex(str[s:], closeTag); se >= 0 { // correct open & close
					jump4close = false

					advance = s + se + len(closeTag)
					token = []byte(str[s:advance])
					return advance, token, nil
				}

				// correct open BUT could not find closeTag
				// bigger than max buf,
				if len(data) >= bufio.MaxScanTokenSize {
					jump4close = true

					advance, token = bufio.MaxScanTokenSize, data
					return advance, token, nil
				}
				return 0, nil, nil
			}

			// couldn't find openTag, reach the max buf, skip max, return 0
			if len(data) >= bufio.MaxScanTokenSize {
				advance, token = bufio.MaxScanTokenSize, []byte{0}
				return advance, token, nil
			}
			return 0, nil, nil
		}

		scanner := bufio.NewScanner(reader)
		scanBuf := make([]byte, bufio.MaxScanTokenSize)
		scanner.Buffer(scanBuf, bufio.MaxScanTokenSize)
		scanner.Split(split)
		sb := sBuilder{}
		sb.Grow(bufio.MaxScanTokenSize)

		for scanner.Scan() {
			str := scanner.Text()
			if len(str) == 1 {
				continue
			}
			sb.WriteString(str)
			if len(str) < bufio.MaxScanTokenSize {
				eleChan <- sb.String()
				sb.Reset()
			}
		}

		// useScanBuf := true
		// for scanner.Scan() {
		// 	str := scanner.Text()
		// 	if len(str) == 1 {
		// 		useScanBuf = false
		// 		continue
		// 	}
		// 	if useScanBuf {
		// 		eleChan <- string(scanBuf)
		// 	} else {
		// 		sb.WriteString(str)
		// 		if len(str) < bufio.MaxScanTokenSize {
		// 			eleChan <- sb.String()
		// 			sb.Reset()
		// 			useScanBuf = true
		// 		}
		// 	}
		// }

		close(eleChan)
	}()

	return eleChan
}
