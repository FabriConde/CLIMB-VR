# CLIMB-VR
## Descripción del problema
Como escalador me hallo con el problema de la falta de información que hay de los sectores de escalada en Vélez-Rubio.
Los sectores de escalada tienen un número de vías y este depende de cada sector, cada vía tiene su nombre y grado que ronda desde V (fácil) hasta 9c (muy difícil).
El problema es que los escaladores de las zonas son los que tienen las direcciones de los sectores de escalada, el número de vías y el grado de cada vía, pero no tienen nigún sitio donde compartirla.
Tener esta información es necesaria para prevenir accidentes.

> Descripción detallada del problema
La recopilación información será mediante textos existentes: como guías de escalada, publicaciones en blogs, foros o grupos de WhatsApp.

La nomenclatura es siempre la misma:
- Ubicación: Puede ser una ubicación de google o un texto que explica la localización del sector y como llegar hasta él. 
- Sector: Nombre del sector.
- Vías: Número de vías.
- Nombre, grado y metros: Nombre de la vía, su grado y los metros de la vía.
  
Ejemplo:
- Ubicación: Ubicación de google
- Sector Avispa
- Numero de vías: 2
- Avispa 6b
- El barrenazo, 6b+, 25 metros

Para que la lógica de negocio no sea trivial se prodría calcular la dificultad promedio de las rutas en una zona o filtrar las rutas que son aptas para cierto nivel de experiencia. Para que la información procesada sea valiosa para un escalador que busca rutas que se ajusten a su nivel o que quiere explorar nuevas zonas.

## Información adicional
### Documentación
[Documentación](https://github.com/FabriConde/CLIMB-VR/tree/main/docs)
### Configuración del repositorio

### Historias de usuario
#### [HU001] Obtener información general de un sector de escalada
#### [HU002] Ver vías de un sector ordenadas por dificultad
Para más información: [Historias de usuario](https://github.com/FabriConde/IV-2024-2025/blob/objetivo_1/docs/historias-usuario.md)

### Milestones
#### [M0] Creación del modelo de datos
#### [M1] Módulo de la lógica de negocio
#### [M2] Componente de persistencia de datos
#### [M1] API de consulta de datos
Para más información: [Milestones](https://github.com/FabriConde/IV-2024-2025/blob/objetivo_1/docs/milestones.md)

### User journey
Para más información: [Usuarios](https://github.com/FabriConde/IV-2024-2025/blob/objetivo_1/docs/user%20journey.md)

## Lenguaje de programación
El lenguaje de programación para este proyecto es: Go
