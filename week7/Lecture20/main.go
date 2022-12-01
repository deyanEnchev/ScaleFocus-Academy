package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type CocktailBartender struct {
	Drinks []struct {
		IDDrink                     string      `json:"idDrink"`
		StrDrink                    string      `json:"strDrink"`
		StrDrinkAlternate           interface{} `json:"strDrinkAlternate"`
		StrTags                     string      `json:"strTags"`
		StrVideo                    interface{} `json:"strVideo"`
		StrCategory                 string      `json:"strCategory"`
		StrIBA                      string      `json:"strIBA"`
		StrAlcoholic                string      `json:"strAlcoholic"`
		StrGlass                    string      `json:"strGlass"`
		StrInstructions             string      `json:"strInstructions"`
		StrInstructionsES           interface{} `json:"strInstructionsES"`
		StrInstructionsDE           string      `json:"strInstructionsDE"`
		StrInstructionsFR           interface{} `json:"strInstructionsFR"`
		StrInstructionsIT           string      `json:"strInstructionsIT"`
		StrInstructionsZHHANS       interface{} `json:"strInstructionsZH-HANS"`
		StrInstructionsZHHANT       interface{} `json:"strInstructionsZH-HANT"`
		StrDrinkThumb               string      `json:"strDrinkThumb"`
		StrIngredient1              string      `json:"strIngredient1"`
		StrIngredient2              string      `json:"strIngredient2"`
		StrIngredient3              string      `json:"strIngredient3"`
		StrIngredient4              string      `json:"strIngredient4"`
		StrIngredient5              interface{} `json:"strIngredient5"`
		StrIngredient6              interface{} `json:"strIngredient6"`
		StrIngredient7              interface{} `json:"strIngredient7"`
		StrIngredient8              interface{} `json:"strIngredient8"`
		StrIngredient9              interface{} `json:"strIngredient9"`
		StrIngredient10             interface{} `json:"strIngredient10"`
		StrIngredient11             interface{} `json:"strIngredient11"`
		StrIngredient12             interface{} `json:"strIngredient12"`
		StrIngredient13             interface{} `json:"strIngredient13"`
		StrIngredient14             interface{} `json:"strIngredient14"`
		StrIngredient15             interface{} `json:"strIngredient15"`
		StrMeasure1                 string      `json:"strMeasure1"`
		StrMeasure2                 string      `json:"strMeasure2"`
		StrMeasure3                 string      `json:"strMeasure3"`
		StrMeasure4                 interface{} `json:"strMeasure4"`
		StrMeasure5                 interface{} `json:"strMeasure5"`
		StrMeasure6                 interface{} `json:"strMeasure6"`
		StrMeasure7                 interface{} `json:"strMeasure7"`
		StrMeasure8                 interface{} `json:"strMeasure8"`
		StrMeasure9                 interface{} `json:"strMeasure9"`
		StrMeasure10                interface{} `json:"strMeasure10"`
		StrMeasure11                interface{} `json:"strMeasure11"`
		StrMeasure12                interface{} `json:"strMeasure12"`
		StrMeasure13                interface{} `json:"strMeasure13"`
		StrMeasure14                interface{} `json:"strMeasure14"`
		StrMeasure15                interface{} `json:"strMeasure15"`
		StrImageSource              string      `json:"strImageSource"`
		StrImageAttribution         string      `json:"strImageAttribution"`
		StrCreativeCommonsConfirmed string      `json:"strCreativeCommonsConfirmed"`
		DateModified                string      `json:"dateModified"`
	} `json:"drinks"`
}

func (cb *CocktailBartender) Start() {
	fmt.Println("Hello, Sir. What are you going to drink?")
	var result string
	for result != "nothing" {
		_, err := fmt.Scanln(&result)

		if err != nil {
			log.Fatal(err)
		}

		u := ParseURL(&result)

		res, err := http.Get(u.String())
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(body, &cb)

		if err != nil {
			log.Fatal(err)
		}
		// s, _ := json.MarshalIndent(cb,"","\t")
		// fmt.Println(string(s))

		for _, instr := range cb.Drinks {

			if strings.Contains(instr.StrInstructions, ".") {
				splitted := strings.Split(instr.StrInstructions, ".")
				for i, v := range splitted {
					if i == len(splitted)-1 {
						break
					}
					fmt.Println(v + ".")

				}

			} else {
				fmt.Println(instr.StrInstructions)
			}
			break

		}
	}

}

func ParseURL(result *string) *url.URL {
	u, err := url.Parse("https://www.thecocktaildb.com")
	if err != nil {
		log.Fatal(err)
	}
	u.Path = "/api/json/v1/1/search.php"
	q := u.Query()
	q.Add("s", *result)
	u.RawQuery = q.Encode()

	return u
}

func main() {
	cb := CocktailBartender{}
	cb.Start()

}
