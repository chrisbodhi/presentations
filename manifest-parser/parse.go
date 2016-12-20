{
  "rows": [{
    "Account: Account ID": "1111",
    "Account: Account Name": "Chris Store",
    "Production: Production Number": "chris666",
    "Account: Phone": "(609) 999-6415"
  }, {
    "Account: Account ID": "2222",
    "Account: Account Name": "Boette Store",
    "Production: Production Number": "boette666",
    "Account: Phone": "(609) 666-6415"
  }],
  "mappings": {
    "ad": { "custom_id": "Account: Account ID", "file_name": "Production: Production Number" },
    "business": { "name": "Account: Account Name", "contact_info": {
      "phone": "Account: Phone" }}},
  "transformations": {
    "ad": { "file_name": "appendPdf" },
    "business": { "contact_info": {
      "phone": "removeNonAlphanumeric"
  }}}
}
