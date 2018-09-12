'use strict';

const Path = require('path');
const Webpack = require('webpack');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const ImageminPlugin = require('imagemin-webpack-plugin').default;

const dest = Path.join(__dirname, 'dist');

var config = {
  entry: [
    Path.resolve(__dirname, 'src/scripts/polyfills'),
    Path.resolve(__dirname, 'src/scripts/index')
  ],
  output: {
    path: dest,
    filename: 'bundle.js'
  },
  optimization: {
    minimize: true
  },
  plugins: [
    new CleanWebpackPlugin([dest], {
      root: Path.resolve(__dirname, 'dist')
    }),
    new CopyWebpackPlugin([{
      from: Path.resolve(__dirname, 'public'),
      to: 'public'
    }]),
    new MiniCssExtractPlugin({
      filename: 'bundle.css'
    }),
    new ImageminPlugin({
      //disable: process.env.NODE_ENV !== 'production', // Disable during development
      pngquant: {
        quality: '95-100'
      },
      test: /\.(jpe?g|png|gif|svg)$/i
    })
  ],
  /*resolve: {
    alias: {
      '~': Path.resolve(__dirname, '../src')
    }
  },*/
  module: {
    rules: [{
      test: /\.(ico|jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2)(\?.*)?$/,
      use: {
        loader: 'file-loader',
        options: {
          name: '[path][name].[ext]'
        }
      }
    }, {
      test: /\.(js)$/,
      exclude: /(node_modules|bower_components)/,
      use: 'babel-loader'
    }, {
      test: /\.s?css/i,
      use: [
        MiniCssExtractPlugin.loader,
        'css-loader',
        'sass-loader'
      ]
    }]
  }
};

module.exports = (env, argv) => {
  // Disable image compress during development
  config.plugins.push(new ImageminPlugin({
    disable: argv.mode !== 'production';
    pngquant: {
      quality: '95-100'
    }
  }));
  return config;
};
