# Commit Instructions for GitHub Copilot Agent

This file provides instructions for the GitHub Copilot coding agent when making commits to this repository.

## Commit Message Format

All commits **must** follow the Conventional Commits standard:

```
<type>(<scope>): <short description>
```

### Required Components

- **type**: The kind of change (required)
- **scope**: Single-word identifier for the affected subpackage/domain (optional but recommended)
- **description**: Brief summary in lowercase, no ending punctuation (required)

## Commit Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **refactor**: Code changes that neither fix bugs nor add features
- **test**: Adding or updating tests
- **chore**: Maintenance tasks
- **style**: Code style/formatting changes
- **perf**: Performance improvements
- **build**: Build system or dependency changes
- **ci**: CI configuration changes
- **revert**: Revert previous commit

## Scope Guidelines

Use a **single word** identifying the primary affected subpackage or domain:

Common scopes: `api`, `apiserver`, `cli`, `storage`, `model`, `controller`, `agent`, `database`, `cmd`, `core`, `cloud`

Omit scope only if no clear primary domain applies.

## Examples

```
feat(api): add user authentication endpoint
```

```
fix(storage): resolve race condition in volume attachment
```

```
docs: update contributing guidelines
```

## Critical Requirements

- **Format correctly on first attempt** - you cannot rewrite history after pushing
- **PRs with non-compliant commits will be blocked** by commitlint in CI
- Validation runs automatically via `.github/commitlint.config.mjs`

## References

- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [Project guidelines](../docs/contributor/reference/conventional-commits.md)
- [Contributing guide](../CONTRIBUTING.md)
