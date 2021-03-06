# Instructions

## Running python docker for OCI

1) Make sure you have Docker installed.

2) Make sure that your OCI configuration in ~/.oci/config does not have any absolute path references to your OCI key.  For example:
```
[DEFAULT]
key_file=~/.oci/oci_api_key.pem
user=<USER_OCID>
fingerprint=<USER_FINGERPRINT>
tenancy=<TENANCY_OCID>
region=us-ashburn-1
```

3) Add this to your aliases (can add to a profile / startup script):
```
alias ocipython='docker run -it --rm --mount type=bind,source=$HOME/.oci,target=/root/.oci --name my-running-script -v "$PWD":/usr/src/myapp -w /usr/src/myapp oci-python:3 python "$@"'
```

## Building oci-python

```
docker build --tag oci-python:3 .
# optionally push
docker push my_hub_user/oci-python
```
