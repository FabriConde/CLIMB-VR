# [Gestor de tareas](https://github.com/FabriConde/CLIMB-VR/issues/12)

Para seleccionar el gestor de tareas más adecuado para el proyecto, se sigue un proceso estructurado que incluye la definición de criterios y la evaluación de diferentes herramientas en base a esos criterios.

## Criterios de evaluación 

- **Estándares y conformidad**: Seguir los estándares establecidos para el lenguaje de programación o el área específica del proyecto.
- **Mejores prácticas**: Investigar y aplicar las mejores prácticas recomendadas para el tipo de herramienta que se necesita.
- **Deuda técnica**: Considerar cómo la elección de una herramienta puede afectar la deuda técnica del proyecto a largo plazo.
- **Seguridad**: Evaluar las características de seguridad de la herramienta y cómo puede ayudar a proteger el proyecto contra vulnerabilidades.
- **Rendimiento**: Analizar las métricas de rendimiento de la herramienta y cómo puede afectar la eficiencia del proyecto.
- **Mantenimiento y soporte**: La herramienta debe recibir actualizaciones regulares y tener un buen soporte de la comunidad o del equipo de desarrollo.

## Comparación de herramientas

### Estándares: Seguir los estándares establecidos para el lenguaje de programación o el área específica del proyecto.

**Make**: Sigue los estándares de la industria para la automatización de tareas en proyectos de software. Es ampliamente utilizado en proyectos de código abierto y comerciales, lo que garantiza su compatibilidad y fiabilidad.

**Mage**: Compatible con los estándares de Go, pero no es tan ampliamente adoptado como Make. Mage es más específico para proyectos en Go, lo que puede limitar su aplicabilidad en proyectos que utilizan múltiples lenguajes.

**Task**: Compatible con los estándares de Go, similar a Mage. Task es más específico para proyectos en Go, pero ofrece una sintaxis más moderna y flexible.

### Mejores prácticas: Investigar y aplicar las mejores prácticas recomendadas para el tipo de herramienta que se necesita.

**Make**: Facilita la implementación de mejores prácticas en la automatización de tareas, como la separación de tareas en objetivos específicos y la reutilización de reglas. Por ejemplo, un Makefile bien estructurado puede tener objetivos como `build`, `test`, `clean`, y `deploy`, cada uno con reglas claras y reutilizables.

**Mage**: También facilita la implementación de mejores prácticas, pero su sintaxis y estructura están más alineadas con Go. Esto puede ser una ventaja en proyectos exclusivamente en Go, pero menos flexible en proyectos con múltiples lenguajes.

**Task**: Facilita la implementación de mejores prácticas en Go, con una sintaxis moderna que permite una estructura clara y reutilizable de las tareas.

### Deuda técnica: Considerar cómo la elección de una herramienta puede afectar la deuda técnica del proyecto a largo plazo.

**Make**: Para evitar deuda técnica, es crucial documentar cada objetivo y regla, utilizar variables para evitar duplicación de código, y mantener los Makefiles organizados y fáciles de leer. Por ejemplo, usar variables como BINARY_NAME y SOURCE_DIR ayuda a mantener el Makefile limpio y fácil de mantener.

**Mage**: La claridad y el mantenimiento de Go ayudan a reducir la deuda técnica, pero la falta de una comunidad tan grande como la de Make puede dificultar encontrar soluciones a problemas específicos.

**Task**: La claridad y el mantenimiento de Go ayudan a reducir la deuda técnica, y su sintaxis moderna puede facilitar la escritura y mantenimiento de tareas.

### Seguridad: Evaluar las características de seguridad de la herramienta y cómo puede ayudar a proteger el proyecto contra vulnerabilidades.

**Make**: Make se considera una herramienta generalmente segura, pero su nivel de seguridad real depende de la forma en que se escriban los Makefiles. Es crucial seguir las mejores prácticas de seguridad al crear Makefiles para evitar vulnerabilidades.

**Mage**: Al estar escrito en Go, un lenguaje conocido por su seguridad, Mage también ofrece un alto nivel de seguridad. Sin embargo, la seguridad en Mage depende en gran medida de cómo se escriben los scripts de Mage.

**Task**: Al estar escrito en Go, Task también ofrece un alto nivel de seguridad. La seguridad en Task depende en gran medida de cómo se escriben los scripts de Task.

### Mantenimiento y soporte: La herramienta debe recibir actualizaciones regulares y tener un buen soporte de la comunidad o del equipo de desarrollo.

**Make**: Recibe actualizaciones ocasionales y tiene una comunidad activa.

**Mage**: No recibe muchas actualizaciones, la última fue de hace un año, la comunidad de Mage es menos activa que la de Make, lo que puede dificultar el acceso a soporte y recursos.

**Task**: Recibe actualizaciones regulares y tiene una comunidad activa, aunque no tan grande como la de Make.

## Conclusión

No conozco ningún gestor de tareas, pero mi opción elegida es `Make` ya que me resulta particularmente interesante y deseo aprender a utilizarlo. A pesar de que su sintaxis puede parecer complicada para algunos, estoy motivado por el desafío y confío en que su capacidad para manejar las tareas será suficiente para mis necesidades. Por tanto, `Make`se presenta como una herramienta robusta y ampliamente utilizada, ideal para aprender los fundamentos de la automatización de tareas.