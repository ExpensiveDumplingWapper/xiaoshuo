
all: ucloud aws

ucloud:
	kustomize build deploy/overlays/ucloud/dev > deploy/overlays/ucloud/dev/deploy.yaml
	kustomize build deploy/overlays/ucloud/prod > deploy/overlays/ucloud/prod/deploy.yaml

aws:
	kustomize build deploy/overlays/aws/prod > deploy/overlays/aws/prod/deploy.yaml

#tencent:
#	kustomize build deploy/overlays/tencent/dev > deploy/overlays/tencent/dev/deploy.yaml
#	kustomize build deploy/overlays/tencent/prod > deploy/overlays/tencent/prod/deploy.yaml