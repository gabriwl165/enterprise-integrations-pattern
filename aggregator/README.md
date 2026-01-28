# Aggregator

Usually, in distribuited systems, we would like to process message as a whole. When designing an aggregator we have some specific key points to thing about:

- Correlation: How do we group the incoming messages?
- Completeness Condition: When the group of messages are ready to be published foward?
- Aggregation Algorithm: How do we group the valid incoming message?

### Details

We must use the previous steps to be able to implemented our expected Aggregator. When aggregator receives the events, it must check if we need to group this data (Correlation), if it's the first one, we must sabe inside a set with it's id, we already exists we add the new message inside it using an algorithm (Aggregation Algorithm). After add the current message inside the set, we should check if it's already completed (Completeness Condition).

### Installation
In your terminal, inside aggreagator folder, you can type `make compose-up`