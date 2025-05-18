package main
import (
	"fmt"
	"math/rand"
	"time"
	"github.com/bxcodec/faker/v3"
	"github.com/inancgumus/screen"
	"strings"
)

func refresh(){
	time.Sleep(500*time.Millisecond)
		screen.Clear()
		screen.MoveTopLeft()
	var i int
	i = rand.Intn(40)
	fmt.Println("Cleaning in progress...", i, "%")
	time.Sleep(100*time.Millisecond)
	i = 40 + rand.Intn(30) 
	fmt.Println("Cleaning in progress...", i, "%")
	time.Sleep(100*time.Millisecond)
	i = 70 + rand.Intn(30)
	fmt.Println("Cleaning in progress...", i, "%")
	time.Sleep(500*time.Millisecond)
		screen.Clear()
		screen.MoveTopLeft()
}

func qrefresh(){
	screen.Clear()
		screen.MoveTopLeft()
}

func safeScanInt(opt *int) {
	var input string
	for {
		fmt.Scan(&input)
		_, err := fmt.Sscanf(input, "%d", opt)
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please input a number.")
	}
}

func healthBar(current, max, barLength int) string {
	temphp := current
	if current >= max {
		current = max
	}
	if current <= 0 {
		current = 0
	}
	healthRatio := float64(current) / float64(max)
	filledLength := int(healthRatio * float64(barLength))
	bar := strings.Repeat("█", filledLength) + strings.Repeat("░", barLength-filledLength)
	return fmt.Sprintf("HP: [%s] %d/%d", bar, temphp, max)
}


func battle(name string, hp, dmg, ehp, edmg, opt, heal, eheal, perk int, slot [4] string, start int)(bool, [4] string, int, int ){
	var enemy string
	var status bool
	var info bool
	var doom int
	inf, phn, dsc, thanos, snap, vnm := false, false, false, false, false, false
	maxhp := hp
	maxehp := ehp
	
	opt = 0
    for i := 0; i < 4; i++ {
	if slot[i] != "" {
	fmt.Println("Keep equipment unchanged? 1 for keep, 0 to Edit")
	safeScanInt(&opt)}
	break
}
	if opt == 1 {
		status = true
	} else {
		status = false
	}
	for start < 4{
	
		if status == true || perk <= 0 {
			break
		}
		fmt.Println("Welcome to the equipment store")
		fmt.Println("0. Skip")
	    fmt.Println("1. Oracle for 1 perk")
		fmt.Println("Oracle, once per stage heal 50% of current hp")
	    fmt.Println("2. Infinity power for 2 perk")
		fmt.Println("Infinity power temporarily boost damage each turn")
		fmt.Println("3. Phoenix pendant for 3 perk")
		fmt.Println("Revive with 1 hp after receiving fatal damage for first time")
		fmt.Println("4. Discipline whip for 3 perk ")
		fmt.Println("Grant chance critical effect")
		fmt.Println("5. Thanos hand for 2 perk")
		fmt.Println("After 3 turn, a snap will come and kill enemy when their hp is lower than you, else youre death")
		fmt.Println("6. Venomous blade for 4 perk")
		fmt.Println("Blade that will guide you to do inrecoverable bleed")
		fmt.Println("Current perk budget", perk)
		fmt.Println("Current slot", slot)
		fmt.Println("Input equipment", start+1)
		safeScanInt(&opt)
		if opt == 1 && perk >= 1 {
			perk -= 1
			slot[start] = "Oracle"
			start++
		}
		 if opt == 2 && perk >= 2 {
			perk -= 2
			slot[start] = "Infinity power"
			start++
		}
		if opt == 3 && perk >= 3 {
           perk -= 3 
		   slot[start] = "Phoenix pendant"
		   start++
		}
		if opt == 4 && perk >= 3 {
			perk -= 3 
			slot[start] = "Discipline whip"
			start++
		 }
		 if opt == 5 && perk >= 2 {
			perk -= 2
			slot[start] = "Thanos hand"
			start++
		 }
		 if opt == 6 && perk >= 4 {
			perk -= 4
			slot[start] = "Venomous blade"
			start++
		 }
		if opt == 0 {
			refresh()
			break
		}
		refresh()
	}
	for i := 0; i < 4; i++ {
		if slot[i] == "Oracle"{
			maxhp += (hp/2)
			hp += (hp/2)
			fmt.Println("Oracle granted")
			fmt.Println("You've received blessing from oracle")
		}
		if slot[i] == "Infinity power"{
			inf = true
			fmt.Println("Infinity power granted")
		}
	    if slot[i] == "Phoenix pendant" {
            phn = true
			fmt.Println("Phoenix pendant granted")
		}
		if slot[i] == "Discipline whip" {
            dsc = true
			fmt.Println("Discipline whip granted")
		}
	    if slot[i] == "Thanos hand" {
            thanos = true
			fmt.Println("Thanos hand granted")
		}
		if slot[i] == "Venomous blade" {
            vnm = true
			fmt.Println("Venomous blade granted")
		}


		
	 }
	
	enemy = faker.Name()
	var auto bool
	var first bool
	for hp > 0 && ehp > 0 {
		if first == false{
			fmt.Println(healthBar(hp, maxhp, 20))
			fmt.Println(healthBar(ehp, maxehp, 20))
			first = true
		}
		
		fmt.Println(name+"'s", " stat is")
		fmt.Println(hp, "Hp", dmg, "Dmg", heal, "Heal")
		fmt.Println(slot)
		fmt.Println(enemy+"s", "stat is")
		fmt.Println(ehp, "Hp", edmg, "Dmg", eheal, "Heal")
		fmt.Println("Try to auto battle with '99' ")
		fmt.Println("To Attack press 1", "to Heal press 2")
		if thanos == true && doom == 0 {
			fmt.Println("Thanos snap press 3")
		}
		if auto == true {
			opt = 1
		} else {
			safeScanInt(&opt)
		}
		if opt == 99 {
			auto = true
		}
		refresh()
		if opt == 1 || auto == true {
			if dsc {
               opt = rand.Intn(4)
			   if opt == 0 {
				ehp = ehp - (dmg * 2)
				fmt.Println("Double damage activated")
				fmt.Println(name+"'s", "have dealt", dmg*2, "Dmg ")
			   } else  if opt == 1 {
				ehp = ehp - (dmg * 3)
				fmt.Println("Triple DAMAGE activated")
			fmt.Println(name+"'s", "have dealt", dmg*3, "Dmg ")
			   } else  {
				ehp = ehp - dmg
				fmt.Println("No critical for you")
			fmt.Println(name+"'s", "have dealt", dmg, "Dmg")
			   }
			  
			} else {
				ehp = ehp - dmg
			fmt.Println(name+"'s", "have dealt", dmg, "Dmg")
			}
			
		} else if opt == 2 {
			hp = hp + heal
			fmt.Println(name+"'s", "healed", heal, "Hp")
		}else if opt == 3 && thanos == true && doom == 0 {
			doom += 1
			} else {
			fmt.Println("Invalid option lmao")
		}
        if doom >= 1 {
			fmt.Println("Thanos count at", 4 - doom)
			doom += 1
		}
		if doom == 5 && ehp < hp {
			snap = true
			ehp = 0
			fmt.Println("Opponent snapped by Thanos")
		   } else if doom == 5 {
			snap = true
			hp = 0
			fmt.Println("You lied to thanos") }


		if snap == false {
			opt = rand.Intn(5)
			if opt == 0 || opt >= 1 && opt <= 3 {
				hp = hp - edmg
			fmt.Println(enemy," hit you", edmg, "Dmg")
			} else if opt == 4{
				if vnm == true {
                    fmt.Println("Enemy suffered from inrecoverable bleed")
				} else {
					ehp = ehp + eheal
					fmt.Println(enemy, "healed", eheal, "Hp")
				}
			}
			if phn && hp <= 0 && info == false{
				info = true
				hp = 1
				fmt.Println("You reborn as phoenix with 1 Hp")
			}
			if inf  {
				dmg = dmg + 1
				dmg = dmg + (dmg/2)
				fmt.Println("Damage upgraded to", dmg, "Thanks to infinity power")
			}
			
		} else {
        fmt.Println("Thanos returned")
		}
		if hp > maxhp {
			hp = maxhp
			fmt.Println("Maximum Hp reached")
		}
		if ehp > maxehp {
			ehp = maxehp
			fmt.Println("Enemy maximum Hp reached")
		}
        fmt.Println(healthBar(hp, maxhp, 20))
		fmt.Println(healthBar(ehp, maxehp, 20))

		
	
	}
	if hp > 0 && ehp <= 0 {
		return true, slot, perk, start
	}

	return false, slot, perk, start
}

func shop(perk, hp, dmg, ehp, edmg, heal, eheal int)(int, int, int, int, int, int, int){
	var opt int
	var gb int
	opt = 1
	for perk > 0 && opt != 0{
		fmt.Println("Your perk is", perk)
		fmt.Println("Option 1 Hp (1 perk each), option 2 Dmg (2 perk each), option 3 Heal (1 perk each), option 0 Skip")
		fmt.Println("Type 99 to gamble with 66% of winning, and 34% of losing")
		safeScanInt(&opt)
		if opt == 1 {
			fmt.Println("How many to buy?")
			fmt.Scan(&opt)
			for i := 0; i < opt; i++ {
			hp = hp + rand.Intn(2) + 1
			perk = perk - 1
			if opt >= 5 {
				qrefresh()
			} else {
				refresh()
			}
			fmt.Println("Your hp is", hp, "now")
			if perk <= 0 {
				break
			}
			}
		} else if opt == 2  && perk >= 2{
			fmt.Println("How many to buy?")
			fmt.Scan(&opt)
			for i := 0; i < opt; i++ {
				dmg = dmg + rand.Intn(4) + 1
			perk = perk - 2
			if opt >= 5 {
				qrefresh()
			} else {
				refresh()
			}
			fmt.Println("Your damage is", dmg, "now")
			if perk <= 1 {
				break
			}
			}
			
		} else if opt == 3 {
			fmt.Println("how many to buy?")
			fmt.Scan(&opt)
			for i := 0; i < opt; i++ {
				heal = heal + rand.Intn(2) + rand.Intn(3)
			perk = perk - 1
			if opt >= 5 {
				qrefresh()
			} else {
				refresh()
			}
			fmt.Println("Your heal is", heal, "now")	
			if perk <= 0 {
				break
			}
			}
		} else if opt == 0{
			refresh()
			break
		} else if opt == 99 {
			fmt.Println("How many perks you wanna risk?")
           opt = rand.Intn(3)
		   fmt.Scan(&gb)
		   if perk < gb {
			fmt.Println("You little fraud eh?")
		   } else {
			if opt == 0{
				perk = perk - gb
				fmt.Println("You lost", gb, "perks")
				refresh()
				opt = 99
			   } else{
				perk = perk + gb
				fmt.Println("You doubled", gb, "the perks")
				refresh()
				opt = 99
			   }
		   }
		 
		} else	{
			refresh()
			fmt.Println("Not enough balance or invalid options")
		}
	
	}
	
	
	return perk, hp, dmg, ehp, edmg, heal, eheal
}

type player struct{
	nama string
	perk int
	slot [4] string
	start int
}

type stat struct{
	hp,dmg, heal int
}

type pvp struct{
	hp1, hp2, dmg1, dmg2, heal1, heal2, perk1, perk2 int
	nama1, nama2 string
}

func main(){
	var pvp pvp
	var islose bool
	var opt int
	var opt2 string
	var rng int
	var isfinish bool
	jumlah := 0
    for opt <= 0 {
    fmt.Println("How many player?")
	safeScanInt(&opt)
	refresh()
	}
	arrdmg := make([]int, opt)
	arrhp := make([]int, opt)
	arrheal := make([]int, opt)
	arrayplayer := make([]player, opt)
	arrayplayer[jumlah].perk = 0
	arraystat := make([]stat, opt)
	arrayplayer[jumlah].nama = ""
	arraystat[jumlah].hp = 8
	arraystat[jumlah].dmg = 1
	arraystat[jumlah].heal = 1
	ehp, edmg, eheal := 6, 1, 1
	condition := false
	status := false
	jumlah = -1

	for isfinish != true {
		status = false
		
		if jumlah+1 >= len(arrayplayer){
			fmt.Println("All player have finished")
		}

	if jumlah+1 < len(arrayplayer){
    fmt.Println("Hi player", jumlah+2)
	 fmt.Println("Wanna play? press 1")
	}
	fmt.Println("Past records press 2")
	fmt.Println("Pvp press 3")
	fmt.Println("Finish press 99")
	opt = 99
    fmt.Scan(&opt)
	if opt == 1 && jumlah+1 < len(arrayplayer){
		jumlah++
		arrayplayer[jumlah].slot = [4]string{"", "", "", ""}
		arrayplayer[jumlah].start = 0
		arrayplayer[jumlah].perk = 0
        arraystat[jumlah].hp = 8
        arraystat[jumlah].dmg = 1
        arraystat[jumlah].heal = 1
		ehp = 6
		edmg = 1
		eheal = 1
    for status == false {
    fmt.Println("Welcome, please input your name")
	fmt.Println("Type random to get random name")
	fmt.Scan(&arrayplayer[jumlah].nama)
	if arrayplayer[jumlah].nama == "random" {
		arrayplayer[jumlah].nama = faker.Name()
	}
	fmt.Println("Your name is", arrayplayer[jumlah].nama +"," , "proceed?")
	fmt.Println("To proceed please press 1")
	fmt.Scan(&opt2)
	if opt2 == "1" {
		status = true
		refresh()
	}
	refresh()
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Current level is", i+1)
		if i+1 == 10 {
			fmt.Println("Boss battle alert")
		}
		condition, arrayplayer[jumlah].slot, arrayplayer[jumlah].perk, arrayplayer[jumlah].start = battle(arrayplayer[jumlah].nama,arraystat[jumlah].hp, arraystat[jumlah].dmg, ehp, edmg, opt, arraystat[jumlah].heal, eheal,arrayplayer[jumlah].perk, arrayplayer[jumlah].slot, arrayplayer[jumlah].start)
	if condition == true {
		fmt.Println("You won")
		rng = rand.Intn(4)
		arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng + 3
		arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng
		ehp  += rand.Intn(3) + 1
		edmg += 1
		eheal += 1
		if i >= 3 {
		rng = rand.Intn(9)
	   arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng + 5
        ehp += i
		edmg += i + 1
		eheal += i + 2
		}
		if i >= 4 {
			rng = rand.Intn(20) + 5
	   arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng + 12
			ehp += rand.Intn(95) + 55
			edmg +=  rand.Intn(13) + 12
			eheal += rand.Intn(8) + 29
		}
		if i >= 7 {
			rng = rand.Intn(35) + 5
	    arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng + 24
		ehp += rand.Intn(245) + 80
			edmg +=  rand.Intn(12) + 95
			eheal += rand.Intn(6) + 40
		}
		if i == 9 {
			rng = rand.Intn(73) + 27
	    arrayplayer[jumlah].perk = arrayplayer[jumlah].perk+ rng + 82
		ehp += rand.Intn(999) + 333 + rand.Intn(666)
		ehp += ehp * 2
		edmg +=  rand.Intn(46) + 40
		eheal += rand.Intn(35) + 45
		}
		edmg  += rand.Intn(3)
		eheal  += rand.Intn(2) + 1
		if i < 8 {
        arrayplayer[jumlah].perk,arraystat[jumlah].hp, arraystat[jumlah].dmg, ehp, edmg, arraystat[jumlah].heal,eheal = shop(arrayplayer[jumlah].perk, arraystat[jumlah].hp, arraystat[jumlah].dmg, ehp, edmg, arraystat[jumlah].heal, eheal)
		}
	} else {
		fmt.Println("You lose")
		arrdmg[jumlah] = arraystat[jumlah].dmg
		arrhp[jumlah] = arraystat[jumlah].hp
		arrheal[jumlah] = arraystat[jumlah].heal
		islose = true
		refresh()
		break
	}
	}
	if islose == false {
		arrdmg[jumlah] = arraystat[jumlah].dmg
		arrhp[jumlah] = arraystat[jumlah].hp
		arrheal[jumlah] = arraystat[jumlah].heal
		fmt.Println("Congratulations you won")
		refresh()
	}
	} else if opt == 99 {
		isfinish = true
		fmt.Println("Finished")
	} else if opt == 2 {
	   for i := 0; i <= jumlah; i++ {
		fmt.Println("Player", i+1)
		fmt.Println("Name", arrayplayer[i].nama)
		fmt.Println(arraystat[i].dmg, "Dmg")
		fmt.Println(arraystat[i].hp, "Hp")
		fmt.Println(arraystat[i].heal, "Heal")
		fmt.Println(arrayplayer[i].slot)
	   }
	   fmt.Println("Press 1 to sort by Dmg")
	   fmt.Println("Press 2 to sort by Hp")
	   fmt.Println("Press 3 to sort by Heal")
	   fmt.Println("Press 4 to check for player name")
	   safeScanInt(&opt)
	   if opt == 1 {
		selectionsort(arrdmg)
		fmt.Println(arrdmg)
	   } else if opt == 2 {
        selectionsort(arrhp)
		fmt.Println(arrhp)
	   } else if opt == 3 {
        selectionsort(arrheal)
		fmt.Println(arrheal)
	   } else if opt == 4 {
			fmt.Println("Enter the name you want to check")
			found := false
			fmt.Scan(&opt2)
       for i := 0; i <= jumlah ; i++ {
		   namalower := strings.ToLower(arrayplayer[i].nama)
		if opt2 == arrayplayer[i].nama || strings.HasPrefix(namalower, opt2) {
			refresh()
			fmt.Println("Found player", opt2, "at player", i+1)
			fmt.Println("Name", arrayplayer[i].nama)
		fmt.Println(arraystat[i].dmg, "Dmg")
		fmt.Println(arraystat[i].hp, "Hp")
		fmt.Println(arraystat[i].heal, "Heal")
		fmt.Println(arrayplayer[i].slot)
		found = true
		} 
	   }
       if found == false {
		fmt.Println("No player found")
	   }
		} else {
		fmt.Println("Invalid option")
	   }
	   fmt.Println("Please wait for 5 seconds for refresh")
	   time.Sleep(5*time.Second)
	   refresh()
	} else if opt == 3 {
	fmt.Println("Enter name for player 1")
	fmt.Scan(&pvp.nama1)
	fmt.Println("Enter name for player 2")
	fmt.Scan(&pvp.nama2)
	fmt.Println("Enter amount perks for both player")
	safeScanInt(&opt)
	pvp.perk1, pvp.perk2 = opt, opt
	fmt.Println("Welcome", pvp.nama1, "and", pvp.nama2)
	time.Sleep(1*time.Second)
	refresh()
	pvp.hp1, pvp.hp2, pvp.dmg1, pvp.dmg2, pvp.heal1, pvp.heal2 = 8, 8, 1, 1, 1, 1
	fmt.Println(pvp.nama1, "'s shop")
	pvp.perk1, pvp.hp1, pvp.dmg1, ehp, edmg, pvp.heal1, eheal = shop(pvp.perk1, pvp.hp1, pvp.dmg1, ehp, edmg, pvp.heal1, eheal)
	fmt.Println(pvp.nama2, "'s shop")
	pvp.perk2, pvp.hp2, pvp.dmg2, ehp, edmg, pvp.heal2, eheal = shop(pvp.perk2, pvp.hp2, pvp.dmg2, ehp, edmg, pvp.heal2, eheal)
	maxhp1 := pvp.hp1
	maxhp2 := pvp.hp2
	for pvp.hp1 > 0 && pvp.hp2 > 0 {
    fmt.Println(healthBar(pvp.hp1, maxhp1, 20))
	fmt.Println(healthBar(pvp.hp2, maxhp2, 20))
	fmt.Println("Press action for ", pvp.nama1)
	fmt.Println("Attack press 1, Heal press 2")
	fmt.Scan(&opt)
	place := opt
	fmt.Println("Press action for", pvp.nama2)
	fmt.Println("Attack press 1, Heal press 2")
	fmt.Scan(&opt)
	refresh()
	if place == 1 {
		fmt.Println(pvp.nama2, " is hit by", pvp.dmg1, "Dmg")
		pvp.hp2 = pvp.hp2 - pvp.dmg1
	} else if place == 2 {
		fmt.Println(pvp.nama1, " healed by", pvp.heal1, "Heal")
		pvp.hp1 = pvp.hp1 + pvp.heal1
	} else {
		fmt.Println("Invalid option lmao")
	}
	if opt == 1 {
		fmt.Println(pvp.nama1, " is hit by", pvp.dmg2, "Dmg")
		pvp.hp1 = pvp.hp1 - pvp.dmg2
	} else if opt == 2 {
		fmt.Println(pvp.nama2, " healed by", pvp.heal2, "Heal")
		pvp.hp2 = pvp.hp2 + pvp.heal2
	} else {
		fmt.Println("Invalid option lmao")
	}
	if pvp.hp1 > maxhp1 {
		fmt.Println(pvp.nama1, " overhealed")
		pvp.hp1 = maxhp1
	}
	if pvp.hp2 > maxhp2 {
		fmt.Println(pvp.nama2, " overhealed")
		pvp.hp2 = maxhp2
	}
	}
	if pvp.hp1 <= 0 && pvp.hp2 <= 0 {
        fmt.Println(pvp.nama1, "and", pvp.nama2, " is draw")
	} else if pvp.hp1 <= 0 {
		fmt.Println(pvp.nama1, "defeat")
		fmt.Println(pvp.nama2, " win")
	} else if pvp.hp2 <= 0 {
		fmt.Println(pvp.nama1," win")
		fmt.Println(pvp.nama2,"defeat")
	}
	} else {
	fmt.Println("Invalid option")
	refresh()
	}
	
	}
    
}


func selectionsort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
