"What if i created the worst chance program ever?", is essentially what i thought of when i created this program and i thought it would be pretty fun.

The way the program works, is that you have an array of integers from 1-10 and for 1,000 times, we iterate through the array until the variable "throw" is equal to 1, which is randomly acuired from the statement: rand.Intn(2). So, a 1/2 or 50% chance, of getting a 1.

The problem with this, as you can imagine, is that every integer after the first in the array, gets its chance of a succesful throw cut in half. So, the second integer in our array has a 25% chance of succeding, the third has a 12.5% chance of succeding and so on.
