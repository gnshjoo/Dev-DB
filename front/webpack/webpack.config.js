"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.config = void 0;
const path = __importStar(require("path"));
const tsconfigPath = path.join(__dirname, '../tsconfig.json');
const babelrc = path.join(__dirname, '../babel.config.js');
const babelSetting = require(babelrc);
babelSetting.babelrc = false;
exports.config = {
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
                        options: Object.assign({}, babelSetting),
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
module.exports = exports.config;
