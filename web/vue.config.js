module.exports = {
    runtimeCompiler: true,
    devServer: {
        progress: false,
        host: '0.0.0.0',
        port: 8080,
        proxy: {
            '/api/trainings': {
                target: 'http://trainings-http:3000',
                changeOrigin: true
            },
            '/api/trainer': {
                target: 'http://trainer-http:3000',
                changeOrigin: true
            },
            '/api/users': {
                target: 'http://users-http:3000',
                changeOrigin: true
            }
        }
    }
}