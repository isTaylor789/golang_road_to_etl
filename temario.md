# **Golang Road to ETL**

## **1. Bases de Go (Main)**
### Objetivo:
Dominar las bases del lenguaje Go: tipos de variables, estructuras de control, arrays, slices, mapas y structs. 

### Temas:
1. **Introducción a Go**
   - Instalación y configuración del entorno.
   - Filosofía del lenguaje.
2. **Sintaxis básica**
   - Variables y constantes.
   - Tipos de datos (primitivos y compuestos).
   - Operadores y control de flujo (`if`, `switch`, `for`).
3. **Estructuras de datos básicas**
   - Arrays, slices y mapas.
   - Manipulación de strings.
4. **Funciones**
   - Declaración, parámetros, valores de retorno.
   - Funciones anónimas y closures.

### Proyecto Final:
**Implementar dos estructuras de datos:**
1. Un árbol binario con operaciones básicas (insertar, buscar, recorrer).
2. Una cola simple con operaciones de encolar y desencolar.

---

## **2. Programación Orientada a Objetos en Go**
### Objetivo:
Comprender cómo se implementa la POO en Go usando structs, interfaces y métodos.

### Temas:
1. **Structs y métodos**
   - Declaración y uso de structs.
   - Métodos asociados a structs.
2. **Interfaces**
   - Definición y uso.
   - Polimorfismo con interfaces.
3. **Principios de diseño**
   - Composición sobre herencia.
   - Encapsulación.

### Proyecto Final:
**Implementar dos patrones de diseño:**
1. Patrón Singleton para gestionar una instancia única.
2. Patrón Factory para crear instancias dinámicas.

---

## **3. Bases de Datos**
### Objetivo:
Trabajar con bases de datos usando Go, con y sin ORM.

### Temas:
1. **Conexión con bases de datos**
   - Introducción a SQL y drivers de Go (ej. `database/sql`).
   - Configuración de bases de datos (ej. PostgreSQL o MySQL).
2. **ORM en Go**
   - Uso de GORM o Ent.
   - Ventajas y desventajas de un ORM.
3. **CRUD básico**
   - Crear, leer, actualizar y eliminar registros con y sin ORM.
4. **Migraciones de bases de datos.**

### Proyecto Final:
**Crear un CRUD funcional:**
- Implementación del CRUD con y sin ORM para comparar enfoques.

---

## **4. Creación de APIs Web**
### Objetivo:
Construir APIs RESTful en Go utilizando un framework como Gin o Echo.

### Temas:
1. **Introducción a frameworks web**
   - Instalación y configuración de Gin/Echo.
   - Conceptos básicos de MVC.
2. **Rutas y controladores**
   - Gestión de rutas dinámicas y controladores.
   - Validación de datos y middlewares.
3. **Consumo de datos**
   - Fetch de APIs externas y parseo de JSON.
4. **Middleware y autenticación básica.**

### Proyecto Final:
**Crear una API funcional:**
- Tema libre, siguiendo arquitectura MVC.
- Debe consumir datos de una API externa y procesarlos.

---

## **4.5 [Opcional] Proyecto Especial**
### Opciones:
1. **Server-side rendering (SSR):**
   - Usar Go para generar páginas HTML dinámicas (ej. con `html/template`).
   - Tema libre, como un blog básico.
2. **Proyecto en Linux:**
   - Crear una herramienta CLI estilo `fastfetch`.
   - Debe mostrar información del sistema (CPU, RAM, etc.) usando librerías de Go.

---

## **5. Testing en Go**
### Objetivo:
Implementar pruebas unitarias, de integración y de rendimiento.

### Temas:
1. **Introducción a testing en Go**
   - Uso del paquete `testing`.
   - Estructura de los tests.
2. **Pruebas unitarias**
   - Tablas de prueba.
   - Mocking.
3. **Pruebas de integración**
   - Validación de flujos completos.
4. **Benchmarks**
   - Medir rendimiento de funciones.

### Proyecto Final:
**Mejorar la API del módulo 4:**
- Agregar pruebas unitarias e integración.
- Incluir benchmarks para endpoints críticos.

---

## **6. Concurrencia en Go**
### Objetivo:
Aprender a manejar la concurrencia y paralelismo en Go.

### Temas:
1. **Goroutines y canales**
   - Creación y uso.
   - Comunicación entre goroutines.
2. **Sincronización**
   - Mutex y WaitGroups.
   - Problemas comunes: data race, deadlocks.
3. **Patrones de concurrencia**
   - Worker Pools.
   - Contextos (`context` package).
4. **Manejo de errores concurrentes.**

### Proyecto Final:
**Implementar un ETL básico:**
1. Leer datos desde un archivo o API.
2. Procesar los datos usando goroutines (transformación).
3. Cargar los datos en una base de datos o archivo.

---

## **Materiales y Recursos**
1. **Documentación oficial de Go:** [https://go.dev/doc/](https://go.dev/doc/)
2. **Tutoriales y ejemplos:**
   - [Tour of Go](https://go.dev/tour/)
   - [Gophercises](https://gophercises.com/)
3. **Libros recomendados:**
   - _"The Go Programming Language"_ de Donovan y Kernighan.
4. **Videos y cursos:**
   - Go por FreeCodeCamp (YouTube).
   - Cursos en plataformas como Udemy o Platzi.

---
