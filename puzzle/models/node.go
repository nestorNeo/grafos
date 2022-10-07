package models

import "log"

type Direction uint8

const (
	Origin Direction = iota
	Right
	Left
	Down
	Up
)

type Node[T string | uint | int] struct {
	Values [][]T
	Level  uint
	Cost   int
	Move   Direction
	Cursor T
}

type Point struct {
	X int
	Y int
}

func (p *Point) InLimits(limit int) bool {
	return p.X >= 0 && p.X < limit && p.Y >= 0 && p.Y < limit
}

func (n *Node[T]) FindValue(value T) (int, int, error) {
	for i := 0; i < len(n.Values); i++ {
		for j := 0; j < len(n.Values[i]); j++ {
			if n.Values[i][j] == value {
				return i, j, nil
			}
		}
	}
	return -1, -1, &PuzzleProcessError{"unable to find directions for the value in the puzzle"}
}

func (n *Node[T]) ComputeDiffValuesWith(targetValues [][]T) (int, error) {
	diff := 0

	if len(n.Values) != len(targetValues) {
		return diff, &PuzzleProcessError{"values from source and target have not the same rows"}
	}

	for i := 0; i < len(n.Values); i++ {
		for j := 0; j < len(n.Values[i]); j++ {
			if n.Values[i][j] != targetValues[i][j] {
				diff++
			}
		}
	}

	return diff, nil

}

func (n *Node[T]) CopyAndShuffle(A, B Point) (out [][]T) {

	if !B.InLimits(len(n.Values)) {
		log.Println("not in limits omitting copy")
		return
	}

	out = make([][]T, len(n.Values))
	copy(out, n.Values)

	out[A.X][A.Y], out[B.X][B.Y] = out[B.X][B.Y], out[A.X][A.Y]
	return
}