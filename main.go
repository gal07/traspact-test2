package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Menu Aplikasi : \n")
	fmt.Print("1. Pencarian kalimat tertentu \n")
	fmt.Print("2. Ganti kalimat \n")
	fmt.Print("3. Urutkan Kalimat Berdasarkan Abjad \n")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		SearchSentences()
		break
	case "2":
		ReplaceSentences()
		break
	case "3":
		SortingSentences()
		break
	default:
		fmt.Print("Out of syntax")
		break
	}

}

func SortingSentences() {
	fmt.Print("SortingSentences")
}

func ReplaceSentences() {

	fmt.Print("Masukan kata yang ingin diganti : \n")
	var target string
	fmt.Scanln(&target)

	fmt.Print("Masukan kata pengganti : \n")
	var replace string
	fmt.Scanln(&replace)

	r, err := os.Open("og_article.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	w, err := os.Create("replaceword.txt")
	if err != nil {
		panic(err)
	}
	defer w.Close()
	w.ReadFrom(r)

	repl, err := os.ReadFile("replaceword.txt")

	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	tr := []byte(target)
	rp := []byte(replace)

	repl = bytes.Replace(repl, tr, rp, -1)
	if err = os.WriteFile("replaceword.txt", repl, 0666); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Replace Berhasil Lihat file replaceword.txt \n")
	fmt.Print("Program selesai \n")

}

func SearchSentences() {

	fmt.Print("Masukan kata yang ingin dicari : \n")
	var input string
	fmt.Scanln(&input)

	// initiate file-handle to read from
	fileHandle, err := os.Open("og_article.txt")

	// check if file-handle was initiated correctly
	if err != nil {
		panic(err)
	}

	// make sure to close file-handle upon return
	defer fileHandle.Close()

	// initiate scanner from file handle
	fileScanner := bufio.NewScanner(fileHandle)

	// tell the scanner to split by words
	fileScanner.Split(bufio.ScanWords)

	// initiate counter
	stringFounf := 0

	// for looping through results
	for fileScanner.Scan() {
		if fileScanner.Text() == input {
			stringFounf++
		}
	}

	// check if there was an error while reading words from file
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}

	// print total word count
	fmt.Printf("total kata yang dicari : '%d' \n", stringFounf)
	fmt.Printf("Program selesai")
}
