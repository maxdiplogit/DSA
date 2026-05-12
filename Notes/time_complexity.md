# Time & Space Complexity — Building Real Intuition

A complete walkthrough built from the ground up. Every question, every diagram, every "wait, why?" moment.

---

## Part 1: What is time complexity, really?

Time complexity is a way of describing **how the number of steps an algorithm takes grows as the input gets bigger**.

That's the whole concept. Everything else — Big O, log n, n², all of it — is just notation for talking about this one idea precisely.

### Why "max number of steps"?

When we describe an algorithm's complexity, we use the **worst case** — the maximum number of steps it might take. Why? Because that's the guarantee. The best case (we got lucky) isn't useful. The worst case tells us what we're committing to: "no matter how bad the input is, the algorithm won't take more steps than this."

### Why we care about how it *grows*, not the exact number

We don't care that an algorithm takes 10 steps on an array of 10 elements. We care about: *what happens as the array gets bigger?*

- Double the input → does the work double? (linear)
- Double the input → does the work stay the same? (constant)
- Double the input → does the work go up by just 1 step? (logarithmic)
- Double the input → does the work quadruple? (quadratic)

That growth pattern is what we name with Big O.

### What counts as a "step"?

Any single basic operation — arithmetic, comparison, array access, variable assignment — is one step. The exact number doesn't matter, because Big O ignores constants. An algorithm that does 3n steps and one that does 5n steps are both O(n). We only care about the **shape** of growth.

---

## Part 2: The simplest complexities — O(1), O(n), O(n²)

### O(1) — Constant time

```python
def get_first_element(arr):
    return arr[0]
```

One step, no matter how big the array is. The "1" doesn't literally mean one step — it means some fixed number that doesn't grow with the input.

### O(n) — Linear time

```python
def sum_array(arr):
    total = 0
    for x in arr:
        total += x
    return total
```

If the array has n elements, that's n additions. Double the array, double the work.

### O(n²) — Quadratic time

```python
def all_pairs(arr):
    for i in arr:
        for j in arr:
            print(i, j)
```

For every element, the inner loop runs through *all* of arr again. Array of 10 → 100 operations. Array of 100 → 10,000. **Doubling the input quadruples the work.**

### Mental shortcut for reading code

- No loops, just a few operations → **O(1)**
- One loop over the input → **O(n)**
- Two nested loops over the input → **O(n²)**
- Three nested loops → **O(n³)**

This covers about 70% of code you'll see.

### Why this matters in practice

Each step takes ~1 microsecond. With 1 million inputs:

- O(1): done instantly
- O(n): 1 second
- O(n²): about 11 days

This is why interviewers ask "can you do better?"

---

## Part 3: O(log n) — the magic of halving

**O(log n) means: double the input, the steps only go up by 1.**

How? By throwing away half the input at each step.

### The dictionary example

Find "mango" in a 1000-page dictionary:

1. Open page 500. Mango comes before this → throw away pages 500-1000. (500 pages left.)
2. Open page 250. Mango comes after → throw away 1-250. (250 left.)
3. Open page 375. And so on.

`1000 → 500 → 250 → 125 → 62 → 31 → 15 → 7 → 3 → 1` — about 10 steps.

### What log₂(n) actually means

It's just a fancy way of asking **"how many times do I need to halve n to get down to 1?"**

- log₂(8) = 3 (because 8 → 4 → 2 → 1)
- log₂(1000) ≈ 10
- log₂(1,000,000) ≈ 20

### Code shape: binary search

```python
def binary_search(arr, target):
    lo, hi = 0, len(arr) - 1
    while lo <= hi:
        mid = (lo + hi) // 2
        if arr[mid] == target: return mid
        elif arr[mid] < target: lo = mid + 1
        else: hi = mid - 1
```

The loop runs, but each iteration cuts the search space in half. That's the signature of O(log n).

### Comparing back to basics

For an input of 1,000,000:
- **O(1)**: 1 step
- **O(log n)**: ~20 steps
- **O(n)**: 1,000,000 steps

Log n is *almost* as fast as constant time for any realistic input.

---

## Part 4: O(n log n) — merge sort and the levels mystery

This is where most people get stuck. Let me unpack it slowly.

### The recursion tree for merge sort with n = 8

```
Level 0:                  [5,3,8,1,9,2,7,4]              ← 1 piece of size 8
                          /                \
Level 1:           [5,3,8,1]              [9,2,7,4]      ← 2 pieces of size 4
                   /     \                  /     \
Level 2:       [5,3]    [8,1]          [9,2]    [7,4]    ← 4 pieces of size 2
               / \      / \            / \      / \
Level 3:     [5][3]   [8][1]         [9][2]   [7][4]     ← 8 pieces of size 1
```

### What is a "level"?

A level is one horizontal row in the tree. When you call merge_sort on the full array, that's level 0. When that call splits into two recursive calls, those two calls together are level 1. And so on.

### Why log n levels?

You start with 8 elements and keep halving: `8 → 4 → 2 → 1`. That's 3 halvings. The number of levels = number of halvings = **log₂(n)**.

### Why is n multiplied in n log n? (The key question)

**Every single level processes n total elements.** Count it:

- Level 0: 1 merge of size 8 → 8 work
- Level 1: 2 merges of size 4 each → 4 + 4 = 8 work
- Level 2: 4 merges of size 2 each → 2 + 2 + 2 + 2 = 8 work

The number of pieces **doubles** each level, but each piece is **half the size**. Those effects cancel out perfectly. **Every level does exactly n total work.**

So: `log n levels × n work per level = n log n total`.

### Why did someone invent merge sort?

The trick: merging two already-sorted arrays is cheap (just walk through them with two pointers, O(n) for total size). So the question became: "Can I get sorted arrays for free?"

The answer: keep splitting until each piece has 1 element. A list of 1 element is *already sorted by definition*. Then you merge your way back up.

### Binary search vs merge sort — what's the difference?

Both halve the input. But:

- **Binary search** throws away one half at each level. Only does constant work per level. → **O(log n)**
- **Merge sort** keeps both halves and works on all of them. Does n total work per level. → **O(n log n)**

| | Binary Search | Merge Sort |
|---|---|---|
| Halves the input? | Yes | Yes |
| Number of levels | log n | log n |
| Work per level | constant (look at one piece) | n (process all pieces) |
| Total work | log n | n × log n |
| Reaches at the bottom | One index (the answer) | All n elements as singletons |

**The halving is identical. What differs is how much work you do per level.**

---

## Part 5: O(2ⁿ) — fibonacci and the catastrophic case

### The fibonacci recursion tree for fib(5)

```
Level 0:                       f(5)                      ← 1 call
                              /    \
Level 1:                   f(4)    f(3)                  ← 2 calls
                          /   \    /   \
Level 2:                f(3) f(2) f(2) f(1)              ← 4 calls
                       /   \  / \  / \
Level 3:             f(2) f(1) ...                       ← 8 calls
                    /  \
Level 4:          f(1) f(0)                              ← ~16 calls
```

### The crucial difference from merge sort

Both functions call themselves twice. Both have base cases. So why is one log n levels deep and the other n levels deep?

**It's how fast the input shrinks.**

- Merge sort: array of 8 → calls on arrays of size **4**. The input is **halved**.
- Fibonacci: fib(8) → calls fib(**7**) and fib(**6**). The input shrinks by **1**.

How many steps from 1000 to 1?
- Halving: ~10 steps (1000 → 500 → 250 → ... → 1)
- Subtracting 1: 999 steps

That's the gap between log n and n levels.

### Why fibonacci becomes 2ⁿ

Calls double at each level (because each function makes 2 recursive calls):

- Level 0: 1 call
- Level 1: 2 calls
- Level 2: 4 calls
- Level 3: 8 calls
- Level k: 2ᵏ calls

In **merge sort**, this doubling stops after log n levels because the input becomes size 1. Max calls at the bottom = 2^(log n) = n. The doubling gets canceled by the halving.

In **fibonacci**, this doubling continues for n levels because the input only shrinks by 1. Max calls at the bottom = **2ⁿ**. Nothing cancels the doubling.

### The mental picture

- Merge sort's tree: **short and wide.** log n levels deep, n leaves at the bottom.
- Fibonacci's tree: **tall and explosively wide.** n levels deep, 2ⁿ leaves.

### Real numbers

For n = 30:
- Merge sort: ~5 levels deep, ~150 operations total
- Naive fibonacci: 30 levels deep, ~1 *billion* operations

### Memoization fixes this

With memoization, fib(5) is only computed once. The tree *looks* like it has 2ⁿ nodes, but you only do real work for n unique inputs. Memoized fibonacci is **O(n)** despite the recursive structure.

---

## Part 6: O(n!) — permutations and the worst case

n! (n-factorial) = n × (n-1) × (n-2) × ... × 2 × 1

- 5! = 120
- 10! = 3,628,800
- 15! ≈ 1.3 trillion
- 20! ≈ 2.4 quintillion

### The permutation tree for 4 elements

```
Level 0:                Pick 1st element            ← 4 choices
                       / | | \
Level 1:           A   B   C   D                    ← × 3 choices
                  /|\ /|\ /|\ /|\
Level 2:         (each branches into 3)             ← × 2 choices
                  ...
Level 3:        (each branches into 2)              ← × 1 choice
                  ...
                                                    Total leaves: 4! = 24
```

### Where does n! come from?

If you have n items and want to count all orderings:
- For slot 1: n choices
- For slot 2: n-1 choices remaining
- For slot 3: n-2 choices
- ...down to 1

Total: n × (n-1) × (n-2) × ... × 1 = **n!**

### Example: arranging A, B, C

3 × 2 × 1 = 6 orderings: ABC, ACB, BAC, BCA, CAB, CBA.

### Code shape

```python
def permutations(arr):
    if len(arr) <= 1: return [arr]
    result = []
    for i in range(len(arr)):
        rest = arr[:i] + arr[i+1:]
        for p in permutations(rest):
            result.append([arr[i]] + p)
    return result
```

---

## Part 7: The unified framework

For any algorithm, ask **two questions**:

1. **How deep is the tree?** (How many levels of recursion / how many times do we shrink the input?)
2. **How much work happens at each level?**

Then multiply.

| Complexity | Depth | Work per level | Where it shows up |
|------------|-------|----------------|-------------------|
| O(1) | 1 | constant | Direct access (arr[0]) |
| O(log n) | log n | constant | Binary search, balanced BST operations |
| O(n) | 1 (or n with constant work) | linear | Single pass, two pointers, sliding window |
| O(n log n) | log n | n | Merge sort, quicksort, sort + linear scan |
| O(n²) | n | n | Nested loops, brute-force pair finding |
| O(2ⁿ) | n | doubling | Naive recursion with 2 branches, no memo |
| O(n!) | n | shrinking factor | Permutations, traveling salesman brute force |

### The "halving vs subtracting" insight

- **Halving the input at each step** → log n levels deep (the tree is short)
- **Subtracting a constant at each step** → n levels deep (the tree is tall)

Combine this with the branching factor:

- Halving + 2 branches = 2^(log n) = n total work → manageable
- Subtracting 1 + 2 branches = 2ⁿ total work → catastrophic

### The summary you wrote yourself (with refinements)

**O(log n):** halving + only working on one piece per level (throwing away the other half) = log n levels × constant work = **log n total**.

**O(n log n):** halving + working on all pieces per level (which sum to n) = log n levels × n work = **n log n total**.

**O(2ⁿ):** doubling calls without halving the input = n levels × doubling at each level = **2ⁿ total**.

### Bonus: O(log² n) does exist

If you nest a binary search inside another binary search (e.g., searching a 2D sorted matrix where each row is sorted), you get log n × log n = **O(log² n)**. This comes from *nested halvings*, not from a halving tree where each level does log n work.

---

## Part 8: Space complexity follows the same trees

For space, ask: **how many things are alive at the same time?**

In recursion, only one path from root to leaf is active at any moment (depth-first execution). So space used = depth of the tree, not total nodes.

- **O(1) space**: a few variables, regardless of input. (Two pointers, in-place reversal.)
- **O(log n) space**: recursion stack log n deep. (Recursive binary search.)
- **O(n) space**: a structure proportional to input. (Hash map, recursion n deep, memo table.)
- **O(n²) space**: a 2D table. (DP problems with state (i, j) like edit distance.)

### Subtle point about recursion

A "two pointer" iterative solution might be O(1) space, but the recursive version of the same algorithm is O(n) space because of the call stack. Interviewers ask about this.

---

## Part 9: Recognizing complexity from problem constraints

Problem constraints telegraph the expected solution complexity:

| Constraint | Expected complexity |
|-----------|---------------------|
| n ≤ 10⁹ | O(log n) or O(1) |
| n ≤ 10⁶ | O(n) or O(n log n) |
| n ≤ 10⁵ | O(n log n) or O(n √n) |
| n ≤ 10⁴ | O(n²) might be okay |
| n ≤ 500 | O(n³) is acceptable |
| n ≤ 20 | O(2ⁿ) is the intended solution |
| n ≤ 10 | O(n!) is acceptable |

Reading the constraints before solving is one of the most important skills in competitive programming and interviews.

---

## Part 10: The bigger picture (from the conversation)

### How to actually build intuition (process)

- Always try the brute force first, then ask what's wasteful about it
- Sit with problems for 25-40 minutes before looking at solutions
- After solving, write three things: pattern, signal that should have told you, what brute force was missing
- Re-solve problems from scratch a week later — re-reading isn't learning

### The interleaving problem

Doing problems organized by pattern means your brain already knows the answer category. To fix:

- Mix problems across patterns randomly (shuffle a list)
- Use LeetCode's "random" feature instead of category lists
- Force yourself to articulate the problem in plain terms before reaching for a known technique
- Build the within-pattern reflex first, then switch to interleaved random practice

### Reading the problem (the blur problem)

Externalize the parsing — write down:

1. **Input**: what am I given, exactly? (Type, size, constraints, edge cases.)
2. **Output**: what am I returning?
3. **The rule**: what defines a valid answer? (Custom definitions in your own words.)
4. **A concrete example**: traced by hand, with your own small input if needed.

Three reads with different goals: gist, then input/output/constraints, then rule and examples.

If you can't construct a small example by hand, you don't understand the problem yet. Don't start coding.

---

## Quick reference cheat sheet

```
O(1)        Constant      | Direct access                          | 1 step always
O(log n)    Logarithmic   | Halving + discard one half             | ~20 steps for n=1M
O(n)        Linear        | Single pass                            | n steps
O(n log n)  Linearithmic  | Halving + work on both halves          | n × log n steps
O(n²)       Quadratic     | Nested loops                           | n × n steps
O(2ⁿ)       Exponential   | Recursion: 2 branches, shrink by 1     | doubles per element
O(n!)       Factorial     | All orderings / permutations           | n × (n-1) × (n-2) ...
```

The two questions to ask any algorithm: **how deep is the tree, and how much work per level?**
