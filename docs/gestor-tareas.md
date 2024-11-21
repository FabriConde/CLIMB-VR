# [Gestor de tareas](https://github.com/FabriConde/CLIMB-VR/issues/12)

Para seleccionar el gestor de tareas más adecuado para el proyecto, se sigue un proceso estructurado que incluye la definición de criterios y la evaluación de diferentes herramientas en base a esos criterios.

## Criterios de evaluación

- **Configuración y mantenimiento**: La herramienta debe requerir una configuración mínima para empezar a usarla y un mantenimiento continuo reducido.
- **Claridad de la sintaxis**: La sintaxis debe ser clara y concisa.
- **Funcionalidades básicas**: Debe permitir ejecutar tareas comunes como compilación, pruebas, limpieza, etc.
- **Actualizaciones**: La herramienta debe recibir actualizaciones regulares para asegurar mejoras y corrección de errores.
- **Seguridad**: La herramienta debe ser segura y no introducir vulnerabilidades en el proyecto.
- **Prestaciones**: La herramienta debe ser eficiente y no afectar negativamente el rendimiento del proyecto.

## Evaluación de herramientas

### Make

- **Configuración y mantenimiento**: Moderada, requiere escribir Makefiles. Aunque no es extremadamente complejo, puede llevar tiempo configurarlo correctamente y requiere un mantenimiento moderado, especialmente si los Makefiles se vuelven complejos.
- **Claridad de la sintaxis**: Moderada, la sintaxis de Make puede ser compleja y difícil de leer, especialmente para usuarios nuevos.
- **Funcionalidades básicas**: Sí, permite ejecutar tareas comunes como compilación, pruebas y limpieza.
- **Actualizaciones**: Recibe actualizaciones ocasionales.
- **Seguridad**: Generalmente segura, pero depende de cómo se escriban los Makefiles.
- **Prestaciones**: Eficiente, pero puede ser menos flexible que otras herramientas.

### Mage

- **Configuración y mantenimiento**: Alta, se escribe en Go, lo que facilita la configuración para desarrolladores familiarizados con el lenguaje y requiere poco mantenimiento, especialmente si ya se tiene experiencia con Go.
- **Claridad de la sintaxis**: Alta, se usa Go, un lenguaje conocido por su claridad y simplicidad.
- **Funcionalidades básicas**: Sí, permite ejecutar tareas comunes como compilación, pruebas y limpieza.
- **Actualizaciones**: Recibe actualizaciones regulares.
- **Seguridad**: Alta, al estar escrito en Go, un lenguaje conocido por su seguridad.
- **Prestaciones**: Muy eficiente y flexible.

### Sage

- **Configuración y mantenimiento**: Alta, se escribe en Python, lo que facilita la configuración para desarrolladores familiarizados con el lenguaje y requiere poco mantenimiento.
- **Claridad de la sintaxis**: Alta, Python es conocido por su claridad y simplicidad.
- **Funcionalidades básicas**: Sí, permite ejecutar tareas comunes como compilación, pruebas y limpieza.
- **Actualizaciones**: Recibe actualizaciones regulares.
- **Seguridad**: Alta, al estar escrito en Python, un lenguaje conocido por su seguridad.
- **Prestaciones**: Muy eficiente y flexible.

## Conclusión

No conozco ningún gestor de tareas, pero mi opción elegida es `Make` ya que me resulta particularmente interesante y deseo aprender a utilizarlo. A pesar de que su sintaxis puede parecer complicada para algunos, estoy motivado por el desafío y confío en que su capacidad para manejar las tareas será suficiente para mis necesidades. Además, Make cumple con todos los criterios de evaluación mencionados, lo que lo convierte en una opción viable para comenzar.