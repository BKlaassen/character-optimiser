package main

import (
    "sort"
    "fmt"
)

type Character struct {
    stats [6]int
    race, class string
}

func (c Character) String() string {
    return fmt.Sprintf("Race: %s, Class: %s, \nStr: %d,\nDex: %d,\nCon: %d,\nInt: %d,\nWis: %d,\nCha: %d\n", c.race, c.class, c.stats[0], c.stats[1], c.stats[2], c.stats[3], c.stats[4], c.stats[5])
}

func main() {
    optimised := charopt([]int{8,12,15,14,13,10}, "Fighter", "High Elf")
    fmt.Println(optimised)
}

func charopt (statline []int, class string, race string) (optimised Character){
    var bonus []int
    switch race{
    case "Human":
        bonus = []int{1,1,1,1,1,1}
    case "Mountain Dwarf":
        bonus = []int{2,0,2,0,0,0}
    case "High Elf":
        bonus = []int{0,2,0,1,0,0}
    }
    optimised.race = race
    optimised.class = class
    sort.Ints(statline)
    
    optimised.stats[0] = statline[5]+bonus[0]
    optimised.stats[1] = statline[3]+bonus[1]
    optimised.stats[2] = statline[4]+bonus[2]
    optimised.stats[3] = statline[0]+bonus[3]
    optimised.stats[4] = statline[2]+bonus[4]
    optimised.stats[5] = statline[1]+bonus[5]
    return
}