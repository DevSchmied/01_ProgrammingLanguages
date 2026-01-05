package main

/*
Interview Task Description:

You are given a sequence of elements.

Iterate over this sequence and output the position and value of each element.

During iteration, modify the sequence by changing some of its elements and adding new ones.

Consider and demonstrate the difference in behavior between two iteration approaches:
- one where the number of iterations is determined at the beginning,
- and one where the termination condition depends on the current size of the sequence.

Ensure that the program avoids runtime errors caused by invalid access
and prevents an infinite loop when the sequence is dynamically extended.

Explain or demonstrate how modifying a collection while iterating over it can affect
control flow, termination conditions, and data consistency.
*/

import "fmt"

func main() {
	x := []string{"A", "B", "C"}

	for idx, v := range x {
		fmt.Printf("%d: %s\n", idx, v)
		x[idx+1] = "M"
		x = append(x, "Z")
		x[idx+1] = "Z"
	}
	fmt.Println("array x:", x)

	x = []string{"A", "B", "C"}
	for i := 0; i < len(x); i++ {
		if len(x) >= 10 {
			fmt.Println("Otherwise infinite loop")
			break
		}
		x[i+1] = "M"
		fmt.Printf("%d: %s\n", i, x[i])
		x = append(x, "Z")
		x[i+1] = "Z"

	}
}
