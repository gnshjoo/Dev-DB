import * as webpack from 'webpack';
import * as path from 'path';

const tsconfigPath = path.join(__dirname, '../tsconfig.json');

const babelrc = path.join(__dirname, '../babel.config.js');
const babelSetting = require(babelrc);
babelSetting.babelrc = false;

export const config: webpack.Configuration = {
  name: 'devdb-setting',
  mode: 'development',
  devtool: 'inline-source-map',

  context: path.join(__dirname, '../src'),
  entry: {
    app: ['./index.tsx'],
  },
  module: {
    rules: [
      {
        test: /\.tsx$/,
        use: [
          {
            loader: 'babel-loader',
            options: {
              ...babelSetting,
            },
          },
          {
            loader: 'ts-loader',
            options: {
              transpileOnly: false,
              configFile: tsconfigPath,
            },
          },
        ],
        exclude: [/node_modules/],
      },
    ],
  },
  output: {
    path: path.resolve(__dirname, '../build'),
    filename: 'bundle.js',
  },
};

module.exports = config;
