package Legality

type Legality string

type Legalities struct {
	Standard        Legality `json:"standard"`
	Furure          Legality `json:"future"`
	Historic        Legality `json:"historic"`
	Gladiator       Legality `json:"gladiator"`
	Pioneer         Legality `json:"pioneer"`
	Modern          Legality `json:"modern"`
	Legacy          Legality `json:"legacy"`
	Pauper          Legality `json:"pauper"`
	Vintage         Legality `json:"vintage"`
	Penny           Legality `json:"penny"`
	Commander       Legality `json:"commander"`
	Brawl           Legality `json:"brawl"`
	HistoricBrawl   Legality `json:"historicbrawl"`
	Alchemy         Legality `json:"alchemy"`
	PauperCommander Legality `json:"paupercommander"`
	Duel            Legality `json:"duel"`
	Oldschool       Legality `json:"oldschool"`
	Premodern       Legality `json:"premodern"`
}
