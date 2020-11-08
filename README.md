# Daily Coding Problem: Problem #703 [Hard]

This problem was asked by Two Sigma.

A knight is placed on a given square on an 8 x 8 chessboard.
It is then moved randomly several times,
where each move is a standard knight move.
If the knight jumps off the board at any point,
however, it is not allowed to jump back on.

After k moves,
what is the probability that the knight remains on the board?

## Analysis

### Data structure

8x8 board - 2D array-of-arrays?

### Positions and moves

Knight positioned on `<x,y>`

Possible moves:
```
<x-1, y+2>
<x+1, y+2>
<x-1, y-2>
<x+1, y-2>
<x-1, y-2>
<x+1, y-2>
<x+2, y+1>
<x+2, y-1>
<x-2, y+1>
<x-2, y-1>
```
