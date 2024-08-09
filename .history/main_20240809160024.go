package main

type todo struct {
	ID        string 'jso'
	Item      string 'jso'
	Completed bool 'jso'
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}
