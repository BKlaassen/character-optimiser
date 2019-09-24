package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Character struct {
	stats       [6]int
	calc        [6]string
	race, class string
}

func (c Character) String() string { //weird behaviour with "Ra" being " ,"
	return fmt.Sprintf("Race: %s, Class: %s, \nStr: %d (%s),\nDex: %d (%s),\nCon: %d (%s),\nInt: %d (%s),\nWis: %d (%s),\nCha: %d (%s)", c.race, c.class, c.stats[0], c.calc[0], c.stats[1], c.calc[1], c.stats[2], c.calc[2], c.stats[3], c.calc[3], c.stats[4], c.calc[4], c.stats[5], c.calc[5])
}

func main() { //TODO: argpass for --help --races --classes
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Race (as 'High Elf'): ")
	race, _ := reader.ReadString('\n')
	race = strings.TrimSpace(race)
	fmt.Print("Enter Class (as 'Fighter' or 'Tempest Cleric'): ")
	class, _ := reader.ReadString('\n')
	class = strings.TrimRight(class, "\n")
	fmt.Print("Enter Statline (format: 'a b c d e f g'): ")
	var statline [6]int
	n, err := fmt.Fscanf(os.Stdin, "%d %d %d %d %d %d", &statline[0], &statline[1], &statline[2], &statline[3], &statline[4], &statline[5])
	if err != nil {
		fmt.Println("An unknown Error has occurred while reading from Terminal. Exiting.")
		fmt.Fscan(os.Stdin)
		return
	}
	if n < 6 {
		fmt.Printf("Not enough arguments, expected 6, got %d. Exiting.\n", n)
		fmt.Fscan(os.Stdin)
		return
	}
	fmt.Println("Converted Input. Optimising....")
	optimised, err := charopt(statline[0:6], class, race)
	if err != nil {
		fmt.Println(err)
		fmt.Fscan(os.Stdin)
		return
	}
	fmt.Println("Optimised. Here is your Character:")
	fmt.Println(optimised)
	fmt.Fscan(os.Stdin)
	return
}

func charopt(statline []int, class string, race string) (optimised Character, err error) {
	bonus, err := bonus(strings.ToLower(race))
	if err != nil {
		return optimised, err
	}
	optimised.race = race
	optimised.class = class
	sort.Ints(statline)
	optimised.stats = [6]int{0, 0, 0, 0, 0, 0}
	optimised.stats[0] = statline[5] + bonus[0] //TODO: Add Class switch, Even/Odd Checker
	optimised.calc[0] = fmt.Sprintf("%d + %d", statline[5], bonus[0])
	optimised.stats[1] = statline[3] + bonus[1]
	optimised.calc[1] = fmt.Sprintf("%d + %d", statline[3], bonus[1])
	optimised.stats[2] = statline[4] + bonus[2]
	optimised.calc[2] = fmt.Sprintf("%d + %d", statline[4], bonus[2])
	optimised.stats[3] = statline[0] + bonus[3]
	optimised.calc[3] = fmt.Sprintf("%d + %d", statline[0], bonus[3])
	optimised.stats[4] = statline[2] + bonus[4]
	optimised.calc[4] = fmt.Sprintf("%d + %d", statline[2], bonus[4])
	optimised.stats[5] = statline[1] + bonus[5]
	optimised.calc[5] = fmt.Sprintf("%d + %d", statline[1], bonus[5])
	return optimised, err
}

func bonus(race string) (bonus [6]int, err error) {
	switch race {
	case "hill dwarf":
		bonus = [6]int{0, 0, 2, 0, 1, 0}
	case "mountain dwarf":
		bonus = [6]int{2, 0, 2, 0, 0, 0}
	case "duergar":
		bonus = [6]int{1, 0, 2, 0, 0, 0}
	case "high elf":
		bonus = [6]int{0, 2, 0, 1, 0, 0}
	case "wood elf":
		bonus = [6]int{0, 2, 0, 0, 1, 0}
	case "drow elf":
		bonus = [6]int{0, 2, 0, 0, 0, 1}
	case "eladrin":
		bonus = [6]int{0, 2, 0, 0, 0, 1}
	case "sea elf":
		bonus = [6]int{0, 2, 1, 0, 0, 0}
	case "shadar-kai":
		bonus = [6]int{0, 2, 1, 0, 0, 0}
	case "lightfoot halfling":
		bonus = [6]int{0, 2, 0, 0, 0, 1}
	case "stout halfling":
		bonus = [6]int{0, 2, 1, 0, 0, 0}
	case "standard human":
		bonus = [6]int{1, 1, 1, 1, 1, 1}
	case "variant human": //Hell is other humans.
		bonus = [6]int{0, 0, 0, 0, 0, 0}
	case "dragonborn":
		bonus = [6]int{2, 0, 0, 0, 0, 1}
	case "forest gnome":
		bonus = [6]int{0, 1, 0, 2, 0, 0}
	case "rock gnome":
		bonus = [6]int{0, 0, 1, 2, 0, 0}
	case "deep gnome":
		bonus = [6]int{0, 1, 0, 2, 0, 0}
	case "half elf":
		bonus = [6]int{0, 0, 0, 0, 0, 2}
	case "half orc":
		bonus = [6]int{2, 0, 1, 0, 0, 0}
	case "tiefling":
		bonus = [6]int{0, 0, 0, 1, 0, 2}
	case "asmodeus tiefling":
		bonus = [6]int{0, 0, 0, 1, 0, 2}
	case "baalzebul tiefling":
		bonus = [6]int{0, 0, 0, 1, 0, 2}
	case "dispater tiefling":
		bonus = [6]int{0, 1, 0, 0, 0, 2}
	case "fierna tiefling":
		bonus = [6]int{0, 0, 0, 0, 1, 2}
	case "glasya tiefling":
		bonus = [6]int{0, 1, 0, 0, 0, 2}
	case "levistus tiefling":
		bonus = [6]int{0, 0, 1, 0, 0, 2}
	case "mammon tiefling":
		bonus = [6]int{0, 0, 0, 1, 0, 2}
	case "mephistopheles tiefling":
		bonus = [6]int{0, 0, 0, 1, 0, 2}
	case "zariel tiefling":
		bonus = [6]int{1, 0, 0, 0, 0, 2}
	case "protector aasimar":
		bonus = [6]int{0, 0, 0, 0, 1, 2}
	case "scourge aasimar":
		bonus = [6]int{0, 0, 1, 0, 0, 2}
	case "fallen aasimar":
		bonus = [6]int{1, 0, 0, 0, 0, 2}
	case "firbolg":
		bonus = [6]int{1, 0, 0, 0, 2, 0}
	case "goliath":
		bonus = [6]int{2, 0, 1, 0, 0, 0}
	case "kenku":
		bonus = [6]int{0, 2, 0, 0, 1, 0}
	case "lizardfolk":
		bonus = [6]int{0, 0, 2, 0, 1, 0}
	case "tabaxi":
		bonus = [6]int{0, 2, 0, 0, 0, 1}
	case "triton":
		bonus = [6]int{1, 0, 1, 0, 0, 1}
	case "bugbear":
		bonus = [6]int{2, 1, 0, 0, 0, 0}
	case "goblin":
		bonus = [6]int{0, 2, 1, 0, 0, 0}
	case "hobgoblin":
		bonus = [6]int{0, 0, 2, 1, 0, 0}
	case "kobold":
		bonus = [6]int{-2, 2, 0, 0, 0, 0}
	case "orc":
		bonus = [6]int{2, 0, 1, 0, 0, 0}
	case "yuan-ti pureblood":
		bonus = [6]int{0, 0, 0, 1, 0, 2}
	case "githyanki":
		bonus = [6]int{2, 0, 0, 1, 0, 0}
	case "githzerai":
		bonus = [6]int{0, 0, 0, 1, 2, 0}
	default:
		bonus = [6]int{0, 0, 0, 0, 0, 0}
		err = errors.New(strings.Join([]string{race, " is not a race. Check --races"}, ""))
	}
	return bonus, err
}
