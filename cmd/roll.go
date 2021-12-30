/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"math/rand"
	"os"
	"strconv"
	"time"
)

// rollCmd represents the roll command
var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("roll called")
		diceEntered, _ := rootCmd.PersistentFlags().GetString("dice")
		rollDice(diceEntered)
	},
}

func init() {
	rootCmd.AddCommand(rollCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rollCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rollCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().String("dice", "1d20", "select dice to roll")
}

func rollDice(dice string) {

	diceSplitString := strings.Split(dice, "d")
	//fmt.Println(diceSplitString)
	diceCount := diceSplitString[0]
	diceValue := diceSplitString[1]

	diceCountInt, err := strconv.Atoi(diceCount)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	diceValueInt, err := strconv.Atoi(diceValue)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	rollValue := 0

	i := 1
	for i <= diceCountInt {
		rand.Seed(time.Now().UnixNano())
		rollValue = rollValue + rand.Intn(diceValueInt)
		i = i + 1
	}

	//fmt.Println("Dice Count: ", diceCount)
	//fmt.Println("Dice Value: ", diceValue)
	fmt.Println(rollValue)
}
