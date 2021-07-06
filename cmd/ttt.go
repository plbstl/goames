/*
Copyright © 2021 Paul Ebose <paulebose@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/paulebose/goames/ticTacToe"
	"github.com/spf13/cobra"
)

// tttCmd represents the ttt command.
var tttCmd = &cobra.Command{
	Use:   "ttt",
	Short: "tic-tac-toe game",
	Long: `Tic-tac-toe is a game for two players, who take turns marking the
spaces in a 3×3 grid. The player who succeeds in placing three of
their marks in a diagonal, horizontal, or vertical row is the winner.`,
	Run: tttRun,
}

func init() {
	rootCmd.AddCommand(tttCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tttCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tttCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func tttRun(cmd *cobra.Command, args []string) {
	fmt.Println(ticTacToe.Logo, "\n", ticTacToe.HelpMsg)

	p1, p2 := ticTacToe.NewPlayers()
	b := ticTacToe.NewBoard(p1, p2)
	b.Show()

	fmt.Printf("\n--> select a tile to play | %s: ", b.CurrentPlayer)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		switch t := s.Text(); t {
		case "":
			fmt.Println("\n please select a tile to play!")

		case "s", "e", "h":
			ticTacToe.RunHotkey(t, &b)

		case "a1", "a2", "a3", "b1", "b2", "b3", "c1", "c2", "c3":
			b.PlayTile(t)

		default:
			fmt.Println("\n invalid selection!")
		}

		fmt.Printf("\n--> select a tile to play | %s: ", b.CurrentPlayer)
	}
}
