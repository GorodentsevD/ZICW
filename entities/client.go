package entities

import (
	. "../color"
	. "../graph"
	"fmt"
	google_coloring "github.com/gookit/color"
	"math/big"
)

type Client struct {
	EdgesArr     []int
	EncryptNodes [][]EncryptNode
}

func (client *Client) Initialize(size int) *Client {
	for i := 0; i < size; i++ {
		client.EdgesArr = append(client.EdgesArr, i)
	}

	return client
}

func (client *Client) GetRandEdge() int {
	edges := ShuffleColors(client.EdgesArr)
	edge := edges[0]
	client.EdgesArr = edges[1:]

	return edge
}

func (client *Client) CheckColors(edge int, c1 *big.Int, c2 *big.Int) {
	u := client.EncryptNodes[edge]
	u1 := u[0]
	IZu1 := new(big.Int).Exp(u1.Z, c1, u1.N)
	u2 := u[1]
	IZu2 := new(big.Int).Exp(u2.Z, c2, u2.N)
	color1 := GetColorFromR(IZu1)
	color2 := GetColorFromR(IZu2)

	var printColor1, printColor2 string
	red := google_coloring.FgRed.Render
	green := google_coloring.FgGreen.Render
	blue := google_coloring.FgBlue.Render

	if color1 == 0 {
		printColor1 = red("RED  ")
	} else if color1 == 1 {
		printColor1 = green("GREEN")
	} else {
		printColor1 = blue("BLUE ")
	}
	if color2 == 0 {
		printColor2 = red("RED  ")
	} else if color2 == 1 {
		printColor2 = green("GREEN")
	} else {
		printColor2 = blue("BLUE ")
	}

	var result string
	if color1 != color2 {
		result = "true"
	} else {
		result = "false"
	}

	fmt.Printf("color1: %5s color2: %5s check: %s\n", printColor1, printColor2, result)
}

func (client *Client) UpdateEncryptNodes(encryptNodes [][]EncryptNode) {
	client.EncryptNodes = encryptNodes
}
