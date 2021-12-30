/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"math/rand"

	"github.com/goombaio/namegenerator"
)

// createNpcCmd represents the createNpc command
var createNpcCmd = &cobra.Command{
	Use:   "createNpc",
	Short: "creates randomly generated NPC",
	Long:  `creates randomly generated NPC with name and stats`,
	Run:   createNpcRun,
}

func createNpcRun(cmd *cobra.Command, args []string) {
	getName()
	getNpcStats()
}

func init() {
	rootCmd.AddCommand(createNpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createNpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//createNpcCmd.Flags().BoolP("toggle", "t", true, "Help message for toggle")
	rootCmd.PersistentFlags().String("name", "", "Set name of character")
}

func getNpcStats() {
	str := rand.Intn(24-4) + 4
	dex := rand.Intn(24-4) + 4
	con := rand.Intn(24-4) + 4
	inte := rand.Intn(24-4) + 4
	wis := rand.Intn(24-4) + 4
	cha := rand.Intn(24-4) + 4

	fmt.Println("Strength: ", str)
	fmt.Println("Dexterity: ", dex)
	fmt.Println("Constitution: ", con)
	fmt.Println("Intelligence: ", inte)
	fmt.Println("Wisdom: ", wis)
	fmt.Println("Charisma: ", cha)
}

func getRandomName() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()
	fmt.Println("Name: ", name)
}

func getName() {
	name := ""
	givenName, _ := rootCmd.PersistentFlags().GetString("name")
	if givenName == "" {
		seed := time.Now().UTC().UnixNano()
		nameGenerator := namegenerator.NewNameGenerator(seed)

		name = nameGenerator.Generate()
	} else {
		name = givenName
	}

	fmt.Println("Name: ", name)
}
