package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//#define N 4
//
//// Function to get cofactor of A[p][q] in temp[][]. n is current
//// dimension of A[][]
//void getCofactor(int A[N][N], int temp[N][N], int p, int q, int n)
//{
//	int i = 0, j = 0;
//
//	// Looping for each element of the matrix
//	for (int row = 0; row < n; row++)
//	{
//		for (int col = 0; col < n; col++)
//		{
//			// Copying into temporary matrix only those element
//			// which are not in given row and column
//			if (row != p && col != q)
//			{
//				temp[i][j++] = A[row][col];
//
//				// Row is filled, so increase row index and
//				// reset col index
//				if (j == n - 1)
//				{
//					j = 0;
//					i++;
//				}
//			}
//		}
//	}
//}
//
///* Recursive function for finding determinant of matrix.
//n is current dimension of A[][]. */
//int determinant(int A[N][N], int n)
//{
//	int D = 0; // Initialize result
//
//	// Base case : if matrix contains single element
//	if (n == 1)
//		return A[0][0];
//
//	int temp[N][N]; // To store cofactors
//
//	int sign = 1; // To store sign multiplier
//
//	// Iterate for each element of first row
//	for (int f = 0; f < n; f++)
//	{
//		// Getting Cofactor of A[0][f]
//		getCofactor(A, temp, 0, f, n);
//		D += sign * A[0][f] * determinant(temp, n - 1);
//
//		// terms are to be added with alternate sign
//		sign = -sign;
//	}
//
//	return D;
//}
//
//// Function to get adjoint of A[N][N] in adj[N][N].
//void adjoint(int A[N][N],int adj[N][N])
//{
//	if (N == 1)
//	{
//		adj[0][0] = 1;
//		return;
//	}
//
//	// temp is used to store cofactors of A[][]
//	int sign = 1, temp[N][N];
//
//	for (int i=0; i<N; i++)
//	{
//		for (int j=0; j<N; j++)
//		{
//			// Get cofactor of A[i][j]
//			getCofactor(A, temp, i, j, N);
//
//			// sign of adj[j][i] positive if sum of row
//			// and column indexes is even.
//			sign = ((i+j)%2==0)? 1: -1;
//
//			// Interchanging rows and columns to get the
//			// transpose of the cofactor matrix
//			adj[j][i] = (sign)*(determinant(temp, N-1));
//		}
//	}
//}
//
//// Function to calculate and store inverse, returns false if
//// matrix is singular
//bool inverse(int A[N][N], float inverse[N][N])
//{
//	// Find determinant of A[][]
//	int det = determinant(A, N);
//	if (det == 0)
//	{
//		cout << "Singular matrix, can't find its inverse";
//		return false;
//	}
//
//	// Find adjoint
//	int adj[N][N];
//	adjoint(A, adj);
//
//	// Find Inverse using formula "inverse(A) = adj(A)/det(A)"
//	for (int i=0; i<N; i++)
//		for (int j=0; j<N; j++)
//			inverse[i][j] = adj[i][j]/float(det);
//
//	return true;
//}

func random(min, max int) float64 {
	return float64(rand.Intn(max-min) + min)
}

func getCofactor(A [][]float64, temp [][]float64, p int, q int, n int) {
	i := 0
	j := 0

	// Looping for each element of the matrix
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			// Copying into temporary matrix only those element
			// which are not in given row and column
			if row != p && col != q {
				temp[i][j] = A[row][col]
				j++
				if j == n-1 {
					j = 0
					i++
				}
			}
		}
	}
}

func determinant(A [][]float64, n int) float64 {
	D := float64(0)
	if n == 1 {
		return A[0][0]
	}

	temp := createMatrix(n, n)
	sign := 1

	for f := 0; f < n; f++ {
		// Getting Cofactor of A[0][f]
		getCofactor(A, temp, 0, f, n)
		D += float64(sign) * A[0][f] * determinant(temp, n-1)
		sign = -sign
	}

	return D
}

func adjoint(A [][]float64) ([][]float64, error) {
	N := len(A)
	adj := createMatrix(N, N)
	if N == 1 {
		adj[0][0] = 1
		return adj, nil
	}

	// temp is used to store cofactors of A[][]
	sign := 1
	var temp = createMatrix(N, N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			// Get cofactor of A[i][j]
			getCofactor(A, temp, i, j, N)

			// sign of adj[j][i] positive if sum of row
			// and column indexes is even.
			if (i+j)%2 == 0 {
				sign = 1
			} else {
				sign = -1
			}

			// Interchanging rows and columns to get the
			// transpose of the cofactor matrix
			adj[j][i] = float64(sign) * (determinant(temp, N-1))
		}
	}
	return adj, nil
}

func inverseMatrix(A [][]float64) ([][]float64, error) {
	N := len(A)
	var inverse = createMatrix(N, N)
	det := determinant(A, N)
	if det == 0 {
		fmt.Println("Singular matrix, can't find its inverse")
		return nil, nil
	}

	adj, err := adjoint(A)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	fmt.Println("Determinant:", det)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			inverse[i][j] = float64(adj[i][j]) / float64(det)
		}
	}

	return inverse, nil
}

func multiplyMatrices(m1 [][]float64, m2 [][]float64) ([][]float64, error) {
	if len(m1[0]) != len(m2) {
		return nil, errors.New("Cannot multiply the given matrices!")
	}

	result := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		result[i] = make([]float64, len(m2[0]))
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return result, nil
}

func createMatrix(row, col int) [][]float64 {
	r := make([][]float64, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r[i] = append(r[i], random(-5, i*j))
		}
	}
	return r
}

func main() {
	rand.Seed(time.Now().Unix())
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Wrong number of arguments!")
		return
	}

	var row int
	row, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Need an integer:", arguments[1])
		return
	}
	col := row
	if col <= 0 {
		fmt.Println("Need positive matrix dimensions!")
		return
	}

	m1 := createMatrix(row, col)
	m2 := createMatrix(row, col)
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)

	inverse, err := inverseMatrix(m2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\t\t\tPrinting inverse matrix!")
	for i := 0; i < len(inverse); i++ {
		for j := 0; j < len(inverse[0]); j++ {
			fmt.Printf("%.2f ", inverse[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\t\t\tPrinting result!")
	r1, err := multiplyMatrices(m1, inverse)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(r1); i++ {
		for j := 0; j < len(r1[0]); j++ {
			fmt.Printf("%.3f ", r1[i][j])
		}
		fmt.Println()
	}
}
