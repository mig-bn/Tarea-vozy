package user_service_test

import (
	m "Tarea-vozy/models"
	userService "Tarea-vozy/services/user.service"
	"testing"
	"time"
)

//var userID string

func TestCreate(t *testing.T) {

	// Para crear una prueba unitaria para realizar un CRUD completo
	//oid := primitive.NewObjectID()
	//userID = oid.Hex()

	user := m.User{
		Name:      "miguel",
		Email:     "miguel@ejemplo.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		t.Error("La prueba de persistencia de datos del usuario a fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito!")
	}

}

func TestRead(t *testing.T) {

	users, err := userService.Read()

	if err != nil {
		t.Error("se ah presentado un error en la consulta")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("la consulta no retorno datos")
		t.Fail()
	} else {
		t.Log("prueba Finalizada con exito")
	}

}

func TestUpdate(t *testing.T) {

	user := m.User{
		Name:  "Lory",
		Email: "Lory@ejemplo.com",
	}

	err := userService.Update(user, "6272f5f4bd7e61f06422cb4b")

	if err != nil {
		t.Error("Error al tratar de actualizar el usuario")
		t.Fail()
	} else {
		t.Log("la prueba de actualizacion fue un exito")
	}
}

func TestDelete(t *testing.T) {

	err := userService.Delete("6272f5f4bd7e61f06422cb4b")

	if err != nil {
		t.Error("error al tratar de eliminarl el usuario")
	} else {
		t.Log("la prueba de eliminacion fue un exito")
	}

}
