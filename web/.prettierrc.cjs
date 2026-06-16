/**
 * @see https://prettier.io/docs/en/configuration.html
 * @type {import('prettier').Config}
 */
module.exports = {
    // 格式化选项
    semi: true,
    endOfLine: 'lf',
    singleQuote: true,
    trailingComma: 'all',
    bracketSpacing: true,
    bracketSameLine: false,
    vueIndentScriptAndStyle: true,
    htmlWhitespaceSensitivity: 'ignore',

    // 代码宽度和缩进
    printWidth: 100,
    tabWidth: 2,
    useTabs: false,

    // 语言和解析器选项
    plugins: [
        'prettier-plugin-tailwindcss',
    ],

    // 覆盖配置
    overrides: [
        {
            files: '*.html',
            options: {
                parser: 'html',
            },
        },
        {
            files: '*.json',
            options: {
                parser: 'json',
            },
        },
        {
            files: '*.md',
            options: {
                parser: 'markdown',
            },
        },
        {
            files: '*.yml',
            options: {
                parser: 'yaml',
            },
        },
    ],

    // 实验性选项
    experimentalTernaries: false,

    // 其他优化选项
    arrowParens: 'always',
    quoteProps: 'as-needed',
    proseWrap: 'preserve',
    embeddedLanguageFormatting: 'auto',
}
