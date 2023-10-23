package models

type MinCost struct {
	Mincost_found int   `json:"mincost_found"`
	Cost          []int `json:"integer_array"`
}

type HighEarner struct {
	Department string `json:"department"`
	Employee   string `json:"employee"`
	Salary     int    `json:"salary"`
}

type Transaction struct {
	Tool string `json:"tool"`
}

type CustomData struct {
	BoolField    bool    `json:"boolField"`
	IntField     int     `json:"intField"`
	StringField  string  `json:"stringField"`
	PointerField *string `json:"pointerField"`
	ExtraField   struct {
		Head int `json:"head"`
		Tail int `json:"tail"`
	} `json:"extra-field"`
}
