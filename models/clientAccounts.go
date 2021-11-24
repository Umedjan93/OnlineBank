package models

import (
	"encoding/xml"
	"time"
)

type XmlForClientAccounts struct {
	XMLName        xml.Name         `xml:"xml_name"`
	ClientAccounts []ClientAccounts `xml:"client_accounts"`
}

type ClientAccounts struct {
	XMLName      xml.Name  `xml:"xml_name"`
	ID           int       `json:"id,omitempty" xml:"id"`
	AccNumber    string    `json:"acc_number,omitempty" xml:"acc_number"`
	ClientID     int       `json:"client_id,omitempty" xml:"client_id"`
	Balance      int       `json:"balance,omitempty" xml:"balance"`
	Locked       bool      `json:"locked,omitempty" xml:"locked"`
	CreationDate time.Time `json:"creation_date" xml:"creation_date"`
	UpDate       time.Time `json:"up_date" xml:"up_date"`
	Deleted      bool      `json:"deleted,omitempty" xml:"deleted"`
}
