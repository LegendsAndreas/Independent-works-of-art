package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type recipe struct {
	recipeNumber uint
	mealType     rune // B = Breakfast, L = Lunch, D = Dinner, S = Side, K = Snack
	recipeName   string
	ingredients  []ingredient
	total        totalMacros
}

type ingredient struct {
	name          string
	kilograms     float32
	calories      float32
	fats          float32
	carbohydrates float32
	protein       float32
	multiplier    float32

	/* EXAMPLE
	Tulip Pulled Pork	// Name
	125					// Measured in grams
	133					// Calories/Kcal
	6					// Fats
	2					// Carbs
	17					// Protein
	125/100 = 1.25 		// Multiplier. So the amount of kilograms divided with 100.
	*/
}

type totalMacros struct {
	tCalories      float32
	tFats          float32
	tCarbohydrates float32
	tProteins      float32
}

func main() {
	var input string
	var recipes []recipe

	recipes = initializeRecipes(recipes)

	for {
		// Asks what the user wants to do and stores the input in 'input'.
		fmt.Printf("Enter your action\n" +
			"	new\n" +
			"	edit\n" +
			"	remove\n" +
			"	print\n" +
			"	plan\n" +
			"	q\n>")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal("Scanning input error:", err)
			return
		}

		// If the user wants to add a new recipe.
		if input == "new" {
			recipes = newRecipe(recipes)

			// If the user wants to edit an existing recipe.
		} else if input == "edit" {
			recipes = editRecipe(recipes)

			// If the user wants to remove a recipe.
		} else if input == "remove" {
			recipes = removeRecipe(recipes)
			// Updates the file "recipes_go"
			updateFile(recipes)

		} else if input == "print" {
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

			if printInput == "basic" {
				printBasicRecipes(recipes)
			} else if printInput == "single" {
				printSingleRecipe(recipes)
			} else if printInput == "everything" {
				printEverything(recipes)
			} else {
				fmt.Println("Invalid input")
			}

		} else if input == "plan" {
			planDay(recipes)
		} else if input == "q" {
			fmt.Println("Exiting the program...")
			break
		} else {
			fmt.Println("Invalid input")
		}

	}
}

func initializeRecipes(recipeArray []recipe) []recipe {
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
			// Perhaps, i could rewrite this with fprintf(name, kilo, cal, fats, carbs, pro)
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

		// Create a new recipe and append it to the recipe array
		newRecipe := recipe{
			recipeNumber: uint(recipeNumber),
			mealType:     mealType,
			recipeName:   recipeName,
			ingredients:  ingredients,
			total:        total,
		}
		recipeArray = append(recipeArray, newRecipe)
	}

	err = recipesFile.Close()
	if err != nil {
		return nil
	}

	return recipeArray
}

func newRecipe(recipeSlice []recipe) []recipe {
	fmt.Println("Creating new recipe...")

	var tempRecipe recipe                                     // temptRecipe will be added to slice 'recipeSlice'
	tempRecipe.recipeNumber = uint(len(recipeSlice)) + 1      // Gets recipe number
	tempRecipe.mealType = getMealType()                       // Gets the meal type
	tempRecipe.recipeName = getRecipeName()                   // Gets recipe name
	tempRecipe.ingredients = getIngredients()                 // Gets ingredients
	tempRecipe.total = getTotalMacros(tempRecipe.ingredients) // Gets total macros

	addRecipeToFile(tempRecipe) // Adds the recipe to the file, so that it is saved and does not disappear upon restarting the program.

	fmt.Println("Inserting recipe into array... ")
	recipeSlice = append(recipeSlice, tempRecipe) // Appends the recipe in the recipe array 'rArr'
	return recipeSlice
}

func getMealType() rune {
	reader := bufio.NewReader(os.Stdin)
	var tMeal rune

	for {
		// If we were to use fmt.Scan, it would an integer, since a rune is almost the same as an int32.
		// So instead, we use read a string from the user, that we can convert to a rune.
		fmt.Print("Enter meal type>")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Reading from standard input error:", err)
		}

		// One might think that this line is unnecessary, since our delimiter is a newline, BUT the way that
		//.ReadString() Works, is that it reads up until and INCLUDING the delimiter. So, we remove the newline.
		input = strings.TrimSpace(input)

		// If the entered input is longer than a single character, we skip this iteration.
		if len(input) > 1 {
			fmt.Println("Too long input, try again!")
			continue
		}

		if input != "B" && input != "L" && input != "D" && input != "S" && input != "K" {
			fmt.Println("Invalid meal type.")
			continue
		}

		// Assuming that the input is not nothing, we convert our entered string to a rune.
		// Else, if the user for example just presses enter or the space bar, the else-statement would execute.
		if len(input) > 0 {
			tMeal = rune(input[0])
			break
		} else {
			fmt.Println("Input is nothing, try again!")
			continue
		}
	}

	return tMeal
}

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
func getIngredients() []ingredient {
	tIngredients := make([]ingredient, 0)
	var input string
	for {
		var ingre ingredient
		fmt.Print("Type your ingredient (or q)>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Fatal("Scanning ingredient name error:", err)
		}

		input = scanner.Text()

		if input == "q" {
			fmt.Println("Exiting the ingredients program...")
			break
		}

		ingre.name = input
		fmt.Println("Enter ingredient kilograms>")
		_, err := fmt.Scan(&ingre.kilograms)
		if err != nil {
			log.Fatal("Scanning ingredient kilograms error:", err)
		}

		fmt.Println("Enter ingredient calories (pr. 100g)>")
		_, err = fmt.Scan(&ingre.calories)
		if err != nil {
			log.Fatal("Scanning ingredient calories error:", err)
		}

		fmt.Println("Enter ingredient fats (pr. 100g)>")
		_, err = fmt.Scan(&ingre.fats)
		if err != nil {
			log.Fatal("Scanning ingredient fats error:", err)
		}

		fmt.Println("Enter ingredient carbohydrates (pr. 100g)>")
		_, err = fmt.Scan(&ingre.carbohydrates)
		if err != nil {
			log.Fatal("Scanning ingredient carbohydrates error:", err)
		}

		fmt.Println("Enter ingredient protein (pr. 100g)>")
		_, err = fmt.Scan(&ingre.protein)
		if err != nil {
			log.Fatal("Scanning ingredient protein error:", err)
		}

		ingre.multiplier = ingre.kilograms / 100

		tIngredients = append(tIngredients, ingre)

	}

	return tIngredients
}

func getTotalMacros(rArr []ingredient) totalMacros {
	var total totalMacros
	for _, r := range rArr {
		total.tCalories += r.calories * r.multiplier
		total.tFats += r.fats * r.multiplier
		total.tCarbohydrates += r.carbohydrates * r.multiplier
		total.tProteins += r.protein * r.multiplier
	}
	return total
}

func printBasicRecipes(recipes []recipe) {
	fmt.Println("Printing recipe...")
	for _, r := range recipes {
		fmt.Printf("Recipe %d: %s, %.2f, %.2f, %.2f, %.2f\n", r.recipeNumber, r.recipeName, r.total.tCalories, r.total.tFats, r.total.tCarbohydrates, r.total.tProteins)
	}
}

func printEverything(recipes []recipe) {
	fmt.Println("Printing recipe...")
	// Prints the recipe number, meal type, name and all the macros for it.
	for _, r := range recipes {
		fmt.Printf("%d, %c: %s\n", r.recipeNumber, r.mealType, r.recipeName)
		fmt.Printf("	- Total Calories: %.2f\n", r.total.tCalories)
		fmt.Printf("	- Total Proteins: %.2f\n", r.total.tProteins)
		fmt.Printf("	- Total Fats    : %.2f\n", r.total.tFats)
		fmt.Printf("	- Total Carbs   : %.2f\n", r.total.tCarbohydrates)

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

	} else if part == "ingredients" {
		// If the user wants to change the ingredients.

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

	}
	// Updates the file.
	updateFile(repSli)

	return repSli

}

func printSingleRecipe(recipeArray []recipe) {
	fmt.Println("Printing recipe...")
	var recipeNumber uint
	fmt.Print("Enter the recipe number>")
	_, err := fmt.Scan(&recipeNumber)
	if err != nil {
		log.Fatal("Scanning recipe number error:", err)
	}
	for _, r := range recipeArray {
		if r.recipeNumber == recipeNumber {
			fmt.Println(r)
			break
		}
	}
}

func planDay(recipeArray []recipe) {
	var macrosForToday totalMacros
	input := ""
	for {
		fmt.Printf("Current macros:\n"+
			"	Calories: 	%.2f\n"+
			"	Fats: 		%.2f\n"+
			"	Carbs: 		%.2f\n"+
			"	Protein: 	%.2f\n", macrosForToday.tCalories, macrosForToday.tFats, macrosForToday.tCarbohydrates, macrosForToday.tProteins)

		fmt.Print("Enter the recipe number, that you want to eat today (or q)> ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Println("Scanning error:", err)
			return
		}

		if input == "q" {
			fmt.Println("Your final macros are for the day are:")
			fmt.Printf("Current macros:\n"+
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

		recipeIndex := recipeNum - 1
		if recipeIndex >= 0 && recipeIndex < len(recipeArray) { // Checks that 'recipeIndex' is actually within range of our array.
			fmt.Println("You added:", recipeArray[recipeIndex].recipeName)
			macrosForToday.tCalories += recipeArray[recipeIndex].total.tCalories
			macrosForToday.tFats += recipeArray[recipeIndex].total.tFats
			macrosForToday.tCarbohydrates += recipeArray[recipeIndex].total.tCarbohydrates
			macrosForToday.tProteins += recipeArray[recipeIndex].total.tProteins
		} else {
			fmt.Println("Invalid recipe number")
		}
	}
}

func removeRecipe(recipeArray []recipe) []recipe {
	fmt.Println("Removing recipe...")
	var recipeNum int
	fmt.Print("Enter the recipe number you wish to remove>")
	_, err := fmt.Scan(&recipeNum)
	if err != nil {
		log.Println("Scanning recipe number error:", err)
		return nil
	}
	fmt.Println("Removing:", recipeArray[recipeNum-1].recipeName)

	// Creates the left side of the new slice, without the appropriate element.
	leftSli := make([]recipe, recipeNum)
	leftSli = recipeArray[0 : recipeNum-1]

	// Creates the right side of the new slice, without the appropriate element.
	rightSli := make([]recipe, recipeNum)
	rightSli = recipeArray[recipeNum:]

	// Creates a slice with the left and right slices.
	leftAndRightSli := make([]recipe, 0)
	leftAndRightSli = append(leftAndRightSli, leftSli...)
	leftAndRightSli = append(leftAndRightSli, rightSli...)

	// Updates the numbers of the recipe.
	leftAndRightSli = updateRecipeNumbers(leftAndRightSli)

	return leftAndRightSli
}

/*
*
When a recipe gets removed, the recipe numbers past the removed recipe, the number will have to be updated.
updateRecipeNumbers is used for this, by starting from the beginning and giving each recipe an increasing number, from 1 to the length of the list.
*/
func updateRecipeNumbers(recipeArray []recipe) []recipe {
	for i := 0; i < len(recipeArray); i++ {
		recipeArray[i].recipeNumber = uint(i + 1)
	}

	return recipeArray
}

func updateFile(recipeArray []recipe) {
	fmt.Println("Updating file...")
	// Opens the file in append mode with 0644 permissions.

	// os.O_WRONLY is used for write-only access, os.O_TRUNC truncates the file before opening, and os.O_CREATE creates
	//the file if it does not exist yet. When using these flags together, the file will be either created if it does not
	//exist or wiped clean if it does, ready for overwriting. You can then use the *os.File value to write data into the file.
	file, err := os.OpenFile("recipes_go.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Opening file error:", err)
	}

	for i := 0; i < len(recipeArray); i++ {

		// Adds the number, meal type and name to the file.
		_, err2 := fmt.Fprintf(file, "%d %c %s ", recipeArray[i].recipeNumber, recipeArray[i].mealType, recipeArray[i].recipeName)
		if err2 != nil {
			log.Fatal("Writing to file error:", err2)
		}

		// Adds the ingredients part of our recipe to the file.
		for _, j := range recipeArray[i].ingredients {
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

func removeIngredient(recip recipe) []ingredient {
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

	fmt.Println("Removing:", recip.ingredients[ingredientIndex].name)

	// Creates the left side of the new slice, without the appropriate element.
	leftSli := make([]ingredient, ingredientIndex)
	//leftSli = recipeArray[recipeNum-1].ingredients[0 : recipeNum-1]
	leftSli = recip.ingredients[0 : ingredientIndex-1] // We minus one, because we don't expect the user to take that fact that we start at 0, into account.

	// Creates the right side of the new slice, without the appropriate element.
	rightSli := make([]ingredient, ingredientIndex)
	//rightSli = recipeArray[recipeNum-1].ingredients[recipeNum:]
	rightSli = recip.ingredients[ingredientIndex:]

	// Creates a slice with the left and right slices.
	leftAndRightSli := make([]ingredient, 0)
	leftAndRightSli = append(leftAndRightSli, leftSli...)
	leftAndRightSli = append(leftAndRightSli, rightSli...)

	// Sets the recipes ingredients to be 'leftAndRightSli'.
	// recipeArray[recipeNum-1].ingredients = leftAndRightSli
	recip.ingredients = leftAndRightSli

	return recip.ingredients
}

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

func addExtraIngredient(recip recipe) []ingredient {
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

	recip.ingredients = append(recip.ingredients, ingre)

	return recip.ingredients
}
