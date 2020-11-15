package domain

type ItemClass int

const (
	Food ItemClass = iota
	Snack
	Soda
	Shot
	Cocktail
	Longdrink
)

type Item struct {
	ItemID int       `json:"itemId" db:"itemId"`
	Class  ItemClass `json:"class" db:"class"`

	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	Ingredients []struct {
		Ingredient string
	} `json:"ingredients" db:"ingredients"`
	Price float64 `json:"price" db:"price"`
}
