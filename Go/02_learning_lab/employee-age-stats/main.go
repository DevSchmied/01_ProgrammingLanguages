package main

/*
Task. Employee Ages

You are given data about company employees.
You must determine:
1- the minimum age of an employee;
2- the median age of all employees;
3- the maximum age.

The median is the central number in a sorted list.
If the number of employees is even, the median is calculated as the average of the two middle values, rounded down (integer division).

==============

Input format

A single line containing employee data separated by ;.
Each employee is described as:
Name,Age,Department

Age is an integer from 20 to 100.

Output format

Three integers separated by spaces:
MinAge MedianAge MaxAge

==============

Example 1

Input:

John,28,Engineering;Alex,34,HR;Dennis,45,Marketing;Anna,30,Engineering;Bob,24,HR


Output:

24 30 45

==============

Example 2

Input:

Paul,28,Engineering;Elena,34,Marketing


Output:

28 31 34


(Median = (28 + 34) / 2 = 31)
*/
