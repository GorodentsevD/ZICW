package main

import (
	. "./entities"
	"strconv"
)

func main() {
	server := new(Server).Initialize()
	//server.GenerateGraph(10000)
	client := new(Client).Initialize(len(server.Graph))

	len := len(client.EdgesArr)

	server.GeneratedR()
	server.EncryptNodes()
	for i := 0; i < len; i++ {

		client.UpdateEncryptNodes(server.CreatedZu())
		edge := client.GetRandEdge()

		print("#" + strconv.Itoa(edge) + " ")
		c1, c2 := server.GetDecryptKey(edge)
		client.CheckColors(edge, c1, c2)
	}
}
