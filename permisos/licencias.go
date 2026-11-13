package permisos

import "fmt"

func AsignarLicencia(idEfectivo int, mes string, diaInicio int) string {
	return fmt.Sprintf("Licencia de 30 días asignada al efectivo %d a partir del %d de %s.", idEfectivo, diaInicio, mes)
}
