source .env
  
APOLLO_KEY=$APOLLO_KEY APOLLO_GRAPH_REF=$APOLLO_GRAPH_REF ./router \
  --dev \
  --log trace \
  --anonymous-telemetry-disabled \
  --config router-config.yaml \
  --supergraph schema.graphql