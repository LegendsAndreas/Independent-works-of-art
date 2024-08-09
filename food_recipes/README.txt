This program is a macro tracker, that specifically tracks your calories, fats, carbs and protein. It works by you inserting your own recipes into the program, which includes the name of the recipe, along with every ingredient measured in grams and its macros pr. 100 grams.

When you start the program, it makes a text file where it stores all your recipes. So make sure that the program and the text file is in the same directory (just make a folder with them).

After many moments of absolute despair, I finally return with a 1.1 version of this program!
The reason for this, is that I found out that "bufio.NewScanner(os.stdin)" does not work well with the terminal, when you execute the .exe file. So for this, i have reworked it to just use "fmt.Scan()"
