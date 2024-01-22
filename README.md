### Catalog Indexing Engine
A simple and customizable Catalog Indexing Engine that can index catalog data and search through data at very high speed. Supports both structured and unstructured queries (will be added in upcoming days).

#### Tech Stack: 
1. <a href="https://go.dev/">Go</a>
2. <a href="https://www.elastic.co/elasticsearch">Elasticsearch</a>

#### How to setup: 
1. Create a new file and name is `.env`
2. Create a new deployment of Elasticsearch through elasticsearch cloud portal.
3. Get your deployment endpoint and API Key.
4. Add following code to `.env`

```
ElasticSearchEndpoint=<YOUR_ENDPOINT_URL>

ElasticSearchAPIKey=<YOUR_API_KEY>
```
5. Run the following command `go run main.go`
6. Application will automatically start the server on port 5000.


#### API Endpoints:

<b>1. Index Catalog Data:</b><br>
<b>localhost:5000/index-data</b><br>
<b>Method: </b> POST <br>
<b>Request Body: </b> You can also index any data according to you. Below is just and example.
```
{
    "name" : "Bread",
    "item_type" : "Grocery",
    "price" : 10,
    "category" : "Bakery",
    "quantity" : 20,
    "measuring_unit" : "Packets"
}
```

<br>

<b>2. Search Catalog Data:</b><br>
<b>localhost:5000/search</b><br>
<b>Method: </b> GET <br>
<b>Request Body: </b> Currently it only support structured query. Unstructured query support will be added in few days.

```
{
  "query": {
    "bool": {
      "must": [
        { "match": { "name": "Bread" } },
        { "match": { "quantity": 20 } }
      ]
    }
  }
}
```

<b>Response: </b><br>
```
{
    "data": [
        {
            "category": "Bakery",
            "item_type": "Grocery",
            "measuring_unit": "Packets",
            "name": "Bread",
            "price": 10,
            "quantity": 20
        }
    ],
    "message": "Search Successfull"
}
```
