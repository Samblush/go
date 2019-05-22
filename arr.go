package main
import "fmt"
func main() {
	var i,j,m,n int
	var a[10][10]int
	var b[10][10]int
	fmt.Print("Enter the no of rows of matrix :")
	fmt.Scanln(&m)
	fmt.Print("Enter the no of column : ")
	fmt.Scanln(&n)
	
	fmt.Println("Enter the Matrix A")
	for i=0;i<m;i++ {
	for j=0;j<n;j++{
	fmt.Scan(&a[i][j])
	}
	}

	fmt.Println("Enter the Matrix B element: ")
	for i=0;i<m;i++ {
	for j=0;j<n;j++{
	fmt.Scan(&b[i][j])
	}
	}

	fmt.Println("The results addition of matrix A & B:")
	for i=0;i<m;i++ {
	for j=0;j<n;j++ {
	fmt.Print(a[i][j]+b[i][j],"\t")
	}
	fmt.Println()
	}
}

