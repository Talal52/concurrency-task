# filereading: 
 1. This is a Go program that reads the contents of a file named "text.txt" using the os.ReadFile() function, which returns the file's contents as a byte slice.
2. The if err != nil statement checks if an error occurred during the file reading process, and if so, prints the error message to the console using fmt.Println().
3. Otherwise, the program converts the byte slice to a string using string(talal) and prints it to the console using fmt.Println().

## version:
As for the version of Go that supports this code, the `os.ReadFile()` function was introduced in Go `version 1.16`, which was released in February 2021. So, any version of Go that is 1.16 or later should support this code.

# main.go:
1. The following Go program demonstrates how to list the files in a directory and print their names to the console using the os.ReadDir() function.
2. The `os.ReadDir()` function returns a slice of DirEntry values representing the files and subdirectories in the specified directory. In this example, the current directory is specified using ".".

3. If an error occurs while reading the directory, such as the directory not existing, the error value will be non-nil and the program will print the error message to the console.

4. The program then uses a for loop to print the names of the first four files in the directory to the console. The *Name()* method of the DirEntry type is used to get the name of each file.

# wordCount:
Imports the necessary packages:`fmt` for printing to the console, `os` for reading files from the operating system, and `strings` for string manipulation functions.

The `main()` function is defined as the entry point of the program.

```
os.ReadFile("text.txt")
```
reads the contents of the file named "text.txt" into a byte array called `talal`. If there is an error during the read operation, the error message will be printed to the console using
```
fmt.Println(err)
```
*strings.Fields(string(talal))* converts the contents of talal from a byte array to a string, and then splits the string into an array of substrings using whitespace as the delimiter. However, the result of this operation is not stored or used in any way.

Finally, the length of the `talal` array is printed to the console using
 ```
fmt.Printf("length = %d",len(talal))
```