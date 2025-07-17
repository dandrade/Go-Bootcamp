# Ejercicios - Semana 4: Interfaces y Manejo de Errores

## Descripción
Esta carpeta contiene ejercicios prácticos para dominar interfaces y manejo de errores en Go. Los ejercicios están diseñados para ser completados en orden, construyendo sobre los conceptos de structs y methods de la Semana 3.

## Ejercicios

### Ejercicio 1: Interfaces Básicas
**Archivo**: `ejercicio1.go`
- Definir interfaces simples
- Implementar interfaces implícitamente
- Usar polimorfismo básico
- Crear un sistema de formas geométricas

### Ejercicio 2: Polimorfismo Avanzado
**Archivo**: `ejercicio2.go`
- Interfaces más complejas
- Múltiples implementaciones
- Composición de interfaces
- Sistema de procesamiento de datos

### Ejercicio 3: Custom Errors
**Archivo**: `ejercicio3.go`
- Crear errores personalizados
- Implementar la interface error
- Error wrapping y unwrapping
- Sistema de validación con errores ricos

### Ejercicio 4: Error Handling Patterns
**Archivo**: `ejercicio4.go`
- Patrones avanzados de manejo de errores
- Error accumulation
- Recuperación de panics
- Sistema de operaciones con retry

## Instrucciones Generales

1. **Lee los TODOs**: Cada ejercicio contiene comentarios TODO que indican qué implementar
2. **No modifiques main()**: La función main ya está implementada para probar tu código
3. **Ejecuta frecuentemente**: Prueba tu código a medida que avanzas
4. **Verifica errores**: Asegúrate de manejar todos los casos de error

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
- [ ] Las interfaces están correctamente definidas
- [ ] Los errores proporcionan información útil
- [ ] Se usan patrones idiomáticos de Go

## Conceptos Clave

### Interfaces
- Definen comportamiento, no implementación
- Implementación implícita
- Permiten polimorfismo
- Deben ser pequeñas y enfocadas

### Error Handling
- Siempre verificar errores
- Proporcionar contexto útil
- Usar error wrapping cuando sea apropiado
- Crear custom errors para casos específicos

## Tips

1. **Interfaces pequeñas**: Prefiere interfaces con pocos métodos
2. **Nombres descriptivos**: Usa convenciones como Writer, Reader, etc.
3. **Error context**: Siempre agrega contexto a los errores
4. **Type assertions**: Usa type switches para manejar diferentes tipos

## Recursos de Ayuda
- Revisa el material teórico en `../teoria/04-interfaces-errors.md`
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Go by Example - Errors](https://gobyexample.com/errors)

---

¡Estas son las herramientas que te convertirán en un desarrollador Go más avanzado! 🚀