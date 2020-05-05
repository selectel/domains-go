/*
Package record provides the ability to interact with domain records
through the Selectel Domains API V1.

Example of getting a single domain record by its id

  domainRecord, _, err := record.Get(ctx, serviceClient, domainID, recordID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", domainRecord)

Example of getting a list of domain records by domain id

  domainRecords, _, err := record.ListByDomainID(ctx, serviceClient, domainID)
  if err != nil {
    log.Fatal(err)
  }
  for _, r := range domainRecords {
    fmt.Printf("%+v\n", r)
  }

Example of getting a list of domain records by domain name

  domainRecords, _, err := record.ListByDomainName(ctx, serviceClient, domainName)
  if err != nil {
    log.Fatal(err)
  }
  for _, r := range domainRecords {
    fmt.Printf("%+v\n", r)
  }
*/
package record
