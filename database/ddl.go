// данный файл содержит команды для создания sql-таблиц

package database

//	ClientsDataDDL таблица с данными клиентов банка(физ.лиц)
const ClientsDataDDL = `create table if not exists clients_data (
	id serial primary key,
	name varchar(30) not null,
	login text not null unique,
	password text not null,
	phone text not null,
	locked boolean not null default false,
	creation_date date not null default current_timestamp,
	upDate date not null default current_timestamp,
	deleted boolean default false
);`

//	MerchantsDataDDL - таблица с данными клиентов банка (юр.лиц)
const MerchantsDataDDL = `create table if not exists merchants_data (
	id serial primary key,
	name varchar(30) not null,
	company text not null,
	login text not null unique,
	password text not null,
	phone text not null,
	locked boolean not null default false,
	creation_date date not null default current_timestamp,
	upDate date not null default current_timestamp,
	deleted boolean default false
);`

//	ManagersDataDDL - таблица с данными менеджеров банка
const ManagersDataDDL = `create table if not exists managers_data (
	id serial primary key,
	name varchar(30) not null,
	login text not null unique,
	password text not null,
	phone text not null,
	locked boolean not null default false,
	creation_date date not null default current_timestamp,
	upDate date not null default current_timestamp,
	deleted boolean default false
);`

// ClientAccountsDDL - таблица аккаунтов клиентов банка (физ.лиц)
const ClientAccountsDDL = `create table if not exists client_accounts (
	id serial primary key,
	acc_number text not null,
	client_ID integer,
	balance integer,
	locked boolean not null default false,
	creation_date date not null default current_timestamp,
	upDate date not null default current_timestamp,
	deleted boolean default false
);`

// MerchantAccountsDDL - таблица аккаунтов клиентов банка (юр.лиц), занимающихся бизнесом.
const MerchantAccountsDDL = `create table if not exists merchant_accounts (
	id serial primary key,
	acc_number text not null,
	merchant_ID integer,
	balance integer,
	locked boolean not null default false,
	creation_Date date not null default current_timestamp,
	upDate date not null default current_timestamp,
	deleted boolean default false
);`

//	ServicesDDL - таблица онлайн услуг
const ServicesDDL = `create table if not exists services (
	id serial primary key,
	service_name text unique not null,
	merchant_ID integer,
	merchant_name text not null,
	price integer not null,
	creation_date date not null default current_timestamp,
	upDate date not null default current_timestamp,
	deleted boolean default false
);`

//	TransactionsDDL - таблица данных об осуществляемых трансакциях
const TransactionsDDL = `create table if not exists transactions (
	id serial primary key,
	payer_Acc text not null,
	beneficiary_Acc text not null,
	transfer_Amount integer not null,
	creation_Date date not null default current_timestamp
);`

//	atmsDDL - таблица с данными о банкоматах
const ATMsDDL = `create table if not exists ATMs(
	id serial primary key,
	address varchar(30) not null,
	balance integer,
	max_Cash_Limit integer,
	commission text,
	owner text
);`
