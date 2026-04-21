package structspractice

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go.com/structspractice/note"
	"go.com/structspractice/todo"
)

func StartPractice() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = OutputData(todo)

	if err != nil {
		return
	}

	OutputData(userNote)
}

func printSomething(value interface{}) { //any kind of value
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	default:
		fmt.Println(value)
	}
}

func checkTypeAndIncrement(value interface{}) {
	typedVal, ok := value.(int)

	if ok {
		typedVal += 1
	}
}

func OutputData(data Outputtable) error {
	data.Display()
	return SaveData(data)
}

func SaveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Printf("Saving the %v failed!\n", data.Name())
		return err
	}

	fmt.Printf("Saving the %v was successful\n", data.Name())
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

type Saver interface {
	Save() error
	Name() string
}

type Outputtable interface {
	Saver
	Display() // embedding interfaces
}
