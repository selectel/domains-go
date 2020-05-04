/*
Package record provides the ability to interact with domain records
through the Selectel Domains API V1.

Example of getting a single domain record by its id

  domainRecord, _, err := record.Get(ctx, serviceClient, domainID, recordID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", domainRecord)
*/
package record
