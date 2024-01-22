### Catalog Indexing Engine

```
curl -X GET "https://8fa4a0da88f948278910552683db741e.us-central1.gcp.cloud.es.io/name/_search" -H 'Content-Type: application/json' -d'     
{
    "query": {
        "match": {
            "name": "Bread"
        }
    }
}' -u 'elastic:o2PE0mohdwggfKU4fTKnbTvz'

```