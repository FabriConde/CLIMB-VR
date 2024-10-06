# [[M0] Creación del modelo de datos](https://github.com/FabriConde/IV-2024-2025/milestone/1)
- **Objetivo**: Crear el modelo de datos (entidades y atributos) y definir las relaciones entre las entidades.
- **Entregable**:   
  - Entidades:
    - Sector:
      - Atributos:
        - Nombre (String): Identificador único del sector.
        - Ubicación (Objeto Ubicación): Información sobre la ubicación del sector.
        - Número de vías (Integer): Cantidad total de vías en el sector.
    - Ubicación:
      - Atributos:
        - Coordenadas GPS (String): Latitud y longitud.
        - Descripción (String): Explicación detallada de cómo llegar al sector.
    - Vía:
      - Atributos:
        - Nombre (String): Nombre de la vía.
        - Grado de dificultad (String): Escala de dificultad (YDS, Francés, etc.).
        - Metros (Integer): Longitud de la vía en metros.
        - Descripción (String, opcional)
  - Relaciones:
    - Un sector tiene muchas vías (1:N).
    - Un sector tiene una ubicación (1:1).
- **Viabilidad**: Las relaciones entre las entidades deben ser implementadas de forma que reflejen con precisión la realidad del problema.

# [[M1] Módulo de la lógica de negocio](https://github.com/FabriConde/IV-2024-2025/milestone/3)
- **Objetivo**: Crear la lógica para extraer la información de los textos y poblar el modelo de datos definido en el M0.
- **Entregable**: Una clase o módulo que recibe un texto como entrada y devuelve una instancia del modelo de datos con la información extraída.
- **Viabilidad**: Tests unitarios que demuestren que el código produce los resultados esperados para diferentes textos.

# [[M2] Componente de persistencia de datos](https://github.com/FabriConde/IV-2024-2025/milestone/4)
- **Objetivo**: Implementar un mecanismo básico para almacenar las instancias del modelo de datos y poder recuperarlas posteriormente.
- **Entregable**: Un controlador simple que permita guardar y cargar la información extraída.
- **Viabilidad**:
    - La correcta integración con el módulo de lógica de negocio
    - La capacidad de almacenar y recuperar información sin pérdida ni corrupción
    - Validación mediante pruebas de integración.

# [[M3] API de consulta de datos](https://github.com/FabriConde/IV-2024-2025/milestone/2)
- **Objetivo**: Implementar las funciones básicas para que los usuarios puedan buscar sectores y filtrar vías.
- **Entregable**: Una lógica funcional que permita buscar sectores por nombre o ubicación, mostrar la información básica del sector y filtrar las vías por grado.
- **Viabilidad**:
    - Eficiencia en la ejecución de consultas.
    - Correcta gestión de errores.
    - Ejemplos de uso que demuestren su funcionalidad.
