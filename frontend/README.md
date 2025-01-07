# Chinese checkers Frontend

- [Chinese checkers Frontend](#chinese-checkers-frontend)
  - [Dependencies](#dependencies)
  - [Documentation](#documentation)
  - [Linting](#linting)
  - [Testing](#testing)
    - [Run workflow via act](#run-workflow-via-act)
  - [Project Structure](#project-structure)
  

## Dependencies

- [TypeScript](https://www.typescriptlang.org/)
- [React](https://reactjs.org/)
- [Redux](https://redux.js.org/)
- [Styled Components](https://styled-components.com/)
- [Jest](https://jestjs.io/)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro)
- [ESLint](https://eslint.org/)
- [Prettier](https://prettier.io/)

Install via

```bash
bun install
```

## Documentation

Run JSDoc to generate documentation.

```bash
bun run typedoc
```

## Linting

```bash
bun lint
```

## Testing 

```bash
bun test
```

### Run workflow via act

```bash
act -j react
```

## Project Structure

Frontend project structure is presented below.

```plaintext
.
├── README.md
├── package.json
├── public
│   ├── index.html
│   └── manifest.json
├── src
│   ├── App.tsx
│   ├── components
│   │   ├── AppContainer.tsx
│   │   ├── Board.tsx
│   │   ├── Cell.tsx
│   │   ├── Game.tsx
│   │   ├── Home.tsx
│   │   ├── Menu.tsx
│   │   ├── Player.tsx
│   │   ├── PlayerList.tsx
│   │   └── PlayerStats.tsx
│   ├── config
│   │   └── index.ts
│   ├── index.tsx
│   ├── react-app-env.d.ts
│   ├── serviceWorker.ts
│   ├── setupTests.ts
│   ├── store
│   │   ├── actions.ts
│   │   ├── index.ts
│   │   ├── reducer.ts
│   │   └── types.ts
│   └── utils
│       ├── index.ts
│       └── test.ts
├── tsconfig.json
└── yarn.lock
```
