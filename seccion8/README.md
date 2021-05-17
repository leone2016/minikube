## SERVICIOS
[DOC](https://kubernetes.io/docs/concepts/services-networking/service/)
¿Por qué necesitamos un servicio ? ó que es en esencia un servicio

> un servicio es un objeto aislado, lo que hace es observar los pods con cierto label, por ejemplo: digo, servicio por favor observa
> los pods con el label -> app: web

¿De que sirve observar los pods desde otro objeto totalmente distinto?

> este servicio me va entregar una ip única que se va ha mantener en el tiempo y que el cluster va garantizar, cuando un usuario realiza un request a la petición del 
>servicio (IP ó DNS), este servicio de K8  es consultar a los pods que tengan el label que hemos definido, luego el pod devuelve la información 

¿ Cual es la ventaja de trabajar de esta forma?

> si un servicio de K8 esta consultando un POD y este muere, tranquilamente pueden consultar otro pod
>el servicio de k8 utiliza  **round robin balancer** para distribuir la carga en los pods, ya que todos los pods tienen la misma versión
> por lo que el cliente va tener la misma version  

* en resumen el servicio de k8, funciona como balanceador sobre los pods que este observando, sin importar si los pods estan dentro de un mismo replicaset

### ENDPOINT en un servicio K8
 
> si tenemos varios pods con label X y si Ip, luego creamos un servicio que va observar los pods con label X, los servicios tienen una Ip única, esto lo va garantizar K8 

¿Como hace un servicio K8 cuando un pods cambia su Ip?

> Puede haber muchos motivos por el cual la Ip del pod pueda cambiar, para  esto hay otro servicio llamado ENDPOINT que lo crea K8 cuando creamos un servicio K8 que tiene label 
> el servicio sabe a que Ip consultar gracias a los endpoints, los ENDPOINTS son básicamente  una lista de Ips

## Primer servicio
command  | Explanation 
------------- | -------------
`kubectl apply -f seccion8/servicios/service.yaml` | en ese archivo se encuentra el service -> deployment -> replica -> pod 
`kubectl get svc -l app=front  ` | consulta un servicio en base a un label, el label viene de la metadata que aplicamos de la metadata servicio, por defecto aparece el servicio por defecto de K8
`kubectl describe svc my-service` | se describe el servicio con el nombre del mismo, en el mismo se puede ver la Ip y puerto del servicio y `Endpoints:  172.17.0.2:80,172.17.0.3:80,172.17.0.4:80` con una serie de Ips que seran utilizadas por el servicio para acceder al los pods 
`kubectl get endpoints` | cuando se crea el servicio de k8 automaticamente se crea el endpoint
`kubectl get pods -o wide` | muestra los pods con su respectiva ip 

## Servicios y DNS

Cuando se crea un servicio K8 este hereda una Ip y un DNS

## Tipos de Servicios 

(DOC)[https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types]

### ClusterIp

al ejecutar `kubectl get svc` el tipo del servicio creado es de typo `ClusterIP` que es una IP virtual que K8 asigna al servicio, esta ip es permanente en el tiempo, K8 se va encargar de mantener esta IP, 
el problema que la ip es internal cluster (no se puede acceder fuera de ella   )


### NodePort

Es un servicio similar a ClusterIp pero este nos permite exponer el cluster

> **Para recordar: ** un cluster de Kubernetes está conformado por un master y por uno o varios nodos

* dentro del nodo tenemos el servicio de tipo ClusterIP y este está haciendo a la  vez de un tipo balanceador para un grupo de PODS

El nodePort es una exposición del servicio por medio  de un puerto en el nodo 
* el node port es básicamente un puerto que  se abre a nivel del nodo, una vez que ingrese al nodo pasa al clusterIP y este último podrá acceder a los PODS

command  | Explanation 
------------- | -------------
`kubectl apply -f seccion8/servicios/nodeport.yaml` | en ese archivo se encuentra el service:NODEPORT -> deployment -> replica -> pod
