package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var challengeTokens = []string{
	"0:33", "0:2", "0:25", "0:3", "0:30", "0:15", "0:28", "0:5", "0:28", "0:0",
	"1:36", "0:70", "2:18", "0:5", "0:6", "1:36", "0:15", "0:6", "0:25", "0:3", "0:5", "0:10", "0:5", "0:33", "0:15", "0:28", "0:5", "0:10", "0:14", "0:0",
	"0:33", "0:2", "0:25", "0:21", "0:14", "4:52", "0:0",
	"0:28", "0:5", "0:5", "0:0",
	"1:36", "0:5", "0:28", "0:0",
	"0:5", "0:25", "0:0",
	"0:30", "0:14", "0:28", "0:5", "0:0",
	"0:5", "0:1", "0:21", "0:0",
	"0:70", "0:30", "0:30", "0:14", "0:3", "0:14", "0:21", "0:14", "y:x",
	"1:36", "0:25", "0:6", "0:480", "0:15", "0:3", "0:14", "0:0",
	"1:36", "0:14", "0:1", "0:21", "0:25", "0:21", "0:5", "0:5", "0:6", "0:25", "0:67", "0:25", "0:28",
	"0:14", "0:33", "0:14", "0:30", "0:21", "0:15", "0:28", "0:5", "0:0",
	"0:67", "0:2", "0:25", "0:28", "0:3", "0:14", "0:3", "0:15", "0:28", "0:5", "0:0",
	"0:30", "0:14", "0:4", "0:5", "0:1", "0:21", "0:15", "0:28", "0:5", "1:36", "0:14", "0:28", "0:3", "0:25",
}

const (
	firstParaLen = 183
	url          = "https://raw.githubusercontent.com/greyscalepress/manifestos/refs/heads/master/content/manifestos/1986-hacker-manifesto.txt"
	start        = 17
	end          = 72
)

func main() {
	text, err := getText()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result []string
	for _, t := range challengeTokens {
		result = append(result, decode(t, text))
	}
	fmt.Println(strings.Join(result, ""))
}

func getText() (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var builder strings.Builder

	i := 0
	for scanner.Scan() {
		i++
		if i < start {
			continue
		}
		if i > end {
			break
		}
		line := scanner.Text()
		for _, r := range line {
			if !unicode.IsSpace(r) {
				builder.WriteRune(r)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed during scanning: %w", err)
	}

	return builder.String(), nil
}

func decode(t, text string) string {
	switch t {
	case "y:x":
		return ". "
	case "0:0":
		return " "
	case "0:70":
		return "ü"
	case "0:2":
		return "õ"
	case "4:52":
		t = "2:18"
	}

	parts := strings.Split(t, ":")
	if len(parts) != 2 {
		return "?"
	}

	p, err1 := strconv.Atoi(parts[0])
	i, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return "?"
	}

	if p > 0 {
		i += firstParaLen
	}

	runes := []rune(text)
	if i < 0 || i >= len(runes) {
		return "#"
	}
	return string(runes[i])
}
