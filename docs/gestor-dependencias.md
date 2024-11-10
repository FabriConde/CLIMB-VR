# [Gestor de dependencias](https://github.com/FabriConde/CLIMB-VR/issues/13)
Go Modules es el sistema estándar para la gestión de dependencias en Go, introducido en la versión 1.11 en 2018. Este sistema se ha consolidado como el enfoque recomendado para proyectos en Go.

Para este proyecto no es necesario un gestor de dependencias ya que voy a usar el paquete REGEX de Go y como a diferencia de otros lenguajes de programación, Go no utiliza un gestor de dependencias externo.
Go Modules se integra directamente en el lenguaje, proporcionando una solución nativa y eficiente para la gestión de dependencias.

Go Modules funciona a través de dos archivos clave:

go.mod: Este archivo define las rutas de los módulos, actuando como un manifiesto del proyecto. Contiene información crucial como el nombre del módulo y todas sus dependencias.
go.sum: Este archivo gestiona las versiones de las dependencias y garantiza su integridad. Asegura que las versiones utilizadas sean consistentes y confiables.