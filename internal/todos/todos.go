package todos

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

var TIMELAYOUT = "2006-01-02 15:04:05"

type Todo struct {
	Name        string
	CompletedAt *string
	CreatedAt   string
	Completed   bool
}

type Todos []Todo

func NewTablePrint() *table.Table {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Completed At")

	return table
}

func (tds *Todos) Add(title string) {
	if ok := tds.CheckTitle(title); !ok {
		fmt.Printf("title %v is already in use", title)
		return
	}

	todo := Todo{
		Name:        title,
		CreatedAt:   time.Now().Format(TIMELAYOUT),
		CompletedAt: nil,
		Completed:   false,
	}
	*tds = append(*tds, todo)
	
	table := NewTablePrint()
	table.AddRow(strconv.Itoa(len(*tds)), todo.Name, "❌", "")
	
	table.Render()
}

func (tds *Todos) CheckTitle(title string) bool {

	for _, todo := range *tds {
		if todo.Name == title {
			return false
		}
	}

	return true
}

func (tds *Todos) CheckIndex(index int) bool {
	if 0 <= index && index <= len(*tds) {
		return true
	}
	return false
}

func (tds *Todos) Delete(index int) {
	if ok := tds.CheckIndex(index); !ok {
		fmt.Println("index out of bounds...")
		os.Exit(1)
	}

	t := *tds // adds every item up to index, and after index to a new array
	*tds = append(t[:index], t[index+1:]...)
}

func (tds *Todos) Toggle(index int, enum int) {
	if ok := tds.CheckIndex(index); !ok {
		fmt.Println("index out of bounds...")
		os.Exit(1)
	}

	t := *tds

	todo := &t[index]

	if enum == Untoggle {
		tds.untoggle(todo)
		return
	}

	if todo.Completed {
		fmt.Println("todo already completed, use -untoggle if you want to mark todo as not completed")
		return
	}

	todo.Completed = true
	time := time.Now().Format(TIMELAYOUT)
	todo.CompletedAt = &time
}

func (tds *Todos) untoggle(todo *Todo) {
	if todo.Completed {
		todo.Completed = false
		todo.CompletedAt = nil
	}else {
		fmt.Println("todo is not completed yet...")
	}
}

func (tds *Todos) Edit(index int, newTitle string) {
	if ok := tds.CheckIndex(index); !ok {
		fmt.Println("index out of bounds...")
		os.Exit(1)
	}

	if ok := tds.CheckTitle(newTitle); !ok {
		fmt.Printf("title %v is already in use", newTitle)
		os.Exit(1)
	}

	t := *tds
	t[index].Name = newTitle
}

func (tds *Todos) Print(enum int) {
	table := NewTablePrint() 

	switch enum {

	case Completed:
		tds.addCompleted(table)

	case Waiting:
		tds.addWaiting(table)

	case All:
		tds.addAll(table)
	}

	table.Render()
}

// adds all completed todos to the table
func (tds *Todos) addCompleted(table *table.Table) {

	for index, todo := range *tds {
		cmpl := "✅"
		if todo.Completed {
			table.AddRow(strconv.Itoa(index), todo.Name, cmpl, *todo.CompletedAt)
		}
	}
}

// adds all non completed todos to the table
func (tds *Todos) addWaiting(table *table.Table) {

	for index, todo := range *tds {
		cmpl := "❌"
		if !todo.Completed {
			table.AddRow(strconv.Itoa(index), todo.Name, cmpl, "")
		}
	}
}

// ads all the todos to the table
func (tds *Todos) addAll(table *table.Table) {

	for index, todo := range *tds {
		cmpl := "❌"
		cmplAt := ""
		if todo.Completed {
			cmpl = "✅"
			cmplAt = *todo.CompletedAt
		}
		table.AddRow(strconv.Itoa(index), todo.Name, cmpl, cmplAt)
	}
}

func (tds *Todos) Flush(enum int) {
	switch enum {
	case FlushC:
		tds.flushCompleted()
	case FlushAll:
		tds.flushAll()
	}
}

func (tds *Todos) flushCompleted() {
	for index, todo := range *tds {
		if todo.Completed {
			tds.Delete(index)
			fmt.Println("sucessfully deleted: " + todo.Name)
		}
	}
}

func (tds *Todos) flushAll() {
	t := Todos{}
	*tds = t
}
