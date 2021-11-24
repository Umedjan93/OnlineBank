package models

import (
	"encoding/xml"
	"time"
)

type XmlForClientsData struct {
	XMLName     xml.Name      `xml:"xml_name"`
	ClientsData []ClientsData `xml:"clients_data"`
}
type ClientsData struct {
	XMLName      xml.Name  `xml:"clients_data" json:"-"`
	ID           int64     `json:"id,omitempty" xml:"id"`
	Name         string    `json:"name,omitempty" xml:"name"`
	Login        string    `json:"login,omitempty" xml:"login"`
	Password     string    `json:"password,omitempty" xml:"password"`
	Phone        string    `json:"phone,omitempty" xml:"phone"`
	Locked       bool      `json:"locked,omitempty" xml:"locked"`
	CreationDate time.Time `json:"creation_date" xml:"creation_date"`
	UpDate       time.Time `json:"up_date" xml:"up_date"`
	Deleted      bool      `json:"deleted,omitempty" xml:"deleted"`
}
