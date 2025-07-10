# Proyecto - Sistema de Inventario Simple

## Descripción

Desarrollar un sistema de inventario simple que permita gestionar productos usando las estructuras de datos aprendidas en la semana 2 (maps, slices, strings).

## Objetivos

- Aplicar maps para almacenar información de productos
- Usar slices para listas y reportes
- Manipular strings para búsquedas y formateo
- Implementar un menú interactivo
- Practicar validación de datos

## Funcionalidades Requeridas

### 1. Gestión de Productos
- Agregar nuevos productos
- Modificar productos existentes
- Eliminar productos
- Buscar productos por ID o nombre

### 2. Consultas
- Listar todos los productos
- Buscar productos por categoría
- Filtrar por rango de precios
- Mostrar productos con stock bajo

### 3. Reportes
- Inventario total
- Valor total del inventario
- Productos más caros/baratos
- Estadísticas por categoría

### 4. Interfaz
- Menú interactivo en consola
- Validación de entrada
- Mensajes de error claros
- Formateo de salida

## Estructura del Producto

```go
type Product struct {
    ID          string
    Name        string
    Category    string
    Price       float64
    Stock       int
    Description string
}
```

## Comandos de Ejecución

```bash
# Ejecutar el proyecto
go run main.go

# Ejecutar con datos de ejemplo
go run main.go -demo
```

## Criterios de Evaluación

- **Funcionalidad**: Todas las características funcionan correctamente
- **Uso de Maps**: Correcta implementación para almacenar productos
- **Manipulación de Strings**: Búsquedas y formateo eficientes
- **Validación**: Manejo adecuado de errores
- **Interfaz**: Menú claro y fácil de usar
- **Código**: Organizado y comentado adecuadamente

## Extensiones Opcionales

- Persistencia de datos en archivo
- Importar/exportar productos
- Historial de movimientos
- Categorías dinámicas
- Código de barras automático