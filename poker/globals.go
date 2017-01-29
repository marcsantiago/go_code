package main

import "github.com/fatih/color"

var (
	player   Player
	computer Player
	spade    = color.New(color.FgBlue, color.Bold).SprintFunc()
	clubs    = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	heart    = color.New(color.FgHiRed, color.Bold).SprintFunc()
	diamond  = color.New(color.FgRed, color.Bold).SprintFunc()
	royal    = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	score    = map[string]int{
		"royal flush":       200,
		"flush":             190,
		"straight":          180,
		"full house":        170,
		"4 of a kind ace":   160,
		"4 of a kind king":  155,
		"4 of a kind queen": 150,
		"4 of a kind jack":  145,
		"4 of a kind 9":     139,
		"4 of a kind 8":     138,
		"4 of a kind 7":     137,
		"4 of a kind 6":     136,
		"4 of a kind 5":     135,
		"4 of a kind 4":     134,
		"4 of a kind 3":     133,
		"4 of a kind 2":     132,
		"4 of a kind 1":     131,

		"3 of a kind ace":   130,
		"3 of a kind king":  125,
		"3 of a kind queen": 120,
		"3 of a kind jack":  125,
		"3 of a kind 9":     119,
		"3 of a kind 8":     118,
		"3 of a kind 7":     117,
		"3 of a kind 6":     116,
		"3 of a kind 5":     115,
		"3 of a kind 4":     114,
		"3 of a kind 3":     113,
		"3 of a kind 2":     112,
		"3 of a kind 1":     111,

		"high pair of a kind ace":   110,
		"high pair of a kind king":  105,
		"high pair of a kind queen": 100,
		"high pair of a kind jack":  95,
		"high pair of a kind 9":     89,
		"high pair of a kind 8":     88,
		"high pair of a kind 7":     87,
		"high pair of a kind 6":     86,
		"high pair of a kind 5":     85,
		"high pair of a kind 4":     84,
		"high pair of a kind 3":     83,
		"high pair of a kind 2":     82,
		"high pair of a kind 1":     81,

		"pair of a kind ace":   80,
		"pair of a kind king":  79,
		"pair of a kind queen": 78,
		"pair of a kind jack":  77,
		"pair of a kind 9":     76,
		"pair of a kind 8":     75,
		"pair of a kind 7":     74,
		"pair of a kind 6":     73,
		"pair of a kind 5":     72,
		"pair of a kind 4":     71,
		"pair of a kind 3":     70,
		"pair of a kind 2":     69,
		"pair of a kind 1":     68,
	}
)
