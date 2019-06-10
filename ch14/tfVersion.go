package main

import (
    tf "github.com/tensorflow/tensorflow/tensorflow/go"
    "github.com/tensorflow/tensorflow/tensorflow/go/op"
    "fmt"
)

func main() {
    s := op.NewScope()
    c := op.Const(s, "Using TensorFlow version: " + tf.Version())
    graph, err := s.Finalize()

    if err != nil {
        fmt.Println(err)
		return
    }

    sess, err := tf.NewSession(graph, nil)
    if err != nil {
        fmt.Println(err)
		return
    }

    output, err := sess.Run(nil, []tf.Output{c}, nil)
    if err != nil {
        fmt.Println(err)
		return
    }

    fmt.Println(output[0].Value())
}


