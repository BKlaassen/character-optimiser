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
	if len(os.Args) > 1 && os.Args[1] == "--races" {
		fmt.Println(raceshelp)
		return
	}
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
		return
	}
	if n < 6 {
		fmt.Printf("Not enough arguments, expected 6, got %d. Exiting.\n", n)
		return
	}
	fmt.Println("Converted Input. Optimising....")
	optimised, err := charopt(statline[0:6], class, race)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Optimised. Here is your Character:")
	fmt.Println(optimised)
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
	case "shadar kai":
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
	case "yuan ti pureblood":
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

const raceshelp = "Supported races:\nHill Dwarf (+2 Con, +1 Wis, p. 18-20 PHB)\nMountain Dwarf (+2 Str, +2 Con, p. 18-20 PHB)\nDuergar (+1 Str, +2 Con, p. 81 MTF)\nHigh Elf (+2 Dex, +1 Int, p. 21-24 PHB)\nWood Elf (+2 Dex, +1 Wis, p. 21-24 PHB)\nDrow Elf (+2 Dex, +1 Cha, p. 21-24 PHB)\nEladrin (+2 Dex, +1 Cha, p. p. 61f MTF)\nSea Elf (+2 Dex, +1 Con, p. 62 MTF)\nShadar Kai (+2 Dex, +1 Con, p. 62f MTF)\nLightfoot Halfling (+2 Dex, +1 Cha, p. 26-28 PHB)\nStout Halfling (+2 Dex, +1 Con, p. 26-28 PHB)\nStandard Human (+1 to All, p. 29-31 PHB)\nDragonborn (+2 Str, +1 Cha, p. 32-34 PHB)\nForest Gnome (+1 Dex, +2 Int, p. 35-37 PHB)\nRock Gnome (+1 Con, +2 Int, p. 35-37 PHB)\nDeep Gnome (+1 Dex, +2 Int, p. 113f MTF)\nHalf Orc (+2 Str, +1 Con, p. 40f PHB)\nTiefling (+1 Int, +2 Cha, p.42f PHB)\nAsmodeus Tiefling (+1 Int, +2 Cha, p. 21 MTF)\nBaalzebul Tiefling (+1 Int, +2 Cha, p. 21 MTF)\nDispater Tiefling (+1 Dex, +2 Cha, p. 21 MTF)\nFierna Tiefling (+1 Wis, +2 Cha, p. 21f MTF)\nGlasya Tiefling (+1 Dex, +2 Cha, p. 22 MTF)\nLevistus Tiefling (+1 Con, +2 Cha, p. 22 MTF)\nMammon Tiefling (+1 Int, +2 Cha, p. 22 MTF)\nMephistopheles Tiefling (+1 Int, +2 Cha, p. 23 MTF)\nZariel Tiefling (+1 Str, +2 Cha, p. 23 MTF)\nProtector Aasimar (+1 Wis, +2 Cha, p. 104f Volos)\nScourge Aasimar (+1 Con, +2 Cha, p. 104f Volos)\nFallen Aasimar (+1 Str, +2 Cha, p. 104f Volos)\nFirbolg (+1 Str, +2 Wis, p. 106f Volos)\nGoliath (+2 Str, +1 Con, p. 108f Volos)\nKenku (+2 Dex, +1 Wis, p. 109-111 Volos)\nLizardfolk (+2 Con, +1 Wis, p. 111-113 Volos)\nTabaxi (+2 Dex, +1 Cha, p. 113-115 Volos)\nTriton (+2 Dex, +1 Int, p. 115-118 Volos)\nBugbear (+2 Str, +1 Dex, p. 119 Volos)\nGoblin (+2 Dex, +1 Con, p. 119 Volos)\nHobgoblin (+2 Con, +1 Int, p. 119 Volos)\nKobold (-2 Str, +2 Dex, p. 119 Volos)\nOrc (+2 Str, +1 Con, p. 120 Volos)\nYuan Ti Pureblood (+1 Int, +2 Cha, p. 120 Volos)\nGithyanki (+2 Str, +1 Int, p. 96 MTF)\nGithzerai (+1 Int, +2 Wis, p. 96 MTF)"
