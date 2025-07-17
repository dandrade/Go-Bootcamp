# Ejercicios - Semana 3: Structs y Methods

## Descripción
Esta carpeta contiene ejercicios prácticos para dominar el uso de structs y methods en Go. Los ejercicios están diseñados para ser completados en orden, ya que cada uno construye sobre los conceptos del anterior.

## Ejercicios

### Ejercicio 1: Structs Básicos
**Archivo**: `ejercicio1.go`
- Crear structs para representar diferentes entidades
- Inicializar structs de diferentes formas
- Trabajar con campos y acceso a datos
- Implementar una mini biblioteca

### Ejercicio 2: Methods y Receivers
**Archivo**: `ejercicio2.go`
- Implementar methods con value receivers
- Implementar methods con pointer receivers
- Entender cuándo usar cada tipo
- Crear un sistema de cuentas bancarias

### Ejercicio 3: Composición de Structs
**Archivo**: `ejercicio3.go`
- Componer structs complejos
- Usar composición explícita
- Implementar relaciones entre structs
- Crear un sistema de empresa con empleados

### Ejercicio 4: Embedding
**Archivo**: `ejercicio4.go`
- Usar embedding de structs
- Override de methods
- Embedding múltiple
- Crear una jerarquía de vehículos

## Instrucciones Generales

1. **Lee los TODOs**: Cada ejercicio contiene comentarios TODO que indican qué implementar
2. **No modifiques main()**: La función main ya está implementada para probar tu código
3. **Ejecuta frecuentemente**: Prueba tu código a medida que avanzas
4. **Verifica con go vet**: Asegúrate de que tu código no tenga errores

## Cómo Ejecutar

```bash
# Ejecutar un ejercicio específico
go run ejercicio1.go

# Verificar el código
go vet ejercicio1.go

# Formatear el código
go fmt ejercicio1.go
```

## Criterios de Éxito

Para cada ejercicio:
- [ ] El código compila sin errores
- [ ] Pasa todas las pruebas en main()
- [ ] Sigue las convenciones de Go
- [ ] Los métodos están implementados correctamente
- [ ] Se usan los receivers apropiados

## Tips

1. **Value vs Pointer Receiver**:
   - Usa value receiver cuando no necesites modificar el struct
   - Usa pointer receiver cuando necesites modificar el struct o cuando el struct sea grande

2. **Inicialización**:
   - Prefiere la inicialización con nombres de campos
   - Considera crear funciones constructor `New...`

3. **Composición**:
   - Piensa en "tiene un" en lugar de "es un"
   - El embedding es poderoso pero úsalo con cuidado

## Recursos de Ayuda
- Revisa el material teórico en `../teoria/03-structs-methods.md`
- [Go by Example - Structs](https://gobyexample.com/structs)
- [Go by Example - Methods](https://gobyexample.com/methods)

---

¡Buena suerte con los ejercicios! Recuerda: la práctica hace al maestro. 🚀