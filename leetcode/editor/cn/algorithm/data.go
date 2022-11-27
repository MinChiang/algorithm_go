package main

import (
	"container/list"
	"errors"
	"strconv"
	"strings"
)

var (
	nilValues   = [...]string{"nil", "NIL", "null", "NULL", ""}
	nilValueMap map[string]struct{}
)

func init() {
	nilValueMap = make(map[string]struct{}, len(nilValues))
	for _, value := range nilValues {
		nilValueMap[value] = struct{}{}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(input string) (*TreeNode, error) {
	var root *TreeNode
	input = strings.ReplaceAll(input, " ", "")
	input = strings.TrimPrefix(input, "[")
	input = strings.TrimSuffix(input, "]")
	if input == "" {
		return root, nil
	}
	split := strings.Split(input, ",")
	if split[0] == "" {
		return root, nil
	}
	val, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, errors.New("can not parse int: " + split[0])
	}

	l := list.New()
	root = &TreeNode{Val: val}
	l.PushBack(root)
	index := 1
	for l.Len() != 0 {
		element := l.Front()
		node := l.Remove(element).(*TreeNode)

		if index == len(split) {
			return root, nil
		}
		leftStr := split[index]
		if _, ok := nilValueMap[leftStr]; !ok {
			v, err := strconv.Atoi(leftStr)
			if err != nil {
				return nil, errors.New("can not parse int: " + leftStr)
			}
			left := &TreeNode{Val: v}
			node.Left = left
			l.PushBack(left)
		}
		index++

		if index == len(split) {
			return root, nil
		}
		rightStr := split[index]
		if _, ok := nilValueMap[rightStr]; !ok {
			v, err := strconv.Atoi(rightStr)
			if err != nil {
				return nil, errors.New("can not parse int: " + rightStr)
			}
			right := &TreeNode{Val: v}
			node.Right = right
			l.PushBack(right)
		}
		index++
	}

	return root, nil
}

func (t *TreeNode) PreorderTraversal() []int {
	var res []int
	if t == nil {
		return res
	}

}

func (t *TreeNode) preorderTraversal(res []int) {
	if t == nil {
		return res
	}

}

func main() {
	node, err := NewTreeNode("[1, 2, 3, nil, nil, 4, nil, nil, 5]")
	if err != nil {
		return
	}
	println(node)
}
