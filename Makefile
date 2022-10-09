export NAMESPACE := default
export IMAGE := amrfarid/webhook:0.3.0

install-cert-manager:
	kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.9.1/cert-manager.yaml
build:
	docker build -t ${IMAGE} .
	kind load docker-image ${IMAGE}
deploy:
	envsubst < deployment/certificate.yaml | kubectl apply -f -
	envsubst < deployment/deployment.yaml |  kubectl apply -f -
remove:
	kubectl delete -f deployment/
run-local:
	mkdir -p certs
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout certs/tls.key -out certs/tls.crt -subj "/CN=webhook/O=webhook"
	air server -debug=true