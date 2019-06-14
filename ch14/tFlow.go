package main

import (
	"fmt"
	"os"
	"strconv"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func Add(sum_arg1, sum_arg2 int8) (interface{}, error) {
	sum_scope := op.NewScope()
	input1 := op.Placeholder(sum_scope.SubScope("a1"), tf.Int8)
	input2 := op.Placeholder(sum_scope.SubScope("a2"), tf.Int8)
	sum_result_node := op.Add(sum_scope, input1, input2)

	graph, err := sum_scope.Finalize()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	a1, err := tf.NewTensor(sum_arg1)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	a2, err := tf.NewTensor(sum_arg2)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	session, err := tf.NewSession(graph, nil)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer session.Close()

	sum, err := session.Run(
		map[tf.Output]*tf.Tensor{
			input1: a1,
			input2: a2,
		},
		[]tf.Output{sum_result_node}, nil)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return sum[0].Value(), nil
}

func Multiply(sum_arg1, sum_arg2 int8) (interface{}, error) {
	sum_scope := op.NewScope()
	input1 := op.Placeholder(sum_scope.SubScope("x1"), tf.Int8)
	input2 := op.Placeholder(sum_scope.SubScope("x2"), tf.Int8)

	sum_result_node := op.Mul(sum_scope, input1, input2)
	graph, err := sum_scope.Finalize()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	x1, err := tf.NewTensor(sum_arg1)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	x2, err := tf.NewTensor(sum_arg2)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	session, err := tf.NewSession(graph, nil)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer session.Close()

	sum, err := session.Run(
		map[tf.Output]*tf.Tensor{
			input1: x1,
			input2: x2,
		},
		[]tf.Output{sum_result_node}, nil)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return sum[0].Value(), nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}

	t1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	n1 := int8(t1)

	t2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	n2 := int8(t2)

	res, err := Add(n1, n2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add:", res)
	}

	res, err = Multiply(n1, n2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Multiply:", res)
	}
}
