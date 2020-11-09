# Daily Coding Problem: Problem #703 [Hard]

This problem was asked by Two Sigma.

A knight is placed on a given square on an 8 x 8 chessboard.
It is then moved randomly several times,
where each move is a standard knight move.
If the knight jumps off the board at any point,
however, it is not allowed to jump back on.

After k moves,
what is the probability that the knight remains on the board?

## Algorithm analysis

There's at least 2 ways to get an answer:

* [Probabilistic version](a3.go) for as close as you want to spend time
* [Closed form version](b1.go) that gives an exact solution

The probabilistic version does a specified number of randomly-chosen
knight's moves, starting from each of the 64 squares on the board.
It counts the number of such paths that remain on the board.
This allows me to approximate the probability requested in the problem statement.

An example run:

```sh
% ./a3 3 100000
Looking 3 moves deep, 100000 times
0.1246 0.1979 0.2623 0.2824 0.2855 0.2634 0.1999 0.1239 
0.1980 0.2904 0.3938 0.4264 0.4279 0.3952 0.2887 0.2001 
0.2641 0.3933 0.5346 0.5767 0.5751 0.5352 0.3956 0.2629 
0.2846 0.4278 0.5718 0.6223 0.6222 0.5756 0.4277 0.2858 
0.2867 0.4271 0.5740 0.6215 0.6228 0.5743 0.4290 0.2835 
0.2624 0.3954 0.5375 0.5749 0.5732 0.5354 0.3929 0.2656 
0.1986 0.2888 0.3944 0.4256 0.4281 0.3936 0.2877 0.2008 
0.1238 0.1995 0.2623 0.2878 0.2857 0.2621 0.1998 0.1248 

```

It's easily possible to calculate the probability of a knight's
path for 1 or 2 moves.
I verified a few starting square's probabilities that way.

### Closed form solution

While working out a few two- or three-move probabilities,
I realized that what I was doing was counting the number
of paths that stayed on the board the entire time,
then dividing by 8<sup>k</sup>, where `k` is the number
of moves.

This leads to the closed form solution:
finding all the k-move-long paths from each square
on the chessboard.
I did shorten the recursion by calculating how many
stay-on-the-board moves there are for "the next move".
When the remaining path is only one move long,
the recursion consults a table.
This is a vestige of working out the probability by hand
for a few starting squares.

This diagram shows the number of 
a knight's moves that stay on the board.

```
8 |2  3  4  4  4  4  3  2 
7 |3  4  6  6  6  6  4  3 
6 |4  6  8  8  8  8  6  4 
5 |4  6  8  8  8  8  6  4 
4 |4  6  8  8  8  8  6  4 
3 |4  6  8  8  8  8  6  4 
2 |3  4  6  6  6  6  4  3 
1 |2  3  4  4  4  4  3  2 
  ----------------------
   A  B  C  D  E  F  G  H
```

For a 3-move path, the probablities look like this:

```sh
% ./b1 3
0.1250 0.1992 0.2637 0.2852 0.2852 0.2637 0.1992 0.1250 
0.1992 0.2891 0.3945 0.4277 0.4277 0.3945 0.2891 0.1992 
0.2637 0.3945 0.5352 0.5742 0.5742 0.5352 0.3945 0.2637 
0.2852 0.4277 0.5742 0.6211 0.6211 0.5742 0.4277 0.2852 
0.2852 0.4277 0.5742 0.6211 0.6211 0.5742 0.4277 0.2852 
0.2637 0.3945 0.5352 0.5742 0.5742 0.5352 0.3945 0.2637 
0.1992 0.2891 0.3945 0.4277 0.4277 0.3945 0.2891 0.1992 
0.1250 0.1992 0.2637 0.2852 0.2852 0.2637 0.1992 0.1250 
```

It's left-and-right, top-and-bottom,
and across diagonals symmetrical.
I could optimize further by only considering one,
16-square quadrant, and reflecting it to get the entire board.
The algorithm isn't computationally expensive,
so I chose to do it for all 64 squares to get some
confirmation that the code is correct via symmetry.

The probabilistic and the brute-force, closed form solution
come out pretty close,
allowing me a warm feeling about the correctness of the answers.

## Intervew Analysis

I only got to my closed form solution because I thought there might be
a fancy matrix-multiplication method of getting to the exact probabilities.
This shows up as the algorithm using a table to look up a count of moves
for paths of length 1.
It would be easier to have the algorithm try every k-move path,
calculating probablity based on how many paths end on the board,
no table lookup required.

Either solution requires about the same amount of coding.
The interviewer could get a decent feel for the candidate's knowledge
of the subject programming language via this problem.
Both solutions seem a bit much for writing entirely on a whiteboard, though.

Given that there are 2 divergent solutions,
the interviewer should have a different expectation of candidates for each solution.
I found it easier to conceive of the probabilistic solution,
but this won't be true for every one.

If I were a candidate, I would try to score points by discussing what
the merits (easy to understand and code)
and demerits (not exact, only close enough)
of the probabilistic version.
