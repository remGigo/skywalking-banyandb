# Contributing to Apache SkyWalking Cloud on Kubernetes

Firstly, thanks for your interest in contributing! We hope that this will be a
pleasant first experience for you, and that you will return to continue
contributing.

## Code of Conduct

The Apache software Foundation's [Code of Conduct](http://www.apache.org/foundation/policies/conduct.html) governs this project, and everyone participating in it.
By participating, you are expected to adhere to this code. If you are aware of unacceptable behavior, please visit the
[Reporting Guidelines page](http://www.apache.org/foundation/policies/conduct.html#reporting-guidelines)
and follow the instructions there.

## How to contribute?

Most of the contributions that we receive are code contributions, but you can
also contribute to the documentation or report solid bugs
for us to fix.

## How to report a bug?

* **Ensure no one did report the bug** by searching on GitHub under [Issues](https://github.com/apache/skywalking/issues).

* If you're unable to find an open issue addressing the problem, [open a new one](https://github.com/apache/skywalking/issues/new).
Be sure to include a **title and clear description**, as much relevant information as possible,
and a **code sample** or an **executable test case** demonstrating the expected behavior that is not occurring.

## How to add a new feature or change an existing one

_Before making any significant changes, please [open an issue](https://github.com/apache/skywalking/issues)._
Discussing your proposed changes ahead of time will make the contribution process smooth for everyone.

Once we've discussed your changes and you've got your code ready, make sure that tests are passing and open your pull request. Your PR is most likely to be accepted if it:

* Update the README.md with details of changes to the interface.
* Includes tests for new functionality.
* References the original issue in the description, e.g., "Resolves #123".
* Has a [good commit message](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).

## Building and Testing

Clone the source code and run `make build` in the source directory,
which will build the default binary file in `<sub_project>/build/bin/`.

```shell
make build
```

Please refer to [installation](./docs/installation.md) for more details.

Test your changes before submitting them by

```shell
make test
```

## Linting your codes

We have some rules for the code style and please lint your codes locally before opening a pull request.

```shell
make lint
```

If you found some errors in the output of the above command, try `make format` to fix some obvious style issues. As for the complicated errors, please correct them manually.

## Test your changes before pushing

We have a series of checking tests to verify your changes by

```shell
make pre-push
```

Please fix any errors raised by the above command.

## Update licenses

If you import new dependencies or upgrade an existing one, trigger the licenses generator
to update the license files.

```shell
make license-dep 
```
