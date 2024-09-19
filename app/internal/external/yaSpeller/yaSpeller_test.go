package yaspeller

import (
	"bufio"
	"os"
	"testing"
	"time"
)

func TestCorrectMistakeFunc(t *testing.T) {
	file, err := os.Open("testData.txt")
	if err != nil {
		t.Fatal("failed to open file")
	}

	var totalSb, totalRune time.Duration

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		start := time.Now()
		runeRes, err := correctMistakesRune(sc.Text())
		if err != nil {
			t.Fatalf("error correcting mistake using rune func, err: %v", err)
		}
		runeFuncTime := time.Since(start)
		totalRune += runeFuncTime

		start = time.Now()
		sbRes, err := CorrectMistakes(sc.Text())
		if err != nil {
			t.Fatalf("error correcting mistake using sb func, err: %v", err)
		}
		sbFuncTime := time.Since(start)
		totalSb += sbFuncTime

		t.Logf("sbTime: %v\truneTime: %v\n", sbFuncTime, runeFuncTime)
		t.Logf("sbRes: %s\nruneRes: %s\n", sbRes, runeRes)

		if sbRes != runeRes {
			t.Fatal("results are different")
		}
	}

	t.Logf("Total sb time: %v\nTotal rune time: %v\n", totalSb, totalRune)
}
