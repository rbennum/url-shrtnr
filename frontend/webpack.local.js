const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin');
const dotenv = require('dotenv');
const webpack = require('webpack');

dotenv.config({ path: 'local.env' });

module.exports = {
    mode: 'development',
    entry: './src/js/index.js',
    devtool: 'inline-source-map',
    devServer: {
        static: './dist',
        port: 8081
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: './src/index.html',
            filename: 'index.html',
            minify: true
        }),
        new HtmlWebpackPlugin({
            template: './src/not_found.html',
            filename: '404.html',
            minify: true,
            inject: false
        }),
        new MiniCssExtractPlugin({
            filename: 'styles.css',
        }),
        new webpack.DefinePlugin({
            'process.env': {
                'API_URL': JSON.stringify(process.env.API_URL)
            }
        })
    ],
    output: {
        filename: 'bundle.[contenthash].js',
        path: path.resolve(__dirname, 'dist'),
        clean: true,
    },
    module: {
        rules: [
            {
                test: /\.css$/i,
                include: path.resolve(__dirname, 'src/css'),
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader',
                ]
            }
        ]
    },
    optimization: {
        minimizer: [
            `...`,
            new CssMinimizerPlugin(),
        ],
    }
};