package main

import (
	"fmt"
	"syscall/js"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var store []User = []User{
	{1, "user1"},
	{2, "user2"},
	{3, "user3"},
}

func getRecords() []User {
	return store
}
func addRecord(id int) {
	store = append(store, User{id, fmt.Sprintf("user%d", id)})
}

var window js.Value = js.Global()

func main() {

	go func() {
		id := len(store) + 1
		for {
			time.Sleep(1 * time.Second)
			addRecord(id)
			id++
		}
	}()

	window.Get("document").Call("getElementById", "btn_show").
		Call("addEventListener", "click", js.FuncOf(onShowButtonClick))
	window.Get("document").Call("getElementById", "btn_clear").
		Call("addEventListener", "click", js.FuncOf(onClearButtonClick))

	// To do not finish main
	select {}
}

func addElementByID(id string, elem string) {
	dom := window.Get("document").Call("getElementById", id)
	inner_html := dom.Get("innerHTML").String()
	inner_html = inner_html + elem
	dom.Set("innerHTML", inner_html)
}

func onShowButtonClick(js.Value, []js.Value) interface{} {
	showRecords()
	return nil
}
func onClearButtonClick(js.Value, []js.Value) interface{} {
	clearRecords()
	return nil
}

func showRecords() {
	clearRecords()
	users := getRecords()
	for _, user := range users {
		addElementByID("records", fmt.Sprintf("<tr><td>%d</td><td>%s</td></tr>", user.ID, user.Name))
	}
}

func clearRecords() {
	dom := window.Get("document").Call("getElementById", "records")
	dom.Set("innerHTML", "")
}
