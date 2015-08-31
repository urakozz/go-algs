#Problem Statement

You are given a 2D matrix, a, of dimension MxN and a positive integer R. You have to rotate the matrix R times and print the resultant matrix. Rotation should be in anti-clockwise direction.

Rotation of a 4x5 matrix is represented by the following figure. Note that in one rotation, you have to shift elements by one step only (refer sample tests for more clarity).

matrix-rotation

It is guaranteed that the minimum of M and N will be even.

###Input 
First line contains three space separated integers, M, N and R, where M is the number of rows, N is number of columns in matrix, and R is the number of times the matrix has to be rotated. 
Then M lines follow, where each line contains N space separated positive integers. These M lines represent the matrix.

###Output 
Print the rotated matrix.

###Constraints 
```
2 <= M, N <= 300 
1 <= R <= 109 
min(M, N) % 2 == 0 
1 <= aij <= 108, where i ∈ [1..M] & j ∈ [1..N]
```

Sample Input #00
```
4 4 1
1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16
```

Sample Output #00
```
2 3 4 8
1 7 11 12
5 6 10 16
9 13 14 15
```


Sample Input #02
```
5 4 7
1 2 3 4
7 8 9 10
13 14 15 16
19 20 21 22
25 26 27 28
```

Sample Output #02

```
28 27 26 25
22 9 15 19
16 8 21 13
10 14 20 7
4 3 2 1
```
