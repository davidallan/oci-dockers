import oci

import dis

from oci.object_storage.models import CreateBucketDetails

# What version of the OCI CLI is it?
print("Version is " + oci.__version__)

#
# Make sure you have your .oci/config file defined and setup to your tenancy
#
config = oci.config.from_file()

# Let's print the tenancy
compartment_id = config["tenancy"]
print("Compartment is "+compartment_id)

# Let's get the Object Storage Namespace - must have privs for this
object_storage = oci.object_storage.ObjectStorageClient(config)
namespace = object_storage.get_namespace().data
print("Object storage namespace is "+namespace)
