# [Gestor de tareas](https://github.com/FabriConde/CLIMB-VR/issues/12)

Se van a comparar 7 herramientas en base a unos criterios de evaluación objetivos que me ayudarán a elegir el gestor de tareas más adecuado para el proyecto.

## Criterios de evaluación

- **Estándares y conformidad**: Seguir los estándares establecidos para el lenguaje de programación o el área específica del proyecto.
- **Deuda técnica**: Considerar cómo la elección de una herramienta puede afectar la deuda técnica del proyecto a largo plazo.
- **Seguridad**: Evaluar las características de seguridad de la herramienta y cómo puede ayudar a proteger el proyecto contra vulnerabilidades.
- **Actualizaciones**: Saber si la herramienta recibe actualizaciones.

### Make
- **Estándares y conformidad**:
Make sigue los estándares tradicionales de la automatización de tareas en proyectos de software, especialmente en sistemas UNIX y proyectos en lenguajes como C y C++. Es la herramienta de referencia en muchos ecosistemas, especialmente en proyectos que no dependen de un solo lenguaje.

- **Deuda técnica**
Make es muy flexible, pero su uso incorrecto o la falta de documentación puede conducir a una deuda técnica significativa. Esto es especialmente cierto en proyectos grandes y complejos, donde el mantenimiento de Makefiles puede volverse tedioso si no se tiene cuidado.

- **Seguridad**
La seguridad en Make depende principalmente de las dependencias de los comandos ejecutados. Como Make ejecuta comandos arbitrarios, los riesgos de seguridad surgen si no se validan adecuadamente los comandos o las rutas de archivos.

- **Actualizaciones**
Make es una herramienta consolidada que recibe pocas actualizaciones, la ultima fue hace más de un año.

### [Mage](https://github.com/magefile/mage)

- **Estándares y conformidad**:
Mage sigue las mejores prácticas de Go, utilizando el propio lenguaje Go para definir las tareas. Esto lo hace particularmente adecuado para proyectos Go, pero no necesariamente para otros lenguajes. No es una herramienta "estándar" para otros lenguajes o entornos de construcción como Make.

- **Deuda técnica**
Mage puede reducir la deuda técnica a largo plazo, especialmente en proyectos Go, al aprovechar el poder del lenguaje Go. Sin embargo, si el proyecto no es de Go, puede generar cierta deuda técnica por depender de Go, que puede no ser ideal en entornos heterogéneos.

- **Seguridad**
Mage ejecuta tareas definidas en Go, lo que puede aumentar la seguridad si se siguen las mejores prácticas del lenguaje (por ejemplo, validación de entradas y control de excepciones). Sin embargo, también depende de cómo se gestionen las dependencias en Go, ya que las vulnerabilidades de las bibliotecas pueden introducir riesgos.

- **Actualizaciones**
Mage cuenta con pocas actualizaciones, la última fue hace un año.

### [Task](https://github.com/go-task/task)

- **Estándares y conformidad**:
Task sigue una sintaxis declarativa basada en YAML, lo que lo hace fácil de entender y usar. Aunque es bastante moderno, no sigue necesariamente estándares establecidos en otros lenguajes o ecosistemas, pero se adhiere a buenas prácticas de automatización moderna.

- **Deuda técnica**
Task puede ayudar a reducir la deuda técnica al permitir una configuración sencilla y modular en YAML. Debido a su simplicidad, es menos probable que la deuda técnica se acumule en comparación con herramientas más complejas. Sin embargo, en proyectos grandes, la gestión de un archivo YAML podría volverse menos eficiente que otras soluciones más complejas.

- **Seguridad**
Al igual que Make, Task depende de la seguridad de los comandos que ejecuta. No ofrece mecanismos de seguridad integrados, por lo que es crucial revisar las dependencias y validar las entradas antes de ejecutar comandos.

- **Actualizaciones**
Task está ganando popularidad rápidamente y recibe actualizaciones regulares enfocadas en mejorar la experiencia del usuario y agregar nuevas funcionalidades, la última actualización fue hace una semana.

### [Gotaskr](https://github.com/Roemer/gotaskr)

- **Estándares y conformidad**:
Gotaskr sigue los estándares de Go, ya que está escrito en este lenguaje. Esto lo hace especialmente adecuado para proyectos Go, pero no necesariamente es una herramienta "estándar" en otros ecosistemas.

- **Deuda técnica**
Gotaskr puede ayudar a evitar la deuda técnica en proyectos Go, ya que se basa en el mismo lenguaje del proyecto. Sin embargo, si el proyecto no utiliza Go, es posible que se introduzca deuda técnica por la necesidad de integrar una herramienta de Go en un entorno que no lo utiliza.

- **Seguridad**
La seguridad en Gotaskr dependerá de cómo se implementen las tareas. Al estar basado en Go, se puede aprovechar las prácticas de seguridad del lenguaje, pero no ofrece características de seguridad adicionales como las validaciones integradas en algunas otras herramientas.

- **Actualizaciones**
Gotaskr recibe actualizaciones regulares, la última fue hace un mes.

### [Just](https://github.com/casey/just)

- **Estándares y conformidad**:
Just sigue las mejores prácticas de la automatización de tareas modernas. La sintaxis declarativa y la facilidad de uso lo convierten en una opción popular para proyectos que no necesitan la complejidad de herramientas más antiguas como Make.

- **Deuda técnica**
Just tiene una deuda técnica relativamente baja debido a su sencillez. Las tareas son fáciles de definir y entender, y la herramienta no requiere una configuración compleja, lo que ayuda a mantener bajo control la deuda técnica.

- **Seguridad**
La seguridad depende de los comandos que se ejecutan. Just no ofrece características de seguridad adicionales, pero la simplicidad de la herramienta reduce los riesgos de malentendidos que podrían ocurrir en sistemas más complejos.

- **Actualizaciones**
Just recibe actualizaciones regulares, la última actualización fue hace dos días.

### [Bazel](https://github.com/bazelbuild/buildtools)

- **Estándares y conformidad**:
Bazel sigue los estándares más estrictos y modernos en cuanto a herramientas de construcción. Está basado en un enfoque de construcción declarativo y es usado en grandes empresas tecnológicas, como Google, lo que le da una alta confiabilidad.

- **Deuda técnica**
Bazel es muy efectivo para manejar grandes proyectos, pero su complejidad puede generar deuda técnica en proyectos pequeños o medianos. La configuración inicial puede ser costosa en términos de tiempo, pero a largo plazo, ayuda a mantener el proyecto bien organizado.

- **Seguridad**
Bazel tiene algunas características que mejoran la seguridad, como el aislamiento de procesos y la capacidad de gestionar dependencias de forma controlada. Sin embargo, la seguridad también depende de las prácticas y configuraciones implementadas en el proyecto.

- **Actualizaciones**
Bazel recibe actualizaciones regulares respaldadas por su uso en grandes empresas tecnológicas, como Google, que garantizan su estabilidad y evolución constante. La ultima actualización fue hace 2 meses.

### [Sage](https://github.com/einride/sage)

- **Estándares y conformidad**
Sage sigue los estándares modernos de la automatización de tareas al estar escrito en Go, lo que facilita la integración con proyectos que también están escritos en Go. No sigue necesariamente los estándares tradicionales de herramientas de construcción como Make o Bazel, pero al estar basado en Go, se ajusta a las buenas prácticas de este lenguaje.

- **Deuda técnica**
Sage es una herramienta bastante sencilla, por lo que tiene una baja probabilidad de generar deuda técnica, siempre y cuando las tareas sean bien definidas. Sin embargo, si se usa incorrectamente en proyectos complejos, la falta de reglas explícitas podría derivar en tareas difíciles de gestionar a largo plazo. Además, la integración con otros lenguajes no Go podría requerir más esfuerzo y generar deuda técnica si no se mantiene adecuadamente.

- **Seguridad**
Sage, al igual que otras herramientas de automatización de tareas, ejecuta comandos directamente en el sistema. Esto significa que la seguridad depende en gran medida de cómo se escriban y gestionen esos comandos. No proporciona mecanismos de seguridad adicionales como aislamiento de procesos, por lo que es fundamental tener cuidado al usar Sage con fuentes externas o no verificadas.

- **Actualizaciones**
Sage cuenta con actualizaciones regulares, la última fue hace dos días.

## Conclusión

- Sage es una opción sólida para proyectos Go debido a su simplicidad y rendimiento, pero carece del mismo nivel de soporte y madurez que herramientas más establecidas como Make o Bazel. Es excelente para tareas simples en proyectos Go, pero puede ser menos adecuado para proyectos más grandes o para entornos que no utilicen Go.
- Make sigue siendo el estándar en muchos entornos, especialmente para proyectos grandes y complejos, pero su sintaxis y flexibilidad pueden resultar difíciles de mantener a largo plazo.
- Bazel es ideal para proyectos grandes y complejos, especialmente cuando se requiere un alto nivel de optimización y paralelización, pero es más complicado de implementar y mantener en proyectos pequeños.
- Mage, Task y Just son herramientas modernas y fáciles de usar para proyectos más pequeños o medianos, pero no tienen el mismo soporte o flexibilidad que Bazel o Make para proyectos grandes.

No conozco ningún gestor de tareas, pero mi opción elegida es `Make` ya que me resulta particularmente interesante y deseo aprender a utilizarlo. A pesar de que su sintaxis puede parecer complicada para algunos, estoy motivado por el desafío y confío en que su capacidad para manejar las tareas será suficiente para mis necesidades. Por tanto, Make es una herramienta robusta y ampliamente utilizada, ideal para aprender los fundamentos de la automatización de tareas.