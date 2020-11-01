const path = require('path')

module.exports = {
  pluginOptions: {
    'style-resources-loader': {
      preProcessor: 'sass',
      patterns: [
        path.resolve(__dirname, './src/assets/styles/variables.sass')
      ]
    }
  },
  chainWebpack: config => {
    const svgRule = config.module.rule('svg')
    svgRule.uses.clear()
    svgRule
      .test(/\.svg$/)
      .oneOf('raw')
      .resourceQuery(/raw/)
      .use('html')
      .loader('html-loader')
      .end()
      .end()
      .oneOf('file')
      .use('file')
      .loader('file-loader')
  },
  publicPath: process.env.NODE_ENV === 'production' ? '/hack2020/' : '/' // удалить для прода!!!
}
