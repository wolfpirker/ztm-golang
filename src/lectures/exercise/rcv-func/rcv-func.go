//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name string
	health, maxHealth uint
	energy, maxEnergy uint

}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
func (player *Player) changeHealthBy (healthDelta int){
	if healthDelta > 0 {
		player.health += uint(healthDelta)
		if player.health > player.maxHealth{
			player.health = player.maxHealth
		} 	
		fmt.Println("Health restored; add", healthDelta)
	} else {
		if player.health > uint(-healthDelta){
			player.health -= uint(-healthDelta)
		} else{
			player.health = 0
		}
		fmt.Println("Health reduced by", healthDelta)
	}	
}

func (player *Player) changeEnergyBy (healthDelta int){
	// quite the same...
}

func main() {
	hero := Player{
		name: "Lukas",
		health: 100,
		maxHealth: 100,
		energy: 100,
		maxEnergy : 100,		
	}

	hero.changeHealthBy(-20)
	hero.changeEnergyBy(-30)
	hero.changeHealthBy(30)
	hero.changeEnergyBy(10)
}
