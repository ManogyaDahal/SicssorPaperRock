package main

import (
	"fmt"
	"math/rand"
  "os/exec"
  "os"

	"github.com/eiannone/keyboard"
)
const(
  siscors = iota
  paper
  rock
)

func main(){
  greetings();
  var userPoints, botPoints int
  for{
    botValue := randomGen(); 
    userValue := 99 //for invalid key press
    userInput := readingSingleKey() 
    invalid := false 
    //Pressing q exits the program
    if userInput == 'q'{
      break
    }
    if userInput == 'c'{
      clearScreen()
    }
    if userInput == 't'{
      clearScreen()
      fmt.Println("User points: ", userPoints)
      fmt.Println("bot points: " ,botPoints)
    }

    switch userInput{
    case 's':
      userValue = 0
    case 'p':
      userValue = 1
    case 'r':
      userValue = 2
    default:
      invalid = true 
      fmt.Println("Enter a valid input") 
    }
    //If invlid is true this code will not work
    if !invalid {
      switch{
      case userValue == 0 && botValue == 1:
        userPoints += 1
        fmt.Println("User Won") 

      case userValue == 1 && botValue == 2:
        userPoints += 1
        fmt.Println("User Won") 

      case userValue == 2 && botValue == 0:
        userPoints += 1
        fmt.Println("User Won") 

      case botValue == 0 && userValue== 1:
        botPoints += 1
        fmt.Println("Bot Won") 

      case botValue == 1 && userValue == 2:
        botPoints += 1
        fmt.Println("Bot Won") 

      case botValue == 2 && userValue == 0:
        botPoints += 1
        fmt.Println("Bot Won") 
      default:
        fmt.Println("Draw Case")
      }
    }
  }
}

func greetings(){
  fmt.Println(
    "Thank you for playing this game.\n",
    "Instructions for playing the game:\n",
    "-r (rock), -p (paper), -s (siscors)\n",
    )
  fmt.Println("Press 'y' to continue to the game or 'q' to exit.........")
  for{
  char := readingSingleKey()
    if char == 'y'{
      clearScreen()
      fmt.Println("***Game Start***")
      break
    } else if char == 'q'{
      os.Exit(0)
    } else{
      fmt.Println("Please Press y to continue or q to exit only ")
    }
  }
}

func randomGen()int{
  return rand.Intn(3)
}

func readingSingleKey() rune{
  char, _, err := keyboard.GetSingleKey()
  if (err != nil) {
    panic(err)
  }
  return char
}

func clearScreen(){
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}


/*
  things to implement:
    - clearing the screen  (done)
    - displaying the points on the screen (done:after pressing t)
    - displaying the bot's rock paper or sicssors on the screen
      {idea: making random gen func return string too while returning the random value}
    - refactoring the code into functions
    -Making ui tidy 
*/
