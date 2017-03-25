module.exports = {
    entry: './assets/js/application.js',
    output: {
        path: '/Users/bketelsen/src/github.com/bketelsen/teachweb/assets/dist/js',
        filename: 'app.js',
    },
    module: {
        loaders: [{
            test: /\.js$/,
            exclude: /node_modules/,
            loader: 'babel-loader'
        }]
    }
}
