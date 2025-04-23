package utils

type Command int

var CommandNames = []string{"Add", "List", "Done", "Delete"}

const (
	Add Command = iota
	List
	Done
	Delete
)

func AddTask() {

}

func ListTask() {

}

func SetComplete() {

}

func DeleteTask() {

}
