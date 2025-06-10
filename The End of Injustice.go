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


func battle(name string, hp, dmg, ehp, edmg, opt, heal, eheal, perk int, slot [5] string, start int)(bool, [5] string, int, int, bool ){
	var enemy string
	var status bool
	var info bool
	var doom int
	orc, inf, phn, dsc, thanos, snap, vnm, fs, mf := false, false, false, false, false, false, false, false, false
	stk, sf, stun, dp := false, false, false, false
	maxhp := hp
	maxehp := ehp
	
	opt = 0
    for i := 0; i < len(slot); i++ {
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
	for start < len(slot){
	
		if status == true || perk <= 0 {
			break
		}
		fmt.Println("Welcome to the equipment store")
		fmt.Println("0. Skip")
	    fmt.Println("1. Oracle for 1 perk")
		fmt.Println("Oracle will bless your hp temporarily boosting it by 150%")
	    fmt.Println("2. Infinity power for 2 perk")
		fmt.Println("Infinity power temporarily boost damage each turn")
		fmt.Println("3. Phoenix pendant for 3 perk")
		fmt.Println("Revive with 1 hp after receiving fatal damage for first time")
		fmt.Println("4. Discipline whip for 3 perk ")
		fmt.Println("Grant chance for critical effect")
		fmt.Println("5. Thanos hand for 2 perk")
		fmt.Println("After 3 turn, a snap will come and kill enemy when their hp is lower than you, else youre death")
		fmt.Println("Grant 40% of enemy stat afterwards")
		fmt.Println("6. Venomous blade for 4 perk")
		fmt.Println("Blade that will make serious bleed, reducing healing effectiveness by 80%")
		fmt.Println("7. Fragile shield for 3 perk")
		fmt.Println("A fragile shield that will protect you from any harm")
		fmt.Println("8. Mysterious flower for 3 perk")
		fmt.Println("Flower that will randomly grant you heal")
		fmt.Println("9. Slappy fish for 2 perk")
		fmt.Println("Small chance to stun opponent for a turn")
		fmt.Println("10. Dodge potion for 2 perk")
		fmt.Println("Small chance to dodge opponent damage")
		fmt.Println("Current perk budget", perk)
		fmt.Println("Current slot", slot)
		fmt.Println("Input equipment", start+1)
		safeScanInt(&opt)
		if opt == 1 && perk >= 1 && orc == false {
			orc = true
			perk -= 1
			slot[start] = "Oracle"
			start++
		}
		 if opt == 2 && perk >= 2 && inf == false {
			inf = true
			perk -= 2
			slot[start] = "Infinity power"
			start++
		}
		if opt == 3 && perk >= 3 && phn == false {
			phn = true
           perk -= 3 
		   slot[start] = "Phoenix pendant"
		   start++
		}
		if opt == 4 && perk >= 3 && dsc == false {
			dsc = true
			perk -= 3 
			slot[start] = "Discipline whip"
			start++
		 }
		 if opt == 5 && perk >= 2 && thanos == false {
			thanos = true
			perk -= 2
			slot[start] = "Thanos hand"
			start++
		 }
		 if opt == 6 && perk >= 4 && vnm == false {
			vnm = true
			perk -= 4
			slot[start] = "Venomous blade"
			start++
		 }
		 if opt == 7 && perk >= 3 && fs == false {
			fs = true
			perk -= 3
            slot[start] = "Fragile shield"
			start++
		 }
		 if opt == 8 && perk >= 3 && mf == false {
		   mf = true
           perk -= 3
		   slot[start] = "Mysterious flower"
		   start++
		 }
		 if opt == 9 && perk >= 2 && sf == false {
			sf = true
			perk -= 2
			slot[start] = "Slappy fish"
			start++
		 }
		 if opt == 10 && perk >= 2 && dp == false {
			dp = true
			perk -= 2
			slot[start] = "Dodge potion"
			start++
		 }
		if opt == 0 {
			refresh()
			break
		}
		refresh()
	}
	for i := 0; i < len(slot); i++ {
		if slot[i] == "Oracle"{
			orc = true
			maxhp += (hp/2)
			hp += (hp/2)
			fmt.Println("Oracle granted")
			fmt.Println("You've received blessing", hp/2 ," hp from oracle")
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
		if slot[i] == "Fragile shield" {
			fs = true
			fmt.Println("Fragile shield granted")
		}
		if slot[i] == "Mysterious flower"{
			mf = true
			fmt.Println("Mysterious flower granted")
		}
        if slot[i] == "Slappy fish"{
			sf = true
			fmt.Println("Slappy fish granted")
		}
		if slot[i] == "Dodge potion" {
			dp = true
			fmt.Println("Dodge potion granted")
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
			opt = rand.Intn(4)
			if opt == 0 && sf == true { 
				stun = true
			}
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
			stk = true
			snap = true
			ehp = 0
			fmt.Println("Opponent snapped by Thanos")
		   } else if doom == 5 {
			snap = true
			hp = 0
			fmt.Println("You lied to thanos") }

        if stun == false {
		if snap == false {
			opt = rand.Intn(5)
			if opt == 0 || opt >= 1 && opt <= 3 {
				if dp == true {
					opt = rand.Intn(5)
					if opt == 0 {
						fmt.Println("You dodged opponent's attack")
					} else {
                     	if fs == true {
					fmt.Println("Fragile shield took", edmg, "Dmg")
                   fmt.Println("Fragile shield is now broken")
				   fs = false
				} else {
					  hp = hp - edmg
			        fmt.Println(enemy," hit you", edmg, "Dmg")
				}
				
					}
			} else {
				if dp == false {
					if fs == true {
					fmt.Println("Fragile shield took", edmg, "Dmg")
                   fmt.Println("Fragile shield is now broken")
				   fs = false} else {
					  hp = hp - edmg
			        fmt.Println(enemy," hit you", edmg, "Dmg")
				   }
				}
			}
			} else if opt == 4{
				if vnm == true {
                    fmt.Println("Opponent suffered from serious bleed")
					ehp = ehp + eheal * 20/100
					fmt.Println(enemy, "healed", eheal*20/100, "Hp")
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
		}}
		if stun == true {
			fmt.Println("Opponent stunned")
			if snap {
				fmt.Println("Thanos returned")
			}
		}
		if mf == true && snap == false {
			opt = rand.Intn(3)
			if opt == 0 {
			fmt.Println("Mysterious flower bloomed")
			fmt.Println("Mysterious flower healed you for", heal, "Hp")
            hp = hp + heal
			} else {
				fmt.Println("Mysterious flower didnt bloom")
			}
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
		stun = false

		
	
	}
	if hp > 0 && ehp <= 0 {
		return true, slot, perk, start, stk
	}

	return false, slot, perk, start, stk
}

func shop(perk, hp, dmg, heal int)(int, int, int, int){
	var opt int
	var gb int
	opt = 1
	for perk > 0 && opt != 0{
		fmt.Println("Your perk is", perk)
		fmt.Println("Option 1 Hp (1 perk each), option 2 Dmg (2 perk each), option 3 Heal (1 perk each), option 0 Skip")
		fmt.Println("Type 4 to play a minigame")
		fmt.Println("Type 99 to gamble with 66% of winning, and 34% of losing")
		safeScanInt(&opt)
		if opt == 1 && perk >= 1 {
			fmt.Println("How many to buy? (Min 1)")
			fmt.Scan(&opt)
			if opt <= 0 {
				opt = 1
			}
			for i := 0; i < opt; i++ {
			hp = hp + rand.Intn(3) + 1
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
				fmt.Println("How many to buy? (Min 1)")
			fmt.Scan(&opt)
			if opt <= 0 {
				opt = 1
			}
			for i := 0; i < opt; i++ {
				dmg = dmg + rand.Intn(4) + 2
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
			
		} else if opt == 3 && perk >= 1 {
				fmt.Println("How many to buy? (Min 1)")
			fmt.Scan(&opt)
			if opt <= 0 {
				opt = 1
			}
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
		   if perk < gb || gb <= 0 {
			fmt.Println("You little fraud eh?")
			refresh()
		   } else {
			if opt == 0{
				perk = perk - gb
				fmt.Println("You lost", gb, "perks")
				refresh()
				opt = 99
			   } else{
				perk = perk + gb
				fmt.Println("You doubled", gb, "perks")
				refresh()
				opt = 99
			   }
		   }
		 
		} else if opt == 4	{
			fmt.Println("How many perks you wanna risk?")
			fmt.Scan(&gb)
			if perk < gb || gb <= 0 {
				fmt.Println("You little fraud eh?")
				refresh()
			} else {
			arr := make([]int, 5)
			arr = []int {rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}
			guess := make([] int, 5)
			refresh()
			fmt.Println("Remember this")
			fmt.Println(arr)
			time.Sleep(3*time.Second)
			done := insertionSort(arr)
			refresh()
			fmt.Println(done)
			for i := 0; i < len(arr); i++ {
				fmt.Scan(&guess[i])
			}
			same := true
			for i := 0; i < len(arr); i++ {
				if guess[i] != arr[i] {
                same = false
				break
				}
			}
			if same {
				perk += gb
				fmt.Println("Congratulation you won", gb , "Perks")
					fmt.Println("The answer was", arr)
				refresh()
			} else {
				perk -= gb
				fmt.Println("You lost", gb , "Perks")
				fmt.Println("The answer was", arr)
				refresh()
			}}
		} else {
			refresh()
			fmt.Println("Not enough balance or invalid options")
		}
	
	}
	
	
	return perk, hp, dmg, heal
}

type player struct{
	nama string
	perk int
	slot [5] string
	start int
	hp,dmg,heal int
}

type enemies struct{
	ehp, edmg, eheal int
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
    fmt.Println("The end of injustice")
	fmt.Println("Battle against opponent and boss")
	fmt.Println("PVP is finally here")
    fmt.Println("How many player?")
	safeScanInt(&opt)
	refresh()
	}
	arrdmg := make([]int, opt)
	arrhp := make([]int, opt)
	arrheal := make([]int, opt)
	arrayplayer := make([]player, opt)
	var enemy enemies
	arrayplayer[jumlah].nama = ""
	condition := false
	status := false
	jumlah = -1
    stk := false
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
		arrayplayer[jumlah].slot = [5]string{"", "", "", "",""}
		arrayplayer[jumlah].start = 0
		arrayplayer[jumlah].perk = 0
        arrayplayer[jumlah].hp = 8
        arrayplayer[jumlah].dmg = 1
        arrayplayer[jumlah].heal = 1
		enemy.ehp = 6
		enemy.edmg = 1
		enemy.eheal = 1
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
		condition, arrayplayer[jumlah].slot, arrayplayer[jumlah].perk, arrayplayer[jumlah].start, stk = battle(arrayplayer[jumlah].nama,arrayplayer[jumlah].hp, arrayplayer[jumlah].dmg, enemy.ehp, enemy.edmg, opt, arrayplayer[jumlah].heal, enemy.eheal,arrayplayer[jumlah].perk, arrayplayer[jumlah].slot, arrayplayer[jumlah].start)
	if condition == true {
		fmt.Println("You won")
		if stk == true {
			fmt.Println("You have received", enemy.ehp*40/100, "Hp", enemy.edmg*40/100, "Dmg", enemy.eheal*40/100, "Heal from Thanos")
			arrayplayer[jumlah].hp += enemy.ehp * 40 / 100
			arrayplayer[jumlah].dmg += enemy.edmg * 40 / 100
			arrayplayer[jumlah].heal += enemy.eheal * 40 / 100
		}
		rng = rand.Intn(3)
		arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng + 5
		arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng
		enemy.ehp  += rand.Intn(3) + 1
		enemy.edmg += 1
		enemy.eheal += 1
		if i >= 3 {
		rng = rand.Intn(10)
	   arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng + 6
        enemy.ehp += i
		enemy.edmg += i + 1
		enemy.eheal += i + 2
		}
		if i >= 4 {
			rng = rand.Intn(20) + 8
	   arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng + 12
			enemy.ehp += rand.Intn(95) + 55
			enemy.edmg +=  rand.Intn(13) + 12
			enemy.eheal += rand.Intn(8) + 29
		}
		if i >= 7 {
			rng = rand.Intn(35) + 13
	    arrayplayer[jumlah].perk = arrayplayer[jumlah].perk + rng + 24
		enemy.ehp += rand.Intn(245) + 80
			enemy.edmg +=  rand.Intn(12) + 95
			enemy.eheal += rand.Intn(6) + 40
		}
		if i == 9 {
			rng = rand.Intn(73) + 27
	    arrayplayer[jumlah].perk = arrayplayer[jumlah].perk+ rng + 82
		enemy.ehp += rand.Intn(999) + 333 + rand.Intn(666)
		enemy.ehp += enemy.ehp * 2
		enemy.edmg +=  rand.Intn(46) + 40
		enemy.eheal += rand.Intn(35) + 45
		}
		enemy.edmg  += rand.Intn(2) + 1
		enemy.eheal  += rand.Intn(2) + 2
		if i <= 8 {
        arrayplayer[jumlah].perk,arrayplayer[jumlah].hp, arrayplayer[jumlah].dmg, arrayplayer[jumlah].heal = shop(arrayplayer[jumlah].perk, arrayplayer[jumlah].hp, arrayplayer[jumlah].dmg, arrayplayer[jumlah].heal)
		}
	} else {
		fmt.Println("You lose")
		arrdmg[jumlah] = arrayplayer[jumlah].dmg
		arrhp[jumlah] = arrayplayer[jumlah].hp
		arrheal[jumlah] = arrayplayer[jumlah].heal
		islose = true
		time.Sleep(2*time.Second)
		refresh()
		break
	}
	}
	if islose == false {
		arrdmg[jumlah] = arrayplayer[jumlah].dmg
		arrhp[jumlah] = arrayplayer[jumlah].hp
		arrheal[jumlah] = arrayplayer[jumlah].heal
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
		fmt.Println(arrayplayer[i].dmg, "Dmg")
		fmt.Println(arrayplayer[i].hp, "Hp")
		fmt.Println(arrayplayer[i].heal, "Heal")
		fmt.Println(arrayplayer[i].slot)
	   }
	   fmt.Println("Press 1 to sort by Dmg")
	   fmt.Println("Press 2 to sort by Hp")
	   fmt.Println("Press 3 to sort by Heal")
	   fmt.Println("Press 4 to check for player name")
	   var done string
	   safeScanInt(&opt)
	   if opt == 1 {
		done = selectionsort(arrdmg)
		fmt.Println(done)
		fmt.Println(arrdmg)
	   } else if opt == 2 {
        done = selectionsort(arrhp)
		fmt.Println(done)
		fmt.Println(arrhp)
	   } else if opt == 3 {
        done = selectionsort(arrheal)
		fmt.Println(done)
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
		fmt.Println(arrayplayer[i].dmg, "Dmg")
		fmt.Println(arrayplayer[i].hp, "Hp")
		fmt.Println(arrayplayer[i].heal, "Heal")
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
	opt = 0
	for opt <= 0 {
		fmt.Println("Enter amount perks for both player")
		safeScanInt(&opt)
	}
	pvp.perk1, pvp.perk2 = opt, opt
	fmt.Println("Welcome", pvp.nama1, "and", pvp.nama2)
	time.Sleep(1*time.Second)
	refresh()
	pvp.hp1, pvp.hp2, pvp.dmg1, pvp.dmg2, pvp.heal1, pvp.heal2 = 8, 8, 1, 1, 1, 1
	fmt.Println(pvp.nama1, "'s shop")
	pvp.perk1, pvp.hp1, pvp.dmg1, pvp.heal1 = shop(pvp.perk1, pvp.hp1, pvp.dmg1, pvp.heal1)
	fmt.Println(pvp.nama2, "'s shop")
	pvp.perk2, pvp.hp2, pvp.dmg2, pvp.heal2= shop(pvp.perk2, pvp.hp2, pvp.dmg2, pvp.heal2)
	maxhp1 := pvp.hp1
	maxhp2 := pvp.hp2
	for pvp.hp1 > 0 && pvp.hp2 > 0 {
		fmt.Println(pvp.nama1, "'s Healthbar:")
    fmt.Println(healthBar(pvp.hp1, maxhp1, 20))
	fmt.Println("Hp : ", pvp.hp1, "Dmg : ", pvp.dmg1, "Heal : ", pvp.heal1)
	fmt.Println(pvp.nama2, "'s Healthbar:")
	fmt.Println(healthBar(pvp.hp2, maxhp2, 20))
	fmt.Println("Hp : ", pvp.hp2, "Dmg : ", pvp.dmg2, "Heal : ", pvp.heal2)
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
		fmt.Println(pvp.nama1, "'s Healthbar:")
    fmt.Println(healthBar(pvp.hp1, maxhp1, 20))
	fmt.Println(pvp.nama2, "'s Healthbar:")
	fmt.Println(healthBar(pvp.hp2, maxhp2, 20))
        fmt.Println(pvp.nama1, "and", pvp.nama2, " is draw")
		time.Sleep(2*time.Second)
		refresh()
	} else if pvp.hp1 <= 0 {
		fmt.Println(pvp.nama1, "'s Healthbar:")
    fmt.Println(healthBar(pvp.hp1, maxhp1, 20))
	fmt.Println(pvp.nama2, "'s Healthbar:")
	fmt.Println(healthBar(pvp.hp2, maxhp2, 20))
		fmt.Println(pvp.nama1, "defeat")
		fmt.Println(pvp.nama2, " win")
		time.Sleep(2*time.Second)
		refresh()
	} else if pvp.hp2 <= 0 {
		fmt.Println(pvp.nama1, "'s Healthbar:")
    fmt.Println(healthBar(pvp.hp1, maxhp1, 20))
	fmt.Println(pvp.nama2, "'s Healthbar:")
	fmt.Println(healthBar(pvp.hp2, maxhp2, 20))
		fmt.Println(pvp.nama1," win")
		fmt.Println(pvp.nama2,"defeat")
		time.Sleep(2*time.Second)
		refresh()
	}
	} else {
	fmt.Println("Invalid option")
	refresh()
	}
	
	}
    
}
func insertionSort(a []int) string {
	done := "Minigame Started"
	n := len(a)
	for i := 1; i < n; i++ {
		k := a[i]
		j := i - 1
		for j >= 0 && a[j] > k {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = k
	}
	return done
}


func selectionsort(arr []int) string {
	n := len(arr)
	done:= "Stat sorted."
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
	return done
}