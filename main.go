package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/gcjbr/gopacman/ansi"
)

type sprite struct {
	row, col int
}

var maze []string
var player sprite

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

		movePlayer(input)

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
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
		}
	}
}

	return "", nil
}


func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)
}

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
			newRow = newRow - 1
			if newRow < 0 {
					newRow = len(maze) - 1
			}
	case "DOWN":
			newRow = newRow + 1
			if newRow == len(maze) {
					newRow = 0
			}
	case "RIGHT":
			newCol = newCol + 1
			if newCol == len(maze[0]) {
					newCol = 0
			}
	case "LEFT":
			newCol = newCol - 1
			if newCol < 0 {
					newCol = len(maze[0]) - 1
			}
	}

	if maze[newRow][newCol] == '#' {
			newRow = oldRow
			newCol = oldCol
	}

	return
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

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player.row = row
				player.col = col
			}
		}
	}

	return nil
}

func printMaze() {
	ansi.ClearScreen()
	for _, line := range maze {
		for _, chr := range line {
				switch chr {
				case '#':
						fmt.Printf("%c", chr)
				default:
						fmt.Print(" ")
				}
		}
		fmt.Println()
}

ansi.MoveCursor(player.row, player.col)
fmt.Print("😺")

// Move cursor outside of maze drawing area
ansi.MoveCursor(len(maze)+1, 0)
}

