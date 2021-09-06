# Algorithm-HUB
![GitHub Logo](https://raw.githubusercontent.com/Mhdaan/Algorithm-HUB/main/img/logo.png)

Solving Algorithm Problem in Algorithm Hub

My daily solving algorithms with **Golang**
## Problem 1: Check if two arrays are equal or not in Golang
Given two given integer arrays of any length without sorting them, the task is to find if the given arrays are equal or not. Two arrays are said to be equal if both of them contain the same set of elements. *Growth Order*: **O(n^2)**
## Problem 2: MergeSort in Golang
Simple RECURSIVE MergeSort [MergeSort wiki](https://en.wikipedia.org/wiki/Merge_sort). *Growth Order*: **Θ(nlg(n))** by **Master's theorem** case 2
## Problem 3: BinarySearch in Golang
Simple RECURSIVE BinarySearch without loop [BinarySearch wiki](https://en.wikipedia.org/wiki/Binary_search_algorithm). *Growth Order*: **Θ(lg(n))** by **Master's theorem** case 2
## Problem 4: Exponentiation in Golang
Simple RECURSIVE Power for integers [Exponentiation wiki](https://en.wikipedia.org/wiki/Exponentiation). *Growth Order*: **Θ(lg(n))** by **Master's theorem** case 2
## Problem 5: Fibonacci in Golang
Simple RECURSIVE Fibonacci [Fibonacci wiki](https://en.wikipedia.org/wiki/Fibonacci_number). *Growth Order*: **Θ(n)** by **Master's theorem** case 1 (This theta is for calling the Fibonacci function **once** which I explained in the code, but the growth order for the Naive Recursive Algorithm is **Ω(φ^n)**(big Omega of Phi(Golden Ratio) powered by n))
## Problem 6: Matrix multiplication in Golang
Simple Matrix multiplication [Matrix multiplication wiki](https://en.wikipedia.org/wiki/Matrix_multiplication). *Growth Order*: **Θ(n^3)**
## Problem 7: Fibonacci with Matrix multiplication in Golang
RECURSIVE Fibonacci using Matrix multiplication (If we want to compute the nth Fibonacci number we just power the matrix M=[1,1,/n,1,0] n times and we get 
M=[Fn+1,Fn,/n,Fn,Fn−1]). *Growth Order*: **Θ(lg(n))** (Also wrote the algorithm for the nth number of Fibonacci sequence with different beginning numbers )
## Problem 8: The minimum-subarray problem(Maximum subarray problem) in Golang
RECURSIVE finding minimum-subarray(maximum-subarray)[Maximum subarray problem](https://en.wikipedia.org/wiki/Maximum_subarray_problem). *Growth Order*: **Θ(nlg(n))** by **Master's theorem** case 2
## Problem 9: Recursive squared matrix multiplication in Golang
RECURSIVE Matrix multiplication [Matrix multiplication wiki](https://en.wikipedia.org/wiki/Matrix_multiplication). *Growth Order*: **Θ(n^3)** by **Master's theorem** case 3 and this simple divide-and-conquer approach is no faster than the straightforward matrix multiplication procedure.
## Problem 10: Strassen algorithm for squared matrix multiplication in Golang
Strassen implementation [Strassen algorithm](https://en.wikipedia.org/wiki/Strassen_algorithm). *Growth Order*: **Θ(n^2.8)** by **Master's theorem** case 3
It is slightly faster than the other two methods.