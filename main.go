package main

import(
	"fmt"
	"udemy-go/nutrition-calculator/nutriscore"
)

func main() {
	ns := nutriscore.GetNutritionalScore(nutriscore.NutritionalData{
		Energy:              nutriscore.EnergyFromKcal(0),
		Sugars:              nutriscore.SugarGram(10),
		SaturatedFattyAcids: nutriscore.SaturatedFattyAcids(2),
		Sodium:              nutriscore.SodiumMilligram(500),
		Fruits:              nutriscore.FruitsPercent(60),
		Fiber:               nutriscore.FiberGram(4),
		Protein:             nutriscore.ProteinGram(2),
	}, nutriscore.Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
