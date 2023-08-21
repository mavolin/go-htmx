# Contributing

We would love to see the ideas you want to bring in to improve this project.
Before you get started, make sure to read the guidelines below.

## Issues

If you have an idea how to improve this project, or if you find a bug, create an issue to let us know.
Please format your issue titles according to the [conventional commits guidelines](https://www.conventionalcommits.org/en/v1.0.0/).

## Code Contributions

### Committing

Please use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) for your commits.

##### Types
We use the following types:

- **build**: Changes that affect the build system or external dependencies
- **ci**: changes to our CI configuration files and scripts
- **docs**: changes to the documentation
- **feat**: a new feature
- **fix**: a bug fix
- **perf**: an improvement to performance
- **refactor**: a code change that neither fixes a bug nor adds a feature
- **style**: a change that does not affect the meaning of the code
- **test**: a change to an existing test, or a new test

### Fixing a Bug

If you're fixing a bug, if possible, add a test case for that bug to ensure it's gone for good.

### Code Style

Make sure all code passes the golangci-lint checks.
If necessary, add a `//nolint:{{name_of_linter}}` directive to the line or block to silence false positives or exceptions.

### Testing

If possible and appropriate you should fully test the code you submit.
Each function should have a single test, which either tests directly or is split into subtests, preferably table-driven.

#### Table-Driven Tests

If there is a single table, it should be called `testCases`, multiple use the name `{{type}}Cases`, e.g. `successCases` and `failureCases` for tests that test the output for a valid input (a success case), and those that aim to provoke an error (a failure case) and therefore work different from a success case.
The same applies if there is a table that's only testing a portion of a function, and multiple non-table-driven tests in addition.

The structs used in tables should always anonymous.
Use `except` as the name of the field that will hold the expected correct value.

Every sub-test including table-driven ones should have a name that clearly shows what is being done.
For table-driven tests this name is either obtained from a `name` field or computed using the other fields in the table entry.

Every case in a table should run in its own subtest (`t.Run`).
Additionally, if there are multiple tables, each table should have its own subtest, in which it calls its cases:

```
TestSomething
    testCase
    testCase

    additionalTest
```

```
TestSomething
    successCases
        successCase
        successCase
    failureCases
        failureCase
        failureCase

    additionalTest
```

```
TestSomething
    successCases
        successCase
        successCase
    failureCases
        failureCase
        failureCase
        additionalFailureTest
```

## Opening a Pull Request

When opening a pull request, use the title of the issue as PR title.

A Pull Request must pass all tests to be merged.
