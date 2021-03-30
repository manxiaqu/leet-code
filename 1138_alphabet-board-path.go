package leetcode

func alphabetBoardPath(target string) string {
	currRow := byte(0)
	currColumn := byte(0)
	output := make([]byte, 0)

	for i := 0; i < len(target); {
		posRow := (target[i] - 'a') / 5
		posColumn := (target[i] - 'a') % 5

		// deal with 'z' => other
		if currRow == 5 && posRow != 5 {
			output = append(output, 'U')
			currRow = 4
			continue
		}

		// deal with
		if posColumn > currColumn {
			for currColumn < posColumn {
				output = append(output, 'R')
				currColumn++
			}
		} else if posColumn < currColumn {
			for currColumn > posColumn {
				output = append(output, 'L')
				currColumn--
			}
		}

		// step from current position to next position
		if posRow > currRow {
			for currRow < posRow {
				output = append(output, 'D')
				currRow++
			}
		} else if posRow < currRow {
			for currRow > posRow {
				output = append(output, 'U')
				currRow--
			}
		}

		output = append(output, '!')

		i++
	}

	return string(output)
}
