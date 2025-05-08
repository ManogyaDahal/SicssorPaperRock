package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"

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
    clearScreen()
    invalid := false 
    //Pressing q exits the program
    if userInput == 'q'{
      break
    }
    if userInput == 'c'{
      clearScreen()
    }
    switch userInput{
    case 's':
      userValue = 0
			printSicssors()
    case 'p':
      userValue = 1
			printPaper()
    case 'r':
      userValue = 2
			Printrock()
    default:
      invalid = true 
      fmt.Println("Enter a valid input") 
    }
    //If invlid is true this code will not work
    if !invalid {
      switch{
      case userValue == 0 && botValue == 1:
        userPoints += 1
				printPaper()
        fmt.Println("User: Sicssors and  bot: Paper") 
        fmt.Println("User Won") 

      case userValue == 1 && botValue == 2:
        userPoints += 1
				Printrock()
        fmt.Println("User: Paper and  bot: Rock") 
        fmt.Println("User Won") 

      case userValue == 2 && botValue == 0:
        userPoints += 1
				printSicssors()
        fmt.Println("User: Rock and  bot: Sicssors") 
        fmt.Println("User Won") 

      case botValue == 0 && userValue== 1:
        botPoints += 1
				printSicssors()
        fmt.Println("User: Paper and  bot: Siscssors") 
        fmt.Println("Bot Won") 

      case botValue == 1 && userValue == 2:
        botPoints += 1
				printPaper()
        fmt.Println("User: Rock and  bot: Paper") 
        fmt.Println("Bot Won") 

      case botValue == 2 && userValue == 0:
        botPoints += 1
				printSicssors()
        fmt.Println("User: Paper and  bot: Siscsors") 
        fmt.Println("Bot Won") 
      default:
        if userInput == 0 {
					printSicssors()
        fmt.Println("User: Siscors and  bot: Siscsors") 
        } else if userInput == 1 {
					printPaper()
        fmt.Println("User: Paper and  bot: Paper") 
        }
				Printrock()
        fmt.Println("User: Rock and  bot: Rock") 
        fmt.Println("Draw Case")
      }
    }
    showPoints(userPoints, botPoints)
  }
}

func greetings(){
  fmt.Println(
    "************************************\n",
    "Thank you for playing this game.\n",
    "Instructions for playing the game:\n",
    "-r (rock)\n-p (paper)\n-s (siscors)\n",
    "************************************",
    )
  fmt.Println("Press 'y' to continue to the game or 'q' to exit...")
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

func printSicssors(){
fmt.Println(

"    ________ \n",
"---'   ____)____\n",
"          ______)\n",
"       __________)\n",
"      (____)\n",
"---.__(___)\n",
"                   ",
		)
}
func printPaper(){
fmt.Println(

"     _______\n",
"---'    ____)____\n",
"           ______)\n",
"          _______)\n",
"         _______)\n",
"---.__________)\n",
	"",
		)

}
func Printrock(){
fmt.Println(
"    _______\n",
"---'   ____)\n",
"      (_____)\n",
"      (_____)\n",
"      (____)\n",
"---.__(___) \n",
	)
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

func showPoints(userPoints, botPoints int){
      fmt.Printf("**************************\n")
      fmt.Println("Press 'q' to exit")
      fmt.Printf("Press 's', 'p' or 'r' to play next move")
      fmt.Printf("\n**************************\n")
      fmt.Println("*** Total points ****")
      fmt.Println("User points: ", userPoints)
      fmt.Println("bot points: " ,botPoints)
}

/* Make UI Tidy */
