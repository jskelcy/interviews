package main

import (
	"fmt"
	"strconv"
)

func main() {
	expression := "6+9-12"
	val := eval(expression)
	println(val)
}

func eval(exp string) int {
	var (
		op    *string
		left  *string
		right *string
	)
	for _, r := range exp {
		s := string(r)
		switch s {
		case "+":
			maybeEvalOp(op, left, &right)
			op = &s
		case "-":
			maybeEvalOp(op, left, &right)
			op = &s
		default:
			if left == nil {
				left = &s
				continue
			}
			if op == nil && right == nil {
				leftI := *left + s
				left = &leftI
				continue
			}
			if right != nil {
				rightI := *right + s
				right = &rightI
				continue
			}
			right = &s
		}
	}

	return evalOp(*op, *left, *right)
}

func maybeEvalOp(op *string, left *string, right **string) {
	if op != nil && *right != nil {
		newLeft := evalOp(*op, *left, **right)
		sLeft := strconv.Itoa(newLeft)
		*left = sLeft
		*right = nil
	}
}

func evalOp(op, left, right string) int {
	leftN, _ := strconv.Atoi(left)
	rightN, _ := strconv.Atoi(right)
	fmt.Printf("left: %v right: %v op: %v\n", leftN, rightN, op)
	switch op {
	case "+":
		return leftN + rightN
	default:
		return leftN - rightN
	}
}
