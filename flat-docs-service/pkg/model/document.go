package model

type Document struct {
	Docs Docs `json:"docs"`
}

type Docs struct {
	Name       string     `json:"name"`
	Department Department `json:"department"`
	Price      Price      `json:"price"`
	Owner      Owner      `json:"owner"`
	Data       Data       `json:"data"`
	Delivery   Delivery   `json:"delivery"`
	Goods      []Goods    `json:"goods"`
}

type Department struct {
	Code     string   `json:"code"`
	Time     int64    `json:"time"`
	Employee Employee `json:"employee"`
}

type Employee struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Code    string `json:"code"`
}

type Price struct {
	CategoryA string `json:"categoryA"`
	CategoryB string `json:"categoryB"`
	CategoryC string `json:"categoryC"`
}

type Owner struct {
	UUID   string `json:"uuid"`
	Secret string `json:"secret"`
}

type Data struct {
	Transaction Transaction `json:"transaction"`
}

type Transaction struct {
	Type      string `json:"type"`
	UUID      string `json:"uuid"`
	PointCode string `json:"pointCode"`
}

type Delivery struct {
	Company string  `json:"company"`
	Address Address `json:"address"`
}

type Address struct {
	Code      string `json:"code"`
	Country   string `json:"country"`
	Street    string `json:"street"`
	Apartment string `json:"apartment"`
}

type Goods struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Code   string `json:"code"`
}
