buildlocalppc:
	docker build . -f Dockerfile.ppc64le -t raiproject/carml-mxnet:ppc64le-gpu
buildlocalamd64gpu:
	docker build . -f Dockerfile.amd64_gpu -t raiproject/carml-mxnet:amd64-gpu
buildlocalamd64cpu:
	docker build . -f Dockerfile.amd64_cpu -t raiproject/carml-mxnet:amd64-cpu

buildlocalamd64: buildlocalamd64cpu buildlocalamd64gpu
