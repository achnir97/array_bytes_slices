package main

import(
	"fmt"
	"time"
)





func main(){
	var Array [10]int 
	Slices:= make([]int,0,10)
	bytesArray:= []byte{1,2,3, 4,5, 6,7, 8, 9,10}
	
	b:=[]byte("Hello, World!")

	// print the slices of bytes
	fmt.Println(b)
	
	fmt.Println(string(b))

	//Modify the slices of bytes
	b[7]=','
	b[8]=' '
	
	b=append(b, []byte("GO")...)

	//Print the modified slice of bytes
	fmt.Println(b)

	fmt.Println(string(b))


	go func(){
		for i:=0; i<=10;i++{
		Array[i]=i
		fmt.Printf("The value is Array is %d\n", Array[i])
		time.Sleep(1*time.Second)
	}   
}()
go func(){
	for i :=0; i<=10;i++{
		Slices=append(Slices, i)
		fmt.Printf("The value in the slice is %d\n",Slices[i])
		time.Sleep(1*time.Second)
	}
}()

go func(){
	
	for v:=range(bytesArray){
		fmt.Printf("The Value from the Bytes Array is %d\n", v)
		time.Sleep(1*time.Second)
	}
	
}()

time.Sleep(10*time.Second)
fmt.Printf("The execution of the main function is over.")

}
// bytes are used in go programming language because
// they are the smallest unit od data in computer memory
// are commonly used to represent raw data and text. 

// Go provides []byte type which is an alias for []unit8 
// type 
// Portability: Bytes are a universal data type and can be easily
// transfered between different platforms and systems, 
// making them ideal for network and internet applicatio 


// Efficient manipualti: Bytes can be easily manipulated using bitwise operations 
// which are faster than string operations

// encoding and decofing: bytes are used to represt and 
// store encoded data, such as text in differnt character sets, images and audio files, 
// the Encoding and decoding of bytes is an important aspcet 
// data preprocessing and communication. 

// Low level programming: Bytes provide a low-level representation of
// data which is useful for programming at the hardware level 
// for working with binary files and protocols. 


// Interoperability: Bytes are commonly used to exchane data between differn porgrmaming langauges 
// systems, making it easy to integrate with other tools and 
// technologies. 


// In Go, []byte type is often used for reading and writing binary data
// working with encoded text, and processing data at low-level.
// By using the []byte type, Go programmers can write efficient, portable and 
// interoperable application that can handle a wide range of data 
// processing and communication tasks. 

