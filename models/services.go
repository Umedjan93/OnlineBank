package models

import (
	"encoding/xml"
	"time"
)

type xmlForServices struct {
	XMLName  xml.Name   `xml:"xml_name"`
	Services []Services `xml:"services"`
}

type Services struct {
	XMLName      xml.Name  `xml:"xml_name"`
	ID           int       `json:"id,omitempty" xml:"id"`
	ServiceName  string    `json:"service_name,omitempty" xml:"service_name"`
	MerchantID   int       `json:"merchant_id,omitempty" xml:"merchant_id"`
	MerchantName string    `json:"merchant_name,omitempty" xml:"merchant_name"`
	Price        int       `json:"price,omitempty" xml:"price"`
	CreationDate time.Time `json:"creation_date" xml:"creation_date"`
	UpDate       time.Time `json:"up_date" xml:"up_date"`
	Deleted      bool      `json:"deleted,omitempty" xml:"deleted"`
}

type ServiceList struct {
	Id          int    `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	Company     string `json:"company"`
}

type ServiceTxReq struct {
	ServiceId string `json:"service_id" binding:"required"`
	Amount    int    `json:"amount" binding:"required"`
}
