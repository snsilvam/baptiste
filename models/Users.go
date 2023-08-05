package models

type Users struct {
	//Identificador del objeto en la base de datos
	ID string `json:"ID"`
	//Fecha de creacion
	CreatedAt string `json:"createdAt"`
	//Fecha de actualizacion
	UpdatedAt string `json:"updatedAt"`
	//Nombre del usuario
	Name string `json:"name"`
	//Email del usuario
	Email string `json:"email"`
}
