# codwars-range-of-integers-in-an-unsorted-string
Codewars Kata : 5kyu  Range of Integers in an Unsorted String


https://www.codewars.com/kata/5b6b67a5ecd0979e5b00000e

In this kata, your task is to write a function that returns the smallest and largest integers in an unsorted string.

Input
Your function will receive two arguments:
A string comprised of integers in an unknown range; these numbers will be out of order
An integer value representing the size of the range
Output
Your function should return the starting (minimum) and ending (maximum) numbers of the range in the form of an array/list comprised of two integers.
Test Example
inputString := "1568141291110137"
MysteryRange(inputString,10) // [6 15]

// The 10 numbers in this string are:
// 15 6 8 14 12 9 11 10 13 7
// Therefore the range of numbers is from 6 to 15
Technical Details
The maximum size of a range will be 100 integers
The starting number of a range will be: 0 < n < 100
Full Test Suite: 21 fixed tests, 100 random tests
Use Python 3+ for the Python translation
For JavaScript, require has been disabled and most built-in prototypes have been frozen (prototype methods can be added to Array and Function)
All test cases will be valid


Currently failing on 
s = "1721532418565922162558663126649136347436733301144143236653738464135820194215516155541239452852623450572927602348104049" 
l = 60

due to the fact that there are 2x 64's in the string
