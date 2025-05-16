APOLLO_KEY=service:My-Graph-4yys6h:iy0WFgiv6x0JAqwrstoXcA \
  rover subgraph publish My-Graph-4yys6h@current \
  --schema schema.graphql \
  --name testing \
  --routing-url http://products.prod.svc.cluster.local:4001/graphql