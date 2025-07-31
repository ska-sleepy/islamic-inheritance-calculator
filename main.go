package main

import "fmt"

type Inheritance_tree struct {
	fathers_amount float64
	mothers_amount float64
	partner_amount float64
	daughters_amount []float64
	sons_amount []float64
}

func new_tree() *Inheritance_tree{
	return &Inheritance_tree{
		fathers_amount: 0,
		mothers_amount: 0,
		partner_amount: 0
	}

}

func print_tree(tree Inheritance_tree){
		
}

func main(){
	var choice bool;
	var amount float64;
	tree := new_tree(); 

	// get the amount of liquidated inheritence 
	fmt.Print("enter the liquidated amount of the estate: ");
	fmt.Scan(&amount);

	// navigating the decision tree
	fmt.Println("do you have a living parent?");
	fmt.Scan(&choice);

	if choice {
		fmt.Println("are both of them alive?");
		fmt.Scan(&choice)
		if choice {
			tree.fathers_amount = amount/6;
			tree.mothers_amount = amount/6;
		}else{
			fmt.Println("who's alive? (enter the corrispondering number for choice)\n0-Father\n1-Mother\n")
			fmt.Scan(&choice)
			if choice{
				tree.mothers_amount = amount/6;
			} else {
				tree.fathers_amount = amount/6;
			}
		}

	}



	

	
}
