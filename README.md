# Turborepo template

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fhitz-group%2Fmonorepo.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fhitz-group%2Fmonorepo?ref=badge_shield)

## Turborepo

### Pros

- Good task runner performance
- Local and distributed build cache
- Prune command to trim monorepo for a specific app/package (Faster CI builds)
- Less configuration required for setup or adding new packages
- Great developer/teams freedom

### Cons

- No built in schematics support
- Prune works with yarn classic only

### Turborepo vs Nx vs Rush

- Some arguments here and there

## Package manager

### Why yarn classic?

To take advantage of `turbo prune` to speed up CI builds. Turbo prune may support other package managers in the future. It is relatively very easy to switch to npm or yarn berry if needed.

For now the reduced CI build time with turbo prune outweighs a possibly slower full install

## Project structure

- Show directory tree

## Instructions

```bash
# Do a full install of package dependencies
yarn install

# Run a full build
yarn run build

# Experience a full turbo build
yarn run build

# Run dev on all apps
yarn run dev

# Run unit tests
yarn run test

# Clean build output
yarn run clean

# Rebuild without turbo cache
yarn rebuild

# Purge node_modules and build output
yarn run purge

# Build a single package and its dependencies
yarn run build --no-deps --scope=@oolio/<package>
```

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fhitz-group%2Fmonorepo.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fhitz-group%2Fmonorepo?ref=badge_large)
