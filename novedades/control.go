package novedades

import "fmt"

func RegistrarNovedad(idEfectivo int, descripcion string, fecha string) string {
	return fmt.Sprintf("Novedad registrada para el ID %d en fecha %s: '%s'.", idEfectivo, fecha, descripcion)
}
