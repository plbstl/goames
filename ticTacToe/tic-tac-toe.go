package ticTacToe

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Logo is the first thing a user sees when the
// application is started.
var Logo = `
             TAC  
         TIC     TOE 
             TAC`

// HelpMsg shows available hotkeys
var HelpMsg = "\n show board (s), end (e), help (h)"

//  Player represents a single player.
type Player struct {
	name   string
	weapon string
}

// Board represents the board and its properties.
type Board struct {
	CurrentPlayer Player
	tiles         map[string]string
	p1            Player
	p2            Player
	moves         int
}

// NewBoard returns a new board.
func NewBoard(p1 Player, p2 Player) Board {
	tiles := make(map[string]string, 9)
	emptyTile := "___"

	for _, t := range []string{"a1", "a2", "a3", "b1", "b2", "b3", "c1", "c2", "c3"} {
		tiles[t] = emptyTile
	}

	return Board{p1, tiles, p1, p2, 0}
}

// NewPlayers returns two new players.
func NewPlayers() (Player, Player) {
	player := make([]Player, 2)
	weapon := []string{"X", "O"}

	for i := 0; i < 2; i++ {
		var name string
		fmt.Print("\n--> enter player name: ")

		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			if n, b := validateName(i, s.Text()); b {
				name = n
				break
			} else {
				fmt.Println("\n  Player name is not valid. Only alphanumeric characters and underscores are allowed")
				fmt.Print("\n--> enter player name: ")
			}
		}

		player[i] = Player{name, weapon[i]}
		fmt.Printf(" player `%s` created! [%s] \n", name, weapon[i])
	}

	return player[0], player[1]
}

func validateName(i int, name string) (string, bool) {
	name = normalizeString(name)

	if len(name) > 12 {
		name = name[:12]
	}

	if b, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", name); !b {
		return "", false
	}

	if name == "" {
		name = "player" + fmt.Sprint(i)
	}

	return name, true
}

func normalizeString(s string) string {
	s = strings.TrimSpace(s)
	return strings.ToLower(s)
}

// RunHotkey executes a command based on the given hotkey.
func RunHotkey(s string, b *Board) {
	switch s {
	case "s":
		b.Show()
	case "e":
		fmt.Println("\n Game Ended!")
		os.Exit(0)
	default:
		fmt.Println(HelpMsg)
	}
}

// Show prints out the current state of the board.
func (b *Board) Show() {
	fmt.Printf(` 
  	  | 1 | 2 | 3
  	--------------
  	a |%s|%s|%s
  	b |%s|%s|%s
  	c |%s|%s|%s
	`,
		b.tiles["a1"],
		b.tiles["a2"],
		b.tiles["a3"],
		b.tiles["b1"],
		b.tiles["b2"],
		b.tiles["b3"],
		b.tiles["c1"],
		b.tiles["c2"],
		b.tiles["c3"],
	)
}

// PlayTile plays a tile in a given postion. If sucessful,
// it displays the game board and change the current player.
func (b *Board) PlayTile(t string) {
	t = normalizeString(t)

	if b.tiles[t] == "___" {
		b.tiles[t] = fmt.Sprintf(" %s ", b.CurrentPlayer.weapon)
		fmt.Printf("\n  `%s` played \"%s\" in position %s \n",
			b.CurrentPlayer.name,
			b.CurrentPlayer.weapon,
			strings.ToUpper(t),
		)

		b.moves++

		if gw, gd := b.WinOrDraw(); gw || gd {
			b.Show()
			if gw {
				fmt.Printf("\n MATCH WON! by %s \n\n", b.CurrentPlayer)
				os.Exit(0)
			} else {
				fmt.Print("\n MATCH DRAWN! \n\n")
				os.Exit(0)
			}
		} else {
			b.Show()
			b.changePlayer()
		}
	} else {
		fmt.Printf("\n  position %s is not empty! \n", strings.ToUpper(t))
	}
}

// WinOrDraw extensively checks if game is over.
func (b *Board) WinOrDraw() (gameWon bool, gameDrawn bool) {
	// no need to check when win/draw is impossible
	if b.moves < 5 {
		return false, false
	}

	if b.moves >= 9 {
		// match drawn
		return false, true
	}

	cpw := fmt.Sprintf(" %s ", b.CurrentPlayer.weapon)

	// horizontal
	if cpw == b.tiles["a1"] && b.tiles["a1"] == b.tiles["a2"] && b.tiles["a2"] == b.tiles["a3"] {
		b.tiles["a1"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["a2"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["a3"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		fmt.Println("h1")
		return true, false
	}
	if cpw == b.tiles["b1"] && b.tiles["b1"] == b.tiles["b2"] && b.tiles["b2"] == b.tiles["b3"] {
		b.tiles["b1"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["b2"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["b3"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		fmt.Println("h2")
		return true, false
	}
	if cpw == b.tiles["c1"] && b.tiles["c1"] == b.tiles["c2"] && b.tiles["c2"] == b.tiles["c3"] {
		b.tiles["c1"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["c2"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["c3"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		fmt.Println("h3")
		return true, false
	}

	// vertical
	if cpw == b.tiles["a1"] && b.tiles["a1"] == b.tiles["b1"] && b.tiles["b1"] == b.tiles["c1"] {
		b.tiles["a1"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["b1"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["c1"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		fmt.Println("v1")
		return true, false
	}
	if cpw == b.tiles["a2"] && b.tiles["a2"] == b.tiles["b2"] && b.tiles["b2"] == b.tiles["c2"] {
		b.tiles["a2"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["b2"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["c2"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		fmt.Println("v2")
		return true, false
	}
	if cpw == b.tiles["a3"] && b.tiles["a3"] == b.tiles["b3"] && b.tiles["b3"] == b.tiles["c3"] {
		b.tiles["a3"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["b3"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["c3"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		fmt.Println("v3")
		return true, false
	}

	// diagonal
	if cpw == b.tiles["a1"] && b.tiles["a1"] == b.tiles["b2"] && b.tiles["b2"] == b.tiles["c3"] {
		b.tiles["a1"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["b2"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["c3"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		fmt.Println("d1")
		return true, false
	}

	if cpw == b.tiles["a3"] && b.tiles["a3"] == b.tiles["b2"] && b.tiles["b2"] == b.tiles["c1"] {
		b.tiles["a3"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["b2"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		b.tiles["c1"] = fmt.Sprintf("`%s`", b.CurrentPlayer.weapon)
		fmt.Println("d2")
		return true, false
	}

	return false, false
}

func (b *Board) changePlayer() {
	if b.CurrentPlayer == b.p1 {
		b.CurrentPlayer = b.p2
	} else {
		b.CurrentPlayer = b.p1
	}
}
