# [Herramientas de test](https://github.com/FabriConde/CLIMB-VR/issues/15)

Para las pruebas en Go se puede elegir entre realizar pruebas unitarias, pruebas de integración o pruebas de funcionalidad dependiendo del contexto. Para facilitar el proceso, Go permite usar herramientas como `go test` para ejecutar pruebas de forma automática. También se pueden definir pruebas específicas usando el marco de prueba estándar con funciones que comienzan con Test y recibir el parámetro `*testing.T` para reportar fallos. Esto hace que sea sencillo personalizar y estructurar las pruebas según las necesidades del proyecto, como incluir datos específicos o generar condiciones de prueba más controladas y específicas para cada situación.

# Herramientas de test para Go
Existen varias bibliotecas de prueba para Go disponibles en GitHub que extienden la funcionalidad del marco de pruebas estándar (testing), agregando características adicionales como aserciones, mocks, pruebas de integración, etc.

## [Testify](https://github.com/stretchr/testify)
Testify es una de las bibliotecas de pruebas más populares en Go. Proporciona funciones útiles como aserciones y mocks para hacer las pruebas más legibles y fáciles de escribir. Incluye una amplia variedad de aserciones para comparar valores, comprobar condiciones y manejar errores de forma más intuitiva que con la biblioteca estándar.

**Características clave:**
- Aserciones más legibles (por ejemplo, assert.Equal(t, expected, actual)).
- Mocks para pruebas de integración.
- Hooks para pruebas de setup y teardown.

## [GoMock](https://github.com/golang/mock)
GoMock es la herramienta oficial de Google para la creación de mocks en pruebas unitarias en Go. Se utiliza junto con el paquete testing y es especialmente útil cuando se necesita simular comportamientos de interfaces en pruebas de unidad, lo que es común en aplicaciones Go que emplean interfaces.

**Características clave:**
- Generación automática de mocks a partir de interfaces.
- Compatible con testing para integrarse fácilmente en tu flujo de trabajo.
- Útil para probar componentes que dependen de interfaces externas.

## [Ginkgo](https://github.com/onsi/ginkgo)
Ginkgo es un framework de pruebas BDD (Behavior Driven Development) para Go. Ofrece una estructura de pruebas más descriptiva y legible, similar a otros frameworks BDD como Jasmine o RSpec en JavaScript y Ruby, respectivamente.

**Características clave:**
- Sintaxis tipo BDD para escribir pruebas más legibles y descriptivas.
- Soporte para pruebas paralelas y pruebas de integración.
- Ideal para pruebas de comportamiento y pruebas más complejas.

## [GoTestTools](https://github.com/gotesttools/gotest.tools)
Esta biblioteca proporciona herramientas y utilidades adicionales para mejorar la experiencia de las pruebas en Go. Incluye funciones útiles para el manejo de la salida de las pruebas y la manipulación de datos de prueba.

**Características clave:**
- Funciones adicionales para la manipulación de pruebas y errores.
- Mejor manejo de las salidas de prueba.
- Soporte para escenarios de pruebas avanzadas.

## [TestContainers-Go](https://github.com/testcontainers/testcontainers-go)
Si se necesita hacer pruebas de integración que dependan de contenedores, como bases de datos o servicios externos, TestContainers-Go permite lanzar contenedores Docker desde las pruebas de Go. Esto es útil cuando se necesita un entorno aislado y reproducible para pruebas de integración.

**Características clave:**
- Crear y gestionar contenedores Docker en las pruebas.
- Pruebas de integración de servicios como bases de datos, cachés, etc.
- Aislado y controlado para pruebas más confiables.

# Conclusión

Testify es una de las bibliotecas más populares y poderosas para pruebas unitarias en Go, especialmente útil cuando el enfoque principal es probar la lógica de negocio. Proporciona herramientas como aserciones y mocks, que simplifican la escritura de pruebas claras, comprensibles y efectivas. Dado que Testify se integra bien con el sistema de pruebas estándar de Go, aprovechar ambas herramientas, la biblioteca de pruebas integrada en Go y Testify, será una excelente opción para garantizar la calidad y fiabilidad del código en mi proyecto.