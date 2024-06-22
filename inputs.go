package main

import (
	"fmt"
	"strconv"
)

func getActionId() int {
	var input string
	fmt.Println("0: Get leagues and produce message")
	fmt.Println("1: Get all countries")
	fmt.Println("2: Get all leagues")
	fmt.Println("3: Get all the players for a given team")
	fmt.Println("4: Get all the teams for a given league")
	fmt.Println("5: Get standing for a given league")
	fmt.Println("Enter action id (press enter to default to 0):")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0
	}
	actionId := 0
	if input != "" {
		var err error
		actionId, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Defaulting teamId to 0")
			return 0
		}
	}

	return actionId
}

func getLeagueId() int {
	var leagueId int
	leagues := getLeagues()
	fmt.Println("Select a league")
	fmt.Println("----------------")
	for i, l := range leagues {
		fmt.Printf("%s - %d\n", i, l)
	}

	_, err := fmt.Scan(&leagueId)
	if err != nil {
		fmt.Println("Invalid input")
	}

	return leagueId
}

func getTeamId() int {
	var input string
	fmt.Println("Enter team id (press enter to default to 541):")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 541
	}
	teamId := 541
	if input != "" {
		var err error
		teamId, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Defaulting teamId to 541")
			return 541
		}
	}

	return teamId
}

func getSeason() int {
	var input string
	fmt.Println("Enter the season: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 2023
	}
	season := 2023
	if input != "" {
		var err error
		season, err = strconv.Atoi(input)
		if err != nil {
			return 2023
		}
	}

	return season
}
