#!/bin/bash

ELASTICSEARCH_URL="http://localhost:9200/shakespeare/_search"
HEADERS="Content-Type: application/json"

# Full-Text Search Query
QUERY1='{
    "query": {
        "match": {
            "text_entry": "to be or not to be"
        }
    }
}'

# Aggregation Query
QUERY2='{
    "size": 0,
    "aggs": {
        "speakers": {
            "terms": {
                "field": "speaker.keyword"
            }
        }
    }
}'

# Combined Query
QUERY3='{
    "query": {
        "bool": {
            "must": [
                { "match": { "text_entry": "love" } },
                { "match": { "text_entry": "hate" } }
            ],
            "filter": [
                { "term": { "play_name.keyword": "Hamlet" } }
            ]
        }
    },
    "aggs": {
        "speeches_per_speaker": {
            "terms": { "field": "speaker.keyword" }
        }
    }
}'

# Parameters for hey
REQUESTS=100_000_000  # Increase number of requests
CONCURRENCY=15  # Increase number of concurrent requests

echo "Running Full-Text Search Load Test"
hey -n $REQUESTS -c $CONCURRENCY -m POST -H "$HEADERS" -d "$QUERY1" $ELASTICSEARCH_URL

echo "Running Aggregation Load Test"
hey -n $REQUESTS -c $CONCURRENCY -m POST -H "$HEADERS" -d "$QUERY2" $ELASTICSEARCH_URL

echo "Running Combined Query Load Test"
hey -n $REQUESTS -c $CONCURRENCY -m POST -H "$HEADERS" -d "$QUERY3" $ELASTICSEARCH_URL