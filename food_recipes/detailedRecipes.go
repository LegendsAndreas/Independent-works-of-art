package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// A recipe, which holds a recipe number, meal type, recipe name, a slice of ingredient and a total amount of macros of totalMacros.
type recipe struct {
	recipeNumber uint
	mealType     rune // B = Breakfast, L = Lunch, D = Dinner, S = Side, K = Snack
	recipeName   string
	ingredients  []ingredient
	total        totalMacros
}

// An ingredient to a recipe.
type ingredient struct {
	name          string
	kilograms     float32
	calories      float32
	fats          float32
	carbohydrates float32
	protein       float32
	multiplier    float32

	/* EXAMPLE
	Tulip-Pulled-Pork	// Name
	125					// Measured in grams
	133					// Calories/Kcal
	6					// Fats
	2					// Carbs
	17					// Protein
	125/100 = 1.25 		// Multiplier. So the amount of kilograms divided with 100.
	*/
}

// A struct that holds all the total macros for a recipe.
type totalMacros struct {
	tCalories      float32
	tFats          float32
	tCarbohydrates float32
	tProteins      float32
}

func main() {
	var input string
	var recipes []recipe

	// Initializes 'recipes' with all the current recipes in recipes_go.txt
	recipes = initializeRecipes(recipes)

	for {
		// Asks what the user wants to do and stores the input in 'input'.
		fmt.Printf("Enter your action\n" +
			"	new\n" +
			"	edit\n" +
			"	remove\n" +
			"	print\n" +
			"	plan\n" +
			"	random\n" +
			"	q\n>")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal("Scanning input error:", err)
			return
		}

		if input == "new" { // If the user wants to add a new recipe.
			recipes = newRecipe(recipes)

		} else if input == "edit" { // If the user wants to edit an existing recipe.
			recipes = editRecipe(recipes)

		} else if input == "remove" { // If the user wants to remove a recipe.
			recipes = removeRecipe(recipes)
			// Updates the file "recipes_go"
			updateFile(recipes)

		} else if input == "print" { // If the user wants to print recipes.
			var printInput string
			fmt.Print("Enter what you want to print\n" +
				"	basic\n" +
				"	single\n" +
				"	everything\n" +
				">")
			_, err2 := fmt.Scan(&printInput)
			if err2 != nil {
				log.Fatal("Scanning print input error:", err2)
				return
			}

			if printInput == "basic" { // Prints a basic version of all the recipes.
				printBasicRecipes(recipes)
			} else if printInput == "single" { // Print a single recipe of all the recipes.
				printSingleRecipe(recipes)
			} else if printInput == "everything" { // Prints every single recipe.
				printEverything(recipes)
			} else {
				fmt.Println("Invalid input")
			}

		} else if input == "plan" { // If the user wants to plan a day of eating.
			var planInput string
			fmt.Print("Enter day or week plan\n" +
				"	day\n" +
				"	week\n" +
				">")
			_, err3 := fmt.Scan(&planInput)
			if err3 != nil {
				fmt.Println("Scanning plan input error:", err3)
				return
			}

			if planInput == "day" {
				planDay(recipes)
			} else if planInput == "week" {
				planWeek(recipes)
			}

		} else if input == "random" {
			randomMeal(recipes)
		} else if input == "q" { // If the user wants to quit the program.
			fmt.Println("Exiting the program...")
			break
		} else {
			fmt.Println("Invalid input")
		}

	}
}

// initializeRecipes occurs once during the startup of the program and collects all the recipes from 'recipes_go.txt'
func initializeRecipes(recipeSli []recipe) []recipe {
	fmt.Println("Initializing recipes...")

	// We open the file in os.O_CREATE|os.O_RDWR mode, which makes it, so that if the text file does not exist, it creates
	// a new file called "recipes_go.txt" and opens it in read and write mode. If the file does exist, it just opens
	// it in read and write mode.
	recipesFile, err := os.OpenFile("recipes_go.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Println("Initializing file error:", err)
		return nil
	}

	fileScanner := bufio.NewScanner(recipesFile)
	fileScanner.Split(bufio.ScanLines) // This basically tells the scanner to read the file, until it comes to a newline.

	for fileScanner.Scan() {
		// Reads the file, one line at a time and stores the content in 'line'.
		line := fileScanner.Text()
		// Splits 'line' and an array of substrings, with the delimiter being a space.
		splitLineArray := strings.Split(line, " ")

		// Extract the recipe number from the line and convert it to uint
		recipeNumber, err := strconv.ParseUint(splitLineArray[0], 10, 32)
		if err != nil {
			log.Fatal("Parsing recipe number error:", err)
		}

		// Extract the meal type from the line and convert it to rune
		mealType := rune(splitLineArray[1][0])

		// Extract the recipe name from the line
		recipeName := splitLineArray[2]

		// Extract the ingredients from the line
		ingredients := make([]ingredient, 0)
		for i := 3; i+6 <= len(splitLineArray); i += 6 { // "i+6 <= len(splitLineArray)" checks that there actually are the appropriate amount of elements, so that we don't get an out of bound error. Set it to "i < len(splitLineArray)" and you are done for.
			// Perhaps, I could rewrite this with fprintf(name, kilo, cal, fats, carbs, pro)
			ingredientName := splitLineArray[i]
			kilograms, err := strconv.ParseFloat(splitLineArray[i+1], 32)
			if err != nil {
				log.Fatal("Parsing ingredient kilograms error:", err)
			}
			calories, err := strconv.ParseFloat(splitLineArray[i+2], 32)
			if err != nil {
				log.Fatal("Parsing ingredient calories error:", err)
			}
			fats, err := strconv.ParseFloat(splitLineArray[i+3], 32)
			if err != nil {
				log.Fatal("Parsing ingredient fats error:", err)
			}
			carbohydrates, err := strconv.ParseFloat(splitLineArray[i+4], 32)
			if err != nil {
				log.Fatal("Parsing ingredient carbohydrates error:", err)
			}
			protein, err := strconv.ParseFloat(splitLineArray[i+5], 32)
			if err != nil {
				log.Fatal("Parsing ingredient protein error:", err)
			}

			// Inserts all the gathered info into 'fileIngredient' and appends it to 'ingredients'.
			fileIngredient := ingredient{
				name:          ingredientName,
				kilograms:     float32(kilograms),
				calories:      float32(calories),
				fats:          float32(fats),
				carbohydrates: float32(carbohydrates),
				protein:       float32(protein),
				multiplier:    float32(kilograms / 100),
			}

			ingredients = append(ingredients, fileIngredient)
		}

		// Calculate the total macros
		total := getTotalMacros(ingredients)

		// Create a new recipe and append it to the recipe slice
		newRecipe := recipe{
			recipeNumber: uint(recipeNumber),
			mealType:     mealType,
			recipeName:   recipeName,
			ingredients:  ingredients,
			total:        total,
		}
		recipeSli = append(recipeSli, newRecipe)
	}

	err = recipesFile.Close()
	if err != nil {
		return nil
	}

	return recipeSli
}

// newRecipe creates a new recipe and fills it out.
func newRecipe(recipeSlice []recipe) []recipe {
	fmt.Println("Creating new recipe...")

	var tempRecipe recipe                                     // temptRecipe will be added to slice 'recipeSlice'
	tempRecipe.recipeNumber = uint(len(recipeSlice)) + 1      // Gets recipe number
	tempRecipe.mealType = getMealType()                       // Gets the meal type
	tempRecipe.recipeName = getRecipeName()                   // Gets recipe name
	tempRecipe.ingredients = getIngredients()                 // Gets ingredients
	tempRecipe.total = getTotalMacros(tempRecipe.ingredients) // Gets total macros

	addRecipeToFile(tempRecipe) // Adds the recipe to the file, so that it is saved and does not disappear upon restarting the program.

	fmt.Println("Inserting recipe into slice... ")
	recipeSlice = append(recipeSlice, tempRecipe) // Appends the recipe in the recipe slice 'recipeSlice'
	return recipeSlice
}

// getMealType prompts the user to enter a meal type (B, L, D, S, K) and returns it as a rune.
func getMealType() rune {
	var tMeal string
	var runeMealType rune

	for {
		fmt.Print("Enter meal type>")
		_, err := fmt.Scan(&tMeal)
		if err != nil {
			log.Fatal("Reading from standard input error:", err)
		}

		// Checks if the input is longer than 1.
		if len(tMeal) > 1 {
			fmt.Println("The meal type is too long, try again!")
			continue
		}

		// Checks if the input length is not 0. Don't know how you can achieve this, with the current setup, but better safe than sorry.
		if len(tMeal) < 1 {
			fmt.Println("Input is nothing, try again!")
			continue
		}

		// Checks that the input is a valid meal type.
		if tMeal != "B" && tMeal != "L" && tMeal != "S" && tMeal != "K" && tMeal != "D" {
			fmt.Println("Invalid meal type, try again!")
			continue
		}

		// We convert the first element into a rune and the loop breaks.
		runeMealType = rune(tMeal[0])
		break
	}

	return runeMealType
}

// getRecipeName prompts the user to enter a recipe name and returns it as a string.
// It uses bufio.Scanner to read the input from os.Stdin.
// If there is an error while scanning, it logs a fatal error.
func getRecipeName() string {
	fmt.Print("Enter recipe name>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// We use the .Err() method to check if 'scanner' has an error.
	// You could move the statement: err := scanner.Err(), above the if statement and it would function the same.
	if err := scanner.Err(); err != nil {
		log.Fatal("Scanning recipe name error:", err)
	}

	input := scanner.Text()
	return input
}

// getIngredients retrieves the list of ingredients from the user input. It prompts the user to enter the ingredient name,
// kilograms, calories, fats, carbohydrates, and protein for each ingredient. It calculates the multiplier by dividing the
// kilograms by 100. The function returns a slice of ingredient structs.
// It returns a slice of ingredient, with all the entered ingredients.
func getIngredients() []ingredient {
	tIngredients := make([]ingredient, 0)
	var input string
	for {
		// Asks the user for ingredient name.
		var ingre ingredient
		fmt.Print("Type your ingredient (or q)>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Fatal("Scanning ingredient name error:", err)
		}

		input = scanner.Text()

		// Breaks the loop.
		if input == "q" {
			fmt.Println("Exiting the ingredients program...")
			break
		}

		// Asks the user for the amount of kilograms.
		ingre.name = input
		fmt.Println("Enter ingredient kilograms>")
		_, err := fmt.Scan(&ingre.kilograms)
		if err != nil {
			log.Fatal("Scanning ingredient kilograms error:", err)
		}

		// Asks the user for the amount of calories.
		fmt.Println("Enter ingredient calories (pr. 100g)>")
		_, err = fmt.Scan(&ingre.calories)
		if err != nil {
			log.Fatal("Scanning ingredient calories error:", err)
		}

		// Asks the user for the amount of fats.
		fmt.Println("Enter ingredient fats (pr. 100g)>")
		_, err = fmt.Scan(&ingre.fats)
		if err != nil {
			log.Fatal("Scanning ingredient fats error:", err)
		}

		// Asks the user for the amount of carbs.
		fmt.Println("Enter ingredient carbohydrates (pr. 100g)>")
		_, err = fmt.Scan(&ingre.carbohydrates)
		if err != nil {
			log.Fatal("Scanning ingredient carbohydrates error:", err)
		}

		// Asks the user for the amount of proteins.
		fmt.Println("Enter ingredient protein (pr. 100g)>")
		_, err = fmt.Scan(&ingre.protein)
		if err != nil {
			log.Fatal("Scanning ingredient protein error:", err)
		}

		// Calculates how much each macro will be valued.
		ingre.multiplier = ingre.kilograms / 100

		// Finally appends the ingredient to 'tIngredients'
		tIngredients = append(tIngredients, ingre)

	}

	return tIngredients
}

// getTotalMacros, calculates all the macros for a recipe and returns it as a totalMacros struct.
func getTotalMacros(ingredientsSlice []ingredient) totalMacros {
	var total totalMacros
	for _, r := range ingredientsSlice {
		total.tCalories += r.calories * r.multiplier
		total.tFats += r.fats * r.multiplier
		total.tCarbohydrates += r.carbohydrates * r.multiplier
		total.tProteins += r.protein * r.multiplier
	}
	return total
}

// printBasicRecipes, prints only the recipe number and name, alongside its macros.
func printBasicRecipes(recipes []recipe) {
	fmt.Println("Printing recipe...")
	for _, r := range recipes {
		fmt.Printf("Recipe %d: %s, %.2f, %.2f, %.2f, %.2f\n", r.recipeNumber, r.recipeName, r.total.tCalories, r.total.tFats, r.total.tCarbohydrates, r.total.tProteins)
	}
}

// printEverything prints all the recipes along with their recipe number, meal type, name, and all the macro information.
// It also prints all the individual ingredients for each recipe, along with their macro information.
func printEverything(recipes []recipe) {
	fmt.Println("Printing recipe...")
	// Prints the recipe number, meal type, name and all the macros for it.
	for _, r := range recipes {
		fmt.Printf("%d, %c: %s\n", r.recipeNumber, r.mealType, r.recipeName)
		fmt.Printf("	- Total Calories: %.2f\n", r.total.tCalories)
		fmt.Printf("	- Total Proteins: %.2f\n", r.total.tProteins)
		fmt.Printf("	- Total Fats    : %.2f\n", r.total.tFats)
		fmt.Printf("	- Total Carbs   : %.2f\n", r.total.tCarbohydrates)

		fmt.Println("----------------------------------------------------")

		// Prints all the ingredients needed and the amount along with their macros.
		for _, ing := range r.ingredients {
			fmt.Printf("   %s: %.2fg\n", ing.name, ing.kilograms)
			fmt.Printf("	- Calories pr. 100g: %.2f\n", ing.calories)
			fmt.Printf("	- Proteins pr. 100g: %.2f\n", ing.protein)
			fmt.Printf("	- Fats     pr. 100g: %.2f\n", ing.fats)
			fmt.Printf("	- Carbs    pr. 100g: %.2f\n", ing.carbohydrates)
		}
		fmt.Println("----------------------------------------------------")
	}
}

func addRecipeToFile(copiedRecipe recipe) {
	fmt.Println("Inserting recipe into text file...")

	// Opens the file in append mode.
	file, err := os.OpenFile("recipes_go.txt", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Opening file error:", err)
	}

	// Adds the number, meal type and name to the file.
	_, err2 := fmt.Fprintf(file, "%d %c %s ", copiedRecipe.recipeNumber, copiedRecipe.mealType, copiedRecipe.recipeName)
	if err2 != nil {
		log.Fatal("Writing to file error:", err2)
	}

	// Adds the ingredients part of our recipe to the file.
	for _, i := range copiedRecipe.ingredients {
		_, err3 := fmt.Fprintf(file, "%s %.2f %.2f %.2f %.2f %.2f ", i.name, i.kilograms, i.calories, i.fats, i.carbohydrates, i.protein)
		if err3 != nil {
			log.Fatal("Writing to file error:", err3)
		}
	}

	// Adds a newline at the end.
	_, err4 := fmt.Fprintf(file, "\n")
	if err4 != nil {
		log.Fatal("Adding newline to text file error:", err4)
		return
	}

	err5 := file.Close()
	if err5 != nil {
		log.Fatal("Closing file error:", err4)
	}
}

func editRecipe(repSli []recipe) []recipe {
	// Prints all the recipes, so that the user can easily identify the ones they want to change.
	printBasicRecipes(repSli)

	fmt.Println("Editing recipe...")

	// Asks for which recipe to change and stores the input in 'recipeIndex'
	var recipeIndex int
	fmt.Print("Enter recipe number you wish to edit>")
	_, err := fmt.Scan(&recipeIndex)
	if err != nil {
		log.Println("Scanning input error for editing recipe:", err)
		return nil
	}

	// Checks that the entered value is not out of bound.
	if (recipeIndex - 1) > len(repSli) {
		log.Println("Recipe number out of range!")
		return nil
	}
	if recipeIndex < 0 {
		log.Println("Recipe number is negative!")
		return nil
	}

	// Asks what part of the recipe wants to be changed.
	var part string
	fmt.Print("Enter which part you want to edit\n" +
		"	- type\n" +
		"	- name\n" +
		"	- ingredients\n>")
	_, err2 := fmt.Scan(&part)
	if err2 != nil {
		log.Println("Scanning input error for editing part of recipe:", err)
		return nil
	}

	// If the user wants to change the meal type.
	if part == "type" {
		reader := bufio.NewReader(os.Stdin)

		for {
			// Asks what the meal type should be changed to.
			fmt.Print("Correct type to>")
			correctedType, err3 := reader.ReadString('\n')
			if err3 != nil {
				log.Println("Reading input error for editing part of recipe:", err)
				return nil
			}

			// Makes sure there is no whitespace.
			correctedType = strings.TrimSpace(correctedType)

			// If the user enters more than one character, we skip the iteration.
			if len(correctedType) > 1 {
				fmt.Println("Too long input, try again!")
				continue
			}

			// If the input does not match any of the relevant meal types, we skip the iteration.
			if correctedType != "B" && correctedType != "L" && correctedType != "D" && correctedType != "S" && correctedType != "K" {
				fmt.Println("Invalid meal type.")
				continue
			}

			// If the input is valid, we set 'repSli[recipeIndex].mealType' to equal to 'correctedType'.
			// If the user enters nothing (e.g. just presses "enter"), we skip the iteration.
			if len(correctedType) > 0 {
				// When we convert 'correctedType' to a rune, we must use the brackets, since the rune is technically
				// a single character and since 'correctedType' is a string, multiple characters, we have to specify one specific element.
				repSli[recipeIndex-1].mealType = rune(correctedType[0])
			} else {
				fmt.Println("Input is nothing, try again!")
				continue
			}
		}

		// If the user wants to change the recipe name.
	} else if part == "name" {
		fmt.Print("Enter corrected name>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		// We use the .Err() method to check if 'scanner' has an error.
		// You could move the statement: err := scanner.Err(), above the if statement and it would function the same.
		if err3 := scanner.Err(); err3 != nil {
			log.Println("Scanning recipe name error:", err)
			return nil
		}

		repSli[recipeIndex-1].recipeName = scanner.Text()

		// If the user wants to change the ingredients.
	} else if part == "ingredients" {
		// ToDo: Add a function that prints every ingredients, either by altering "printSingleRecipe" to work, without asking for input, or making a new function.

		// Asks the user for what they want to do with the ingredient.
		var input string
		fmt.Print("Do you want to remove, add or change a recipe?>")
		_, err12 := fmt.Scan(&input)
		if err12 != nil {
			log.Println("Input error:", err12)
			return nil
		}

		if input == "remove" {
			// If the user wants to remove an ingredient.
			repSli[recipeIndex-1].ingredients = removeIngredient(repSli[recipeIndex-1])

		} else if input == "add" {
			// If the user wants to add an ingredient.
			repSli[recipeIndex-1].ingredients = addExtraIngredient(repSli[recipeIndex-1])

		} else if input == "change" {
			// If the user wants to change an ingredient
			repSli[recipeIndex-1].ingredients = changeIngredient(repSli[recipeIndex-1])

		}

		// Updates the total macros.
		repSli[recipeIndex-1].total = getTotalMacros(repSli[recipeIndex-1].ingredients)

	}
	// Updates the file.
	updateFile(repSli)

	return repSli

}

// printSingleRecipe prints the details of a single recipe specified by the user's input recipe number.
// It takes in a slice of recipes and prompts the user to enter a recipe number.
// It iterates through the recipe slice and if a recipe with a matching recipe number is found, it prints the details of that recipe.
// If no matching recipe is found, it does nothing.
func printSingleRecipe(recipeSli []recipe) {
	fmt.Println("Printing recipe...")
	var recipeNum uint
	fmt.Print("Enter the recipe number>")
	_, err := fmt.Scan(&recipeNum)
	if err != nil {
		log.Fatal("Scanning recipe number error:", err)
	}

	recipeNum = recipeNum - 1

	fmt.Printf("%d, %c: %s\n", recipeSli[recipeNum].recipeNumber, recipeSli[recipeNum].mealType, recipeSli[recipeNum].recipeName)
	fmt.Printf("	- Total Calories: %.2f\n", recipeSli[recipeNum].total.tCalories)
	fmt.Printf("	- Total Proteins: %.2f\n", recipeSli[recipeNum].total.tProteins)
	fmt.Printf("	- Total Fats    : %.2f\n", recipeSli[recipeNum].total.tFats)
	fmt.Printf("	- Total Carbs   : %.2f\n", recipeSli[recipeNum].total.tCarbohydrates)

	fmt.Println("----------------------------------------------------")

	for _, element := range recipeSli[recipeNum].ingredients {
		fmt.Printf("   %s: %.2fg\n", element.name, element.kilograms)
		fmt.Printf("	- Calories pr. 100g: %.2f\n", element.calories)
		fmt.Printf("	- Proteins pr. 100g: %.2f\n", element.protein)
		fmt.Printf("	- Fats     pr. 100g: %.2f\n", element.fats)
		fmt.Printf("	- Carbs    pr. 100g: %.2f\n", element.carbohydrates)

		fmt.Println("----------------------------------------------------")
	}
}

// planDay takes a slice of recipes as input and allows the user to plan their meals for the day
// while keeping track of the total macros consumed. The function prompts the user to enter a recipe number
// until they choose to quit ('q'), and adds the macros of the selected recipe to the totalMacros for the day.
// Once the user quits, the function displays the final sum of macros for the entire day.
func planDay(recipeSli []recipe) {
	var macrosForToday totalMacros
	input := ""
	for {
		fmt.Printf("Current macros:\n"+
			"	Calories: 	%.2f\n"+
			"	Fats: 		%.2f\n"+
			"	Carbs: 		%.2f\n"+
			"	Protein: 	%.2f\n", macrosForToday.tCalories, macrosForToday.tFats, macrosForToday.tCarbohydrates, macrosForToday.tProteins)

		// Asks the user for what recipe the user wants to eat today.
		fmt.Print("Enter the recipe number, that you want to eat today (or q)> ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println("Scanning error:", err)
			return
		}

		if input == "q" { // Prints the final sum of macros for an entire day and breaks the loop.
			fmt.Println("Your final macros are for the day are:")
			fmt.Printf(
				"	Calories: 	%.2f\n"+
					"	Fats: 		%.2f\n"+
					"	Carbs: 		%.2f\n"+
					"	Protein: 	%.2f\n", macrosForToday.tCalories, macrosForToday.tFats, macrosForToday.tCarbohydrates, macrosForToday.tProteins)
			break
		}

		recipeNum, err2 := strconv.Atoi(input) // Convert the input to integer
		if err2 != nil {
			log.Println("Conversion error:", err2)
			return
		}

		recipeIndex := recipeNum - 1                          // We subtract 1 to get the recipe according to the user.
		if recipeIndex >= 0 && recipeIndex < len(recipeSli) { // Checks that 'recipeIndex' is actually within range of our array.
			fmt.Println("You added:", recipeSli[recipeIndex].recipeName)
			macrosForToday.tCalories += recipeSli[recipeIndex].total.tCalories
			macrosForToday.tFats += recipeSli[recipeIndex].total.tFats
			macrosForToday.tCarbohydrates += recipeSli[recipeIndex].total.tCarbohydrates
			macrosForToday.tProteins += recipeSli[recipeIndex].total.tProteins
		} else {
			fmt.Println("Invalid recipe number")
		}
	}
}

// removeRecipe removes a recipe from the recipeSli based on user input. It asks for
// the recipe number to be deleted, subtracts 1 from the value, and then removes the
// corresponding recipe from the slice. It then updates the recipe numbers for the
// remaining recipes and returns the updated slice.
func removeRecipe(recipeSli []recipe) []recipe {
	fmt.Println("Removing recipe...")

	// Asks for which recipe to be deleted. We will subtract 1 from the value, because we dont expect the user to take
	// into account the fact that we actually start at 0.
	var recipeNum int
	fmt.Print("Enter the recipe number you wish to remove>")
	_, err := fmt.Scan(&recipeNum)
	if err != nil {
		log.Println("Scanning recipe number error:", err)
		return nil
	}
	fmt.Println("Removing:", recipeSli[recipeNum-1].recipeName)

	// Creates the left side of the new slice, without the appropriate element.
	leftSli := make([]recipe, recipeNum)
	leftSli = recipeSli[0 : recipeNum-1]

	// Creates the right side of the new slice, without the appropriate element.
	rightSli := make([]recipe, recipeNum)
	rightSli = recipeSli[recipeNum:]

	// Creates a slice with the left and right slices.
	leftAndRightSli := make([]recipe, 0)
	leftAndRightSli = append(leftAndRightSli, leftSli...)
	leftAndRightSli = append(leftAndRightSli, rightSli...)

	// Updates the numbers of the recipe.
	leftAndRightSli = updateRecipeNumbers(leftAndRightSli)

	return leftAndRightSli
}

// When a recipe gets removed, the recipe numbers past the removed recipe, the number will have to be updated.
// updateRecipeNumbers is used for this, by starting from the beginning and giving each recipe an increasing number, from 1 to the length of the list.
func updateRecipeNumbers(recipeSli []recipe) []recipe {
	for i := 0; i < len(recipeSli); i++ {
		recipeSli[i].recipeNumber = uint(i + 1)
	}

	return recipeSli
}

// updateFile is a function that updates the "recipes_go.txt" file with the recipes from the recipeSli.
// It opens the file in write-only mode with truncate and create options.
// Then, it loops through the recipeSli and writes the recipe number, meal type, name, and ingredients to the file.
// Finally, it adds a newline character at the end of each recipe.
func updateFile(recipeSli []recipe) {
	fmt.Println("Updating file...")

	// Opens the file in append mode with 0644 permissions.
	// os.O_WRONLY is used for write-only access, os.O_TRUNC truncates the file before opening, and os.O_CREATE creates
	//the file if it does not exist yet. When using these flags together, the file will be either created if it does not
	//exist or wiped clean if it does, ready for overwriting. You can then use the *os.File value to write data into the file.
	file, err := os.OpenFile("recipes_go.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Opening file error:", err)
	}

	for i := 0; i < len(recipeSli); i++ {

		// Adds the number, meal type and name to the file.
		_, err2 := fmt.Fprintf(file, "%d %c %s ", recipeSli[i].recipeNumber, recipeSli[i].mealType, recipeSli[i].recipeName)
		if err2 != nil {
			log.Fatal("Writing to file error:", err2)
		}

		// Adds the ingredients part of our recipe to the file.
		for _, j := range recipeSli[i].ingredients {
			_, err3 := fmt.Fprintf(file, "%s %.2f %.2f %.2f %.2f %.2f ", j.name, j.kilograms, j.calories, j.fats, j.carbohydrates, j.protein)
			if err3 != nil {
				log.Fatal("Writing to file error:", err3)
			}
		}

		// Adds a newline at the end.
		_, err4 := fmt.Fprintf(file, "\n")
		if err4 != nil {
			log.Fatal("Adding newline to text file error:", err4)
			return
		}
	}
	err5 := file.Close()
	if err5 != nil {
		log.Fatal("Closing file error:", err5)
	}
}

// removeIngredient removes an ingredient from a recipe and returns the updated list of ingredients.
// It prompts the user for the index of the ingredient to remove, creates left and right slices excluding the
// ingredient, combines them into a new slice, and updates the recipe's list of ingredients.
func removeIngredient(parsedRecipe recipe) []ingredient {
	// I already ask which recipe the user wants to change in "editRecipe". Change this:
	fmt.Println("Removing ingredient...")

	// Asks the user for which ingredient to remove.
	var ingredientIndex int
	fmt.Print("Enter the ingredient you wish to remove>")
	_, err := fmt.Scan(&ingredientIndex)
	if err != nil {
		log.Println("Scanning ingredient number error:", err)
		return nil
	}

	fmt.Println("Removing:", parsedRecipe.ingredients[ingredientIndex].name)

	// Creates the left side of the new slice, without the appropriate element.
	leftSli := make([]ingredient, ingredientIndex)
	leftSli = parsedRecipe.ingredients[0 : ingredientIndex-1] // We minus one, because we don't expect the user to take that fact that we start at 0, into account.

	// Creates the right side of the new slice, without the appropriate element.
	rightSli := make([]ingredient, ingredientIndex)
	rightSli = parsedRecipe.ingredients[ingredientIndex:]

	// Creates a slice with the left and right slices.
	leftAndRightSli := make([]ingredient, 0)
	leftAndRightSli = append(leftAndRightSli, leftSli...)
	leftAndRightSli = append(leftAndRightSli, rightSli...)

	// Sets the recipes ingredients to be 'leftAndRightSli'.
	parsedRecipe.ingredients = leftAndRightSli

	return parsedRecipe.ingredients
}

// changeIngredient modifies a specific ingredient within a recipe. It takes a recipe as input and prompts the user
// to select the index of the ingredient in the recipe's ingredients slice that they want to change. Then, it prompts the
// user to specify which part of the ingredient they want to change, such as name, grams, calories, fats, carbs, or protein.
// Depending on the user's input, the corresponding field of the selected ingredient is updated. The function continues to
// prompt the user for input until a valid input is received. Finally, it returns the updated ingredients slice for the recipe.
func changeIngredient(recip recipe) []ingredient {
	fmt.Println("Changing ingredients...")
	for {
		// Asks the user for which in the slice to change.
		// Also, whenever 'ingredIndx' is called, we minus one, because the user is not expected to know that we start at 0.
		var ingredIndx int
		fmt.Print("Enter which ingredient you want to change (in numbers)>")
		_, err4 := fmt.Scan(&ingredIndx)
		if err4 != nil {
			log.Println("Ingredient index error:", err4)
			return nil
		}

		// Asks the user for which part of the ingredient to change.
		var ingredPart string
		fmt.Print("Enter what part of the ingredient you want to change\n" +
			"	name\n" +
			"	grams\n" +
			"	calories\n" +
			"	fats\n" +
			"	carbs\n" +
			"	protein\n" +
			">")
		_, err5 := fmt.Scan(&ingredPart)
		if err5 != nil {
			log.Println("Ingredient part error:", err5)
			return nil
		}

		// If the user wants to correct the name.
		if ingredPart == "name" {
			var correctedName string
			fmt.Print("Enter corrected name>")
			_, err6 := fmt.Scan(&correctedName)
			if err6 != nil {
				log.Println("Corrected name error:", err6)
				return nil
			}

			recip.ingredients[ingredIndx-1].name = correctedName
			break

			// If the user wants to correct the grams.
		} else if ingredPart == "grams" {
			var correctedGrams float32
			fmt.Print("Enter corrected grams>")
			_, err7 := fmt.Scan(&correctedGrams)
			if err7 != nil {
				log.Fatal("Corrected grams error:", err7)
			}

			recip.ingredients[ingredIndx-1].kilograms = correctedGrams
			break

			// If the user wants to correct the calories.
		} else if ingredPart == "calories" {
			var correctedCalories float32
			fmt.Print("Enter corrected calories>")
			_, err8 := fmt.Scan(&correctedCalories)
			if err8 != nil {
				log.Fatal("Corrected calories error:", err8)
			}

			recip.ingredients[ingredIndx-1].calories = correctedCalories
			break

			// If the user wants to correct the fats.
		} else if ingredPart == "fats" {
			var correctedFats float32
			fmt.Print("Enter corrected fats>")
			_, err9 := fmt.Scan(&correctedFats)
			if err9 != nil {
				log.Fatal("Corrected fats error:", err9)
			}

			recip.ingredients[ingredIndx-1].fats = correctedFats
			break

			// If the user wants to correct the carbs
		} else if ingredPart == "carbs" {
			var correctedCarbs float32
			fmt.Print("Enter corrected carbs>")
			_, err10 := fmt.Scan(&correctedCarbs)
			if err10 != nil {
				log.Fatal("Corrected carbs error:", err10)
			}

			recip.ingredients[ingredIndx-1].carbohydrates = correctedCarbs
			break

			// If the user wants to correct the protein
		} else if ingredPart == "protein" {
			var correctedProtein float32
			fmt.Print("Enter corrected protein>")
			_, err11 := fmt.Scan(&correctedProtein)
			if err11 != nil {
				log.Fatal("Corrected protein error:", err11)
			}

			recip.ingredients[ingredIndx-1].protein = correctedProtein
			break

			// If the input is invalid.
		} else {
			fmt.Println("Invalid input, try again!")
		}
	}

	return recip.ingredients
}

// addExtraIngredient prompts the user to enter an additional ingredient and adds it to the given recipe.
// Otherwise, it functions the same way as getIngredients, except it is not in a loop.
func addExtraIngredient(parsedRecipe recipe) []ingredient {
	fmt.Println("Adding more ingredients...")

	var ingre ingredient
	fmt.Print("Type your ingredient (or q)>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal("Scanning ingredient name error:", err)
	}

	input := scanner.Text()

	ingre.name = input
	fmt.Print("Enter ingredient kilograms>")
	_, err := fmt.Scan(&ingre.kilograms)
	if err != nil {
		log.Fatal("Scanning ingredient kilograms error:", err)
	}

	fmt.Print("Enter ingredient calories (pr. 100g)>")
	_, err = fmt.Scan(&ingre.calories)
	if err != nil {
		log.Fatal("Scanning ingredient calories error:", err)
	}

	fmt.Print("Enter ingredient fats (pr. 100g)>")
	_, err = fmt.Scan(&ingre.fats)
	if err != nil {
		log.Fatal("Scanning ingredient fats error:", err)
	}

	fmt.Print("Enter ingredient carbohydrates (pr. 100g)>")
	_, err = fmt.Scan(&ingre.carbohydrates)
	if err != nil {
		log.Fatal("Scanning ingredient carbohydrates error:", err)
	}

	fmt.Print("Enter ingredient protein (pr. 100g)>")
	_, err = fmt.Scan(&ingre.protein)
	if err != nil {
		log.Fatal("Scanning ingredient protein error:", err)
	}

	ingre.multiplier = ingre.kilograms / 100

	parsedRecipe.ingredients = append(parsedRecipe.ingredients, ingre)

	return parsedRecipe.ingredients
}

// randomMeal, gets you a random meal for you to make, in case you don't know what to make.
func randomMeal(recipeSli []recipe) {
	fmt.Println("Printing random recipe to make...")
	var randomNumber int

	for {
		randomNumber = rand.Intn(len(recipeSli))

		// In case the meal type is not 'L' or 'D' (lunch or dinner), we skip this iteration to get one.
		if recipeSli[randomNumber].mealType == 'K' || recipeSli[randomNumber].mealType == 'S' || recipeSli[randomNumber].mealType == 'B' {
			fmt.Println("Snack, side or breakfast meal has been picked, finding different meal...")
			continue
		} else {
			fmt.Printf("Recipe %d: %s, %.2f, %.2f, %.2f, %.2f\n", recipeSli[randomNumber].recipeNumber, recipeSli[randomNumber].recipeName, recipeSli[randomNumber].total.tCalories, recipeSli[randomNumber].total.tFats, recipeSli[randomNumber].total.tCarbohydrates, recipeSli[randomNumber].total.tProteins)
			break
		}
	}
}

func planWeek(recipeSli []recipe) {
	fmt.Println("Planning out week...")
	var mondayMacros totalMacros
	var tuesdayMacros totalMacros
	var wednesdayMacros totalMacros
	var thursdayMacros totalMacros
	var fridayMacros totalMacros
	var satudayMacros totalMacros
	var sundayMacros totalMacros
	var weekMacros []totalMacros
	var recipeInput string

	weekMacros = append(weekMacros, mondayMacros, tuesdayMacros, wednesdayMacros, thursdayMacros, fridayMacros, satudayMacros, sundayMacros)

	for i := 0; i < 7; i++ {
		for {
			// Prints the current macros, for all the recipes that the user has added (The values are always zero on first iteration).
			fmt.Printf("Current macros day %d:\n"+
				"	Calories: 	%.2f\n"+
				"	Fats: 		%.2f\n"+
				"	Carbs: 		%.2f\n"+
				"	Protein: 	%.2f\n", i+1, weekMacros[i].tCalories, weekMacros[i].tFats, weekMacros[i].tCarbohydrates, weekMacros[i].tProteins)

			// Asks the user for what recipes they wish to eat.
			fmt.Printf("Enter recipe to eat for day %d (or q)>", i+1)
			_, err := fmt.Scan(&recipeInput)
			if err != nil {
				fmt.Println("Error input!")
				return
			}

			// When the user has added all the desired recipes for a day, input "p" stops the loop, prints the total macros
			// for all the added recipes and changes to the next day.
			if recipeInput == "q" {
				fmt.Printf("Your final macros are for day %d are:\n", i+1)
				fmt.Printf(
					"	Calories: 	%.2f\n"+
						"	Fats: 		%.2f\n"+
						"	Carbs: 		%.2f\n"+
						"	Protein: 	%.2f\n", weekMacros[i].tCalories, weekMacros[i].tFats, weekMacros[i].tCarbohydrates, weekMacros[i].tProteins)
				break
			}

			// Converts "recipeInput" to an integer and stores the value in "convertedRecipeInput".
			convertedRecipeInput, err2 := strconv.Atoi(recipeInput)
			if err2 != nil {
				fmt.Println("Conversion error:", err2)
				return
			}

			// We add the recipe values to the total macros.
			convertedRecipeInput--                                                  // We subtract 1 to get the recipe according to the user.
			if convertedRecipeInput >= 0 && convertedRecipeInput < len(recipeSli) { // Checks that 'recipeIndex' is actually within range of our array.
				fmt.Println("You added:", recipeSli[convertedRecipeInput].recipeName)
				weekMacros[i].tCalories += recipeSli[convertedRecipeInput].total.tCalories
				weekMacros[i].tFats += recipeSli[convertedRecipeInput].total.tFats
				weekMacros[i].tCarbohydrates += recipeSli[convertedRecipeInput].total.tCarbohydrates
				weekMacros[i].tProteins += recipeSli[convertedRecipeInput].total.tProteins
			} else {
				fmt.Println("Invalid recipe number")
			}

		}
	}

	// Prints the total macros for an entire week.
	for i, macro := range weekMacros {
		fmt.Printf("Total macros for day %d: \n", i+1)
		fmt.Printf(
			"	Calories: 	%.2f\n"+
				"	Fats: 		%.2f\n"+
				"	Carbs: 		%.2f\n"+
				"	Protein: 	%.2f\n", macro.tCalories, macro.tFats, macro.tCarbohydrates, macro.tProteins)
	}
}
