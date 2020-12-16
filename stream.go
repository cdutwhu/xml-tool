package xmltool

import (
	"bufio"
	"io"
	"regexp"
)

// StreamEle :
func StreamEle(reader *bufio.Reader, elements ...string) chan string {
	eleChan := make(chan string, 2048)

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
				if len(data) >= maxScanToken {
					jump4close = true

					advance, token = maxScanToken, data
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
				if len(data) >= maxScanToken {
					jump4close = true

					advance, token = maxScanToken, data
					return advance, token, nil
				}
				return 0, nil, nil
			}

			// couldn't find openTag, reach the max buf, skip max, return 0
			if len(data) >= maxScanToken {
				advance, token = maxScanToken, []byte{0}
				return advance, token, nil
			}
			return 0, nil, nil
		}

		scanner := bufio.NewScanner(reader)
		scanner.Split(split)
		// scanner.Buffer(make([]byte, 4096), bufio.MaxScanTokenSize)

		sb := sBuilder{}
		sb.Grow(8192)
		for scanner.Scan() {
			str := scanner.Text()
			if len(str) == 1 {
				continue
			}

			sb.WriteString(str)

			if len(str) < maxScanToken {
				eleChan <- sb.String()
				sb.Reset()
			}
		}

		close(eleChan)
	}()

	return eleChan
}
