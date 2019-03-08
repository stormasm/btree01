package btree

import (
	"fmt"
	"strings"
)

func (n *node) createNodeName() string {
	num := len(n.items)
	itemNames := make([]string, num)
	for i := 0; i < num; i++ {
		item := n.items[i]
		itemName := item.(Int)
		itemNameStr := fmt.Sprintf("%d", itemName)
		itemNames[i] = itemNameStr
	}
	itemNamesStr := strings.Join(itemNames, "-")
	return itemNamesStr
}

func (n *node) printChildren() bool {
	fmt.Println("numofchildren=", len(n.children))
	nodeName := n.createNodeName()
	fmt.Println("node name =", nodeName)
	return true
}
