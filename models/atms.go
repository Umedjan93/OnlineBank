package models

import "encoding/xml"

type XmlForATMs struct {
	XMLName xml.Name `xml:"atms_list"`
	ATMs    []ATMs   `xml:"atms"`
}

type ATMs struct {
	XMLName      xml.Name `xml:"atm"`
	ID           int      `json:"id" xml:"id"`
	Address      string   `json:"address,omitempty" xml:"address"`
	Balance      int      `json:"balance,omitempty" xml:"balance"`
	MaxCashLimit int      `json:"max_cash_limit,omitempty" xml:"max_cash_limit"`
	Commission   string   `json:"commission" xml:"commission"`
	Owner        string   `json:"owner" xml:"owner"`
}

type AtmList struct {
	Address      string `json:"address,omitempty"`
	Balance      int    `json:"balance,omitempty"`
	MaxCashLimit int    `json:"max_cash_limit,omitempty"`
	Commission   string `json:"commission,omitempty"`
	Owner        string `json:"owner,omitempty"`
}
