### TypeScript setup

[markdown doc: https://www.markdownguide.org/cheat-sheet/]: # 

My laptop already had node and npm set up, so I do not include installation instructions for that here.

> `npm` comes with `Node.js`. To learn more about working with `npm`, check out this [How To Use Node.js Modules with npm and package.json tutorial](https://www.digitalocean.com/community/tutorials/how-to-use-node-js-modules-with-npm-and-package-json).

## TypeScript install

I used the following as my guide to getting started in TypeScript:
[How To Set Up a New TypeScript Project](https://www.digitalocean.com/community/tutorials/typescript-new-project)

To start I created a typescript directory and cd-ed to it
```
% mkdir typescript typescript/build
% cd typescript
```

and then installed typescript via `npm` and initialized the TypeScript project:
```
% npm i typescript --save-dev

% npx tsc --init

% ls -l
total 48
drwxr-xr-x  6 bwest  staff    192 Dec 12 12:50 ./
drwxr-xr-x  8 bwest  staff    256 Dec 12 12:29 ../
drwxr-xr-x  5 bwest  staff    160 Dec 12 12:47 node_modules/
-rw-r--r--  1 bwest  staff    597 Dec 12 12:47 package-lock.json
-rw-r--r--  1 bwest  staff     58 Dec 12 12:47 package.json
-rw-r--r--  1 bwest  staff  12598 Dec 12 12:50 tsconfig.json
```

The `node_modules` directory contains the external node modules that TypeScript requires.  The `package.json` will contain the package dependencies.

The `tsconfig.json` file contains the TypeScript configurations.  I edited this file to force the compiler to created the JS files in the build directory:
```
    "outDir": "./build", /* Specify an output folder for all emitted files. */
```

I created a simple Hello Word TS progam
```
% cat index.ts
const world = 'world';

export function hello(who: string = world): string {
  return `Hello ${who}! `;
}
```

and compled and ran it:
```
% npx tsc

% ls -lt build/index.js 
-rw-r--r--  1 bwest  staff  184 Dec 12 12:59 build/index.js

% node build/index.js 
```

Important things to note:  
The TypeScript complier trans-compile the TS into JS
Node.JS can be used to run the compiled JS file
Need `npx` to access `tsc` since `tcs` is not "globally" defined.

## GTS (Google TypeScript Style) install

To install`gts`:
```
% npm i gts --save-dev
```

To initialize `gts`:
```
% npx gts init
```
which will update the files: `tsconfig.json` and `package.json`

The `package.json` file should look like this:
```
% at package.json 
{
  "devDependencies": {
    "@types/node": "^22.7.5",
    "gts": "^6.0.2",
    "typescript": "^5.6.3"
  },
  "scripts": {
    "lint": "gts lint",
    "clean": "gts clean",
    "compile": "tsc",
    "fix": "gts fix",
    "prepare": "npm run compile",
    "pretest": "npm run compile",
    "posttest": "npm run lint"
  },
  "dependencies": {
    "rand-seed": "^2.1.7"
  }
}
```

To run the gts linter and fixer:
```
% npx gts lint
% npx gts lint index.ts
% npx gtx fix
% npx gts fix index.ts
```

> As GTS provides an opinionated, no-configuration approach, it will use its own sensible linting and fixing rules. These follow many best practices, but if you find yourself needing to modify the rules in any way, you can do so by extending the default eslint rules. To do so, create a file in your project directory named .eslintrc which extends the style rules:

I really hate the two spaces for an indent, so I hacked the `.prettierrc.js` and `.prettierrc.json`
```
% cp node_modules/gts/.prettierrc.json .
# edit .prettierrc.json
% cat .prettierrc.json
{
  "tabWidth": 4,
  "bracketSpacing": false,
  "singleQuote": true,
  "trailingComma": "all",
  "arrowParens": "avoid"
}

# edit .prettierrc.js
% cat .prettierrc.js
module.exports = {
  ...require('.prettierrc.json'),
  //   ...require('gts/.prettierrc.json'),
}
```
