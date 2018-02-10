const mix = require('laravel-mix')

mix.js('src/_js/blog.js', 'src/js')
mix.sass('src/_scss/blog.scss', 'src/css')
    .options({ processCssUrls: false })
