package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type contextColor struct {
	R int
	G int
	B int
	A int
}

var colorStringMap = map[string]contextColor{"default": {255, 255, 255, 255}, "red": {255, 0, 0, 255}, "blue": {0, 0, 255, 255}, "dark blue": {0, 0, 139, 255}, "green": {0, 255, 0, 255}, "dark green": {0, 100, 0, 255}, "yellow": {255, 255, 0, 255}, "pink": {255, 192, 203, 255}, "magenta": {255, 0, 255, 255}, "orange": {255, 165, 0, 255}, "purple": {128, 0, 128, 255}, "brown": {165, 42, 42, 255}, "black": {0, 0, 0, 255}}

func GetRGBA(s string) contextColor {
	switch strings.ToLower(strings.Trim(s, " \n\t\r")) {
	case "red":
		return colorStringMap["red"]
	case "blue":
		return colorStringMap["blue"]
	case "dark blue":
		return colorStringMap["dark blue"]
	case "green":
		return colorStringMap["green"]
	case "dark green":
		return colorStringMap["dark green"]
	case "yellow":
		return colorStringMap["yellow"]
	case "pink":
		return colorStringMap["pink"]
	case "magenta":
		return colorStringMap["magenta"]
	case "orange":
		return colorStringMap["orange"]
	case "purple":
		return colorStringMap["purple"]
	case "brown":
		return colorStringMap["brown"]
	case "black":
		return colorStringMap["black"]
	default:
		return colorStringMap["default"]
	}
}

func GetRGBAInput(r *bufio.Reader) contextColor {

	fmt.Println("LoGo offers the Following Color Strings or you could enter a RGBA Color: ")
	fmt.Println("Red, Blue, Dark Blue, Green, Dark Green, Yellow, Pink, Magenta, Orange, Purple, Brown, Black, RGBA for Entering a RGBA String")
	fmt.Printf("Enter Color: ")
	str, err := r.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: Unkown Color %s for bg, Try Using RGBA Color", str)
		return colorStringMap["default"]
	}
	switch strings.ToLower(strings.Trim(str, " \n\t\r")) {
	case "red":
		return colorStringMap["red"]
	case "blue":
		return colorStringMap["blue"]
	case "dark blue":
		return colorStringMap["dark blue"]
	case "green":
		return colorStringMap["green"]
	case "dark green":
		return colorStringMap["dark green"]
	case "yellow":
		return colorStringMap["yellow"]
	case "pink":
		return colorStringMap["pink"]
	case "magenta":
		return colorStringMap["magenta"]
	case "orange":
		return colorStringMap["orange"]
	case "purple":
		return colorStringMap["purple"]
	case "brown":
		return colorStringMap["brown"]
	case "black":
		return colorStringMap["black"]
	case "rgba":
		fmt.Printf("Enter R of RGBA Color: ")
		R, _ := r.ReadString('\n')
		R_int, _ := strconv.Atoi(R)
		fmt.Printf("Enter G of RGBA Color: ")
		G, _ := r.ReadString('\n')
		G_int, _ := strconv.Atoi(G)
		fmt.Printf("Enter B of RGBA Color: ")
		B, _ := r.ReadString('\n')
		B_int, _ := strconv.Atoi(B)
		fmt.Printf("Enter A of RGBA Color: ")
		A, _ := r.ReadString('\n')
		A_int, _ := strconv.Atoi(A)
		return contextColor{R_int, G_int, B_int, A_int}
	}
	return colorStringMap["default"]
}
