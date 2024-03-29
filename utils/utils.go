package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Printf("failed to execute 'clear screen' command: %v", err.Error())
	}
}

func PrintIntro() {
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("-\tAutoRPG - An experiment in software engineering and idle games\t    -")
	fmt.Println("-\tSimply create a character, and let the game play itself, battling,  -")
	fmt.Println("-\tcollecting gear and leveling up, all on its own.\t\t\t-")
	fmt.Println("-----------------------------------------------------------------------------")
}

func AwaitInput() {
	fmt.Println("\nPress enter key to continue...")
	fmt.Scanln()
}

func ReadString(io io.Reader) (string, error) {
	reader := bufio.NewReader(io)
	fmt.Print("-> ")

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	text = strings.Replace(text, "\n", "", -1)

	return text, nil
}

func ReadInt(io io.Reader) (int, error) {
	str, err := ReadString(io)
	if err != nil {
		return -1, err
	}
	str = strings.TrimSpace(str)

	num, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}

	return num, nil
}

func GetRandomNumberInRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	if min < 0 {
		min = 0
	}

	if max < 0 {
		max = 0
	}

	return rand.Intn(max-min+1) + min
}

func GetRandomFloatInRange(min, max float32) float32 {
	rand.Seed(time.Now().UnixNano())

	if min < 0 {
		min = 0
	}

	if max < 0 {
		max = 0
	}

	return float32(math.Round(float64(rand.Float32()+max-min)*100) / 100)
}
