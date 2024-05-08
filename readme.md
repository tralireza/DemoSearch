DemoSearch
==========

BFS/DFS Search Demo - Terminal Vs iTerm2
----------------------------------------

<img alt="Ghosty Logo" src="https://github.com/tralireza/DemoSearch/blob/main/screenshots/GhostyLogo.png">


# Why?
This was supposed to be a benchmark test for **Terminal** and **iTerm2**, but <abbr title="Depth First Search / Breadth First Search">**DFS/BFS**</abbr> searches are fun! So it demos how a clever! _"ghost"_ finds its way in a maze. All and all I think I carry on using Terminal as I find iTerm2 a tiny bit slower for unicode (however it has lots of good features and even more to customize!)

# Build
``` bash
$ git clone https://github.com/tralireza/DemoSearch
$ cd DemoSearch
$ go build .
```

"Demo" binary should be in the current directory now. Run it! _(Side NOTE: But always remember to only run executables you built or you trust.)_

# Run (from Source)
``` bash
$ cd DemoSearch
$ go run . -h
Usage of /var/folders/8w/tjs_qsdj4x3b8y8ycn2cfb680000gn/T/go-build269568008/b001/exe/Demo:
  -BFS
    	Do a BFS search (otherwise DFS)
  -drawMaze
    	Only draw the maze & exit!
  -exits int
    	Number of doors to get out of maze (default 16)
  -i int
    	Ghosty's start row, between 2 & m-1 (default 5)
  -j int
    	Ghosty's start column, between 2 & n-1 (default 26)
  -m int
    	Number of rows in the maze (default 10)
  -n int
    	Number of columns in the maze (default 48)
  -walls int
    	Extra bricks inside of maze (default 128)
```

# Run (binary produced by build in Build above)
``` bash
$ ./Demo -h
...<^^^output as above^^^>
```

# Screenshots
Take a look at the "screenshots" folder for more or create a PR for your screenshots to go there... Would love to know how the performance is like on your setup.

Terminal:
(On my Mac, Terminal runs almost 5% to 7% cooler on CPU while the tiny "Ghosty" is finding its way out to a door!)
<img width="1678" alt="Terminal2" src="https://github.com/tralireza/DemoSearch/blob/main/screenshots/Terminal2.png">

(Waiting to be "Looking" into, already found doors/exits (so thumbsup to our "Ghosty") and the shortest path found (so far) stpes for the Beeline...

<img width="1678" alt="Terminal3" src="https://github.com/tralireza/DemoSearch/blob/main/screenshots/Terminal3.png">

---

iTerm2:
(It doesn't render all unicode the same width! Well Apple double width emoji to be exact but still either render all single/double width)
<img width="1678" alt="iTerm2" src="https://github.com/tralireza/DemoSearch/blob/main/screenshots/iTerm2.png">


# Bugs, Imprevements or More Screenshots
Yes please fire a PR away. Thank you!

# Terminal app on macOS > iTerm2 app on macOS
Well, I rest my case :-)
... And I think I now carry on using both! Gosh!

