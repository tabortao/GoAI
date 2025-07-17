本站favicon使用[realfavicongenerator](https://realfavicongenerator.net) 生成的，非常感谢。

1. Extract this package in <web site>/web/favicon/. If your site is http://www.example.com, you should be able to access a file named http://www.example.com/web/favicon/favicon.ico

2. Insert the following code in the <head> section of your pages:
	
HTML代码：
```html
<link rel="icon" type="image/png" href="/web/favicon/favicon-96x96.png" sizes="96x96" />
<link rel="icon" type="image/svg+xml" href="/web/favicon/favicon.svg" />
<link rel="shortcut icon" href="/web/favicon/favicon.ico" />
<link rel="apple-touch-icon" sizes="180x180" href="/web/favicon/apple-touch-icon.png" />
<meta name="apple-mobile-web-app-title" content="GoAI" />
<link rel="manifest" href="/web/favicon/site.webmanifest" />
```
React代码：
```html
<link rel="icon" type="image/png" href="%PUBLIC_URL%/web/favicon/favicon-96x96.png" sizes="96x96" />
<link rel="icon" type="image/svg+xml" href="%PUBLIC_URL%/web/favicon/favicon.svg" />
<link rel="shortcut icon" href="%PUBLIC_URL%/web/favicon/favicon.ico" />
<link rel="apple-touch-icon" sizes="180x180" href="%PUBLIC_URL%/web/favicon/apple-touch-icon.png" />
<meta name="apple-mobile-web-app-title" content="GoAI" />
<link rel="manifest" href="%PUBLIC_URL%/web/favicon/site.webmanifest" />
```