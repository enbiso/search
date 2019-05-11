# Enbiso Search

Elastic search server based on golang blevesearch

*powered by blevesearch*

## Usage

```
go get -d .
go build
./search --data ./data --addr :8080
```

## Docker

- Simple
```
docker run -p 8080:8080 enbiso/search
```

- Advance
```
docker run -p 8080:8080 -v /data:/data enbiso/search --data /data --addr :8080
```

## API Endpoints

Handlers provided by the bleve.http package is used. The handlers are attached to the following URLs:

`PUT /{indexName}` - create new index

`GET /{indexName}` - get index details

`DELETE /{indexName}` - delete index

`GET /` - list indexes

`PUT /{indexName}/{docID}` - index document

`GET /{indexName}/_count` - count documents in index

`GET /{indexName}/{docID}` - return stored fields of document

`DELETE /{indexName}/{docID}` - delete document

`POST /{indexName}/_search` - search index

`GET /{indexName}/_fields` - list fields used by documents in index

`GET /{indexName}/{docID}/_debug` - return rows in index related to document
