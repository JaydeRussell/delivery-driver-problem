The problem input contains a list of loads. Each load is formatted as an id followed by pickup and dropoff locations in (x,y) floating point coordinates. An example input with four loads is:

```
loadNumber pickup dropoff
1 (-50.1,80.0) (90.1,12.2)
2 (-24.5,-19.2) (98.5,1.8)
3 (0.3,8.9) (40.9,55.0)
4 (5.3,-61.1) (77.8,-5.4)
```
Your program must write a solution to stdout. The solution should list, on separate lines, each driver’s ordered list of loads as a schedule. An example solution to the above problem could be:
```
[1]
[4,2]
[3]
```
This solution means one driver does load 1; another driver does load 4 followed by load 2; and a final driver does load 3.

All problems we provide will be solvable. That is, all loads are possible to complete within the duration of one 12-hour shift. Your program does not have to assess problem feasibility.
No problem will contain more than 200 loads.

