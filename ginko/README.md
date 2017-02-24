## Ginkgo

https://github.com/onsi/ginkgo

### How to run this example?

```console

 ginkgo -v
```

### The design of the project:


The design is similar to cucumber.

You have the specification of the  ```features``` in features dir.

The actual implementation of the features go in ```step```. (see the auth.login function)

The file spacego_suite_test.go is used to run the global suite. 
