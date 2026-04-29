# `gh namespace-clone`

## Installation

```shell
gh extension install spenserblack/gh-namespace-clone
```

## Description

This [`gh`][github-cli] extension wraps `gh repo clone`, namespacing your cloned
repositories by prepending the repository's owner to the clone target path.

## Usage examples

```shell
# Clone to the current directory. You will be considered the owner.
# Clones to `<YOUR_USERNAME>/my-repo`.
gh namespace-clone my-repo
```

```shell
# Clone to `owner-name/repo-name` in the current directory.
gh namespace-clone owner-name/repo-name
gh namespace-clone https://github.com/owner-name/repo-name
gh namespace-clone git@github.com:owner-name/repo-name.git
```

```shell
# Clone to `~/Development/owner-name/repo-name`
gh namespace-clone -P ~/Development owner-name/repo-name
```

Call `gh namespace-clone --help` for more information.

## Aliasing

`namespace-clone` is a lot to write. Additionally, you may want some of the CLI flags
to be the default behavior. For convenience, it's recommended to create an alias with
`gh alias set`.

### Example

This would make the alias `gh ns-clone`, which will clone the repository to
`~/Development` and always include the domain in the clone's target path.

```shell
gh alias set 'ns-clone' 'namespace-clone --domain --prefix ~/Development "$1"'
```

[github-cli]: https://github.com/cli/cli
