package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	fmt.Println(records)
	nodes := make([]*Node, len(records))
	if len(nodes) == 0 {
		return nil, nil
	}
	for _, record := range records {
		if nodes[record.ID] != nil {
			return nil, errors.New("duplicate node")
		}
		if len(nodes) <= record.ID {
			return nil, errors.New("wrong tree structure")
		}
		if record.ID <= record.Parent && record.ID != 0 {
			return nil, errors.New("parent ID must be lower then node ID")
		}
		newNode := &Node{ID: record.ID}
		nodes[record.ID] = newNode
		if newNode.ID == 0 {

		} else {
			parentNode := nodes[record.Parent]
			fmt.Println(newNode, parentNode)
			if parentNode.Children == nil {
				parentNode.Children = []*Node{}
			}
			parentNode.Children = append(parentNode.Children, newNode)
			sort.Slice(parentNode.Children, func(i, j int) bool {
				return parentNode.Children[i].ID < parentNode.Children[j].ID
			})
			// }
			// for i := len(parentNode.Children) - 1; i > 0; i-- {
			// 	if parentNode.Children[i].ID < parentNode.Children[i-1].ID {
			// 		t := parentNode.Children[i-1]
			// 		parentNode.Children[i-1] = parentNode.Children[i]
			// 		parentNode.Children[i] = t
			// 	}
			// }
		}
		//fmt.Println(nodes)
	}
	//fmt.Println(nodes)
	return nodes[0], nil
}
