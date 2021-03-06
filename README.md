# Hackaton BBVA - Transformación Adquirente

```

Bienvenidos a la version 0.1.0 de el proyecto HACKBBVA donde ofrecemos una API 
para comunicacion entre terminal, negocio y seguridad, 
con el fin de detectar, analizar y almacenar todos los errores a los cuales BBVA se enfrenta 
durante el proceso de transacciones en la terminal y negocio.

```

## Uso

~~~
para el uso de esta api, se requiere un token que usted tendrá que solicitar con una cuenta. 
una vez ese token sea obtenido, su uso es enviando  un parametro Authorization y colocando el token. 
~~~
accede a la api desde el siguiente enlace: [https://hackbbva.herokuapp.com/](https://hackbbva.herokuapp.com/)
  
---

# Acceder a la terminal

para acceder a datos de la terminal, como nombre, marca, modelo, banco emisor, etc, accedemos al endpoint **GET** **/terminal** para obtener todas las terminales pero también contamos con el endpoint **/terminal/{id}** para obtener una terminal en específico.
De esta forma, obtenemos a que terminal notificar la información de error, y la transacción que se intentó realizar, en caso de que ocurriera un fallo.

En el flujo natural, si no sucede ningún error, entonces se continua con la operación que se está procesando. En este caso, los pagos por terminal.

---

# Obtener info del banco

Para obtener información del banco perteneciente a la terminal, basta con acceder al endpoint **GET** **/banco/{id}** este endpoint ya está preparado para relacionarse con la info de la terminal, permitiendo saber al banco que terminal de que banco, tuvo problemas, y a cuál notificar el error que se generara o en su defecto, la operación exitosa siguiendo el flujo natural de la terminal.

---

# Obtener información del Cliente ( Empresa )

Para obtener información del cliente que tiene cierta terminal y está con cierto banco, accedemos al endpoint **GET** **/cliente/{id}**. De esta manera, obtenemos información del negocio, de su terminal específica, y demás aspectos de la transacción, para la cual cuenta con sus ciertos endpoints.

---

# Obtener información de las transacciones

Obtener información de la transacción que está siendo enviada entre terminal/negocio es bastante sencillo, enviando una petición al siguiente endpoint **GET** **/transaccion/{id}**. 
De esta forma, manejamos toda la info de la transacción, permitiendonos saber si hay algo que en este proceso esté equivocado, a fin, de obtener el analisis concreto y los errores que deriven de los datos que están pasando a traves de la terminal.

---

# Obtener información general de cada categoría

para obtener información general de cada categoría ( Terminales, Transacciones, Clientes, Bancos ), entonces accedemos a su endpoint plural específico.
Por ejemplo: 
 **GET** **/bancos** para toda la información de todos los bancos.

 **GET** **/terminales** para la información de todas las terminales.

 **GET** **/transacciones** para toda la info de todas las transacciones realizadas.

 **GET** **/clientes** para obtener todos los clientes que están afiliados al banco.

---

# INFORMACIÓN EXTRA - RELEVANTE



Cada endpoint cuenta con sus metodos de acceso **POST, PUT, DELETE**, es decir, cada categoría, podrá **añadir** nueva información y almacenarla, así como **actualizarla** y **eliminarla** de ser necesario.

---

~~~
Te invitamos a conocer mas acerca de nosotros, los que proponemos esta arquitectura de backend para el manejo de errores de las terminales del banco.
~~~

[Ver avances de la arquitectura](https://drive.google.com/drive/u/0/folders/1Cu_YYYWX0FysRqRC6_x1pkequwAnVNPe)

[Ver tiktok de los colaboradores](https://vm.tiktok.com/ZM8AuDD2J/)

[LinkedIn de Roberto De Santiago FullStack Developer](https://www.linkedin.com/in/roberto-de-santiago-nájera-a7258820b/)

[LinkedIn de Carlos Vazquez UI/UX Designer, Java y web Devoloper](https://www.linkedin.com/in/carlos-v%C3%A1zquez-538271161/)

[LinkedIn de Felix Martínez FullStack Developer](https://www.linkedin.com/in/felix-martinez/)