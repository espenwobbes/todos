package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/espenwobbes/todo/internal/flags"
	"github.com/espenwobbes/todo/internal/storage"
	"github.com/espenwobbes/todo/internal/todos"
)

type TodoHandler struct {
	CmdFlags flags.CmdFlags
	Storage  storage.Storage[todos.Todos]
	Todos    todos.Todos
	Default  func(enum int)
}

func main() {
	cf := flags.NewFlags()
	fileName := os.Getenv("TODOS")
	storage := storage.NewStorage[todos.Todos](fileName)

	handler := TodoHandler{
		CmdFlags: *cf,
		Storage:  *storage,
		Todos:    make(todos.Todos, 0),
	}
	handler.Default = handler.Todos.Print

	if err := handler.InitStorage(); err != nil {
		fmt.Println("error while loading todos")
		os.Exit(1)
	}

	if err := handler.Run(); err != nil {
		fmt.Println("error while running: " + err.Error())
		os.Exit(1)
	}
}

func (h *TodoHandler) InitStorage() error {
	return h.Storage.Load(&h.Todos)
}

func (h *TodoHandler) Run() error {

	switch {
	// cases for listing out different todos
	case h.CmdFlags.List:
		h.Todos.Print(todos.All)
	case h.CmdFlags.Completed:
		h.Todos.Print(todos.Completed)
	case h.CmdFlags.Waiting:
		h.Todos.Print(todos.Waiting)

	// cases for deleting diffrent todos
	case h.CmdFlags.Flush:
		h.Todos.Flush(todos.FlushAll)
		h.sendStateUpdate(true)
	case h.CmdFlags.FlushC:
		h.Todos.Flush(todos.FlushC)
		h.sendStateUpdate(true)
	case h.CmdFlags.Del != -1:
		h.Todos.Delete(h.CmdFlags.Del)
		h.sendStateUpdate(true)

	// case for adding new todo
	case h.CmdFlags.Add != "":
		h.Todos.Add(h.CmdFlags.Add)
		h.sendStateUpdate(true)

	// case for editing a todo
	case h.CmdFlags.Edit != "":
		parts := strings.SplitN(h.CmdFlags.Edit, ":", 2)
		if len(parts) != 2 {
			return errors.New("wrong format in edit command, use format <index>:<newTitle>")
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		h.Todos.Edit(index, parts[1])
		h.sendStateUpdate(true)

	// case for toggling a todo as completed
	case h.CmdFlags.Toggle != -1:
		h.Todos.Toggle(h.CmdFlags.Toggle, todos.Toggle)
		h.sendStateUpdate(true)

	case h.CmdFlags.Untoggle != -1:
		h.Todos.Toggle(h.CmdFlags.Untoggle, todos.Untoggle)
		h.sendStateUpdate(true)

	

	default: h.Default(todos.All)
	}

	return h.Storage.Save(h.Todos)
}

func (h *TodoHandler) sendStateUpdate(changed bool) {
	h.Storage.StateChanged(changed)
}

func (h *TodoHandler) updateDefault(newFunc func(enum int)) {
	h.Default = newFunc
}
