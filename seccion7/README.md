## DEPLOYMENT
[DOC](https://kubernetes.io/es/docs/concepts/workloads/controllers/deployment/)
> hasta el momento se a visto que la unidad minima es un pod, y los que manejan el numero de replicas y que el pod se encuentre activo es un replicaset

>un problema hasta ahora es que si se edita el pod en un manifiesto del replicaset, este cambio no se vera afectado, para eso existe deployment
* ver video de la seccion 7 - clase 44 :: explicación del funcionamiento

## Primer deployment
command  | Explanation 
------------- | -------------
`kubectl apply -f seccion7/deployment/deployment.yaml` | en ese archivo se encuentra el deployment -> replica -> pod 
`kubectl get deployment` | obtiene las replicaset activas ó `kubectl get replicaset`
`kubectl get deployment --show-labels` | muestra los labels de los deployments
`kubectl rollout status deployment deployment-test` | nos muestra como estuvo el status del deployment
`kubectl rollout history deployment deployment-test` | nos muestra un historial del deployment
`kubectl rollout undo deployment deployment-test --to-revision=3` | se realiza un rollback, a la revision 3, por defecto  K8 guarda 10 deployments anteriores en el tiempo
`kubectl describe deployments.apps deployment-test` 
 
 
* Revisar owner references clase47
