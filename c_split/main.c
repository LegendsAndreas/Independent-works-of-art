#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct {
    char word[100];
} str;

// Due to the way C works, we cant just have a function that returns an array and get the size with ye'old "sizeof".
// So, to accurately determine the length of the array, we have "count".
typedef struct {
    str* subStrings;
    unsigned int count;
} subString;

subString splitString(const char *words, const char *splicer);

int main(void) {
    char *words = "I am going to the toilet";
    subString subStrings = splitString(words, " ");

    printf("Count of substrings: %d\n", subStrings.count);

    for (int i = 0; i < subStrings.count; ++i) {
        printf("Substring %d: %s\n", i, subStrings.subStrings[i].word);
    }

    free(subStrings.subStrings);
    return 0;
}

/**
 * @brief Splits a given string into multiple substrings based on a splicer character.
 *
 * The input string will be split wherever the splicer character is found. Each
 * substring is stored in a dynamically allocated array, which is returned in a
 * struct containing both the array and the count of substrings.
 *
 * @param words The input string to be split.
 * @param splicer The character used to split the input string.
 * @return A struct containing the array of substrings and the count of these substrings.
 *
 * @note The function will terminate the program with an error message if:\n
 *       - The input string is empty.\n
 *       - The length of a collected substring exceeds 49 characters.\n
 *       - The total number of substrings exceeds 100.\n
 *       - Memory allocation for substrings fails.\n
 */
subString splitString(const char *words, const char *splicer) {
    if (strlen(words) == 0) {
        printf("Empty string\n");
        exit(1);
    }

    str tempSubStrings[100]; // Where our sub strings will be temporarily stored.
    char tempString[50]; // Where the current sub string will be stored.
    unsigned int tempStringLen = 0; // The length of the current temporary sub string.
    unsigned int wordCount = 0; // The amount of counted sub strings.

    // We +1, otherwise the loop would end before we could add the last string.
    for (int i = 0; i < strlen(words) + 1; ++i) {
        if (words[i] != splicer[0] && words[i] != '\0') {
            tempString[tempStringLen++] = words[i];
        }
        else
        {
            // We check if the length of the string will be too long.
            if (tempStringLen > 49) {
                printf("The collected string is too long\n");
                exit(1);
            }

            // When we encounter the designated splicer, we add the collected string to "tempSubString", at the assigned wordCount.
            // We also empty the temporary string, and set "temptStringLen" to 0, to start over.
            tempString[tempStringLen] = '\0'; // Null-terminate the substring
            strcpy(tempSubStrings[wordCount].word, tempString);
            strcpy(tempString, "");
            tempStringLen = 0;
            wordCount++;
        }
    }

    if (wordCount > 100) {
        printf("You have exceeded the max amount of strings (100)\n");
        exit(1);
    }

    // Allocate memory for the sub strings
    str *subStrngs = (str *)malloc(wordCount * sizeof(str));
    if (subStrngs == NULL) {
        printf("Memory allocation failed\n");
        exit(1);
    }

    // Copy substrings to dynamically allocated memory
    for (int i = 0; i < wordCount; ++i) {
        strcpy(subStrngs[i].word, tempSubStrings[i].word);
    }

    subString result;
    result.subStrings = subStrngs;
    result.count = wordCount;

    return result;
}

/*
#include <stdio.h>
#include "string.h"
#include "stdlib.h"

typedef struct {
    char word[100];
}str;

str* splitString(char *words, const char *splicer);

int main(void) {
    char *words = "Hello gamer!";
    str *subStrings = splitString(words, " ");

    printf("%llu\n", sizeof(*subStrings)/ sizeof(str));

    free(subStrings);

    return 0;
}
