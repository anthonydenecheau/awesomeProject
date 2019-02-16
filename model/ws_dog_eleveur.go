package model

import "time"

// https://github.com/go-pg/pg/wiki/Model-Definition
type Ws_dog_eleveur struct {
	TableName           struct{} `sql:"ws_dog_eleveur"`
	Id                  int64
	Civilite            string
	Nom                 string
	Prenom              string
	Typ_profil          string
	Professionnel_actif string
	Raison_sociale      string
	On_suffixe          string
	Pays                string
	Id_chien            int64
	Date_maj            time.Time `sql:"-"`
}
