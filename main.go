package main

import (
	. "./entities"
	. "./graph"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	server := new(Server).Initialize()
	client := new(Client).Initialize(len(server.Graph))

	len := len(client.EdgesArr)
	for i := 0; i < len; i++ {
		server.GeneratedR()
		server.EncryptNodes()

		client.UpdateEncryptNodes(server.CreatedZu())
		edge := client.GetRandEdge()

		print("#" + strconv.Itoa(edge) + " ")
		c1, c2 := server.GetDecryptKey(edge)
		client.CheckColors(edge, c1, c2)
	}
}

func CreatedGraph() {
	graph := [][]Node{
		{Node{
			Edge:       0,
			Color:      0,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}, Node{
			Edge:       0,
			Color:      2,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}},
		{Node{
			Edge:       1,
			Color:      2,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}, Node{
			Edge:       1,
			Color:      0,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}},
		{Node{
			Edge:       2,
			Color:      0,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}, Node{
			Edge:       2,
			Color:      1,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}},
		{Node{
			Edge:       3,
			Color:      2,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}, Node{
			Edge:       3,
			Color:      1,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}},
		{Node{
			Edge:       4,
			Color:      2,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}, Node{
			Edge:       4,
			Color:      0,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}},
		{Node{
			Edge:       5,
			Color:      1,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}, Node{
			Edge:       5,
			Color:      0,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}},
		{Node{
			Edge:       6,
			Color:      0,
			R:          nil,
			PrivateKey: nil,
			PublicKey:  nil,
		}, Node{
			Edge:       6,
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
