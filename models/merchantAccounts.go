package models

import (
	"encoding/xml"
	"time"
)

type XmlForMerchantAccounts struct {
	XMLName          xml.Name           `xml:"xml_name"`
	MerchantAccounts []MerchantAccounts `xml:"merchant_accounts"`
}

type MerchantAccounts struct {
	XMLName      xml.Name  `xml:"xml_name"`
	ID           int       `json:"id,omitempty" xml:"id"`
	AccNumber    string    `json:"acc_number,omitempty" xml:"acc_number"`
	MerchantID   int       `json:"merchant_id,omitempty" xml:"merchant_id"`
	Balance      int       `json:"balance,omitempty" xml:"balance"`
	Locked       bool      `json:"locked,omitempty" xml:"locked"`
	CreationDate time.Time `json:"creation_date" xml:"creation_date"`
	UpDate       time.Time `json:"up_date" xml:"up_date"`
	Deleted      bool      `json:"deleted,omitempty" xml:"deleted"`
}
