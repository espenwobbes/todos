package flags

import "flag"

type CmdFlags struct {
	Add           string
	Del           int
	Edit          string
	Toggle        int
	Untoggle      int
	UpdateDefault string
	Completed     bool
	Waiting       bool
	List          bool
	Flush         bool
	FlushC        bool
}

func NewFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "add a new todo")
	flag.IntVar(&cf.Del, "del", -1, "delete a todo")
	flag.StringVar(&cf.Edit, "edit", "", "edit a todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "mark a todo as completed")
	flag.IntVar(&cf.Untoggle, "untoggle", -1, "mark a todo as not completed (used after toggle)")
	flag.StringVar(&cf.UpdateDefault, "update-default", "", "change the default behaviour when no argument is specified")
	flag.BoolVar(&cf.Completed, "completed", false, "list all the completed todos")
	flag.BoolVar(&cf.Waiting, "waiting", false, "list all the non-completed todos")
	flag.BoolVar(&cf.List, "list", false, "list all the todos")
	flag.BoolVar(&cf.Flush, "flush", false, "deletes all todos")
	flag.BoolVar(&cf.FlushC, "del-completed", false, "deletes all the completed todos")

	flag.Parse()

	return &cf
}
