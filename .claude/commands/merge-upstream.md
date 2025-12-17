# Merging Upstream to Downstream

When merging changes from upstream to downstream, follow these steps:

## 1. Preparation

Ensure your local repository is up to date by fetching all remote branches from
git.

Check the current branch. Here you need to make sure that:
- There are no local changes you may overwrite
- You are on the right branch that is based on the `main` branch of
  `openshift/cluster-api-provider-metal3`
- You're on the latest commit in that branch.

Create a new branch for the merge named `merge-upstream-$(date +%Y-%m-%d)`.

## 2. Merge Upstream Changes

```bash
git merge --no-ff upstream/main
```

## 3. Handle Common Conflicts

**GitHub Workflows**: Downstream typically removes `.github/` directory files. If conflicts occur:
- Keep the downstream version (usually removal)
- Verify with `git diff downstream/main -- .github/`

**Downstream-Specific Changes**: Look for commit messages starting with:
- `DOWNSTREAM:` - Changes specific to OpenShift
- `OCPBUGS-*:` - OpenShift bug fixes

**OWNERS**: Revert any changes to `OWNERS`, you must keep the downstream
version of this file.

These must be preserved during the merge, unless replaced by an equivalent
upstream commit.

Commit the merge; the commit message should be "Merge upstream".

## 4. Downstream specific actions

**Go Dependencies**: After merging, update dependencies:
```bash
# Clean up and tidy all go.mod files across all modules
make modules

# Update vendor directories for all modules
make vendor

# These commands handle:
# - Root module (go.mod)
# - apis/ module
# - pkg/hardwareutils/ module
# - hack/tools/ module
# - test/ module
```

Note that all dependencies must be vendored in OpenShift. You cannot rely on
downloading modules. Do not use `-mod=mod` with `go build`.

Commit the changes to `vendor` directories.  The commit message should be "Update vendor".

Run `make ocp-manifests` in the `openshift` directory and commit the result.

## 5. Test the Changes

Before creating a PR, ensure:
- Code compiles: `make build`
- Tests pass: `make test`
- Manifests are up to date: `make manifests`
- Linting passes: `make lint`

You might run into bugs related to changes in Golang version.  Try to fix the bug and commit the changes.

## 6. Push

Push to the user's fork (replacing `origin` with the name of the user's
personal fork):

```bash
git push -u origin merge-upstream-$(date +%Y-%m-%d)
```

Tell the user how to create the pull request using the base repository
`openshift/cluster-api-provider-metal3`.  Ask them if they want to open the New Pull
Request page on GitHub:
`https://github.com/openshift/cluster-api-provider-metal3/compare/main...<user
fork>:cluster-api-provider-metal3:<merge branch>`.
