package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
  "regexp"
	"strings"
  "math/rand"
)

func main() {
  if len(os.Args) < 2 {
		fmt.Println("Verwendung: program Eingangstext.txt")
		os.Exit(1)
	}

	// Dateipfade für Eingangstext und Phrasen
	inputFilePath := os.Args[1]
	// Dateipfade für Eingangstext und Phrasen
	phrasesFilePath := "phrasen.txt"

	// Eingangstext lesen
	inputText, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("Fehler beim Lesen der Eingangstextdatei: %v", err)
	}

	// Phrasen lesen
	phrases, err := readPhrases(phrasesFilePath)
	if err != nil {
		log.Fatalf("Fehler beim Lesen der Phrasendatei: %v", err)
	}

	// Satzanfänge ersetzen
	outputText := replaceSentenceBeginnings(string(inputText), phrases)

	// Ergebnis ausgeben
	fmt.Println(outputText)
}

func readPhrases(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var phrases []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		phrases = append(phrases, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return phrases, nil
}

func replaceSentenceBeginnings(text string, phrases []string) string {
	// Verwende regulären Ausdruck, um Sätze zu finden
	re := regexp.MustCompile(`[.!?]\s*`)
	sentences := re.Split(text, -1)

	for i, sentence := range sentences {
		// Wähle eine zufällige Phrase aus der Liste der Phrasen
		randomPhrase := phrases[rand.Intn(len(phrases))]

    if len(strings.Fields(sentence)) >0 {
      sentences[i] = randomPhrase + sentence[len(strings.Fields(sentence)[0]):] + ". "
    }
	}

	return strings.Join(sentences, "")
}