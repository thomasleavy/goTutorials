/*
Before using this code, you must install Go onto your system. Go to https://go.dev/dl/ and download the version appropriate for your device (i.e. the '.msi' file) and go through the installation process.
*/
/*
To verify if this has been downloaded, open the CMD prompt or PowerShell and type: go version. If a version is returned in the terminal, it is installed on your system.
*/
/*
If necessary, you might need to set up Go Environmental Variables. Search for "Environmental Variables" from the Start menu of you device. In the "System Variables" section, find the "Path" variable and edit it. Click "New" and add "C:\Go\bin". Press OK and close. Open CMD prompt or PowerShell and type "go version" to verify the path has been set up properly.
*/
/*
If you are using Visual Studio Code (VSCode), click the Extensions icon at the left sidebar and search for the "Go Team at Google" extension and install it.
*/
/*
Create a new folder, for example, GoProjects. Create a new file called "01-hello.go". Then you can use the below code.
*/
/*
To run this programme, open a terminal and type: go run <file_name>.go
*/
package main

import "fmt"

/*
In Go, "fmt" is an abbreviation for "format". It is needed to print text and facilitates input/output operations like the below code
*/

func main() {
	fmt.Println("Hello, how are you today?")
}
