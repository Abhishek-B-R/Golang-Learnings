package main

import "fmt"

type Item struct{
	Name string
	Type string
}

type Player struct{
	Name string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item){
	//TODO: Implement function to add an item to inventory
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s\n",p.Name,item.Name)
}

func (p *Player) DropItem(itemName string){
	//TODO: Implement function to remove an item from inventory
	for i := 0; i < len(p.Inventory); i++{
		if p.Inventory[i].Name == itemName{
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s.\n",p.Name,itemName)
			break
		}
	}
	fmt.Printf("%s does not have %s in inventory.\n",p.Name,itemName)
}

func (p *Player) UseItem(itemName string){
	//TODO: Implement function to use an item
	for _, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "portion"{
				fmt.Printf("%s used %s and feels rejuvenated!\n",p.Name,itemName)
				p.DropItem(itemName)
			} else {
				fmt.Printf("%s used %s.\n", p.Name, itemName)
			}
			return
		}
	}
	fmt.Printf("%s does not hae %s in inventory.\n", p.Name, itemName)
}