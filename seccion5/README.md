
### Manifiesto de Kubernetes
* cuando se habla de un manifiesto es un archivo en yaml, define el recurso que queremos crear o actualizar en K8
* Por qué utilizar manifiestos en vez de linea de comandos ? 
    * imaginemos que tenemos que crear 50 pods individuales, tendriamos que ejecutar kubectl run 50 veces y si quisieramos eliminarlos se tendria que ejecutar `kubectl delete ...` 50 veces
    * no es recomendable crear los pods manualmente, ya que no se tiene un buen control de los recursos que estan creados 
    * es complicado la administración individual  

command  | Explanation 
------------- | -------------
`kubectl apply -f pod.yaml` | crea los pods a partir de un archivo  `seccion5/pod/pod.yaml`
`kubectl delete -f pod.yaml` | elimina los pods a partir de un archivo

### Logs de un contenedor
si al crear existe un error, es necesario entrar a cada pod y ver sus los del motivo por el cual no se creo 

command  | Explanation 
------------- | -------------
`kubectl get pods` | ...
`kubectl logs doscont` | ingresa a los logs del pod, en este caso se llama doscont, igual al ejempo del yaml,  `seccion5/pod/doscontenedores.yaml`
`kubectl logs doscont -c` | si el pod tiene mas de un contenedor, es necesario especificar (-c) que contenedor se desea ver los logs
`kubectl exec -ti doscont -c cont1 -- sh ` | ingresa al contenedor de un pod, en el caso que el tenga varios contenedores es necesario eso especificar a cual se desea ingresar

## Labels
`seccion5/pod/label.yaml`

command  | Explanation 
------------- | -------------
`kubectl apply -f seccion5/pod/label.yaml` |  crea los dos pods
`kubectl get pods -l app=back` |  trae todos los pods con el label back (backend) 
`kubectl get pods -l env=dev` |  trae todos los pods con el env (ambiente) desarrollo (dev )

## Problema con los Pods

* si se elimina manualmente o por medio del manifiesto, no se recuperan solos
* si se necesitaria minimo dos replicas de mi pod, el problema es que lo pods no pueden replicarse
* los pods no pueden actualizarse a si mismos, que pasa si quiero cambiar configuraciones: comando, esto no se puede hacer desde el pod mismo 

> una solucion a esto es crear un `replica set `