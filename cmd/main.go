package main

import (
	"Atestproject/pkg/command"
	"Atestproject/pkg/environment"
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type reader *bufio.Reader

func readLine(reader *bufio.Reader) (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return text, err
	}
	text = strings.TrimSuffix(text, "\n")
	return text, nil
}

func readLocation(reader *bufio.Reader) (command.Location, error) {
	text, err := readLine(reader)
	if err != nil {
		return command.Location{}, err
	}
	splitted := strings.Split(text, " ")
	if len(splitted) != 2 {
		return command.Location{}, errors.New("Two inputs are needed for location.")
	}

	x, err := strconv.ParseFloat(splitted[0], 64)
	if err != nil {
		return command.Location{}, err
	}

	y, err := strconv.ParseFloat(splitted[1], 64)
	if err != nil {
		return command.Location{}, err
	}

	return command.Location{x, y}, nil
}

func readCommand(reader *bufio.Reader) (string, error) {
	text, err := reader.ReadString(' ')
	if err != nil {
		return text, err
	}
	return text, nil
}

func readInt(reader *bufio.Reader) (int, error) {
	text, err := readLine(reader)
	if err != nil {
		return 0, err
	}
	text = strings.Replace(text, " ", "", -1)
	res, err := strconv.Atoi(text)
	return res, err
}

func readFloat(reader *bufio.Reader) (float64, error) {
	text, err := readLine(reader)
	if err != nil {
		return 0, err
	}
	res, err := strconv.ParseFloat(text, 64)
	return res, err
}

func main() {
	fmt.Print("Hello, ")
	var numAgents int
	var err error
	var home command.Location
	var waitTime time.Duration
	reader := bufio.NewReader(os.Stdin)
	for { // Reading number of agents
		fmt.Print("Please enter the number of agents i[int]: ")
		// Since we added cycle in commands there is no need for multiple go routines to handle multiple agents
		numAgents, err = readInt(reader)
		if err != nil {
			fmt.Println("Wrong input with error ", err)
			continue
		}
		break
	}

	for { // Reading Location of home
		fmt.Print("Please enter the location of home x y[float float]: ")
		home, err = readLocation(reader)
		if err != nil {
			fmt.Println("Wrong input with error ", err)
			continue
		}
		break
	}

	for { //Reading agent wait time
		fmt.Print("Please enter wait time for agent in seconds t[int]: ")
		waitSecond, err := readInt(reader)
		if err != nil {
			fmt.Println("Wrong input with error", err)
			continue
		}
		waitTime = time.Duration(waitSecond) * time.Second
		break
	}
	env := environment.New(numAgents, waitTime, home)
	fmt.Println("This is our environment, ", env)
	fmt.Println("Please enter commands, write help for help")
	for {
		command, err := readCommand(reader)
		if err != nil {
			fmt.Println("Wrong command in input.", err)
		}
		switch command {
		case "help":
			fmt.Println("for adding location: location x y")
			fmt.Println("for cycling time: cycle t")
		case "cycle":
			t, err := readFloat(reader)
			if err != nil {
				fmt.Println("invalid input", err)
			}
			tm := time.Duration(int64(math.Ceil(t * float64(time.Second))))
			env.Cycle(tm)
		case "location":
			l, err := readLocation(reader)
			if err != nil {
				fmt.Println("invalid input", err)
				env.AddNewLocation(l)
			}
		}
	}
}
