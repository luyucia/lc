module.exports = {
    devServer: {
        proxy: {
            '/': {
                target: 'http://127.0.0.1:8022',
                ws: true
            }
        }
    }
}