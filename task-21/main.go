package main

import (
	"fmt"
	"math"
)

// Интерфейс и реализация цилиндра
type Cylinder interface {
	Radius() float64
}

type cylinder struct {
	radius float64
	height float64
}

func (c *cylinder) Radius() float64 {
	return c.radius
}

func (c *cylinder) Height() float64 {
	return c.height
}

func NewCylinder(radius, height float64) *cylinder {
	return &cylinder{radius: radius, height: height}
}

// Интерфейс и реализация куба
type Cube interface {
	Width() float64
}

type cube struct {
	width float64
}

func (s *cube) Width() float64 {
	return s.width
}

func NewCube(width float64) *cube {
	return &cube{width: width}
}

// Интерфейс и реализация круглого отверстия
type RoundHole interface {
	Fit(c Cylinder) bool
}

type roundHole struct {
	radius float64
}

func (ch *roundHole) Fit(c Cylinder) bool {
	if c.Radius() > ch.radius {
		return false
	}

	return true
}

func NewRoundHole(radius float64) *roundHole {
	return &roundHole{radius: radius}
}

// Адаптер куба к цилиндру
type CubeToCylinderAdapter struct {
	cube Cube
}

func NewCubeToCylinderAdapter(c Cube) *CubeToCylinderAdapter {
	return &CubeToCylinderAdapter{cube: c}
}

func (a *CubeToCylinderAdapter) Radius() float64 {
	return a.cube.Width() * math.Sqrt(2) / 2
}

/*
Паттерн адаптер позволяет делегировать вызовы между разными, несовместимыми интерфейсами.
Он применяется когда изменения интерфейсов нежелательны или невозможны (например: внешние сервисы)

Плюсы:
1. Лёгкая интеграция между интерфейсами
2. Легко реагировать на изменения какого либо из интерфейсов.

Минусы:
1. Адаптер требует написания и поддержания дополнительных структур.
2. Если возможно изменить интерфейс с малыми трудозатратами, то лучше его поменять, чем поддерживать дополнительную абстракцию.
*/

func main() {
	cb := NewCube(5)
	cl := NewCylinder(3.5, 1)

	h := NewRoundHole(3.5)

	fmt.Println("Does the cylinder fit the round hole: ", h.Fit(cl))

	ad := NewCubeToCylinderAdapter(cb)
	fmt.Println("Does the cube fit the round hole: ", h.Fit(ad))
}
