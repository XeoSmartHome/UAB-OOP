package main

import "fmt"

type Memento struct {
	state string
}

type Word struct {
	word string
}

func (w *Word) SetWord(word string) {
	w.word = word
}

func (w *Word) GetWord() string {
	return w.word
}

func (w *Word) SaveState() Memento {
	return Memento{
		state: w.word,
	}
}

func (w *Word) RestoreState(memento Memento) {
	w.word = memento.state
}

type Caretaker struct {
	undoStack []Memento
	redoStack []Memento
}

func (c *Caretaker) SaveState(memento Memento) {
	fmt.Println("State saved")
	c.undoStack = append(c.undoStack, memento)
	c.redoStack = nil
}

func (c *Caretaker) Undo(word *Word) {
	if len(c.undoStack) == 0 {
		fmt.Println("No states to undo")
		return
	}

	// Get the last state from the undo stack
	c.undoStack = c.undoStack[:len(c.undoStack)-1]
	m := c.undoStack[len(c.undoStack)-1]

	// Save the current state to the redo stack
	c.redoStack = append(c.redoStack, word.SaveState())

	// Restore the state from the memento
	word.RestoreState(m)

}

func (c *Caretaker) Redo(word *Word) {
	if len(c.redoStack) == 0 {
		fmt.Println("No states to redo")
		return
	}

	// Get the last state from the redo stack
	m := c.redoStack[len(c.redoStack)-1]
	c.redoStack = c.redoStack[:len(c.redoStack)-1]

	// Save the current state to the undo stack
	c.undoStack = append(c.undoStack, word.SaveState())

	// Restore the state from the memento
	word.RestoreState(m)
}

func main() {
	word := &Word{}

	caretaker := &Caretaker{}

	word.SetWord("First word")
	caretaker.SaveState(word.SaveState())
	fmt.Printf("Current word: %s\n", word.GetWord())

	word.SetWord("Second word")
	caretaker.SaveState(word.SaveState())
	fmt.Printf("Current word: %s\n", word.GetWord())

	word.SetWord("Third word")
	caretaker.SaveState(word.SaveState())
	fmt.Printf("Current word: %s\n", word.GetWord())

	fmt.Printf("Undo\n")
	caretaker.Undo(word)
	fmt.Printf("Current word: %s\n", word.GetWord())

	fmt.Printf("Undo\n")
	caretaker.Undo(word)
	fmt.Printf("Current word: %s\n", word.GetWord())

	fmt.Printf("Redo\n")
	caretaker.Redo(word)
	fmt.Printf("Current word: %s\n", word.GetWord())

	fmt.Printf("Redo\n")
	caretaker.Redo(word)
	fmt.Printf("Current word: %s\n", word.GetWord())
}
