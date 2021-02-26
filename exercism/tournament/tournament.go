package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name    string
	matches int
	win     int
	loss    int
	draw    int
	points  int
}

type resultsTable map[string][]int
type rankingsTable []team

func readResultsTable(r io.Reader) (resultsTable, error) {

	resT := resultsTable{}
	b := make([]byte, 1024)

	n, err := r.Read(b)
	if err != nil {
		return nil, err
	}

	// trim the last character from input
	s := strings.TrimSpace(string(b[:n]))
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		// trim space characters in each line
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		data := strings.Split(line, ";")
		if len(data) < 3 || (data[2] != "win" && data[2] != "loss" && data[2] != "draw") {
			return nil, errors.New("Wrong format of the input data.")
		}

		h := strings.TrimSpace(data[0])
		v := strings.TrimSpace(data[1])
		r := strings.TrimSpace(data[2])

		if resT[h] == nil {
			resT[h] = make([]int, 3)
		}
		if resT[v] == nil {
			resT[v] = make([]int, 3)
		}

		if r == "win" {
			resT[h][0] = resT[h][0] + 1
			resT[v][1] = resT[v][1] + 1
		} else if r == "loss" {
			resT[h][1] = resT[h][1] + 1
			resT[v][0] = resT[v][0] + 1
		} else if r == "draw" {
			resT[h][2] = resT[h][2] + 1
			resT[v][2] = resT[v][2] + 1
		}
	}

	return resT, nil
}

func buildRankingsTable(resT resultsTable) rankingsTable {
	rankT := rankingsTable{}
	for k, v := range resT {
		rankT = append(rankT, team{
			name:    k,
			matches: v[0] + v[1] + v[2],
			win:     v[0],
			loss:    v[1],
			draw:    v[2],
			points:  v[0]*3 + v[2],
		})
	}
	return rankT
}

func sortRankingsTable(rankT rankingsTable) {
	sort.Slice(rankT, func(i, j int) bool {
		if rankT[i].points == rankT[j].points {
			return rankT[i].name < rankT[j].name
		}
		return rankT[i].points > rankT[j].points
	})
}

func writeRankingsTable(teams []team, w io.Writer) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%-30s |%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P"))
	for _, t := range teams {
		sb.WriteString(fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n", t.name, t.matches, t.win, t.draw, t.loss, t.points))
	}
	w.Write([]byte(sb.String()))
}

func Tally(r io.Reader, w io.Writer) error {
	resT, err := readResultsTable(r)
	if err != nil {
		return err
	}
	rankT := buildRankingsTable(resT)
	sortRankingsTable(rankT)
	writeRankingsTable(rankT, w)

	return nil
}
