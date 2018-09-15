This is just a code in a funny contest at: https://www.facebook.com/groups/335069743983293/permalink/340879163402351/

## Challenge
Write a program that covert a .jpg color picture to a .jpeg greyscale picture.
1. Without external libraries by Golang.
2. Using maximum 5 go routines.
3. Have use command line paramters.

## My solution
Using avarage formula to convert RGB to grey-scale.

`grey = (r + g + b) / 3`

## How to run
The main function recives 2 input:
1. param1: Path to the input file.
2. param2: Path to the ouput file.

**Eg:**

`./main ./color.jpg ./dest.jpg`