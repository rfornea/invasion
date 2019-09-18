# Alien Invasion  

This program reads in a map of cities from a text file and accepts a number of aliens as a command line argument.  It then simulates an invasion of mad aliens, in which aliens can move from city to city (provided the two cities have a link between them).  If two aliens land in the same city, they destroy each other, along with the city and any links to that city from other cities.  An alien can get trapped in a city if all links to that city are removed.  

The program stops if all aliens have moved at least 10,000 times, all aliens are destroyed, all aliens are "trapped", or all cities are destroyed. 

The default file for the city map is `testMap.csv`.  The default number of aliens is `6`.  Both of these values can be changed when you run the program, if you use the appropriate command line flags.  Use the `-aliens` flag to specify the number of aliens as a command line argument and `-fpath` to specify the filepath of the map you want to use.  

For example:

`go run main.go -aliens=100 fpath=/some/other/file.txt`

## Assumptions:  
* You didn't give your cities any bizarre names that you wouldn't see for a real-world city.
* Your map is logically possible and correct, in terms of its layout.  For example, this map makes sense:  

> Foo east=Bar

> Bar west=Foo

   This map doesn't, because if Bar is east of Foo, then Foo should be *west* of Bar:  

>Foo east=Bar

>Bar north=Foo

* Aliens will not wait until all aliens have moved during a turn before attacking each other.  If two aliens appear in the same city (even if this is during the initial distribution of aliens to cities), they will *immediately* fight and destroy the city, and each other.  
* You are using an appropriate amount of aliens relative to the number of cities you have.  For example, if you use 1000 aliens and 7 cities, the simulation will likely end before even 1 turn is completed.  