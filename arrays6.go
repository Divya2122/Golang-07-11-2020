package main 
  
import "fmt"
  
func main() { 
    
myarray:= [...]int{29, 79, 49, 39, 
                   20, 49, 48, 49} 
  

for x:=0; x < len(myarray); x++{ 
fmt.Printf("%d\n", myarray[x]) 
} 
} 