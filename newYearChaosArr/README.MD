# Algorithm
1. In beginning of first for loop if index of item+3 is less than item value, there is no way for that value to achieve its position in less then 2 moves, thus, we return "Too chaotic" and leave the function
2. In case 1. is false, if the current value-2 is greater than 0 use that as a starting index, otherwise use 0
3. Until the second index reaches the first iterate over array and compare with current value, if value iterated is greater, add to moves
4. Return number of moves
