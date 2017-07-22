const path = require('path');
const webpack = require('webpack');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

let plugins = [
  new ExtractTextPlugin('bundle.css'),
];

if (process.env.NODE_ENV === 'development') {
  plugins = plugins.concat([
    new webpack.LoaderOptionsPlugin({
      debug: true
    })
  ]);
}

if (process.env.NODE_ENV === 'production') {
  plugins = plugins.concat([
    new webpack.optimize.UglifyJsPlugin({
      output: {comments: false},
      test: /bundle\.js?$/
    }),
    new webpack.DefinePlugin({
      'process.env': {NODE_ENV: JSON.stringify('production')}
    })
  ]);
};

const config  = {
  entry: {
    bundle: path.join(__dirname, 'client/index')
  },
  output: {
    path: path.join(__dirname, 'server/data/static/build'),
    publicPath: '/static/build/',
    filename: '[name].js'
  },
  plugins: plugins,
  module: {
    rules: [
      {test: /\.css/, loader: ExtractTextPlugin.extract('css-loader')},
      {test: /\.(png|gif)$/, loader: 'url-loader?name=[name]@[hash].[ext]&limit=5000'},
      {test: /\.(pdf|ico|jpg|eot|otf|woff|ttf|mp4|webm)$/, loader: 'file-loader?name=[name]@[hash].[ext]'},
      {test: /\.json$/, loader: 'json-loader'},
      {
        test: /\.jsx?$/,
        include: path.join(__dirname, 'client'),
        loaders: ['babel-loader']
      }
    ]
  },
  resolve: {
    extensions: ['.js', '.jsx', '.css'],
    alias: {
      '#app': path.join(__dirname, 'client'),
      '#c': path.join(__dirname, 'client/components'),
      '#css': path.join(__dirname, 'client/css')
    }
  }
};

module.exports = config;
