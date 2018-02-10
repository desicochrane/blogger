const mix = require('laravel-mix')

mix.js('src/_js/blog.js', 'src/js')
mix.sass('src/_scss/blog.scss', 'src/css')
mix.copy('src/css/blog.css', 'docs/css/blog.css')
mix.copy('src/js/blog.js', 'docs/js/blog.js')
mix.options({ processCssUrls: false })
