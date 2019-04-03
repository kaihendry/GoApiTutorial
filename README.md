If the JSON payload misses a value on POST (INSERT) or PUT (UPDATE), that value if it's optional, becomes null.

Test should use transactions, like so: https://www.reddit.com/r/golang/comments/b6rs6w/testing_database_models_interfaces_without/ejmoaq7/
