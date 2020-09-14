package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

var msg string = "Press ENTER to begin"

func main() {
	fmt.Println("What is your name?")
	name := bufio.NewScanner(os.Stdin)
	name.Scan()

	fmt.Println("What is the server ip?")
	addr := bufio.NewScanner(os.Stdin)
	addr.Scan()

	fmt.Println("What is the port?")
	port := bufio.NewScanner(os.Stdin)
	port.Scan()

	connect("Alex", "localhost", "8081")
	// connect(name.Text(), addr.Text(), port.Text())
}

//Function that connects to the server and waits for a user
//input. Sends the input to the server and waits for a response.
func connect(name string, addr string, port string) {

	c, err := net.Dial("tcp", addr+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(msg)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		if len(message) > 5 && message[:5] == "ERROR" {
			fmt.Print(message)
			fmt.Println("Try again!\n")
			continue
		}		
		if len(message) > 3 && message[:3] == "WIN" {	
			clearTerminal()	
			fmt.Print(message[5:])			
			break
		}
		fmt.Print(formatTable(message))
		fmt.Println("")
		msg = "What position do you want to play?\n"
	}
}

//Clears the terminal
func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//Formats the table so it can look like a 
//tic_tac_toe table.
func formatTable(tbl string) string {
	clearTerminal()
	tbl = tbl[:12] + "\n" + tbl[12:]
	tbl = tbl[:26] + "\n" + tbl[26:]
	tbl = tbl[:39] + "\n" + tbl[39:]
	tbl = tbl[:53] + "\n" + tbl[53:]
	return tbl
}
