{{ define "header" }}
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta charset="utf-8">
    <title>{{ if .customTitle }}{{ .customTitle }}{{ else }}Wisply{{ end }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">

    <meta name="msvalidate.01" content="C26A233D070F9BCB0ADD16FD3F6F3FA0" />
    <meta name="google-site-verification" content="y1ewY3wV1FhvkSa0xTr2PhWLY7W_j8HA1EfXs7KeM2g" />

    <link rel="icon" href="/static/img/wisply/favicon.ico">

    <link rel="stylesheet" href="/static/3rd_party/product/paper/bootstrap.css">
    <link rel="stylesheet" href="/static/3rd_party/product/paper/bootswatch.min.css">

    <link rel="stylesheet" href="/static/css/wisply/print.css" media="print">

    <link rel="stylesheet" href="/static/css/wisply/wisply.css">
</head>
{{ end }}
