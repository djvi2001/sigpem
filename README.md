# SIGPEM — Sistema de Gestión de Personal Militar

El presente sistema modular desarrollado en **Go (Golang)** esta siendo desarrollado bajo un enfoque de programación funcional 
para la administración operativa de personal, novedades, permisos y servicios.

//Arquitectura y Estructura del Proyecto//

El sistema sigue una arquitectura modular en la cual se desglosa un paquete por cada función:

```text
sigpem/
├── go.mod                 # Definición del módulo principal de Go
├── sigpem.go              # Punto de entrada y orquestador de módulos
├── personal/
│   └── registro.go        # Gestión y registro de efectivos militares
├── novedades/
│   └── control.go         # Reporte y control de novedades diarias
├── permisos/
│   └── licencias.go       # Asignación de permisos y licencias
└── servicios/
    └── operaciones.go     # Asignación de guardias, puestos y funciones
