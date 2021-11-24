//данный файл нужен для работы с таблицами

package database

const AddClientQuery = `insert into clients_data (name, phone, login, password) values ($1, $2, $3, crypt($4, gen_salt('bf')))`
const AddClientAccQuery = `insert into client_accounts (acc_number, client_ID, balance) values ($1, $2, $3)`
const AddMerchantQuery = `insert into merchants_data (name, company, phone, login, password) values ($1, $2, $3, $4, crypt($5, gen_salt('bf')))`
const AddMerchantAccQuery = `insert into merchant_accounts (acc_number, merchant_ID, balance) values ($1, $2, $3)`
const AddServiceQuery = `insert into services (service_name, price, merchant_name, merchant_ID) values ($1, $2, $3, $4)`
const AddATMQuery = `insert into ATMs (address, balance, max_cash_limit, commission, owner) values ($1, $2, $3, $4, $5)`
const DecryptPassQuery = `select password from clients_data where password = crypt($1, password)`
const GetServices = `select id, service_name, price, merchant_name from services order by id`
const GetAtmsQuery = `select address, balance, max_cash_limit, commission, owner from atms order by address`
const RegistrationSql = `insert into clients_data (name,  phone, login, password) values ($1, $2, $3, crypt($4, gen_salt('bf')))`