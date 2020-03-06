package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var wheels []wheel

func main() {
	f, err := os.Open("wheels.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		w, err := parseWheel(scanner.Text())
		if err != nil {
			fmt.Print(w, err)
			break
		}
		fmt.Println(w.Name)
		wheels = append(wheels, w)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(wheels))

}

type wheel struct {
	Name        string
	Size        string
	Weight      float64
	Diameter    int
	Width       float64
	Manufacture string
}

func parseWheel(s string) (wheel, error) {
	parts := strings.Split(s, " ")
	sizes := strings.Split(parts[len(parts)-2], "x")
	w := wheel{
		Name:        strings.Join(parts[:len(parts)-3], " "),
		Size:        parts[len(parts)-2],
		Manufacture: parts[len(parts)-3],
	}

	weight, err := strconv.ParseFloat(parts[len(parts)-1], 64)
	if err != nil {
		return w, err
	}
	w.Weight = weight

	diam, err := strconv.Atoi(sizes[0])
	if err != nil {
		return w, err
	}
	w.Diameter = diam

	width, err := strconv.ParseFloat(sizes[1], 64)
	if err != nil {
		return w, err
	}
	w.Width = width

	return w, nil
}
