module.exports = {
    root: true,
    env: {
        browser: true,
        node: true,
        es6: true,
    },
    parser: 'vue-eslint-parser',
    parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
        parser: '@typescript-eslint/parser',
    },
    extends: [
        'eslint:recommended',
        'plugin:vue/vue3-essential',
        'plugin:@typescript-eslint/recommended',
        'prettier',
        'plugin:prettier/recommended',
        './.eslintrc-auto-import.json', //自动导入忽略
    ],
    plugins: ['vue', '@typescript-eslint', 'prettier'],
    globals: {
        $ref: 'readonly',
        $computed: 'readonly',
        $shallowRef: 'readonly',
        $customRef: 'readonly',
        $toRef: 'readonly',
    },
    rules: {
        'prettier/prettier': 'error',
        'no-console': 'off',
        '@typescript-eslint/no-explicit-any': ['off'], // 关闭any类型时的警告
        '@typescript-eslint/no-empty-function': ['off'], // 关闭空函数警告
        "@typescript-eslint/no-var-requires": [0],
        "@typescript-eslint/camelcase": ["off"], // 关闭词组下划线校验
        "@typescript-eslint/ban-ts-ignore": ["off"], // 允许使用ts-ignore
        // 'comma-dangle': ["error", {
        //     "arrays": "never",
        //     "objects": "ignore",
        //     "imports": "ignore",
        //     "exports": "never",
        //     "functions": "ignore"
        // }], // 拖尾逗号
        'vue/multi-word-component-names': [
            'off',
            {
                ignores: ['index'], //需要忽略的组件名
            },
        ],
    },
};
