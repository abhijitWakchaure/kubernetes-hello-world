
### Create deplyment
```
kubectl create deployment example-db-depl --image=abhijitwakchaure/example-db:1.0.0 --port=8081 --replicas=2
```

OR

```
kubectl apply -f db-deployment.yml
```

### Create service
```
kubectl apply -f app-service.yml
```

### Expose deployment
```
kubectl expose deployment example-db-depl --name=example-db-srv --port=8081 --target-port=8081 --type ClusterIP
```

### Scale deployment
```
kubectl scale deployment example-app-depl --replicas=5
```


### Test with multiple requests
```
for n in {1..100}; do printf "%5s" "$n. "; curl http://192.168.49.2:31277; echo ""; sleep 0.3; done
```

### Inject linkerd
```
kubectl get deploy -o yaml | linkerd inject - | kubectl apply -f -
```

## Helm Commands

### Create new chart
```
helm create example-chart
```

### Install chart

Test the rendered/compiled yaml files before install
```
helm install my-chart example-chart --dry-run
```

Now deploy the chart

```
helm install my-chart example-chart
```
OR using custom values
```
helm install my-chart example-chart -f customValues.yml
```

### Delete chart
```
helm delete my-chart
```

