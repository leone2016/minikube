## HANDS-ON

* en esta practica se va crear un backend en Golang y un front, para lo cual el usuario tendrá acceso a front atraves de un nodePort y este (Front) accediera al backend por un ClusterIps

### Backend
* crear api base [doc](https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj)
* `docker pull golang` | descarga la imagen de goLang
* ingresar a backend/src
* `docker run --rm -dti -v $PWD/:/go --net host --name golang-name golang bash` | ejecuta golang en la terminar interna del contenedor
* `docker run --rm -dti -v $PWD/:/go -p 9090:9090 --name golang-name golang bash` | Solucion MACoS
    * **_-dti_**: lo maneja desde el fondo y de forma iterativa
    * **_-v $PWD/:/go_**: comparte la carpeta en la que se edita el código **:** y sé la pasa a una carpeta dentro del contenedor llamado **go**
    * **_--net host_**: se une a la red de host, esto es para no exponer el puerto 9090 
    * **_ --name golang-name_**: se crea el contenedor con el nombre **golang-name**
    * **_ golang_**: utiliza la imagen de golang
    * **_bash_**: utiliza la consola del contenedor
* `docker exec -ti golang-name bash`| se puede pasar el nombre (NAMES) del contenedor o el CONTAINER ID
    * `:/go# cat main.go` se tiene que visualizar el código session9/backend/src/main.go

    
    
    
