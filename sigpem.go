package main

import (
	"fmt"
	"sigpem/novedades"
	"sigpem/permisos"
	"sigpem/personal"
	"sigpem/servicios"
)

func main() {
	fmt.Println("**SISTEMA DE GESTIÓN DE PERSONAL MILITAR (SIGPEM)**")

	//  Módulo 1: Personal
	msgPersonal := personal.RegistrarEfectivo(101, "Carbo Pérez", "Cabo Primero", "CRS Guayas")
	fmt.Println(msgPersonal)

	//  Módulo 2: Novedades
	msgNov := novedades.RegistrarNovedad(101, "S/N", "2026-11-12")
	fmt.Println(msgNov)

	//  Módulo 3: Permisos y Licencias
	msgLic := permisos.AsignarLicencia(101, "Diciembre", 15)
	fmt.Println(msgLic)

	//  Módulo 4: Servicios
	msgServ := servicios.AsignarServicio(101, "Etapa CDP", "Fusilero")
	fmt.Println(msgServ)
}
