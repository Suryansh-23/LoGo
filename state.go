package main

import "image"

type History struct {
	undoStck   [](image.Image)
	redoStck   [](image.Image)
	undoLength int
	redoLength int
}

func NewHistory() *History {
	return &History{undoStck: [](image.Image){}, redoStck: [](image.Image){}, undoLength: 0}
}

func (h *History) Push(cxt image.Image) {
	h.undoStck = append(h.undoStck, cxt)
	h.undoLength += 1
}

func (h *History) Undo() image.Image {
	if h.undoLength != 0 {
		tmp := h.undoStck[h.undoLength-1]
		arr := h.undoStck[:h.undoLength-1]
		h.redoStck = append(h.redoStck, tmp)
		h.undoStck = arr
		h.undoLength -= 1
		h.redoLength += 1
		return tmp
	}
	return nil
}

func (h *History) Redo() image.Image {
	if h.redoLength != 0 {
		tmp := h.redoStck[h.redoLength-1]
		arr := h.redoStck[:h.redoLength-1]
		h.undoStck = append(h.redoStck, tmp)
		h.redoStck = arr
		h.redoLength -= 1
		h.undoLength += 1
		return tmp
	}
	return nil
}

func (h *History) UndoPeek() image.Image {
	return h.undoStck[h.undoLength-1]
}
