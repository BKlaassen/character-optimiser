package main

import (
	"bufio"
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

func main() { //TODO: argpass for --help
	var dex bool
	if len(os.Args) > 1 && os.Args[1] == "--races" {
		fmt.Println(raceshelp)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Race: ")
	race, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An unknown Error has occurred while reading from Terminal. Exiting.")
		return
	}
	race = strings.TrimSpace(race)
	fmt.Print("Enter Class: ")
	class, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An unknown Error has occurred while reading from Terminal. Exiting.")
		return
	}
	class = strings.TrimSpace(class)
	if strings.ToLower(class) == "fighter" || strings.ToLower(class) == "cleric heavy" || strings.ToLower(class) == "ranger" {
		fmt.Println("Do you want to use Dexterity or Strength as primary combat attribute? (Dex/Str, Dexterity/Strength)")
		answ, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An unknown Error has occurred while reading from Terminal. Exiting.")
			return
		}
		answ = strings.TrimSpace(answ)
		if answ == "dex" || answ == "dexterity" {
			dex = true
		} else if answ == "str" || answ == "strength" {
			dex = false
		} else {
			fmt.Println(answ, "is not a valid answer. Exiting")
		}
	} else if strings.ToLower(class) == "rogue" {
		fmt.Println("Do you want to use Charisma or Intelligence as primary utility attribute? (Cha/Int, Charisma/Intelligence)")
		answ, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An unknown Error has occurred while reading from Terminal. Exiting.")
			return
		}
		answ = strings.TrimSpace(answ)
		if answ == "cha" || answ == "charisma" {
			dex = true
		} else if answ == "int" || answ == "intelligence" {
			dex = false
		} else {
			fmt.Println(answ, "is not a valid answer. Exiting")
		}
	}
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
	optimised, err := Charopt(statline[0:6], class, race, dex)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Optimised. Here is your Character:")
	fmt.Println(optimised)
	return
}

func Charopt(statline []int, class string, race string, dex bool) (optimised Character, err error) {
	optimised.race = race
	optimised.class = class
	sort.Ints(statline)
	optimised.stats = [6]int{0, 0, 0, 0, 0, 0}
	var classerr bool
	var dist [6]int
	class = strings.ToLower(class)
	race = strings.ToLower(race)
	if dex {
		dist, classerr = classdistdex[class]
	} else {
		dist, classerr = classdist[class]
	}
	bonus, raceerr := racebonus[race]
	if !raceerr && !classerr {
		fmt.Println("Neither", class, "nor", race, "are valid options, please check against --classes and --races")
		return
	} else if !classerr {
		fmt.Println(class, "is not a valid class option. Please check --classes")
	} else if !raceerr {
		fmt.Println(race, "is not a valid race option. Please check --races")
	}
	if race == "variant human" || (race == "half elf" && dist[5] != 5) {
		optimised.setstats(statline[0:6], dist[0:6], bonus[0:6], 5, 4)
	} else if race == "half elf" {
		optimised.setstats(statline[0:6], dist[0:6], bonus[0:6], 4, 3)
	} else {
		optimised.setstats(statline[0:6], dist[0:6], bonus[0:6], -1, -1)
	}
	return optimised, err
}

func (opt *Character) setstats(statline []int, dist []int, bonus []int, first int, second int) {
	for i := 0; i < 6; {
		if i == first || i == second {
			bonus[i] = 1
		}
		opt.stats[i] = statline[dist[i]] + bonus[i] //TODO: Add Even/Odd Checker
		opt.calc[i] = fmt.Sprintf("%d + %d", statline[dist[i]], bonus[i])
	}
}

var racebonus = map[string][6]int{
	"hill dwarf":              [6]int{0, 0, 2, 0, 1, 0},
	"mountain dwarf":          [6]int{2, 0, 2, 0, 0, 0},
	"duergar":                 [6]int{1, 0, 2, 0, 0, 0},
	"high elf":                [6]int{0, 2, 0, 1, 0, 0},
	"wood elf":                [6]int{0, 2, 0, 0, 1, 0},
	"drow elf":                [6]int{0, 2, 0, 0, 0, 1},
	"eladrin":                 [6]int{0, 2, 0, 0, 0, 1},
	"sea elf":                 [6]int{0, 2, 1, 0, 0, 0},
	"shadar kai":              [6]int{0, 2, 1, 0, 0, 0},
	"lightfoot halfling":      [6]int{0, 2, 0, 0, 0, 1},
	"stout halfling":          [6]int{0, 2, 1, 0, 0, 0},
	"standard human":          [6]int{1, 1, 1, 1, 1, 1},
	"variant human":           [6]int{0, 0, 0, 0, 0, 0}, //Hell is other humans.
	"dragonborn":              [6]int{2, 0, 0, 0, 0, 1},
	"forest gnome":            [6]int{0, 1, 0, 2, 0, 0},
	"rock gnome":              [6]int{0, 0, 1, 2, 0, 0},
	"deep gnome":              [6]int{0, 1, 0, 2, 0, 0},
	"half elf":                [6]int{0, 0, 0, 0, 0, 2},
	"half orc":                [6]int{2, 0, 1, 0, 0, 0},
	"tiefling":                [6]int{0, 0, 0, 1, 0, 2},
	"asmodeus tiefling":       [6]int{0, 0, 0, 1, 0, 2},
	"baalzebul tiefling":      [6]int{0, 0, 0, 1, 0, 2},
	"dispater tiefling":       [6]int{0, 1, 0, 0, 0, 2},
	"fierna tiefling":         [6]int{0, 0, 0, 0, 1, 2},
	"glasya tiefling":         [6]int{0, 1, 0, 0, 0, 2},
	"levistus tiefling":       [6]int{0, 0, 1, 0, 0, 2},
	"mammon tiefling":         [6]int{0, 0, 0, 1, 0, 2},
	"mephistopheles tiefling": [6]int{0, 0, 0, 1, 0, 2},
	"zariel tiefling":         [6]int{1, 0, 0, 0, 0, 2},
	"protector aasimar":       [6]int{0, 0, 0, 0, 1, 2},
	"scourge aasimar":         [6]int{0, 0, 1, 0, 0, 2},
	"fallen aasimar":          [6]int{1, 0, 0, 0, 0, 2},
	"firbolg":                 [6]int{1, 0, 0, 0, 2, 0},
	"goliath":                 [6]int{2, 0, 1, 0, 0, 0},
	"kenku":                   [6]int{0, 2, 0, 0, 1, 0},
	"lizardfolk":              [6]int{0, 0, 2, 0, 1, 0},
	"tabaxi":                  [6]int{0, 2, 0, 0, 0, 1},
	"triton":                  [6]int{1, 0, 1, 0, 0, 1},
	"bugbear":                 [6]int{2, 1, 0, 0, 0, 0},
	"goblin":                  [6]int{0, 2, 1, 0, 0, 0},
	"hobgoblin":               [6]int{0, 0, 2, 1, 0, 0},
	"kobold":                  [6]int{-2, 2, 0, 0, 0, 0},
	"orc":                     [6]int{2, 0, 1, 0, 0, 0},
	"yuan ti pureblood":       [6]int{0, 0, 0, 1, 0, 2},
	"githyanki":               [6]int{2, 0, 0, 1, 0, 0},
	"githzerai":               [6]int{0, 0, 0, 1, 2, 0},
}

var classdist = map[string][6]int{
	"barbarian":     [6]int{5, 3, 4, 0, 2, 1},
	"bard":          [6]int{0, 2, 1, 3, 4, 5},
	"cleric medium": [6]int{0, 4, 3, 2, 5, 1},
	"cleric heavy":  [6]int{4, 1, 3, 2, 5, 0},
	"druid":         [6]int{0, 4, 3, 2, 5, 1},
	"fighter":       [6]int{5, 2, 4, 0, 3, 1},
	"monk":          [6]int{2, 5, 3, 1, 4, 0},
	"paladin":       [6]int{5, 2, 4, 0, 1, 3},
	"ranger":        [6]int{5, 3, 4, 1, 2, 0},
	"rogue":         [6]int{0, 5, 3, 4, 2, 1},
	"sorcerer":      [6]int{0, 3, 4, 1, 2, 5},
	"warlock":       [6]int{0, 3, 4, 1, 2, 5},
	"wizard":        [6]int{0, 3, 4, 5, 2, 1},
}

var classdistdex = map[string][6]int{
	"cleric heavy": [6]int{0, 4, 3, 2, 5, 1},
	"fighter":      [6]int{2, 5, 4, 0, 3, 1},
	"ranger":       [6]int{1, 5, 4, 0, 2, 3},
	"rogue":        [6]int{0, 5, 3, 1, 2, 4},
}

const raceshelp = "Supported races:\nHill Dwarf (+2 Con, +1 Wis, p. 18-20 PHB)\nMountain Dwarf (+2 Str, +2 Con, p. 18-20 PHB)\nDuergar (+1 Str, +2 Con, p. 81 MTF)\nHigh Elf (+2 Dex, +1 Int, p. 21-24 PHB)\nWood Elf (+2 Dex, +1 Wis, p. 21-24 PHB)\nDrow Elf (+2 Dex, +1 Cha, p. 21-24 PHB)\nEladrin (+2 Dex, +1 Cha, p. p. 61f MTF)\nSea Elf (+2 Dex, +1 Con, p. 62 MTF)\nShadar Kai (+2 Dex, +1 Con, p. 62f MTF)\nLightfoot Halfling (+2 Dex, +1 Cha, p. 26-28 PHB)\nStout Halfling (+2 Dex, +1 Con, p. 26-28 PHB)\nStandard Human (+1 to All, p. 29-31 PHB)\nVariant Human (+1 to any 2 Attributes, p. 29-31 PHB)\nDragonborn (+2 Str, +1 Cha, p. 32-34 PHB)\nForest Gnome (+1 Dex, +2 Int, p. 35-37 PHB)\nRock Gnome (+1 Con, +2 Int, p. 35-37 PHB)\nDeep Gnome (+1 Dex, +2 Int, p. 113f MTF)\nHalf Elf (+2 Cha, +1 to any 2 Attributes except Cha, p.38f PHB)\nHalf Orc (+2 Str, +1 Con, p. 40f PHB)\nTiefling (+1 Int, +2 Cha, p.42f PHB)\nAsmodeus Tiefling (+1 Int, +2 Cha, p. 21 MTF)\nBaalzebul Tiefling (+1 Int, +2 Cha, p. 21 MTF)\nDispater Tiefling (+1 Dex, +2 Cha, p. 21 MTF)\nFierna Tiefling (+1 Wis, +2 Cha, p. 21f MTF)\nGlasya Tiefling (+1 Dex, +2 Cha, p. 22 MTF)\nLevistus Tiefling (+1 Con, +2 Cha, p. 22 MTF)\nMammon Tiefling (+1 Int, +2 Cha, p. 22 MTF)\nMephistopheles Tiefling (+1 Int, +2 Cha, p. 23 MTF)\nZariel Tiefling (+1 Str, +2 Cha, p. 23 MTF)\nProtector Aasimar (+1 Wis, +2 Cha, p. 104f Volos)\nScourge Aasimar (+1 Con, +2 Cha, p. 104f Volos)\nFallen Aasimar (+1 Str, +2 Cha, p. 104f Volos)\nFirbolg (+1 Str, +2 Wis, p. 106f Volos)\nGoliath (+2 Str, +1 Con, p. 108f Volos)\nKenku (+2 Dex, +1 Wis, p. 109-111 Volos)\nLizardfolk (+2 Con, +1 Wis, p. 111-113 Volos)\nTabaxi (+2 Dex, +1 Cha, p. 113-115 Volos)\nTriton (+2 Dex, +1 Int, p. 115-118 Volos)\nBugbear (+2 Str, +1 Dex, p. 119 Volos)\nGoblin (+2 Dex, +1 Con, p. 119 Volos)\nHobgoblin (+2 Con, +1 Int, p. 119 Volos)\nKobold (-2 Str, +2 Dex, p. 119 Volos)\nOrc (+2 Str, +1 Con, p. 120 Volos)\nYuan Ti Pureblood (+1 Int, +2 Cha, p. 120 Volos)\nGithyanki (+2 Str, +1 Int, p. 96 MTF)\nGithzerai (+1 Int, +2 Wis, p. 96 MTF)"
const classeshelp = "Supported classes:\nBarbarian (p. 46-50 PHB)\nBard (p. 51-55 PHB)\nCleric (Versions: Cleric Medium, Cleric Heavy by Armor Proficiency. p. 56-63 PHB)\nDruid (p. 64-69)\nFighter (CURRENT CONFIGURATION NOT VIABLE FOR ELDRITCH KNIGHT p. 70-75)\nMonk (p. 76-81 PHB)\nPaladin (p. 82-88 PHB)\nRanger (p. 89-93 PHB)\nRogue (If you want to play Arcane Trickster make sure to choose Int instead of Char, p. 94-98 PHB)\nSorcerer (p. 99-104)\nWarlock (p. 105-111 PHB)\nWizard (p. 112-119 PHB)"
