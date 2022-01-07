# Count word occurrences (exercise)

Write a program that counts occurrences of each word in a given string.

For example, the occurrences of each word in the string "I have a cat. My cat
is grey. Her name is Carrie." are the following:

"I": 1
"have": 1
"a": 1
"cat": 2
"My": 1
"is": 2
"grey": 1
"Her": 1
"name": 1
"Carrie": 1.


Implement a logic as a function

```
func wordOccurrences(s string) (words []string, occurrences []uint)
```

where `s` is an input string, `words` is a slice of unique words in `s` (in any
order), `occurrences` is a slice of occurrences, where the following must hold:

* `len(words) == len(occurrences)`
* `occurrences[i]` is the occurrence of the word `word[i]` in `s`.

For example, `wordOccurrences("I have a cat. My cat is grey. Her name is
Carrie.")` must return `([]string{"I", "have", "a", "cat", "My", "is", "grey",
"Her", "name", "Carrie"}, []uint{1, 1, 1, 2, 1, 2, 1, 1, 1, 1})`.

Implement `main()` function that reads an input string from the keyboard.

Implement tests for the program. The tests must cover all representative inputs.

## Examples

### Example 1

```
Enter a string: Where there is a will, there is a way.
Occurrences:
Where: 1
there: 2
is: 2
a: 2
will: 1
way: 1
```

### Example 2
```
Enter a string: Много снега - много хлеба, много воды - много травы.
Occurrences:
Много: 1
снега: 1
много: 3
хлеба: 1
воды: 1
травы: 1
```
