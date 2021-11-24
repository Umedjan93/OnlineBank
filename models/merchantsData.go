package models

import (
	"encoding/xml"
	"time"
)

type XmlForMerchantData struct {
	XMLName       xml.Name        `xml:"xml_name"`
	MerchantsData []MerchantsData `xml:"merchants_data"`
}

type MerchantsData struct {
	XMLName      xml.Name  `xml:"xml_name"`
	ID           int64     `json:"id,omitempty" xml:"id"`
	Name         string    `json:"name,omitempty" xml:"name"`
	Company      string    `json:"company,omitempty" xml:"company"`
	Login        string    `json:"login,omitempty" xml:"login"`
	Password     string    `json:"password,omitempty" xml:"password"`
	Phone        string    `json:"phone,omitempty" xml:"phone"`
	Locked       bool      `json:"locked,omitempty" xml:"locked"`
	CreationDate time.Time `json:"creation_date" xml:"creation_date"`
	UpDate       time.Time `json:"up_date" xml:"up_date"`
	Deleted      bool      `json:"deleted,omitempty" xml:"deleted"`
}
