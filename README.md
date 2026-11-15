# RestauGest – Sistema de Gestión de Restaurantes

## 1. Descripción del proyecto

RestauGest es un proyecto académico desarrollado en el lenguaje de programación Go, cuyo propósito es diseñar la estructura inicial de un sistema de gestión de restaurantes.

El sistema está organizado en ocho módulos funcionales que permiten representar las principales operaciones administrativas de un restaurante, como la gestión de clientes, empleados, productos, categorías, inventario, pedidos, pagos y reportes.

En esta primera etapa se desarrolla la arquitectura del proyecto mediante paquetes y funciones exportadas. Las funciones cuentan con parámetros y tipos de retorno definidos, pero sus operaciones internas todavía no están implementadas.

## 2. Objetivo general

Diseñar un sistema de gestión de restaurantes utilizando el lenguaje de programación Go, mediante la organización de funciones y paquetes independientes que permitan administrar las principales operaciones del establecimiento, facilitando el registro, consulta y control de la información.

## 3. Alcance del proyecto

El proyecto contempla el diseño de ocho módulos funcionales:

- Gestión de clientes.
- Gestión de empleados.
- Gestión de productos.
- Gestión de categorías.
- Control de inventario.
- Gestión de pedidos.
- Registro de pagos.
- Generación de reportes de ventas.

En esta etapa inicial, el proyecto se limita a la declaración de funciones, definición de parámetros y retornos, organización de paquetes y verificación de compilación.

No incluye interfaz gráfica, aplicación web o móvil, integración con bases de datos, pagos bancarios en línea, gestión de múltiples sucursales ni implementación completa de las funcionalidades.

## 4. Módulos y funciones exportadas

| Paquete | Funciones exportadas |
|---|---|
| clientes | RegistrarCliente, ConsultarCliente |
| empleados | RegistrarEmpleado, ConsultarEmpleado |
| productos | AgregarProducto, ListarProductos |
| categorias | CrearCategoria, ListarCategorias |
| inventario | ActualizarStock, ConsultarStock |
| pedidos | CrearPedido, ConsultarPedido |
| pagos | RegistrarPago |
| reportes | GenerarReporteVentas |

El sistema cuenta con ocho paquetes funcionales y quince funciones exportadas.

Las funciones comienzan con una letra mayúscula para permitir su utilización desde otros paquetes del mismo módulo de Go.

## 5. Arquitectura del proyecto

El repositorio está organizado de la siguiente manera:

```text
restaugest/
├── go.mod
├── main.go
├── README.md
├── clientes/
│   └── clientes.go
├── empleados/
│   └── empleados.go
├── productos/
│   └── productos.go
├── categorias/
│   └── categorias.go
├── inventario/
│   └── inventario.go
├── pedidos/
│   └── pedidos.go
├── pagos/
│   └── pagos.go
└── reportes/
    └── reportes.go
```

Cada carpeta representa un paquete de Go correspondiente a un módulo funcional del sistema.

El archivo `go.mod` identifica el módulo principal del proyecto, mientras que `main.go` contiene el punto de entrada y las llamadas a las funciones de los diferentes paquetes.

## 6. Requisitos técnicos

- Lenguaje de programación: Go.
- Versión de Go declarada en el proyecto: 1.22.
- Módulo principal: `restaugest`.
- Organización mediante paquetes independientes.
- Funciones exportadas con parámetros y retornos definidos.
- Uso de la biblioteca estándar de Go.

## 7. Instrucciones de compilación y ejecución

Para compilar el proyecto, es necesario contar con Go instalado o utilizar un entorno de desarrollo como GitHub Codespaces.

**Compilar el proyecto:**

```bash
go build ./...
```

Este comando permite verificar la compilación de los paquetes del proyecto.

**Ejecutar el programa principal:**

```bash
go run .
```

El programa muestra los siguientes mensajes:

```text
Bienvenido a RestauGest
Sistema de Gestion de Restaurantes
Los ocho paquetes han sido integrados.
```

Las funciones de los módulos tienen implementaciones provisionales, por lo que la ejecución demuestra la integración estructural de los paquetes y no el funcionamiento completo de un restaurante.

## 8. Estado del proyecto

El proyecto se encuentra en su fase inicial de diseño y organización de paquetes.

Se ha definido la estructura modular, declarado las funciones exportadas y verificado satisfactoriamente la compilación mediante el comando `go build ./...` en GitHub Codespaces.

## 9. Contexto académico

**Asignatura:** Programación Orientada a Objetos

**Actividad:** Aprendizaje Autónomo 1

**Proyecto:** RestauGest – Sistema de Gestión de Restaurantes

**Institución:** Universidad Internacional del Ecuador (UIDE)

**Lenguaje:** Go

**Modalidad:** Trabajo individual

**Año:** 2026
