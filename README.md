# KUBERNETES de cero a experto


* [**Seccion 09: Golang, JavaScript y Kubernetes - Aprende a construir aplicaciones reales**](session9)

> `minikube start`
> `minikube stop`
> `minikube status`
> `minikube ssh` dentro del minikube se puede ver las imagenes creadas desde Docker
    *  eval $(minikube docker-env)



### Explorando Pods

> Note: All kubectl run generators are deprecated. See the Kubernetes v1.17 documentation for a list of generators and how they were used.

* un pod es la unidad minima de Kubernetes que podemos crear

[documentación para crear un pod](https://v1-17.docs.kubernetes.io/docs/reference/kubectl/conventions/#generators)

command  | Explanation 
------------- | -------------
`kubectl get pods` | valida que se este ejecutando correectamente, en el caso de no existir pods: `No resources found in default namespace`
`kubectl run --generator=run-pod/v1 <namepod> --image=nginx:alpine` | crea un pod con una imagen 
`kubectl describe pod <namepod>` | se tiene mas detalle del error del pod, o información mas relevante
`kubectl api-resources` | muestra un listado de propiedades que se pueden ejecutar eje: `kubectl get po `
kubectl delete pod [namePOD namePOD]` | elimina un pod, se puede elimanar mas de un pod a la vez  
`kubectl get pod <namepod>` | se obtiene un solo pod específico
`kubectl get pod <namepod> -o yaml` | entrega el manifiesto de kubernetes, es decir toda toda la declaración del pod, de como el pod se creo en la api de K8 

```Console
NAME      READY   STATUS    RESTARTS   AGE
podtest   1/1     Running   0          13s
```
*  NAME: nombre del pod
* READY: estado 1/1, un pod puede tener mas de un contenedor, este caso solo tiene una imagen de la imagen nginx:alpine, 
    en el caso que no se levante el pod se puede visualizar **0/1**


> una vez creado el pod, para visualizar en docker las imagenes que este creo `docker ps`, en este caso no se va manejar el contenedor de  docker, sino desde K8

>al final del dia todo es contenedores de DOCKER

command  | Explanation 
------------- | -------------
`kubectl describe pod podtest` | tomar la ip generada internamente de K8 y pegarla en el navegador, en el caso de no visualizar nada hacer el paso siguiente
`kubectl port-forward <podname> 7000:<podport>`  en el caso de no funcionar la ip interna de k8, ya que nuestra maquina es el cluster se puede utilizar la ip que nos da K8, pero no siempre funciona 
`kubectl exec -ti <namepod> -- sh ` | ingresa a los contenedores dentro de un pod, ingresa al terminal de nginx alpine
`kubectl logs <namepod> -f ` | con **-f** es persistir los logs

  

## DOCKER
* `docker run --net host -ti python:3.6-alpine sh`| ejecuta un contenedor en docker, en la red de host (-ti)  iterativamente, (sh) ingresa a la terminal
* | con el paso anterior, ingresa a la terminal 


## LINUX
* `apk add -U curl` | instala curl en el caso que no este instalado
