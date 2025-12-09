package main

/*
Task. Trip Planning

You are given a list of cities and a list of hotels, where each hotel corresponds to a city by index.
You need to output all cities that have the maximum number of hotels.
If multiple cities share the maximum, output them in alphabetical order.
The city name must be printed with the first letter uppercase and the remaining letters lowercase.

================================

Input format

Two lines:
- A list of cities — strings (Latin alphabet), separated by commas without spaces.
- A list of hotels — strings, separated by commas without spaces.

The number of cities and hotels is the same.
City comparison must be case-insensitive.

================================

Output format

Each line:
- CityName NumberOfHotels

================================

Example 1

Input:
- Paris,London,London,Paris,Lyon,London
- HotelA,HotelB,HotelC,HotelD,HotelE,HotelF


Output:

London 3

(“London” appears 3 times — the maximum.)

================================

Example 2

Input:

- Nice,Nice,Nice,Nice,Nice,Bordeaux,Bordeaux,Bordeaux,Bordeaux,Bordeaux
- Negresco1,Negresco2,Negresco3,Negresco4,Negresco5,Regent1,Regent2,Regent3,Regent4,Regent5


Output:

Bordeaux 5
Nice 5
*/
