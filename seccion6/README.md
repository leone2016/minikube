## Replica set 
[DOC](https://kubernetes.io/es/docs/concepts/workloads/controllers/replicaset/)

* los mismos comandos que se utilizan para el pod se puede utilizar en el replicaset

> En un pod, podemos tener uno o mas contenedores dentro, un problema, cuando se elimine un pod no se levantar y  no se puede garantizar que ciertos numero de pods se ejecuten en cierto periodo de tiempo

 **REPLICA SET** es un objeto separado al pod pero que esta a un nivel mas alto, xq el **replica set**   crea pods, para que pueda crear los pods es necesario especificar un _label_ `label: -> app:` (app)  
 * replica set primero verifica si existe un pod con el label (app) con el valos que nosotros especifiquemos
 * si existe el pod con el label, asigna en la metadata del pod el **owner reference** asignando a un replica ser, eso quiere decir que un pod responde a un replica set 

 ## Primera replica set 
command  | Explanation 
------------- | -------------
`kubectl apply -f seccion6/replica/replicaset.yaml` | en ese archivo se encuentra la replica 
`kubectl get pods -l app=pod-label` | lista los pods `matchLabels: -> app: pod-label`
`kubectl get rs` | obtiene las replicaset activas ó `kubectl get replicaset`

>como un pod referencia a un replicaset

* para identificar que replica set es el dueño de un pod, es necesario imprimir el manifiesto del pod y del replicaset
* con el manifiesto del pod, buscar _ownerReferences -> uid_, este uuid es el mismo que **rs** buscar en _metadata:->annotations->uid_

## Problemas del replicaset
* el concepto genera replicaset es que debe mantener un numero Nº de replicas del mismo pod en todo el tiempo 
* UN PROBLEMA, cuando editemos el yaml (manifiesto) del rs la metadata del pod y apliquemos el cambio al rs, no pasa nada





