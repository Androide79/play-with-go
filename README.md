# Play with GO

Simple project that show how build something with the GO langunage.

## TABLE OF CONTENTS

* **[REQUIREMENTS](#requirements)**
* **[HOW TO START](#how_to_start)**
* **[RUN TESTS](#tests)**
  * [Test Coverage](#tests_coverage)
* **[AUTHOR](#author)**

## REQUIREMENTS <a name="requirements"></a>

- [GO](https://golang.org/) 1.14.4 (or greater)

## HOW TO START <a name="how_to_start"></a>

To run the project go into *greetings* folder from terminal and type

    go run .

## RUN TESTS <a name="tests"></a>

To run tests go into *greetings* folder from terminal and type

    go test

You can output print additional information about test function using verbose *-v* command argument

    go test -v

If you log something in your tests with the ***-v*** will be shown, otherwise no.

### Test Coverage <a name="tests_coverage"></a>

If you want to see the code coverage, you can use

    go test -cover

The output show you the coverage percentage. To see specifically which parts of the code are covered by tests, you should use

    go test -coverprofile=coverage.txt

where coverage.txt is the file in which the covered parts of code will be written.

## AUTHOR <a name="author"></a>
[Pierluigi D'Antona](https://www.pierluigidantona.it)