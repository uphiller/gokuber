AZURE_BASE_GROUP_NAME=테스트
AZURE_LOCATION_DEFAULT=Korea Central
AZURE_SAMPLES_KEEP_RESOURCES=0

# create with:
# `az ad sp create-for-rbac --name 'my-sp' --output json`
# sp must have Contributor role on subscription
AZURE_TENANT_ID=96fe0bf7-70bb-4a02-be1c-a37e898345ea
AZURE_CLIENT_ID=e350e1f4-9cda-4975-b8c3-4721dc24cb5b
AZURE_CLIENT_SECRET=sLEjVIVGLu]@olj/y9:xCq2OTta758cZ
AZURE_SUBSCRIPTION_ID=e8fe9bd4-424f-4247-b815-2c84e7c3d0a8

# create with:
# `az ad sp create-for-rbac --name 'my-sp' --sdk-auth > $HOME/.azure/sdk_auth.json`
# sp must have Contributor role on subscription
AZURE_AUTH_LOCATION=$HOME/.azure/sdk_auth.json

AZURE_STORAGE_ACCOUNT_NAME=
AZURE_STORAGE_ACCOUNT_GROUP_NAME=