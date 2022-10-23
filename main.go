package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"amadeuscommand/auth"
	_ "amadeuscommand/config"
	"amadeuscommand/cryptic"
	"amadeuscommand/helper"
)

func main() {
	sess, err := auth.NewAuth().GetSession()
	if err != nil {
		panic(err)
	}

	cmd := cryptic.NewCryptic(sess)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		scanner.Scan()
		command := scanner.Text()

		ret := helper.ProcCommand(command)
		if ret == helper.BREAK_CODE {
			break
		} else if ret == helper.CONTINUE_CODE {
			continue
		}

		command = strings.ToUpper(command)

		// command := "an20novtyobkk"
		result, err := cmd.GetCommandResult(command)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", *result)
	}
}
