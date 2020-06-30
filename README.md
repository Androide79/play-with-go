# Play with GO

Simple project that show how build something with the GO langunage.

## TABLE OF CONTENTS

* **[REQUIREMENTS](#requirements)**
* **[HOW TO START](#how_to_start)**
  * [Requirements Installation](#how_to_start_requirements_installation)
  * [Environment Configuration](#how_to_start_environment_configuration)
  * [Project Run](#how_to_start_project_run)
  * [Third-Party Packages Update](#third_party_packages_update)
* **[RUN TESTS](#tests)**
  * [Test Coverage](#tests_coverage)
* **[APPENDIX](#appendix)**
  * [.ENV File Configuration](#appendix_env_file_configuration)
* **[AUTHOR](#author)**

## REQUIREMENTS <a name="requirements"></a>

- [GO](https://golang.org/) 1.14.4 (or greater)
- [GO-GREETINGS](https://github.com/Androide79/go-greetings)
- [VIPER](https://github.com/spf13/viper)

## HOW TO START <a name="how_to_start"></a>

### Requirements Installation <a name="how_to_start_requirements_installation"></a>

First you need to install the requirements, to do this type the command

    go mod download

### Environment Configuration <a name="how_to_start_environment_configuration"></a>

GO lang provides several native and non-native tools to work with environment variables.
In this project we will use the [Viper](https://github.com/spf13/viper) library to set the configuration parameters of the project within an *.env* file.
To use the *.env* file rename the ***.env_example*** file to ***.env*** (the file is located in the project root). In this file is possibile to set some environment variables (for a detailed list of what can currently be configured using the *.env* file you can find it in the [Appendix](#appendix_env_file_configuration))

### Project Run <a name="how_to_start_project_run"></a>

To run the project go into *greetings* folder from terminal and type

    go run .

### Third-Party Packages Update <a name="third_party_packages_update"></a>

If you need to update third-party packages, which may have been updated over time, you will need to run the following command from the terminal

    go get -u all

## RUN TESTS <a name="tests"></a>

To run tests go into *greetings* folder from terminal and type

    go test

You can output print additional information about test function using verbose *-v* command argument

    go test -v

If you log something in your tests with the ***-v*** will be shown, otherwise no.
In case you have more than one packages and you want to execute all tests of every package at one time you can run in your base project directory

    go test ./...

with ***./...*** goes through each package and runs the test files.
Or if you need to test a particular package, just run the command

    go test <package_name>

### Test Coverage <a name="tests_coverage"></a>

If you want to see the code coverage, you can use

    go test -cover

The output show you the coverage percentage. To see specifically which parts of the code are covered by tests, you should use

    go test -coverprofile=coverage.txt

where coverage.txt is the file in which the covered parts of code will be written.
To have a more human readable output we can use a built-in GO tool, which generates an output in HTML format

    go tool cover -html=coverage.txt -o coverage.html

using the TXT output, previously seen, creates a more readable HTML page.

## APPENDIX <a name="appendix"></a>

### .ENV File Configuration <a name="appendix_env_file_configuration"></a>

| KEY | VALUE | REQUIRED | NOTE |
| ------ | ------ | ------ | ------ |
| NAME | string || The name that will be passed to receive a greet |
| LANGUAGE | string || User language (it/fr/de/es) - default english |
| LOG_FILE_NAME | string || The name for the log file |


## AUTHOR <a name="author"></a>
[Pierluigi D'Antona](https://www.pierluigidantona.it)