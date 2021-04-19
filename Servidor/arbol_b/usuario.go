package arbol_b

type Users struct {
	Usuarios []User `json:"Usuarios"`
}

type User struct {
	DPI      int    `json:"DPI"`
	Nombre   string `json:"Nombre"`
	Correo   string `json:"Correo"`
	Password string `json:"Password"`
	Cuenta   string `json:"Cuenta"`
}

func NewUser(DPI int, nombre string, correo string, password string, cuenta string) *User {
	user := User{DPI: DPI, Nombre: nombre, Correo: correo, Password: password, Cuenta: cuenta}

	return &user
}

var VectorUsers Users
