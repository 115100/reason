package hoeffding_test

import (
	"fmt"
	"os"

	"github.com/bsm/reason/classifiers/hoeffding"
	"github.com/bsm/reason/core"
)

func ExampleWeather() {
	model := core.NewModel(
		// Target
		&core.Attribute{Name: "play", Kind: core.AttributeKindNominal, Values: core.NewAttributeValues("yes", "no")},

		// Predictores
		&core.Attribute{Name: "outlook", Kind: core.AttributeKindNominal},
		&core.Attribute{Name: "temperature", Kind: core.AttributeKindNumeric},
		&core.Attribute{Name: "humidity", Kind: core.AttributeKindNumeric},
		&core.Attribute{Name: "windy", Kind: core.AttributeKindNominal},
	)

	// Training set data
	trainingSet := []core.MapInstance{
		{"outlook": "sunny", "temperature": 85, "humidity": 85, "windy": "FALSE", "play": "no"},
		{"outlook": "sunny", "temperature": 80, "humidity": 90, "windy": "TRUE", "play": "no"},
		{"outlook": "overcast", "temperature": 83, "humidity": 86, "windy": "FALSE", "play": "yes"},
		{"outlook": "rainy", "temperature": 70, "humidity": 96, "windy": "FALSE", "play": "yes"},
		{"outlook": "rainy", "temperature": 68, "humidity": 80, "windy": "FALSE", "play": "yes"},
		{"outlook": "rainy", "temperature": 65, "humidity": 70, "windy": "TRUE", "play": "no"},
		{"outlook": "overcast", "temperature": 64, "humidity": 65, "windy": "TRUE", "play": "yes"},
		{"outlook": "sunny", "temperature": 72, "humidity": 95, "windy": "FALSE", "play": "no"},
		{"outlook": "sunny", "temperature": 69, "humidity": 70, "windy": "FALSE", "play": "yes"},
		{"outlook": "rainy", "temperature": 75, "humidity": 80, "windy": "FALSE", "play": "yes"},
		{"outlook": "sunny", "temperature": 75, "humidity": 70, "windy": "TRUE", "play": "yes"},
		{"outlook": "overcast", "temperature": 72, "humidity": 90, "windy": "TRUE", "play": "yes"},
		{"outlook": "overcast", "temperature": 81, "humidity": 75, "windy": "FALSE", "play": "yes"},
		{"outlook": "rainy", "temperature": 71, "humidity": 91, "windy": "TRUE", "play": "no"},
	}

	// Train the tree
	tree := hoeffding.New(model, &hoeffding.Config{GracePeriod: 1, EnableTracing: true})
	for _, inst := range trainingSet {
		tree.Train(inst)
		trace := <-tree.Traces()

		if trace != nil {
			fmt.Printf("%#v\n", trace)
		}
	}
	tree.WriteGraph(os.Stdout)

	// Predict
	predicted := tree.Predict(core.MapInstance{"outlook": "sunny", "temperature": 85, "humidity": 85, "windy": "FALSE"})

	fmt.Println(predicted.Top().Value)
	// Output:
	// 0
}
