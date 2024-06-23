package main

import (
	"bufio"
	"os"
	"os/exec"
)

var maze []string

func main() {
	initialize()
	defer cleanup()
	err := loadMaze("mazes/maze01.txt")
	if err != nil {
		println("Error loading maze:", err)
		return
	}

	for {
		printMaze()

		input, err := readInput()
		if err != nil {
			println("Error reading input:", err)
			break
		}

		if input == "ESC" {
			break
		}
		
	}

}

func initialize() {
	// Activate cbreak mode
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		println("Unable to activate cbreak mode:", err)
		return
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)
	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}

	return "", nil
}

func cleanup() {
	// Deactivate cbreak mode
	cbTerm := exec.Command("stty", "-cbreak", "echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		println("Unable to deactivate cbreak mode:", err)
		return
	}
}

func loadMaze(mazeFile string) error {
	file, err := os.Open(mazeFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}

	return nil
}

func printMaze() {
	for _, row := range maze {
		println(row)
	}
}