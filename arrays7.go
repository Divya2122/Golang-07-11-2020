package main 
  
import "fmt"
  
func main() { 
      
 
my_array:= [...]int{100, 200, 300, 400, 500} 
fmt.Println("Original array(Before):", my_array) 
  

new_array := my_array 
  
fmt.Println("New array(before):", new_array) 
  
new_array[0] = 500 
  
fmt.Println("New array(After):", new_array) 
  
fmt.Println("Original array(After):", my_array) 
} 