A Go script file, along its source code.

When the code executes, it asks for 3 things: The folder name, the text file name and the content that will
be added into the text file.

The input supports spaces.

At the moment, the path is hardcoded into the script, so it might not work on every computer. You can
change the path in the source code, by editing the "destination" variable.

I originally wanted to include the script file (the .exe file) itself, but i dont think Git will let me.
So you have to create your own .exe file yourself, or until i figure it out. You just run the:
"go build folderAndTextFileScript.go", from the path that this program is in. You need to install Golang, however.