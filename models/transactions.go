package models

import (
	"encoding/xml"
	"time"
)

type xmlForTransactions struct {
	XMLName      xml.Name       `xml:"xml_name"`
	Transactions []Transactions `xml:"transactions"`
}

type Transactions struct {
	XMLName        xml.Name  `xml:"xml_name"`
	Id             int       `json:"id,omitempty" xml:"id"`
	PayerAcc       string    `json:"payer_acc,omitempty" xml:"payer_acc"`
	BeneficiaryAcc string    `json:"beneficiary_acc,omitempty" xml:"beneficiary_acc"`
	TransferAmount int       `json:"transfer_amount,omitempty" xml:"transfer_amount"`
	CreationDate   time.Time `json:"creation_date" xml:"creation_date"`
}

type PhoneTxReq struct {
	ReceiverPhone string `json:"receiver_phone,omitempty"`
	Amount        int    `json:"amount"`
}

type AccTxReq struct {
	ReceiverAcc string `json:"receiver_acc,omitempty"`
	Amount      int    `json:"amount"`
}
