# IV-2024-2025
# Repositorio para la asignatura Infraestructura Virtual
## OBJETIVO 0
### CLIENTE

Como escalador me hallo con el problema de la falta de información que hay de los sectores de escalada en Vélez-Rubio.
Los sectores de escalada tienen un número de vías y este depende de cada sector, cada vía tiene su nombre y grado que ronda desde V (fácil) hasta 9c (muy difícil).
El problema es que los escaladores de las zonas son los que tienen las direcciones de los sectores de escalada, el número de vías y el grado de cada vía, pero no tienen nigún sitio donde compartirla.
Tener esta información es necesaria para prevenir accidentes.

> Especificación de la solución del problema

La lógica de negocia será capaz de extraer mediante un texto la información que se le ha proporcionado, se extraera ubicación, sector, número de vías y nombre y grado de cada vía.
La recopilación información será mediantes de textos existentes: como guías de escalada en formato físico, publicaciones en blogs o foros, incluso documentos históricos o grupos de WhatsApp.

La nomenclatura es siempre la misma:
- Ubicación: Puede ser una ubicación de google o un texto que explica la localización del sector y como llegar hasta él. 
- Sector: Nombre del sector
- Vías: Número de vías
- Nombre y grado: Nombre de la vía y su grado
  
Ejemplo:

- Ubicación: Ubicación de google
- Sector Avispa
- Numero de vías: 2
- Avispa 6b
- El barrenazo 6b+

Para que la lógica de negocio no sea trivial se prodría calcular la dificultad promedio de las rutas en una zona o filtrar las rutas que son aptas para cierto nivel de experiencia. Para que la información procesada fuera valiosa para un escalador que busca rutas que se ajusten a su nivel o que quiere explorar nuevas zonas.

### Documentación

[Documentación](https://github.com/FabriConde/IV-2024-2025/tree/objetivo_0-v0.0.1/Documentaci%C3%B3n)
