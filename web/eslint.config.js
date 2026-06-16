import antfu from '@antfu/eslint-config'

export default antfu(
    {
        // 启用 Vue 支持
        vue: true,
        // 启用 TypeScript 支持
        typescript: true,
        // 禁用 React 支持
        react: false,
        // 启用 Prettier 集成
        prettier: true,
        // 启用 Unicorn 插件
        unicorn: true,
    },
    {
        // 自定义忽略文件
        ignores: [
            'mock/**',
            'coverage/**',
            '*.config.*',
            '*.d.ts',
        ],
    },
    {
        // 自定义规则覆盖 - 将大部分规则降级为警告
        rules: {
            // ==================== 基础规则 ====================
            // console 和 debugger
            'no-console': 'warn',
            'no-debugger': 'error',
            'no-alert': 'warn',
            // 等号严格模式 - 改为警告
            'eqeqeq': 'warn',
            // throw 字面量 - 改为警告
            'no-throw-literal': 'warn',
            // 未使用的表达式 - 改为警告
            'no-unused-expressions': 'warn',
            // 不要使用 new  for side effects - 改为警告
            'no-new': 'warn',
            // prefer rest params - 改为警告
            'prefer-rest-params': 'warn',
            // prefer promise reject errors - 改为警告
            'prefer-promise-reject-errors': 'warn',

            // ==================== TypeScript 规则 ====================
            '@typescript-eslint/no-explicit-any': 'off',
            '@typescript-eslint/no-empty-function': 'off',
            '@typescript-eslint/no-var-requires': 'off',
            '@typescript-eslint/no-unused-vars': ['warn', {
                argsIgnorePattern: '^_',
                varsIgnorePattern: '^_',
            }],
            '@typescript-eslint/ban-ts-comment': 'warn',
            '@typescript-eslint/no-empty-object-type': 'warn',
            '@typescript-eslint/no-use-before-define': 'warn',

            // ==================== Vue 规则 ====================
            // 组件命名
            'vue/multi-word-component-names': 'off',
            // props 解构
            'vue/no-setup-props-destructure': 'off',
            // prop 命名规范 - 改为警告
            'vue/prop-name-casing': 'warn',
            // 自定义事件命名 - 改为警告
            'vue/custom-event-name-casing': 'warn',
            // v-for 需要 key - 改为警告
            'vue/valid-v-for': 'warn',
            // 禁止修改 props - 改为警告
            'vue/no-mutating-props': 'warn',
            // 未使用的变量 - 改为警告
            'vue/no-unused-vars': 'warn',
            // 未使用的 refs - 改为警告
            'vue/no-unused-refs': 'warn',
            // 计算属性副作用 - 改为警告
            'vue/no-side-effects-in-computed-properties': 'warn',
            // 模板根元素 - 改为警告
            'vue/valid-template-root': 'warn',
            // 保留属性 - 改为警告
            'vue/no-reserved-props': 'warn',
            // 废弃的 native 修饰符 - 改为警告（建议修复但实际不影响功能）
            'vue/no-deprecated-v-on-native-modifier': 'warn',

            // ==================== 代码质量规则 ====================
            // 未使用的变量 - 改为警告
            'no-unused-vars': 'warn',
            'unused-imports/no-unused-vars': 'warn',
            // 使用 before define - 改为警告
            'no-use-before-define': 'warn',
            // 数组回调返回值 - 改为警告
            'array-callback-return': 'warn',
            // 正则捕获组 - 改为警告
            'regexp/no-unused-capturing-group': 'warn',

            // ==================== Unicorn 规则 ====================
            // textContent 替代 innerText - 改为警告
            'unicorn/prefer-dom-node-text-content': 'warn',

            // ==================== 样式规则 ====================
            // 期望换行符
            'expect-expect': 'off',
        },
    }
)
