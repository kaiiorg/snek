# Design
Notes and such about how `snek` works under the hood.

## World
World contains the state of the `snek` world. It contains the boundaries, obstacles, snacks, the player's snek, etc

### Obstacles
Things that will kill your snek like walls and landmines.

### Snacks
Things that will grow your snek and add points to your score.

### The Player's Snek
The curren player's character or a remote player's character. Running into a wall or itself is bad; this kills the snek. 
Desires snacks to grow longer.

