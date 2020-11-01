package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/dgrijalva/jwt-go"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"adress"`
	Token    string `json:"token"`
	ApiToken string `json:"api_token"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Manager  string `json:"manager"`
}

type User struct {
	id       primitive.ObjectID `bson:"_id"`
	Login    string             `bson:"login"`
	Password string             `bson:"password"`
	Token    string             `bson:"token"`
	Address  string             `bson:"address"`
	Phone    string             `bson:"phone"`
	Email    string             `bson:"email"`
	Manager  string             `bson:"manager"`
	Role     string             `bson:"role"`
}

type Voc struct {
	Login string   `bson:"login"`
	Name  string   `bson:"name"`
	Voc   []string `bson:"voc"`
}

type Rvoc struct {
	Request []string `json:"request"`
}

type Wrvoc struct {
	Login string `json:"login"`
	Vocs  []Voc  `json:"voc"`
}

type Animal struct {
	Medicine  Medicine `bson:"medicine"`
	Priute    Priute   `bson:"priute"`
	Params    Params   `bson:"params"`
	Events    []Event  `bson:"events"`
	Socialize bool     `bson:"socialize"`
	Ready     bool     `bson:"ready"`
	Login     string   `bson:"login"`
	Type      string   `bson:"type"`
	Owner     Owner    `bson:"owner"`
	Name      string   `bson:"name"`
}

type Docs struct {
	CardN     string    `bson:"card"`
	Id_metka  string    `bson:"id_m"`
	Act_in    string    `bson:"act_in"`
	Date_in   time.Time `bson:"date_in"`
	Adm_place string    `bson:"adm_place"`
	Act_get   string    `bson:"act_get"`
	Place_get string    `bson:"place_get"`
	Act_out   string    `bson:"act_out"`
}

type Medicine struct {
	Vaccine     []Vaccine   `bson:"vaccine"`
	Aparazite   []Aparazite `bson:"aparazite"`
	Place_in    string      `bson:"place_in"`
	Doctor      string      `bson:"act_out"`
	Status      string      `bson:"status"`
	Status_date time.Time   `bson:"st_date"`
	Sterilized  Sterilized  `bson:"sterilized"`
	Socialize   bool        `bson:"socialize"`
}

type Vaccine struct {
	N_pp  string    `bson:"n_pp"`
	Date  time.Time `bson:"date"`
	Drug  string    `bson:"drug"`
	Serie string    `bson:"serie"`
}

type Aparazite struct {
	N_pp string    `bson:"n_pp"`
	Date time.Time `bson:"date"`
	Drug string    `bson:"drug"`
	Dose string    `bson:"dose"`
}

type Sterilized struct {
	Status  bool      `bson:"status"`
	Date    time.Time `bson:"date"`
	Comment string    `bson:"comment"`
}

type Priute struct {
	Aviary      string      `bson:"place_in"`
	Documents   Docs        `bson:"docs"`
	Responsible Responsible `bson:"responsible"`
}

type Responsible struct {
	Adress   string `bson:"adress"`
	Expl_org string `bson:"expl_org"`
	Manager  string `bson:"manager"`
	Resp_own string `bson:"resp_own"`
}

type Owner struct {
	FL FL `bson:"fl"`
	UL UL `bson:"ul"`
}

type FL struct {
	Name      string    `bson:"name"`
	Id_num    string    `bson:"id_num"`
	Id_date   time.Time `bson:"id_date"`
	Id_plase  string    `bson:"id_plase"`
	Reg_place string    `bson:"place"`
	Phone     string    `bson:"phone"`
}

type UL struct {
	Name       string `bson:"name"`
	UL_place   string `bson:"ul_place"`
	UL_phone   string `bson:"ul_phone"`
	Name_owner string `bson:"name_owner"`
	Place      string `bson:"place"`
	Phone      string `bson:"phone"`
}

type Params struct {
	Type     string  `bson:"type"`
	Year     int64   `bson:"year"`
	Mass     float64 `bson:"mass"`
	Sex      string  `bson:"sex"`
	Name     string  `bson:"name"`
	Sub_type string  `bson:"sub_type"`
	Color    string  `bson:"color"`
	Fur      string  `bson:"fur"`
	Ears     string  `bson:"ears"`
	Size     string  `bson:"size"`
	Tale     string  `bson:"tale"`
	Unique   string  `bson:"unique"`
}

type Event struct {
	Date    time.Time `bson:"date"`
	Comment string    `bson:"comment"`
	Status  string    `bson:"status"`
	Act_n   string    `bson:"act_n"`
}

type CSV struct {
	N             string
	Card_N        string
	Type          string
	Year          string
	Mass          string
	Name          string
	Sex           string
	Sub_type      string
	Color         string
	Fur           string
	Ear           string
	Tale          string
	Size          string
	Uni           string
	Volier_n      string
	Id_m          string
	Sterilized    string
	Doctor        string
	Social        string
	Act_in        string
	Date_in       string
	Adm_place     string
	Act_get       string
	Get_Place     string
	UL_name       string
	UL_adr        string
	UL_phone      string
	Name_own      string
	Addres_own    string
	Phone_own     string
	FL_name       string
	Pasport       string
	P_place       string
	P_date        string
	FL_adress     string
	FL_Phone      string
	Date_in_event string
	Act_in_event  string
	Date_out      string
	Out           string
	Act_out       string
	Priute_adress string
	Expl_org      string
	Manager       string
	Resp_own      string
	N_pp_apar     string
	Date_apar     string
	Drug_apar     string
	Dose_apar     string
	N_pp_vac      string
	Date_vac      string
	Drug_vac      string
	Serie_vac     string
	Date_status   string
	Status        string
}
