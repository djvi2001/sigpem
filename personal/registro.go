package personal

import "fmt"

func RegistrarEfectivo(idEfectivo int, nombre string, jerarquia string, sector string) string {
	return fmt.Sprintf("Efectivo %s con ID %d (%s) registrado exitosamente en el sector %s.", nombre, idEfectivo, jerarquia, sector)
}
