package servicios

import "fmt"

func AsignarServicio(idEfectivo int, puesto string, funcionAsignada string) string {
	return fmt.Sprintf("Efectivo %d asignado al puesto '%s' con función de '%s'.", idEfectivo, puesto, funcionAsignada)
}
