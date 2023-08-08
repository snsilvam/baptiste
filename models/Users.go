package models

import "time"

type Users struct {
	//Identificador del objeto en la base de datos
	ID string `json:"ID"`
	//Fecha de creacion
	CreatedAt time.Time `json:"createdAt"`
	//Fecha de actualizacion
	UpdatedAt time.Time `json:"updatedAt"`
	//Nombre del usuario
	Name string `json:"name"`
	//Email del usuario
	Email string `json:"email"`
	//Estado del usuario en la plataforma, si es falso, quiere decir que esta eliminado o desactivado.
	Status bool `json:"status"`
}

type UserInsert struct {
	//Fecha de creacion
	CreatedAt time.Time `json:"createdAt"`
	//Fecha de actualizacion
	UpdatedAt time.Time `json:"updatedAt"`
	//Nombre del usuario
	Name string `json:"name"`
	//Email del usuario
	Email string `json:"email"`
	//Estado del usuario en la plataforma, si es falso, quiere decir que esta eliminado o desactivado.
	Status bool `json:"status"`
}
