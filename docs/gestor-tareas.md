# [Gestor de tareas](https://github.com/FabriConde/CLIMB-VR/issues/12)

## Criterios de Evaluación

- **Compatibilidad con Go**: La capacidad de la herramienta para integrarse y ejecutar comandos de Go.
- **Facilidad de Configuración**: La simplicidad y claridad del proceso de configuración inicial y mantenimiento.
- **Funcionalidad**: La capacidad de la herramienta para manejar tareas comunes y complejas.
- **Adaptabilidad**: La flexibilidad de la herramienta para adaptarse a diferentes proyectos y entornos.

## Evaluación de Herramientas

### Make

- **Compatibilidad con Go**: Alta. Make es compatible con Go y puede ejecutar comandos de Go como `go build`, `go test` y `go vet` directamente desde el Makefile.
- **Facilidad de Configuración**: Moderada. Requiere escribir un Makefile, pero la sintaxis es bien documentada y hay muchos ejemplos disponibles.
- **Funcionalidad**: Permite definir tareas de comunes y complejas.
- **Adaptabilidad**: Alta. Make es muy flexible y puede adaptarse a una amplia variedad de proyectos y entornos.

**Ventajas**:
- Muy flexible y ampliamente utilizado.
- Gran cantidad de ejemplos y documentación disponibles.
- Amplia comunidad de usuarios y soporte.
- Puede manejar tareas complejas y personalizadas.

**Desventajas**:
- Sintaxis puede ser confusa para principiantes.
- Requiere mantenimiento manual del Makefile.

### Mage

- **Compatibilidad con Go**: Alta. Diseñada específicamente para Go, Mage permite definir tareas usando el propio lenguaje Go, lo que puede ser más intuitivo para los desarrolladores de Go.
- **Facilidad de Configuración**: Alta. Usa Go para definir tareas, eliminando la necesidad de aprender una nueva sintaxis.
- **Funcionalidad**: Su uso de Go para definir tareas lo hace muy intuitivo para desarrolladores Go y puede manejar tareas comunes y complejas.
- **Adaptabilidad**: Alta. Mage permite aprovechar todas las capacidades del lenguaje Go, lo que lo hace muy adaptable a diferentes necesidades.

**Ventajas**:
- Usa Go para definir tareas, lo que puede ser más intuitivo para desarrolladores Go.
- No requiere archivos adicionales como Makefile.
- Puede aprovechar todas las capacidades del lenguaje Go para definir tareas.

**Desventajas**:
- Menos popular que Make, por lo que puede haber menos ejemplos y soporte comunitario.

### Sage

- **Compatibilidad con Go**: Alta. Compatible con Go y puede ejecutar comandos de Go.
- **Facilidad de Configuración**: Moderada. Similar a Make, pero con algunas mejoras en la sintaxis.
- **Funcionalidad**: Permite definir tareas comunes y complejas. Similar a Make, pero con algunas mejoras en la sintaxis que pueden facilitar su uso.
- **Adaptabilidad**: Moderada. Aunque es similar a Make, su menor popularidad puede limitar su adaptabilidad en algunos entornos.

**Ventajas**:
- Similar a Make, pero con algunas mejoras en la sintaxis.
- Puede manejar tareas complejas y personalizadas.

**Desventajas**:
- Menos popular y menos soporte disponible.
- Sintaxis y uso menos conocidos.

### Gotaskr

- **Compatibilidad con Go**: Alta. Compatible con Go y puede ejecutar comandos de Go.
- **Facilidad de Configuración**: Alta. Usa YAML para definir tareas, lo que puede ser más fácil de leer y escribir.
- **Funcionalidad**: Permite definir tareas comunes y complejas. Su uso de YAML para definir tareas lo hace muy legible y fácil de configurar.
- **Adaptabilidad**: Alta. La legibilidad y facilidad de configuración de YAML hacen que Gotaskr sea muy adaptable a diferentes proyectos y entornos.

**Ventajas**:
- Usa YAML para definir tareas, lo que puede ser más fácil de leer y escribir.
- Buena integración con sistemas CI/CD.
- Puede manejar tareas complejas y personalizadas.

**Desventajas**:
- Menos popular que Make, por lo que puede haber menos ejemplos y soporte comunitario.

## Justificación de la Elección de Make

- **Compatibilidad con Go**: Make es altamente compatible con Go y puede ejecutar comandos de Go directamente desde el Makefile, esto permite integrar fácilmente las tareas.
- **Facilidad de Configuración**: Aunque requiere escribir un Makefile, la sintaxis de Make es bien documentada y hay muchos ejemplos disponibles. Esto facilita la configuración inicial y el mantenimiento a largo plazo.
- **Funcionalidad adecuada**: Make ofrece la funcionalidad necesaria para manejar tareas comunes y complejas.
- **Necesidades de automatización a largo plazo**: Make es una herramienta robusta y ampliamente utilizada que puede manejar la automatización de tareas a largo plazo.
- **Adaptabilidad**: Make es muy flexible y puede adaptarse a una amplia variedad de proyectos y entornos, lo que lo hace una opción robusta para diferentes necesidades.
