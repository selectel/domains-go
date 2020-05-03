/*
Package cluster provides the ability to interact with domains
through the Selectel Domains API V!.

Example of getting a single domain by its id

  selectelDomain, _, err := domain.GetByID(ctx, serviceClient, domainID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", selectelDomain)

Example of getting a single domain by its name

  selectelDomain, _, err := domain.GetByName(ctx, serviceClient, domainName)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", selectelDomain)

*/
package domain
