package main

/*
Task Best Candidates

The HR department of a company has conducted interviews and received candidate ratings.
Print the surnames of candidates whose rating is strictly above the average rating, preserving their original order.
If there are no such candidates, output: "no".

==========

Input format

The input contains two lines:
1. A list of candidate ratings — integers from 0 to 10, separated by commas with no spaces.
2. A list of candidate surnames — strings separated by commas with no spaces.

The order of surnames CORRESPONDS to the order of ratings.
It is GUARANTEED that the number of surnames matches the number of ratings.

==========

Output format

Print the surnames of all candidates whose rating is strictly above the average — one surname per line.

If there are no such candidates, print:

no

==========

Example 1

Input:

4,6,8,8,5,8,4,7,9,2
Smith,Johnson,Williams,Brown,Jones,Miller,Davis,Garcia,Rodriguez,Wilson


Output:

Williams
Brown
Miller
Garcia
Rodriguez

==========

Example 2

Input:

7,7,7,7,7
Dupont,Martin,Bernard,Leroy,Петров


Output:

no
*/

import (
	"strconv"
	"strings"
)

func getBestCandidates(ratingsLine, namesLine string) []string {
	ratingStrs := strings.Split(ratingsLine, ",")
	ratings := make([]int, 0, len(ratingStrs))

	sum := 0
	for _, val := range ratingStrs {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil // invalid input
		}
		ratings = append(ratings, num)
		sum += num
	}

	names := strings.Split(namesLine, ",")
	if len(names) != len(ratings) {
		return nil // mismatched input
	}

	avg := float64(sum) / float64(len(ratings))
	results := []string{}

	for i, val := range ratings {
		if float64(val) > avg {
			results = append(results, names[i])
		}
	}

	if len(results) == 0 {
		return nil // no candidates above average
	}

	return results
}
