package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"	
    "math/rand"
)

var pos = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
var start bool
var plays int
var played []int
var winCombo1 = []int{0, 3, 6, 0, 1, 2, 0, 2}
var winCombo2 = []int{1, 4, 7, 3, 4, 5, 4, 4}
var winCombo3 = []int{2, 5, 8, 6, 7, 8, 8, 6}

func main() {

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		c.Write([]byte(setPlay(c, string(netData))))
	}
}

func setTable() string {
	line1 := "  " + pos[0] + " | " + pos[1] + " | " + pos[2] + " "
	line2 := "  " + pos[3] + " | " + pos[4] + " | " + pos[5] + " "
	line3 := "  " + pos[6] + " | " + pos[7] + " | " + pos[8] + "\n"
	lineBreak := " ----------- "
	return line1 + lineBreak + line2 + lineBreak + line3
}

func setPlay(c net.Conn, data string) string {	
	if !start {
		start = true
		return setTable()
	}
	conv, err := strconv.Atoi(string(data[0]))
	if err != nil || conv == 0 {
		return "ERROR: Must enter a number from 1 - 9!\n"
	}
	if contains(conv) {
		return "ERROR: That position already taken... Try again!\n"
	}
	pos[conv-1] = "x"
	plays++
	played = append(played, conv)	
	if verifyWin() == "player" {
		return "WIN: YOU WIN!\n"
	} else if verifyWin() == "tied" {
		return "WIN: TIED!!!! - NOBODY WINS!\n"
	}
	if chkWinMove() {
		return "WIN: YOU LOSE!\n"
	}
	if chkLoseMove(){
		return setTable()
	}
	randomMove()	
	return setTable()
}

//Loops through the available positions to see is if there is any
//wining moves. If so then wins the game.
func chkWinMove() bool {
	for i, v := range pos {
		if v == " "{
			pos[i] = "o"
			if verifyWin() == "computer"{
				return true
			} 
			pos[i] = " "			
		}		
	}
	return false
}

//Checks to see is there is any available position that can lose
//the game. If it finds one then it takes it.
func chkLoseMove() bool{
	for i, v := range pos {
		if v == " "{
			pos[i] = "x"
			if verifyWin() == "player"{
				pos[i] = "o"
				plays++
				played = append(played, i+1)
				return true
			} 
			pos[i] = " "			
		}		
	}
	return false
}

//Chooses a random available position.
func randomMove() {
	for {
		v := rand.Intn(9 - 1) + 1
		if pos[v-1] == " " {
			pos[v-1] = "o"
			plays++
			played = append(played, v)
			break
		}
	}
}

//Goes through all the taken positions to see if anyone has won. If all
//positions are taken then the game is a tied.
func verifyWin() string {	
	for i := 0; i < 8; i++ {
		if pos[winCombo1[i]] == "x" && pos[winCombo2[i]] == "x" && pos[winCombo3[i]] == "x" {
			return "player"
		} else if pos[winCombo1[i]] == "o" && pos[winCombo2[i]] == "o" && pos[winCombo3[i]] == "o" {
			return "computer"
		} else if plays > 8 {
			return "tied"
		}
	}
	return "None"
}

//Helper function that checks if the position being played 
//has not already been played.
func contains(val int) bool {
	for _, p := range played {
		if p == val {
			return true
		}
	}
	return false
}