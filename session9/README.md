## **Seccion 9: Golang, JavaScript y Kubernetes - Aprende a construir aplicaciones reales**

## Tabla de contenidos
- [**Section 04: Orchestrating Collections of Services with Kubernetes**](#seccion-9-golang-javascript-y-kubernetes---aprende-a-construir-aplicaciones-reales)
- [Tabla de contenidos](#tabla-de-contenidos)
  - [60. Introducción](#60-introduccin)
  - [61. Golang: Empieza a escribir tu API](#61-golang-empieza-a-escribir-tu-api)
  - [63. Crea un DockerFile para tu aplicación en golang](#63-crea-un-dockerfile-para-tu-aplicacin-en-golang)
  - [64. Tip: ¿No puedes construir un Dockerfile porque no tienes Docker instalado?](#64-tip-no-puedes-construir-un-dockerfile-porque-no-tienes-docker-instalado)
  - [65. Escribe manifiestos de Kubernetes para desplegar tu aplicación](#65-escribe-manifiestos-de-kubernetes-para-desplegar-tu-aplicacin)
  - [66. Aprender a consumir el servicio que creaste](#66-aprender-a-consumir-el-servicio-que-creaste)
  - [67. Nota: ¿No puedes ver el servicio Backend?](#67-nota-no-puedes-ver-el-servicio-backend)
  - [68. Empieza a escribir el cliente JavaScript que consumirá tu Backend en Go](#68-empieza-a-escribir-el-cliente-javascript-que-consumir-tu-backend-en-go)
  - [69. Despliega una nueva versión de tu BackEnd para resolver errores en el FrontEnd](#69-despliega-una-nueva-versin-de-tu-backend-para-resolver-errores-en-el-frontend)
  - [70. Valida que tu servicio Front esté funcionando como debería](#70-valida-que-tu-servicio-front-est-funcionando-como-debera)
  - [71. Crea los manifiestos de k8s para desplegar tu servicio Front](#71-crea-los-manifiestos-de-k8s-para-desplegar-tu-servicio-front)
  - [72. Crea un Dockerfile para tu aplicación en JavaScript](#72-crea-un-dockerfile-para-tu-aplicacin-en-javascript)
  - [73. Tip: ¿No puedes construir un Dockerfile porque no tienes Docker instalado?](#73-tip-no-puedes-construir-un-dockerfile-porque-no-tienes-docker-instalado)
  - [74. Despliega los servicios y valida su funcionamiento](#74-despliega-los-servicios-y-valida-su-funcionamiento)

## 60. Introducción


> eliminar practicas anteriores

```console 
cd seccion8/servicios
kubectl delete -f nodeport.yaml
kubectl delete -f services.yaml
```

> en esta practica se va crear un backend en Golang y un front, para lo cual el usuario tendrá acceso a front atraves de un nodePort y este (Front) accediera al backend por un ClusterIps
![](printScreen/seccion9-intro.png

**[⬆ volver arriba](#tabla-de-contenidos)**

## 61. Golang: Empieza a escribir tu API

B A C K E N D

```console
mkdir session9
mkdir backend/src
touch main.go

```
* ### paso 1: dentro del directorio backend/src crear un archivo main.go
* ### paso 2: [golang simple rest api](https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj)
* ### paso 3: descargar imagen de golang en docker
```console
docker pull golang
```
* ### paso 4: desde la ruta backend/src, compartir la carpeta y ejecutar en el demonio iteractivo (-dit)

```console 
cd session9/backend/src
docker run --rm -dti -v $PWD/:/go --net host --name golang-name golang bash
docker exec -ti golang-name sh
root@localhost:/go# ls
root@localhost:/go# cat main.go
root@localhost:/go# go run main.go

```
  
| Docker Commands                                                                 | Explanation                                           |
| ------------------------------------------------------------------------------  | ----------------------------------------------------- |
| `docker run --rm -dti -v $PWD/:/go --net host --name golang-name golang bash`   | ejecuta golang en la terminar interna del contenedor  |
| `docker run --rm -dti -v $PWD/:/go -p 9090:9090 --name golang-name golang bash` | Solucion MACoS - FEDORA                               |
| **_-dti_**                                                                      | lo maneja desde el fondo y de forma iterativa         |
| **_-v $PWD/:/go_**                                                              | comparte la carpeta en la que se edita el código **:** y sé la pasa a una carpeta dentro del contenedor llamado **go** |
| **_--net host_**                                                                | se une a la red de host, esto es para no exponer el puerto 9090 |
| **_--name golang-name_**                                                        | se crea el contenedor con el nombre **golang-name**    |
| **_golang_**                                                                    | utiliza la imagen de golang                            |
| **_bash_**                                                                      | utiliza la consola del contenedor                     |

* ### paso 5: ir a un navegador y escribir `localhost:9090` en el caso de existir problemas probar exponiendo los puertos desde docker `-p 9090:9090`
```console
docker stop golang-name

```


**[⬆ volver arriba](#tabla-de-contenidos)**

## 63. Crea un Dockerfile para tu aplicación en Golang
* ### paso 6: crear un archivo Dockerfile, (multistage build) [DOC](https://dev.to/andrioid/slim-docker-images-for-your-go-application-11oo)

```console
cd ../../backend
##touch Dockerfile
docker build -t k8s-hands-on -f Dockerfile .
docker run -d -p 9090:9090 --name k8s-hands-on k8s-hands-on
```

levantar la imagen para validar el perfecto funcionamiento antes de crear el cluster en minikube

| Docker Commands                                                                 | Explanation                                           |
| ------------------------------------------------------------------------------  | ----------------------------------------------------- |
| `docker run -d -p 9090:9090 --name k8s-hands-on k8s-hands-on`                   | ejecutar esto desde session9/backend/src              |
| -d:                                                                             | se va al fonddo los logs                              |
| -p [ puertoOrigen ]:[ puertoDestino ]                                           | expone un puerto del local al contenedor              |
| -name [ nombreAsignar] [nombreImagenUtiliza                                     |_ |
docker run --rm -dti -v $PWD/:/go --net host --name golang-name golang bash
* ### paso 7: eliminar contenedor de prueba creado `docker rm -fv k8s-hands-on`

**[⬆ volver arriba](#tabla-de-contenidos)**

## 64. Tip: ¿No puedes construir un Dockerfile porque no tienes Docker instalado?
¿No puedes construir una imagen local por algún motivo?

No te preocupes, subí la imagen a Docker Hub para que esté disponible a cualquier usuario en Internet. Para usarla, solamente coloca `ricardoandre97/backend-k8s-hands-on:v1` o v2, según corresponda en el spec.image del template del pod.

Kubernetes se encargará de descargar esta imagen para que puedas usarla :)

**[⬆ volver arriba](#tabla-de-contenidos)**

## 65. Escribe manifiestos de Kubernetes para desplegar tu aplicación

* ### paso 8: crear la imagen, si aun no se a creado
  * Why we need `eval $(minikube docker-env)` ? [minikube problem](https://github.com/leone2016/usrv_node/blob/master/seccion-04/minikubeproblem.md)


 #### :stop_sign: Si al ejecutar el siguiente comando da error (CrashLoopBackOff) en el pods de pod/kube-proxy-r445s :rotating_light:
```console
 kubectl get pod,svc -n kube-system
 
 NAME                                   READY   STATUS             RESTARTS   AGE
pod/coredns-74ff55c5b-4x2cc            0/1     Running            3          2d3h
pod/etcd-minikube                      1/1     Running            3          2d3h
pod/kube-apiserver-minikube            1/1     Running            3          2d3h
pod/kube-controller-manager-minikube   1/1     Running            3          2d3h
pod/kube-proxy-r445s                   0/1     CrashLoopBackOff   3          2d3h
pod/kube-scheduler-minikube            1/1     Running            3          2d3h
pod/storage-provisioner                0/1     Error              3          2d3h

NAME               TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
service/kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   2d3h
 ```

### Solucion error pod kubelet

0. Visualizar error
```console
 kubectl -n kube-system logs kube-proxy-r445s 
 
I0527 16:42:25.013647       1 node.go:172] Successfully retrieved node IP: 192.168.49.2
I0527 16:42:25.013716       1 server_others.go:142] kube-proxy node IP is an IPv4 address (192.168.49.2), assume IPv4 operation
W0527 16:42:25.089647       1 server_others.go:578] Unknown proxy mode "", assuming iptables proxy
I0527 16:42:25.089805       1 server_others.go:185] Using iptables Proxier.
I0527 16:42:25.090483       1 server.go:650] Version: v1.20.2
I0527 16:42:25.091805       1 conntrack.go:100] Set sysctl 'net/netfilter/nf_conntrack_max' to 393216
F0527 16:42:25.091836       1 server.go:495] open /proc/sys/net/netfilter/nf_conntrack_max: permission denied

```
1. Delete your local cluster

`minikube delete`

2. en el log capturado **paso 0** el valor máximo que intenta setar es `393216`, agregar el mismo valor en el siguiente comando
```console
sudo sysctl net/netfilter/nf_conntrack_max=393216
cat /proc/sys/net/netfilter/nf_conntrack_max
```
3. Start your local cluster 

`minikube start`

> :warning: continuar si no tienen error en el pod de proxy de kubelet

```console
cd ../../backend
touch backend.yaml
minikube start
minikube status
eval $(minikube docker-env)
docker build -t k8s-hands-on -f Dockerfile .
```

* ### paso 9: crear editar archivo yaml para crear el cluster
  * Usar imagen local: [imagePullPolicy](https://kubernetes.io/docs/concepts/containers/images/#updating-images)
```yaml
## kubectl api-resources permite entender lo que va en el yaml
# DEPLOYMENT
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-k8s-hands-on
  labels:
    app: back # label del deployment
# REPLICASET
spec:
  replicas: 3
  selector:
    matchLabels:
      app: back # label que el replica set va utilizar para encontrar a los pods
  # POD
  template:
    metadata:
      labels:
        app: back # debe ser igual a matchLabels -> app del replicaset
    spec:
      containers:
        - name: back
          image: k8s-hands-on # por defecto docker va traa=tar de desargar la imagen de docker hub
          imagePullPolicy: IfNotPresent #I M P O R T A N T E es necesario agregar IfNotPresent para que no descargue de DOckerHub sino que use la imagen local
---
# SERVICE
apiVersion: v1
kind: Service
metadata:
  name: backend-k8s-hands-on
  labels:
    app: back
spec:
  type: NodePort
  selector:
    app: back # label que utilizar para observar los pods
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9090

```
> creación de cluster

```console
cd session9/backend/
kubectl apply -f backend.yaml
kubectl get deployments.apps backend-k8s-hands-on -o yaml | grep -i Pull
kubectl get pods -l app=backend
```

`kubectl get deployments.apps backend-k8s-hands-on -o yaml | grep -i Pull `una ves ejecutado el deployment, se puede ingresar de esta forma y ver que tipo de policy tiene por defeecto,
en este caso iamgePullPolicy: IfNotPresent


`kubectl get deployments.apps backend-k8s-hands-on -o yaml | grep -i Pull -C 12 `

#### validar funcionamiento de web service en el navegador 

### [usar esta documentacion](https://github.com/leone2016/usrv_node/tree/master/seccion-04#creating-a-nodeport-service)

**[⬆ volver arriba](#tabla-de-contenidos)**

## 66. Aprender a consumir el servicio que creaste

```console
kubectl get svc
minikube ip
```
* Modifica el yaml del servicio backend y colocalo como `type: NodePort`.
* Luego de ello, obtén el nuevo puerto de `NodePort` con `kubectl get svc`.
  ```console
    NAME                   TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
    backend-k8s-hands-on   NodePort    10.102.132.132   <none>        80:31331/TCP   9m3s
    kubernetes             ClusterIP   10.96.0.1        <none>        443/TCP        17m
  ```

* Ahora la IP externa con `kubectl cluster-info` ó `minikube ip`
  ```console
    kubectl cluster-info
    Kubernetes control plane is running at https://192.168.49.2:8443
    KubeDNS is running at https://192.168.49.2:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
    
    To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
  ```
* Coloca la `IP:NodePort` como url en el código JS 

### http://192.168.49.2:31331 :arrow_left:  RESULTADO

**[⬆ volver arriba](#tabla-de-contenidos)**

## 67. Nota: ¿No puedes ver el servicio Backend?

**[⬆ volver arriba](#tabla-de-contenidos)**

## 68. Empieza a escribir el cliente JavaScript que consumirá tu Backend en Go

**[⬆ volver arriba](#tabla-de-contenidos)**

## 69. Despliega una nueva versión de tu BackEnd para resolver errores en el FrontEnd

**[⬆ volver arriba](#tabla-de-contenidos)**

## 70. Valida que tu servicio Front esté funcionando como debería

**[⬆ volver arriba](#tabla-de-contenidos)**

## 71. Crea los manifiestos de k8s para desplegar tu servicio Front

**[⬆ volver arriba](#tabla-de-contenidos)**

## 72. Crea un Dockerfile para tu aplicación en JavaScript

**[⬆ volver arriba](#tabla-de-contenidos)**

## 73. Tip: ¿No puedes construir un Dockerfile porque no tienes Docker instalado?

**[⬆ volver arriba](#tabla-de-contenidos)**

## 74. Despliega los servicios y valida su funcionamiento

**[⬆ volver arriba](#tabla-de-contenidos)**


### Creation Container
   
* 

### Creacion de deploy a partir de la imagen local

* por defecto la imagen que trata de montar kubernetes es desde dockerHub, por lo que 
es necesario utilizar 
* lo que quiere decir imagePullPolicy es que si la imagen no se encuentra en el local, entonces DESCARGA de docker hub   
* `` 

*`` 
