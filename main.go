package main

import (
	"fmt"
	"log"
	"net/http"

	"time"

	m "Tarea-vozy/models"
	userService "Tarea-vozy/services/user.service"
	"encoding/json"
)

func main() {
	//inicio
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/read", Read)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/update", Update)

	log.Println("Server corriendo iniciado...")
	http.ListenAndServe(":8080", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		user := m.User{
			Name:      r.FormValue("Name"),
			Email:     r.FormValue("Email"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		log.Println(user)

		err := userService.Create(user)

		if err != nil {
			panic(err.Error())
		} else {
			log.Println("Insertado con exito!")
		}
	}

}

func Read(w http.ResponseWriter, r *http.Request) {
	var e []byte
	users, err := userService.Read()

	if err != nil {
		panic(err.Error())
	} else {
		e, err = json.Marshal(users)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(e))
		log.Println("Leido con exito!")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(e)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := m.User{
			Name:      r.FormValue("Name"),
			Email:     r.FormValue("Email"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		log.Println(user)

		err := userService.Update(user, r.FormValue("ID"))

		if err != nil {
			panic(err.Error())
		} else {
			log.Println("Modificado con exito!")
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idUsuario := r.FormValue("ID")
	fmt.Println(idUsuario)

	err := userService.Delete(idUsuario)

	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Eliminado con exito!")
	}
}
