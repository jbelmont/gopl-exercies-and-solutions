# Chapter 11 Exercises

Exercise 11.1: Write tests for the charcount program in Section 4.3.

Exercise 11.2: Write a set of tests for IntSet (§6.5) that checks that its behavior after each operation is equivalent to a set based on built-in maps. Save your implementation for benchmarking in Exercise 11.7.

Exercise 11.3: TestRandomPalindromes only tests palindromes. Write a randomized test that generates and verifies non-palindromes.

Exercise 11.4: Modify randomPalindrome to exercise IsPalindrome’s handling of punctuation and spaces.

Exercise 11.6: Write benchmarks to compare the PopCount implementation in Section 2.6.2 with your solutions to Exercise 2.4 and Exercise 2.5. At what point does the table-based approach break even?

Exercise 11.7: Write benchmarks for Add, UnionWith, and other methods of *IntSet (§6.5) using large pseudo-random inputs. How fast can you make these methods run? How does the choice of word size affect performance? How fast is IntSet compared to a set implementation based on the built-in map type?