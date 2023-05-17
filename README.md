
## What is Semankit?

Semankit is an **CI tool** that manages versioning by determining the subsequent version
through the last tag and incrementing it based on the commit from the latest push.


## Why Semankit ?

Many developers mentally determine the version of their project, which causes inconsistencies.
Others use local tools to version their code before pushing it, which wastes time.
**Semankit aims to save time, obtain consistent versions and enforce conventions over commit message**


## How to bump versions ?

To update versions, Semankit relies on the format of commit messages which is
the "[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)" convention and
only commits with messages **containing "feat" and "fix"** increase the versions.

The following table shows the version increase based on the commit message:

| Commit message            | Is a breaking change ? | Version increase |
|---------------------------|------------------------|------------------|
| fix: issue one            | No                     | Patch            |
| fix: issue two            | Yes                    | Minor            |
| fix(signup): issue one    | No                     | Patch            |
| fix(signup): issue two    | Yes                    | Minor            |
| feat: feature one         | No                     | Minor            |
| feat(signup): feature one | Yes                    | Major            |

