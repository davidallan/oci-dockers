package main

import (
  "context"
  "fmt"

  "github.com/oracle/oci-go-sdk/common"
  "github.com/oracle/oci-go-sdk/identity"
)

/*
  Starter Go example. The first part of this connects to Identity using your locally configured OCI using
  .oci/config and the DEFAULT profile. It will just list availability domains to show we are connecting to a tenancy.

*/
func main() {
  c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  // The OCID of the tenancy containing the compartment.
  tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  request := identity.ListAvailabilityDomainsRequest{
    CompartmentId: &tenancyID,
  }

  r, err := c.ListAvailabilityDomains(context.Background(), request)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  fmt.Printf("List of available domains: %v\n", r.Items)

  return
}

