package utils

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintIntro() {
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("-\tAutoRPG - An experiment in distributed systems and idle games\t    -")
	fmt.Println("-\tSimply create a character, and let the game play itself, battling,  -")
	fmt.Println("-\tcollecting gear and leveling up, all on its own.\t\t\t-")
	fmt.Println("-----------------------------------------------------------------------------")
}

func AwaitInput() {
	fmt.Println("\nPress enter key to continue...")
	fmt.Scanln()
}

func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func ReadInt() int {
	var num int
	fmt.Print("-> ")
	fmt.Scanf("%d", &num)

	return num
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
