import js from '@eslint/js'
import tsPlugin from '@typescript-eslint/eslint-plugin'
import unusedImports from 'eslint-plugin-unused-imports'
import prettier from 'eslint-config-prettier'

export default [
  {
    ignores: [
      'node_modules/**',
      'dist/**',
      'coverage/**',
      'bundle/**',
      'runtime/**',
      'vendor/**',
      'eslint.config.js',
      'wasm_exec.js',
      '**/*.pb.ts',
    ],
  },
  js.configs.recommended,
  ...tsPlugin.configs['flat/recommended'],
  {
    plugins: {
      'unused-imports': unusedImports,
    },
    languageOptions: {
      parserOptions: {
        project: './tsconfig.json',
      },
    },
    rules: {
      '@typescript-eslint/explicit-module-boundary-types': 'off',
      '@typescript-eslint/no-non-null-assertion': 'off',
    },
  },
  prettier,
]
