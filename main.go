
package main

import (
    "fmt"

    "restaugest/categorias"
    "restaugest/clientes"
    "restaugest/empleados"
    "restaugest/inventario"
    "restaugest/pagos"
    "restaugest/pedidos"
    "restaugest/productos"
    "restaugest/reportes"
)

func main() {
    fmt.Println("Bienvenido a RestauGest")
    fmt.Println("Sistema de Gestion de Restaurantes")

    clientes.ConsultarCliente(1)
    empleados.ConsultarEmpleado(1)
    productos.ListarProductos()
    categorias.ListarCategorias()
    inventario.ConsultarStock(1)
    pedidos.ConsultarPedido(1)
    pagos.RegistrarPago(1, 25.50, "efectivo")
    reportes.GenerarReporteVentas("2026-01-01", "2026-12-31")

    fmt.Println("Los ocho paquetes han sido integrados.")
}
