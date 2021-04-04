module.exports = {
  presets: [
    [
      '@babel/preset-env',
      {
        targets: {
          browsers: ['> 5% in KR'],
        },
      },
    ],
    '@babel/preset-react',
  ],
  plugins: ['@emotion'],
};
