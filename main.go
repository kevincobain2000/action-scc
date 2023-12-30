package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	limit         string
	limitInt      int
	readmeFile    string
	instachartUrl string
	subtitle      string
	style         string
	width         string
	height        string
	altText       string
)

type SccStat struct {
	Name               string   `json:"Name"`
	Bytes              int      `json:"Bytes"`
	CodeBytes          int      `json:"CodeBytes"`
	Lines              int      `json:"Lines"`
	Code               int      `json:"Code"`
	Comment            int      `json:"Comment"`
	Blank              int      `json:"Blank"`
	Complexity         int      `json:"Complexity"`
	Count              int      `json:"Count"`
	WeightedComplexity int      `json:"WeightedComplexity"`
	Files              []string `json:"Files"`
}
type BarData struct {
	X     []string `json:"x"`
	Y     [][]int  `json:"y"`
	Names []string `json:"names"`
}

func main() {
	flags()

	stats := readStatsFromPipe()
	barChartURL := createBarChartURL(stats, limitInt)
	markdownImgeURL := "![" + altText + "](" + barChartURL + ")"
	newOrReplaceInFile(markdownImgeURL, readmeFile)
	fmt.Println(markdownImgeURL)
}

func newOrReplaceInFile(markdownImgeURL string, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)

	found := false
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if strings.Contains(line, "!["+altText+"]") {
			found = true
			lines[i] = markdownImgeURL
		}
	}
	if found {
		newContent := []byte(strings.Join(lines, "\n"))
		os.WriteFile(readmeFile, newContent, 0644)
	}

	if !found {
		newContent := []byte(markdownImgeURL + "\n\n" + string(content))
		os.WriteFile(readmeFile, newContent, 0644)
	}
}

func normalizeName(name string) string {
	name = strings.ReplaceAll(name, "/", "-")
	name = strings.ReplaceAll(name, " ", "-")
	return name
}

func createBarChartURL(scc []SccStat, limit int) string {
	var chartURL string
	var data BarData

	data.Names = []string{"Code", "Comment", "Files", "Complexity"}
	for idx, stat := range scc {
		if idx >= limit {
			break
		}
		data.X = append(data.X, normalizeName(stat.Name))
	}

	y := make([][]int, len(data.Names))
	for _, x := range data.X {
		for _, stat := range scc {
			if x == normalizeName(stat.Name) {
				y[0] = append(y[0], stat.Code)
				y[1] = append(y[1], stat.Comment)
				y[2] = append(y[2], stat.Count)
				y[3] = append(y[3], stat.Complexity)
			}
		}
	}
	data.Y = y

	json, _ := json.Marshal(data)
	queries := map[string]string{
		"title":    "SCC+-+Sloc,+Cloc+and+Code",
		"width":    width,
		"height":   height,
		"subtitle": subtitle,
		"metric":   "+lines",
		"data":     string(json),
	}

	chartURL = instachartUrl + "/" + style + "?" + createQuery(queries)
	return chartURL
}

func readStatsFromPipe() []SccStat {
	stat, _ := os.Stdin.Stat()

	var stats []SccStat
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		sccJSON, _ := ioutil.ReadAll(os.Stdin)
		err := json.Unmarshal([]byte(sccJSON), &stats)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("No input data")
		os.Exit(1)
	}
	return stats
}

func createQuery(queries map[string]string) string {
	var query string
	for key, value := range queries {
		query += key + "=" + value + "&"
	}
	// strip last &
	query = query[:len(query)-1]
	return query
}

func flags() {
	flag.StringVar(&limit, "limit", "", "limit of the chart")
	flag.StringVar(&readmeFile, "readme-file", "", "readme filename")
	flag.StringVar(&instachartUrl, "instachart-url", "", "instachart url")
	flag.StringVar(&subtitle, "subtitle", "", "chart subtitle")
	flag.StringVar(&altText, "alt-text", "", "alt text for the image")
	flag.StringVar(&style, "style", "", "chart style")
	flag.StringVar(&width, "width", "", "chart width")
	flag.StringVar(&height, "height", "", "chart height")
	flag.Parse()

	// as flags can be empty string, and we want to use default values
	if instachartUrl == "" {
		flag.Set("instachart-url", "https://instachart.coveritup.app")
	}
	if subtitle == "" {
		flag.Set("subtitle", "Source Code Counter")
	}
	if readmeFile == "" {
		flag.Set("readme-file", "README.md")
	}
	if altText == "" {
		flag.Set("alt-text", "action-scc")
	}
	if style == "" {
		flag.Set("style", "bar")
	}
	if width == "" {
		flag.Set("width", "960")
	}
	if height == "" {
		flag.Set("height", "700")
	}
	limitInt, _ = strconv.Atoi(limit)
	if limitInt <= 0 {
		limitInt = 4
	}
}
