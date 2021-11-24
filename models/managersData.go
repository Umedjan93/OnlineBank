package models

import (
	"encoding/xml"
	"time"
)

type XmlForManagersData struct {
	XMLName      xml.Name       `xml:"xml_name"`
	ManagersData []ManagersData `xml:"managers_data"`
}

type ManagersData struct {
	XMLName      xml.Name  `xml:"xml_name"`
	id           int       `json:"id,omitempty" xml:"id"`
	name         string    `json:"name,omitempty" xml:"name"`
	login        string    `json:"login,omitempty" xml:"login"`
	password     string    `json:"password,omitempty" xml:"password"`
	phone        string    `json:"phone,omitempty" xml:"phone"`
	locked       bool      `json:"locked,omitempty" xml:"locked"`
	creationDate time.Time `json:"creation_date" xml:"creation_date"`
	upDate       time.Time `json:"up_date" xml:"up_date"`
	deleted      bool      `json:"deleted,omitempty" xml:"deleted"`
}
