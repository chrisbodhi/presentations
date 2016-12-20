{
  "config": {
    "detail_file_type": {
      "headers": true, "col_separation": ",", "encoding": "utf-8" },
    "mappings": {
      "ad": { "custom_id": "Account: Account ID", "file_name": "Production: Production Number" },
      "business": {
        "name": "Account: Account Name",
        "contact_info": {
          "phone": "Account: Phone", "website": "Account: Website"
        }}},
    "transformations": {
      "ad": { "file_name": "appendPdf" },
      "business": {
        "contact_info": {
          "phone": "removeNonAlphanumeric"
        }}},
    "publisher_uuid": "0f98d4f9-d820-4a4e-bbf7-31b2ec58a771" },
  "file": {
    "type": "Buffer",
    "data": [1, 2, 3, 42]
  },
  "filename": "test.csv"
}
