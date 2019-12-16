package entities

import (
	. "../color"
	. "../graph"
	rsa "../sign/signature_rsa"
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/big"
	rand2 "math/rand"
	"os"
)

type Server struct {
	Graph  [][]Node
	Colors []int
}

func (server *Server) GenerateGraph(num int) {
	var array [][]Node
	for i := 0; i < num; i++ {
		node := []Node{Node{
			Edge:       i,
			Color:      generateColor(),
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}, Node{
			Edge:       i,
			Color:      generateColor(),
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}}

		array = append(array, node)
	}
	toJson := func(any interface{}) []byte {
		bytes, _ := json.Marshal(any)
		return bytes
	}
	_ = ioutil.WriteFile("graph.json", toJson(array), os.ModePerm)
}

func CreatedGraph() {
	graph := [][]Node{
		{Node{
			Edge:       0,
			Color:      0,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		},
			Node{
				Edge:       0,
				Color:      2,
				R:          nil,
				PrivateKey: nil,
				PublicKey:  nil,
			}},
	}
	toJson := func(any interface{}) []byte {
		bytes, _ := json.Marshal(any)
		return bytes
	}
	_ = ioutil.WriteFile("graph.json", toJson(graph), os.ModePerm)
}

func generateColor() int {
	r := rand2.Int63n(4)
	if r == 1 {
		return 0
	} else if r == 2 {
		return 1
	} else {
		return 2
	}
}

func (server *Server) Initialize() *Server {
	gr, _ := ioutil.ReadFile("graph.json")
	graphFromJson := func(any []byte) [][]Node {
		var graph [][]Node
		if err := json.Unmarshal(any, &graph); err != nil {
			panic(err)
		}
		return graph
	}
	server.Graph = graphFromJson(gr)
	server.Colors = []int{0, 1, 10} //Red,Green,Blue
	return server
}

func (server *Server) Recolor() {
	server.Colors = ShuffleColors(server.Colors)
}

func (server *Server) GeneratedR() {
	server.Recolor()
	for i := 0; i < len(server.Graph); i++ {
		for j := 0; j < len(server.Graph[i]); j++ {
			color := server.Colors[server.Graph[i][j].Color]
			c, _ := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(16), nil))
			r := PlanColorBit(c, color)
			server.Graph[i][j].R = r
		}
	}
}

func (server *Server) EncryptNodes() {
	for i := 0; i < len(server.Graph); i++ {
		for j := 0; j < len(server.Graph[i]); j++ {
			server.Graph[i][j].PrivateKey, server.Graph[i][j].PublicKey = rsa.CreateKeys()
		}
	}
}

func (server Server) CreatedZu() [][]EncryptNode {
	var encryptArrayNodes [][]EncryptNode
	for i := 0; i < len(server.Graph); i++ {
		var encryptNodes []EncryptNode
		for j := 0; j < len(server.Graph[i]); j++ {
			encryptNodes = append(encryptNodes, EncryptNode{
				N: server.Graph[i][j].PublicKey.N,
				D: server.Graph[i][j].PublicKey.D,
				Z: new(big.Int).Exp(server.Graph[i][j].R,
					server.Graph[i][j].PublicKey.D,
					server.Graph[i][j].PublicKey.N),
				Edge: server.Graph[i][j].Edge,
			})
		}
		encryptArrayNodes = append(encryptArrayNodes, encryptNodes)
	}
	return encryptArrayNodes
}

func (server Server) GetDecryptKey(edge int) (*big.Int, *big.Int) {
	nodes := server.Graph[edge]
	return nodes[0].PrivateKey.C, nodes[1].PrivateKey.C
}
